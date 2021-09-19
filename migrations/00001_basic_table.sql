-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE data (
    id BIGINT NOT NULL,
    plugin VARCHAR(32) NOT NULL,
    payload JSONB NULL NULL,
    UNIQUE (plugin, id)
) PARTITION BY LIST (plugin);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP TABLE data;
