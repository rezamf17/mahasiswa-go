package model

import "time"

type Mahasiswa struct {
	Id             int
	NPM            string
	Nama           string
	Kelas          string
	Jurusan        string
	NomorHandphone string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
