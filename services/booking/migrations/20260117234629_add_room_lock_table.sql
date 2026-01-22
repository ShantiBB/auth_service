-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS room_lock (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    room_id UUID NOT NULL,
    booking_id UUID NOT NULL REFERENCES booking(id) ON DELETE CASCADE,
    stay_range DATERANGE NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    expires_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_room_lock_room ON room_lock(room_id);
CREATE INDEX IF NOT EXISTS idx_room_lock_booking ON room_lock(booking_id);
CREATE INDEX IF NOT EXISTS idx_room_lock_expires ON room_lock(expires_at) WHERE expires_at IS NOT NULL;

ALTER TABLE room_lock
    ADD CONSTRAINT room_lock_no_overlap
        EXCLUDE USING gist (
        room_id WITH =,
        stay_range WITH &&
        )
        WHERE (is_active = TRUE);

ALTER TABLE room_lock
    ADD CONSTRAINT room_lock_range_valid
        CHECK (upper(stay_range) > lower(stay_range));

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE IF EXISTS room_lock
    DROP CONSTRAINT IF EXISTS room_lock_no_overlap;

ALTER TABLE IF EXISTS room_lock
    DROP CONSTRAINT IF EXISTS room_lock_range_valid;

DROP INDEX IF EXISTS idx_room_lock_room;
DROP INDEX IF EXISTS idx_room_lock_booking;
DROP INDEX IF EXISTS idx_room_lock_expires;

DROP TABLE IF EXISTS room_lock;
-- +goose StatementEnd
