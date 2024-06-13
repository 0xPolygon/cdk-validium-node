-- +migrate Up

CREATE TABLE IF NOT EXISTS state.batch_data_backup
(
    batch_num  BIGINT,
    data       BYTEA,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (batch_num, created_at)
);

-- +migrate StatementBegin
CREATE OR REPLACE FUNCTION backup_batch() RETURNS trigger AS $$
    BEGIN
        INSERT  INTO state.batch_data_backup (batch_num,     data) 
                VALUES                       (OLD.batch_num, OLD.raw_txs_data)
                ON CONFLICT (batch_num, created_at) DO UPDATE SET
                    data = EXCLUDED.data;
        RETURN OLD;
    END;
$$
LANGUAGE plpgsql;
-- +migrate StatementEnd

CREATE TRIGGER backup_batch
   BEFORE DELETE ON state.batch FOR EACH ROW
   EXECUTE PROCEDURE backup_batch();

-- +migrate Down

DROP TRIGGER IF EXISTS backup_batch ON state.batch;
DROP FUNCTION IF EXISTS backup_batch();
DROP TABLE IF EXISTS state.batch_data_backup;
