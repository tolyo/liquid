-- +goose Up

-- +goose StatementBegin
CREATE OR REPLACE FUNCTION generate_tiger_id() 
RETURNS BIGINT AS $$
DECLARE
    next_id BIGINT;
BEGIN
    -- Lock the table for update to avoid race conditions
    SELECT current_id INTO next_id
    FROM tigerbeetle_id_counter
    FOR UPDATE;

    -- Increment the current ID
    next_id := next_id + 1;

    -- Update the table with the new ID
    UPDATE tigerbeetle_id_counter
    SET current_id = next_id;

    -- Return the new ID
    RETURN next_id;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd

-- +goose Down
DROP FUNCTION IF EXISTS generate_tiger_id();
