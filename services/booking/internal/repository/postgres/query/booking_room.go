package query

const (
	BookingRoomCreate = `
		INSERT INTO booking_room (booking_id, room_id, adults, children, price_per_night)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, created_at;`

	BookingRoomGetAll = `
		SELECT
		  id,
		  booking_id,
		  room_id,
		  adults,
		  children,
		  price_per_night,
		  created_at
		FROM booking_room
		WHERE booking_id = $1
		ORDER BY created_at;`

	BookingRoomGetByID = `
		SELECT
		  id,
		  booking_id,
		  room_id,
		  adults,
		  children,
		  price_per_night,
		  created_at
		FROM booking_room
		WHERE id = $1;`

	BookingGuestsUpdateByID = `
		UPDATE booking_room
		SET
		  adults = $2,
		  children = $3
		WHERE id = $1;`

	BookingRoomDeleteByID = `
		DELETE FROM booking_room
		WHERE id = $1;`
)
