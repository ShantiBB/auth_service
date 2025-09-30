package postgres

import (
	"context"
	"fmt"

	"auth_service/internal/domain/models"
)

func (p *Repository) CreateUser(ctx context.Context, u models.UserCreate) (*models.User, error) {
	newUser := u.ToUserRead()
	err := p.db.QueryRow(
		ctx, UserCreate, u.Username, u.FirstName, u.LastName, u.Email, u.Description, u.Password,
	).Scan(&newUser.ID)
	if err != nil {
		return nil, err
	}

	return &newUser, nil
}

func (p *Repository) GetUser(ctx context.Context, id int64) (*models.User, error) {
	u := models.User{ID: id}
	if err := p.db.QueryRow(ctx, UserGet, id).Scan(
		&u.Username, &u.FirstName, &u.LastName, &u.Email, &u.Description,
	); err != nil {
		return nil, err
	}

	return &u, nil
}

func (p *Repository) GetAllUsers(ctx context.Context) ([]models.User, error) {
	var users []models.User

	rows, err := p.db.Query(ctx, UserGetAll)
	if err != nil {
		return nil, err
	}

	var u models.User
	for rows.Next() {
		err = rows.Scan(&u.ID, &u.Username, &u.FirstName, &u.LastName, &u.Email, &u.Description)
		if err != nil {
			return nil, err
		}

		users = append(users, u)
	}

	return users, nil
}

func (p *Repository) UpdateUser(ctx context.Context, u models.User) error {
	rows, err := p.db.Exec(
		ctx, UserUpdate, u.Username, u.FirstName, u.LastName, u.Email, u.Description, u.ID,
	)
	if err != nil {
		return err
	}

	rowsAffected := rows.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("user with id %d is not found", u.ID)
	}

	return nil
}

func (p *Repository) DeleteUser(ctx context.Context, id int64) error {
	rows, err := p.db.Exec(ctx, UserDelete, id)
	if err != nil {
		return err
	}

	rowsAffected := rows.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("user with id %d is not found", id)
	}

	return nil
}
