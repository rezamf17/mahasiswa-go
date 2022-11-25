package request

type MahasiswaRequest struct {
	NPM            string `json:"npm" binding:"required"`
	Nama           string `json:"nama" binding:"required"`
	Kelas          string `json:"kelas" binding:"required"`
	Jurusan        string `json:"jurusan" binding:"required"`
	NomorHandphone string `json:"nomorHandphone" binding:"required"`
}
