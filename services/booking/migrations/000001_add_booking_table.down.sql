DROP TRIGGER IF EXISTS update_hotels_updated_at ON booking;

DROP FUNCTION IF EXISTS update_updated_at_column();

DROP INDEX IF EXISTS idx_booking_user;
DROP INDEX IF EXISTS idx_booking_hotel;
DROP INDEX IF EXISTS idx_booking_status;

DROP TABLE IF EXISTS booking;

DROP TYPE IF EXISTS booking_status;

DROP EXTENSION IF EXISTS btree_gist;
DROP EXTENSION IF EXISTS pgcrypto;
