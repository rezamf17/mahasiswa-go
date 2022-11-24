package response

type MahasiswaResponse struct {
	ID             int    `json:"id"`
	NPM            string `json:"npm"`
	Nama           string `json:"nama"`
	Kelas          string `json:"kelas"`
	Jurusan        string `json:"jurusan"`
	NomorHandphone string `json:"nomorHandphone"`
}
