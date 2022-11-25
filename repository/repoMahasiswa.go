package repository

import (
	"fmt"
	"mahasiswa/model"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]model.Mahasiswa, error)
	FindByID(ID int) (model.Mahasiswa, error)
	Create(mahasiswa model.Mahasiswa) (model.Mahasiswa, error)
	Update(mahasiswa model.Mahasiswa) (model.Mahasiswa, error)
	// Delete(mahasiswa Mahasiswa) (Mahasiswa, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]model.Mahasiswa, error) {
	var mahasiswa []model.Mahasiswa

	err := r.db.Find(&mahasiswa).Error
	fmt.Print(&mahasiswa)
	return mahasiswa, err
}

func (r *repository) FindByID(ID int) (model.Mahasiswa, error) {
	var mahasiswa model.Mahasiswa

	err := r.db.Find(&mahasiswa, ID).Error
	// fmt.Println(&book.Discount)
	return mahasiswa, err
}

func (r *repository) Create(mahasiswa model.Mahasiswa) (model.Mahasiswa, error) {
	err := r.db.Create(&mahasiswa).Error

	return mahasiswa, err
}

func (r *repository) Update(mahasiswa model.Mahasiswa) (model.Mahasiswa, error) {
	err := r.db.Save(&mahasiswa).Error

	return mahasiswa, err
}

// func (r *repository) Delete(book Book) (Book, error) {
// 	err := r.db.Delete(&book).Error

// 	return book, err
// }
