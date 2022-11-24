package repository

import (
	"fmt"
	"mahasiswa/model"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]model.Mahasiswa, error)
	// FindByID(ID int) (Mahasiswa, error)
	// Create(mahasiswa Mahasiswa) (Mahasiswa, error)
	// Update(mahasiswa Mahasiswa) (Mahasiswa, error)
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

// func (r *repository) FindByID(ID int) (Book, error) {
// 	var book Book

// 	err := r.db.Find(&book, ID).Error
// 	// fmt.Println(&book.Discount)
// 	return book, err
// }

// func (r *repository) Create(book Book) (Book, error) {
// 	err := r.db.Create(&book).Error

// 	return book, err
// }

// func (r *repository) Update(book Book) (Book, error) {
// 	err := r.db.Save(&book).Error

// 	return book, err
// }

// func (r *repository) Delete(book Book) (Book, error) {
// 	err := r.db.Delete(&book).Error

// 	return book, err
// }
