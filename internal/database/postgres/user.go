package postgres

import (
	"fmt"

	"auth_service/internal/entity"
)

type UserRepository interface {
	Create(user entity.UserCreate) (*entity.User, error)
	Get(id int64) (*entity.User, error)
	GetAll() ([]entity.User, error)
	Update(user entity.User) (*entity.User, error)
	Delete(id int64) error
}

func (p *Repository) Create(user entity.UserCreate) (*entity.User, error) {
	newUser := user.ToUserRead()
	rows, err := p.db.NamedQuery(UserCreate, user)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	if rows.Next() {
		if err = rows.Scan(&newUser.ID); err != nil {
			return nil, err
		}
	}

	return newUser, nil
}

func (p *Repository) Get(id int64) (*entity.User, error) {
	var user entity.User

	if err := p.db.Get(&user, UserGet, id); err != nil {
		return nil, err
	}

	return &user, nil
}

func (p *Repository) GetAll() ([]entity.User, error) {
	var users []entity.User

	if err := p.db.Select(&users, UserGetAll); err != nil {
		return nil, err
	}

	return users, nil
}

func (p *Repository) Update(user entity.User) (*entity.User, error) {
	rows, err := p.db.NamedExec(UserUpdate, &user)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := rows.RowsAffected()
	if err != nil {
		return nil, err
	} else if rowsAffected == 0 {
		return nil, fmt.Errorf("user with id %d is not found", user.ID)
	}

	updateUser, err := p.Get(user.ID)
	if err != nil {
		return nil, err
	}

	return updateUser, nil
}

func (p *Repository) Delete(id int64) error {
	rows, err := p.db.Exec(UserDelete, id)
	if err != nil {
		return err
	}

	rowsAffected, err := rows.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("user with id %d is not found", id)
	}

	return nil
}
