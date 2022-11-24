package handler

import (
	"mahasiswa/model"
	"mahasiswa/response"
	"mahasiswa/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type mahasiswaHandler struct {
	mahasiswaService services.Service
}

func NewMahasiswaHandler(mahasiswaService services.Service) *mahasiswaHandler {
	return &mahasiswaHandler{mahasiswaService}
}

func (h *mahasiswaHandler) GetAllMahasiswa(ctx *gin.Context) {
	mahasiswaAll, err := h.mahasiswaService.FindAll()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	var mahasiswaResponse []response.MahasiswaResponse

	for _, b := range mahasiswaAll {
		bookResponse := convertResponse(b)
		mahasiswaResponse = append(mahasiswaResponse, bookResponse)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": mahasiswaResponse,
	})
}

func convertResponse(b model.Mahasiswa) response.MahasiswaResponse {
	return response.MahasiswaResponse{
		ID:             b.Id,
		NPM:            b.NPM,
		Nama:           b.Nama,
		Kelas:          b.Kelas,
		Jurusan:        b.Jurusan,
		NomorHandphone: b.NomorHandphone,
	}
}
