package services

import (
	"mahasiswa/model"
	"mahasiswa/repository"
)

type Service interface {
	FindAll() ([]model.Mahasiswa, error)
	// FindByID(ID int) (Book, error)
	// Create(bookRequest BookRequest) (Book, error)
	// Update(ID int, bookRequest BookRequest) (Book, error)
	// Delete(ID int, bookRequest BookRequest) (Book, error)
}

type service struct {
	repository repository.Repository
}

func NewService(repository repository.Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]model.Mahasiswa, error) {
	mahasiswa, err := s.repository.FindAll()
	return mahasiswa, err
}

// func (s *service) FindByID(ID int) (Book, error) {
// 	book, err := s.repository.FindByID(ID)
// 	return book, err
// }

// func (s *service) Create(bookRequest BookRequest) (Book, error) {
// 	price, _ := bookRequest.Price.Int64()
// 	book := Book{
// 		Title:       bookRequest.Title,
// 		Price:       int(price),
// 		Description: bookRequest.Description,
// 		Rating:      bookRequest.Rating,
// 		Discount:    bookRequest.Discount,
// 	}
// 	newBook, err := s.repository.Create(book)
// 	return newBook, err
// }

// func (s *service) Update(ID int, bookRequest BookRequest) (Book, error) {
// 	book, err := s.repository.FindByID(ID)
// 	price, _ := bookRequest.Price.Int64()

// 	book.Title = bookRequest.Title
// 	book.Price = int(price)
// 	book.Description = bookRequest.Description
// 	book.Rating = bookRequest.Rating
// 	book.Discount = bookRequest.Discount

// 	newBook, err := s.repository.Update(book)
// 	return newBook, err
// }

// func (s *service) Delete(ID int, bookRequest BookRequest) (Book, error) {
// 	book, err := s.repository.FindByID(ID)
// 	newBook, err := s.repository.Delete(book)
// 	return newBook, err
// }
