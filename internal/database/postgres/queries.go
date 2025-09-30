package postgres

var UserCreate = `
	INSERT INTO users (username, first_name, last_name, email, description, password)
	VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING id
`

var UserGet = `
	SELECT username, first_name, last_name, email, description
	FROM users 
	WHERE id = $1
`

var UserGetAll = `
	SELECT id, username, first_name, last_name, email, description
	FROM users
`

var UserUpdate = `
	UPDATE users
	SET username = $1, first_name = $2, last_name = $3, email = $4, description = $5
	WHERE id = $6
`

var UserDelete = `
	DELETE FROM users
	WHERE id = $1
`
