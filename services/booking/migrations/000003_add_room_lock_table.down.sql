ALTER TABLE IF EXISTS room_lock
    DROP CONSTRAINT IF EXISTS room_lock_no_overlap;

ALTER TABLE IF EXISTS room_lock
    DROP CONSTRAINT IF EXISTS room_lock_range_valid;

DROP INDEX IF EXISTS idx_room_lock_room;
DROP INDEX IF EXISTS idx_room_lock_booking;
DROP INDEX IF EXISTS idx_room_lock_expires;

DROP TABLE IF EXISTS room_lock;