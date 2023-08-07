package main

import (
	"context"
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/0xPolygon/supernets2-node/config"
	"github.com/0xPolygon/supernets2-node/pool"
	"github.com/0xPolygon/supernets2-node/pool/pgpoolstorage"
	"github.com/ethereum/go-ethereum/common"
	"github.com/urfave/cli/v2"
)

var (
	policyFlag = cli.StringFlag{
		Name:     "policy",
		Aliases:  []string{"p"},
		Usage:    "Name of policy to operate on",
		Required: true,
	}
	csvFlag = cli.StringFlag{
		Name:     "csv",
		Usage:    "CSV file with addresses",
		Required: false,
	}
	allowFlag = cli.BoolFlag{
		Name:     "allow",
		Usage:    "Update policy action to allow/deny by default",
		Required: false,
	}

	policyActionFlags = []cli.Flag{&policyFlag}
)

var policyCommands = cli.Command{
	Name:   "policy",
	Usage:  "View, update, and apply policies",
	Action: describePolicies,
	Flags:  []cli.Flag{&configFileFlag},
	Subcommands: []*cli.Command{
		{
			Name:   "update",
			Action: updatePolicy,
			Flags:  append(policyActionFlags, &allowFlag),
		}, {
			Name:   "describe",
			Action: describePolicies,
		}, {
			Name:   "add",
			Action: addAcl,
			Flags:  append(policyActionFlags, &csvFlag),
		}, {
			Name:   "remove",
			Action: removeAcl,
			Flags:  append(policyActionFlags, &csvFlag),
		}, {
			Name:   "clear",
			Action: clearAcl,
			Flags:  policyActionFlags,
		}, {
			Name:   "list",
			Action: listAcl,
			Flags:  policyActionFlags,
		},
	},
}

func describePolicies(cli *cli.Context) error {
	_, db, err := configAndStorage(cli)
	if err != nil {
		return err
	}
	list, err := db.DescribePolicies(context.Background())
	if err != nil {
		return err
	}

	fmt.Printf("%7s: %s\n", "Policy", "Default")
	for _, p := range list {
		fmt.Printf("%7s: %s\n", p.Name, p.Desc())
	}

	return nil
}

func updatePolicy(cli *cli.Context) error {
	_, db, err := configAndStorage(cli)
	if err != nil {
		return err
	}
	policy, err := resolvePolicy(cli)
	if err != nil {
		return err
	}
	var allow = true
	if !cli.IsSet("allow") {
		return errors.New("supply one policy action [--allow=true or --allow=false]")
	}
	allow = cli.Bool("allow")
	err = db.UpdatePolicy(context.Background(), policy, allow)
	if err != nil {
		return err
	}
	return nil
}

func addAcl(cli *cli.Context) error {
	_, db, err := configAndStorage(cli)
	if err != nil {
		return err
	}
	policy, addresses, err := requirePolicyAndAddresses(cli)
	if err != nil {
		return err
	}
	err = db.AddAddressesToPolicy(context.Background(), policy, addresses)
	if err != nil {
		return err
	}
	return nil
}

func removeAcl(cli *cli.Context) error {
	_, db, err := configAndStorage(cli)
	if err != nil {
		return err
	}
	policy, addresses, err := requirePolicyAndAddresses(cli)
	if err != nil {
		return err
	}
	err = db.RemoveAddressesFromPolicy(context.Background(), policy, addresses)
	if err != nil {
		return err
	}
	return nil
}

func clearAcl(cli *cli.Context) error {
	_, db, err := configAndStorage(cli)
	if err != nil {
		return err
	}
	policy, err := resolvePolicy(cli)
	if err != nil {
		return err
	}
	err = db.ClearPolicy(context.Background(), policy)
	if err != nil {
		return err
	}
	return nil
}

func listAcl(cli *cli.Context) error {
	_, db, err := configAndStorage(cli)
	if err != nil {
		return err
	}
	policyName, err := resolvePolicy(cli)
	if err != nil {
		return err
	}

	policy, err := db.DescribePolicy(context.Background(), policyName)
	if err != nil {
		return err
	}
	fmt.Printf("%s: %s\n", "Policy", policy.Name)
	fmt.Printf("%s: %s\n", "Default", policy.Desc())

	query, err := resolveAddresses(cli, false)
	if err != nil {
		return nil
	}
	list, err := db.ListAcl(context.Background(), policyName, query)
	if err != nil {
		return err
	}
	listAction := "Denied"
	if !policy.Allow {
		listAction = "Allowed"
	}
	fmt.Printf("%s addrs:\n", listAction)
	for _, address := range list {
		fmt.Println(address.Hex())
	}
	return nil
}

func configAndStorage(cli *cli.Context) (*config.Config, *pgpoolstorage.PostgresPoolStorage, error) {
	c, err := config.Load(cli, false)
	if err != nil {
		return nil, nil, err
	}
	setupLog(c.Log)

	db, err := pgpoolstorage.NewPostgresPoolStorage(c.Pool.DB)
	if err != nil {
		return nil, nil, err
	}
	return c, db, nil
}

func requirePolicyAndAddresses(cli *cli.Context) (pool.PolicyName, []common.Address, error) {
	policy, err := resolvePolicy(cli)
	if err != nil {
		return "", nil, err
	}
	addresses, err := resolveAddresses(cli, true)
	if err != nil {
		return "", nil, err
	}
	return policy, addresses, nil
}

func resolvePolicy(cli *cli.Context) (pool.PolicyName, error) {
	policy := cli.String("policy")
	if policy == "" {
		return "", nil
	}
	if !pool.IsPolicy(policy) {
		return "", fmt.Errorf("invalid policy name: %s", policy)
	}
	return pool.PolicyName(policy), nil
}

func resolveAddresses(cli *cli.Context, failIfEmpty bool) ([]common.Address, error) {
	var set = make(map[common.Address]struct{})
	if cli.IsSet("csv") {
		file := cli.String("csv")
		fd, err := os.Open(file)
		if err != nil {
			return nil, err
		}
		defer func(fd *os.File) {
			_ = fd.Close()
		}(fd)

		fileReader := csv.NewReader(fd)
		records, err := fileReader.ReadAll()

		if err != nil {
			return nil, err
		}
		for _, row := range records {
			for _, cell := range row {
				hex := strings.TrimSpace(cell)
				set[common.HexToAddress(hex)] = struct{}{}
			}
		}
	}

	for _, a := range cli.Args().Slice() {
		a = strings.TrimSpace(a)
		a = strings.Trim(a, ",|")
		if !strings.HasPrefix(a, "0x") {
			a = "0x" + a
		}
		set[common.HexToAddress(a)] = struct{}{}
	}
	var ret []common.Address
	for a := range set {
		ret = append(ret, a)
	}
	if failIfEmpty && len(ret) == 0 {
		return nil, errors.New("no addresses given")
	}
	return ret, nil
}
