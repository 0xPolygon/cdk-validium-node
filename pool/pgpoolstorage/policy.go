package pgpoolstorage

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/0xPolygon/supernets2-node/pool"
	"github.com/ethereum/go-ethereum/common"
	"github.com/jackc/pgx/v4"
)

// CheckPolicy returns the rule for the policy if the address is not associated with the rule (default), or the opposite
// of the rule if it is. This allows the rule to act as an allow or deny list.
func (p *PostgresPoolStorage) CheckPolicy(ctx context.Context, policy pool.PolicyName, address common.Address) (bool, error) {
	sql := `SELECT 
				CASE WHEN a.address is null THEN 
					p.allow 
				ELSE 
					NOT p.allow 
				END 
			FROM pool.policy p 
				LEFT JOIN pool.acl a 
					ON p.name = a.policy 
					AND a.address = $1 
			WHERE p.name = $2`

	rows, err := p.db.Query(ctx, sql, address.Hex(), policy)

	if errors.Is(err, pgx.ErrNoRows) {
		return false, pool.ErrNotFound
	} else if err != nil {
		return false, err
	}
	if !rows.Next() { // should always be a row if the policy exists
		return false, nil
	}

	var allow bool
	err = rows.Scan(&allow)
	if err != nil {
		return false, err
	}
	return allow, nil
}

func (p *PostgresPoolStorage) UpdatePolicy(ctx context.Context, policy pool.PolicyName, allow bool) error {
	sql := "UPDATE pool.policy SET allow = $1 WHERE name = $2"
	_, err := p.db.Exec(ctx, sql, allow, string(policy))
	if err != nil {
		return err
	}
	return nil
}

func (p *PostgresPoolStorage) AddAddressesToPolicy(ctx context.Context, policy pool.PolicyName, addresses []common.Address) error {
	sql := "INSERT INTO pool.acl (policy, address) VALUES ($1, $2) ON CONFLICT DO NOTHING"
	tx, err := p.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer func(tx pgx.Tx, ctx context.Context) {
		_ = tx.Rollback(ctx)
	}(tx, ctx)

	for _, a := range addresses {
		_, err = tx.Exec(ctx, sql, policy, a.Hex())
		if err != nil {
			return err
		}
	}
	err = tx.Commit(ctx)
	if err != nil {
		return nil
	}
	return nil
}

func (p *PostgresPoolStorage) RemoveAddressesFromPolicy(ctx context.Context, policy pool.PolicyName, addresses []common.Address) error {
	sql := "DELETE FROM pool.acl WHERE policy = $1 AND address = $2"
	tx, err := p.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer func(tx pgx.Tx, ctx context.Context) {
		_ = tx.Rollback(ctx)
	}(tx, ctx)

	for _, a := range addresses {
		_, err = tx.Exec(ctx, sql, policy, a.Hex())
		if err != nil {
			return err
		}
	}
	err = tx.Commit(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (p *PostgresPoolStorage) ClearPolicy(ctx context.Context, policy pool.PolicyName) error {
	sql := "DELETE FROM pool.acl WHERE policy = $1"
	_, err := p.db.Exec(ctx, sql, policy)
	if err != nil {
		return err
	}
	return nil
}

func (p *PostgresPoolStorage) DescribePolicies(ctx context.Context) ([]pool.Policy, error) {
	sql := "SELECT name, allow FROM pool.policy"
	rows, err := p.db.Query(ctx, sql)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		} else {
			return nil, err
		}
	}
	defer rows.Close()

	var list []pool.Policy
	for rows.Next() {
		var name string
		var allow bool
		err = rows.Scan(&name, &allow)
		if err != nil {
			return nil, err
		}
		if pool.IsPolicy(name) { // skip unknown
			p := pool.Policy{
				Name:  pool.PolicyName(name),
				Allow: allow,
			}
			list = append(list, p)
		}
	}
	return list, nil
}

func (p *PostgresPoolStorage) DescribePolicy(ctx context.Context, name pool.PolicyName) (pool.Policy, error) {
	sql := "SELECT name, allow FROM pool.policy WHERE name = $1 LIMIT 1"
	row := p.db.QueryRow(ctx, sql, name)
	var (
		pName string
		allow bool
	)
	err := row.Scan(&pName, &allow)
	if err != nil {
		return pool.Policy{}, err
	}
	return pool.Policy{
		Name:  pool.PolicyName(pName),
		Allow: allow,
	}, nil
}

func (p *PostgresPoolStorage) ListAcl(
	ctx context.Context, policy pool.PolicyName, query []common.Address) ([]common.Address, error) {
	sql := "SELECT address FROM pool.acl WHERE policy = $1"

	if len(query) > 0 {
		var addrs []string
		for _, a := range query {
			addrs = append(addrs, a.Hex())
		}
		sql = sql + fmt.Sprintf(" IN (%v)", strings.Join(addrs, ","))
	}

	rows, err := p.db.Query(ctx, sql, string(policy))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		} else {
			return nil, err
		}
	}
	defer rows.Close()

	var addresses []common.Address
	for rows.Next() {
		var addr string
		err = rows.Scan(&addr)
		if err != nil {
			return nil, err
		}
		addresses = append(addresses, common.HexToAddress(addr))
	}
	return addresses, nil
}
