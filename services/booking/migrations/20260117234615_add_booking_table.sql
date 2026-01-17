-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS pgcrypto;
CREATE EXTENSION IF NOT EXISTS btree_gist;

CREATE TYPE booking_status AS ENUM ('pending', 'confirmed', 'cancelled');

CREATE TABLE IF NOT EXISTS booking (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    user_id BIGINT NOT NULL,
    hotel_id UUID NOT NULL,

    check_in  DATE NOT NULL,
    check_out DATE NOT NULL CHECK (check_out > check_in),

    status booking_status NOT NULL DEFAULT 'pending',

    guest_name  TEXT NOT NULL,
    guest_email TEXT,
    guest_phone TEXT,

    currency CHAR(3) NOT NULL CHECK (currency ~ '^[A-Z]{3}$'),
    expected_total_amount NUMERIC(12,2) NOT NULL CHECK (expected_total_amount >= 0),
    final_total_amount NUMERIC(12,2) NOT NULL CHECK (final_total_amount >= 0),

    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_booking_user ON booking(user_id, created_at DESC);
CREATE INDEX IF NOT EXISTS idx_booking_hotel ON booking(hotel_id, created_at DESC);
CREATE INDEX IF NOT EXISTS idx_booking_status ON booking(status);

CREATE OR REPLACE FUNCTION update_updated_at_column()
    RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_hotels_updated_at
    BEFORE UPDATE ON booking
    FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TRIGGER IF EXISTS update_hotels_updated_at ON booking;

DROP FUNCTION IF EXISTS update_updated_at_column();

DROP INDEX IF EXISTS idx_booking_user;
DROP INDEX IF EXISTS idx_booking_hotel;
DROP INDEX IF EXISTS idx_booking_status;

DROP TABLE IF EXISTS booking;

DROP TYPE IF EXISTS booking_status;

DROP EXTENSION IF EXISTS btree_gist;
DROP EXTENSION IF EXISTS pgcrypto;
-- +goose StatementEnd
