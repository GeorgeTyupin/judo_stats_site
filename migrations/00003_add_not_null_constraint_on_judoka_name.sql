-- +goose Up
SELECT 'up SQL query';

ALTER TABLE judokas
    ALTER COLUMN name SET NOT NULL;

-- +goose Down
SELECT 'down SQL query';

ALTER TABLE judokas
    ALTER COLUMN name DROP NOT NULL;
