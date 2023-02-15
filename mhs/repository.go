package mhs

import (
	"database/sql"
	"errors"
)

type RepositoryInterface interface {
	Create(mhs Mahasiswa) error
	Update(mhs InputMhs, nim, userID int) error
	Delete(nim, userID int)
	FindAll() ([]Mahasiswa, error)
	FindByNim(nim, userID int) (Mahasiswa, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *repository {
	return &repository{db}
}

func (r *repository) Create(mhs Mahasiswa) error {
	sql := "INSERT INTO data_mhs (nim, nama, email, alamat, created_at, updated_at, user_id) values (?, ?, ?, ?, NOW(), NOW(),?)"
	row := r.db.QueryRow(sql, mhs.Nim, mhs.Nama, mhs.Email, mhs.Alamat, mhs.ID)

	if row.Err() != nil {
		return errors.New("error when insert data")
	}

	return nil
}

func (r *repository) Update(mhs InputMhs, nim, userID int) error {
	sql := "UPDATE data_mhs SET nama = ?, email = ?, alamat = ?, updated_at = NOW() WHERE nim = ? AND user_id = ?"
	row := r.db.QueryRow(sql, mhs.Nama, mhs.Email, mhs.Alamat, nim, userID)

	if row.Err() != nil {
		return errors.New("error when update data")
	}

	return nil
}

func (r *repository) Delete(nim, userID int) {
	sql := "DELETE FROM data_mhs WHERE nim = ? AND user_id = ?"
	r.db.Exec(sql, nim, userID)
}

func (r *repository) FindAll() ([]Mahasiswa, error) {
	sql := "SELECT nim, nama FROM data_mhs"
	rows, err := r.db.Query(sql)
	if err != nil {
		return nil, err
	}

	var data []Mahasiswa

	for rows.Next() {
		var dataTmp Mahasiswa
		err := rows.Scan(&dataTmp.Nim, &dataTmp.Nama)
		if err != nil {
			return data, err
		}

		data = append(data, dataTmp)
	}

	return data, nil
}

func (r *repository) FindByNim(nim, userID int) (Mahasiswa, error) {
	sql := "SELECT nim, nama, email, alamat FROM data_mhs WHERE nim = ? AND user_id = ?"
	row := r.db.QueryRow(sql, nim, userID)

	var data Mahasiswa

	err := row.Scan(&data.Nim, &data.Nama, &data.Email, &data.Alamat)
	if err != nil {
		return data, err
	}

	return data, nil
}
