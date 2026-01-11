package query

const (
	CreateBookingRooms = `
		WITH input AS (
		  SELECT
			$1::uuid AS booking_id,
			unnest($2::uuid[])    AS room_id,
			unnest($3::int[])     AS adults,
			unnest($4::int[])     AS children,
			unnest($5::numeric[]) AS price_per_night
		)
		INSERT INTO booking_room (booking_id, room_id, adults, children, price_per_night)
		SELECT booking_id, room_id, adults, children, price_per_night
		FROM input
		RETURNING id, booking_id, room_id, adults, children, price_per_night, created_at;`

	GetBookingRoomsByBookingID = `
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

	GetBookingRoomByID = `
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

	UpdateBookingRoomGuestCountsByID = `
		UPDATE booking_room
		SET
		  adults = $2,
		  children = $3
		WHERE id = $1;`

	DeleteBookingRoomByID = `
		DELETE FROM booking_room
		WHERE id = $1;`
)
