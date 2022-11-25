package services

import (
	"mahasiswa/model"
	"mahasiswa/repository"
	"mahasiswa/request"
)

type Service interface {
	FindAll() ([]model.Mahasiswa, error)
	FindByID(ID int) (model.Mahasiswa, error)
	Create(mahasiswaRequest request.MahasiswaRequest) (model.Mahasiswa, error)
	Update(ID int, mahasiswaRequest request.MahasiswaRequest) (model.Mahasiswa, error)
	Delete(ID int, mahasiswaRequest request.MahasiswaRequest) (model.Mahasiswa, error)
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

func (s *service) FindByID(ID int) (model.Mahasiswa, error) {
	mahasiswa, err := s.repository.FindByID(ID)
	return mahasiswa, err
}

func (s *service) Create(mahasiswaRequest request.MahasiswaRequest) (model.Mahasiswa, error) {
	// price, _ := mahasiswaRequest.Price.Int64()
	mahasiswa := model.Mahasiswa{
		NPM:            mahasiswaRequest.NPM,
		Nama:           string(mahasiswaRequest.Nama),
		Kelas:          mahasiswaRequest.Kelas,
		Jurusan:        mahasiswaRequest.Jurusan,
		NomorHandphone: mahasiswaRequest.NomorHandphone,
	}
	newMahasiswa, err := s.repository.Create(mahasiswa)
	return newMahasiswa, err
}

func (s *service) Update(ID int, mahasiswaRequest request.MahasiswaRequest) (model.Mahasiswa, error) {
	mahasiswa, err := s.repository.FindByID(ID)
	// price, _ := bookRequest.Price.Int64()

	mahasiswa.NPM = mahasiswaRequest.NPM
	mahasiswa.Nama = mahasiswaRequest.Nama
	mahasiswa.Kelas = mahasiswaRequest.Kelas
	mahasiswa.Jurusan = mahasiswaRequest.Jurusan
	mahasiswa.NomorHandphone = mahasiswaRequest.NomorHandphone

	newMahasiswa, err := s.repository.Update(mahasiswa)
	return newMahasiswa, err
}

func (s *service) Delete(ID int, mahasiswaRequest request.MahasiswaRequest) (model.Mahasiswa, error) {
	mahasiswa, err := s.repository.FindByID(ID)
	newMahasiswa, err := s.repository.Delete(mahasiswa)
	return newMahasiswa, err
}
