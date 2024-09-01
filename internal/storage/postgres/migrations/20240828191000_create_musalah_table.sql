-- +goose Up
-- +goose StatementBegin
CREATE TABLE
    musalahs (
        id serial PRIMARY KEY,
        NAME VARCHAR(50) NOT NULL,
        created_at timestamptz DEFAULT CLOCK_TIMESTAMP(),
        updated_at timestamptz,
        deleted_at timestamptz
    );

CREATE TRIGGER sync_musalah_updated_at BEFORE
UPDATE ON musalahs FOR EACH ROW
EXECUTE PROCEDURE sync_updated_at_column ();

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE musalahs;

-- +goose StatementEnd