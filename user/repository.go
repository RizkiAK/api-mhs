package user

import (
	"database/sql"
	"errors"
)

type RepositoryInterface interface {
	Create(user User) error
	FindByNim(nim int) (User, error)
	UpdatePassword(nim int, password string) error
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *repository {
	return &repository{db}
}

func (r *repository) Create(user User) error {
	sql := "INSERT INTO user (nim, nama, password, created_at) values (?, ?, ?, NOW())"
	row := r.db.QueryRow(sql, user.Nim, user.Nama, user.Password)

	if row.Err() != nil {
		return errors.New("error when insert data")
	}

	return nil
}

func (r *repository) FindByNim(nim int) (User, error) {
	sql := "SELECT nim, nama, password FROM user WHERE nim = ?"
	row := r.db.QueryRow(sql, nim)

	var data User

	err := row.Scan(&data.Nim, &data.Nama, &data.Password)
	if err != nil {
		return data, err
	}

	return data, nil
}

func (r *repository) UpdatePassword(nim int, password string) error {
	sql := "UPDATE user SET password = ? WHERE nim = ?"
	row := r.db.QueryRow(sql, password, nim)

	if row.Err() != nil {
		return errors.New("error when update password")
	}

	return nil
}
