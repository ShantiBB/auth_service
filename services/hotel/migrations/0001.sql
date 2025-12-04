CREATE EXTENSION IF NOT EXISTS postgis;

CREATE TYPE room_type AS ENUM ('single', 'double', 'suite', 'deluxe', 'family', 'presidential');
CREATE TYPE room_status AS ENUM ('available', 'occupied', 'maintenance', 'cleaning');

CREATE TABLE hotel (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) UNIQUE NOT NULL,
    owner_id BIGINT NOT NULL,
    description TEXT,
    address TEXT NOT NULL,
    location GEOGRAPHY(Point, 4326) NOT NULL,
    rating NUMERIC(3,2) CHECK (rating >= 0 AND rating <= 5),
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE room (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    hotel_id UUID NOT NULL REFERENCES hotel(id) ON DELETE CASCADE,
    room_number VARCHAR(10) NOT NULL,
    type room_type NOT NULL,
    status room_status NOT NULL DEFAULT 'available',
    price NUMERIC(10,2) NOT NULL CHECK (price > 0),
    capacity INT NOT NULL CHECK (capacity > 0 AND capacity <= 10),
    area_sqm NUMERIC(6,2),
    floor INT CHECK (floor >= 0),
    description TEXT,
    amenities TEXT[],
    images TEXT[],
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,

    UNIQUE(hotel_id, room_number)
);

CREATE INDEX hotels_location_idx ON hotel USING GIST (location);
CREATE INDEX hotels_owner_idx ON hotel (owner_id);
CREATE INDEX hotels_name_idx ON hotel (name);
CREATE INDEX rooms_hotel_id_idx ON room (hotel_id);
CREATE INDEX rooms_type_idx ON room (type);
CREATE INDEX rooms_status_idx ON room (status) WHERE status = 'available';
CREATE INDEX rooms_price_idx ON room (price);

CREATE OR REPLACE FUNCTION update_updated_at_column()
    RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_hotels_updated_at
    BEFORE UPDATE ON hotel
    FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_rooms_updated_at
    BEFORE UPDATE ON room
    FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();
