-- +goose Up
-- +goose StatementBegin
CREATE TYPE jumaah_status AS ENUM('suggested', 'scheduled');

CREATE TABLE
    jumaahs (
        id SERIAL PRIMARY KEY,
        NAME VARCHAR(255) NOT NULL,
        musalah_id INT NOT NULL REFERENCES musalahs,
        account_id INT NOT NULL REFERENCES accounts,
        begins_at TIMESTAMP WITH TIME ZONE NOT NULL,
        status jumaah_status NOT NULL DEFAULT 'suggested',
        created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP WITH TIME ZONE,
        deleted_at TIMESTAMP WITH TIME ZONE
    );

CREATE TRIGGER sync_jumaah_updated_at BEFORE
UPDATE ON jumaahs FOR EACH ROW
EXECUTE PROCEDURE sync_updated_at_column ();

CREATE TABLE
    jumaah_attendees (
        jumaah_id INT NOT NULL REFERENCES jumaahs,
        account_id INT NOT NULL REFERENCES accounts,
        delayed_eta TIMESTAMP WITH TIME ZONE,
        created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP WITH TIME ZONE,
        deleted_at TIMESTAMP WITH TIME ZONE,
        PRIMARY KEY (jumaah_id, account_id)
    );

CREATE TRIGGER sync_jumaah_attendee_updated_at BEFORE
UPDATE ON jumaah_attendees FOR EACH ROW
EXECUTE PROCEDURE sync_updated_at_column ();

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE jumaah_attendees;

DROP TABLE jumaahs;

DROP TYPE jumaah_status;

-- +goose StatementEnd