-- +goose Up
-- +goose StatementBegin
CREATE TABLE
    accounts (
        id SERIAL PRIMARY KEY,
        NAME TEXT NOT NULL,
        created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP WITH TIME ZONE,
        deleted_at TIMESTAMP WITH TIME ZONE
    );

CREATE TRIGGER sync_account_updated_at BEFORE
UPDATE ON accounts FOR EACH ROW
EXECUTE PROCEDURE sync_updated_at_column ();

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE accounts;

-- +goose StatementEnd