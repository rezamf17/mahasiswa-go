package handler

import (
	"fmt"
	"mahasiswa/model"
	"mahasiswa/request"
	"mahasiswa/response"
	"mahasiswa/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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

func (h *mahasiswaHandler) CreateMahasiswa(ctx *gin.Context) {
	var mahasiswaRequest request.MahasiswaRequest

	err := ctx.ShouldBindJSON(&mahasiswaRequest)
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, Condition %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
			// ctx.JSON(http.StatusBadRequest, errorMessage)
			// return
		}
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}

	mahasiswa, err := h.mahasiswaService.Create(mahasiswaRequest)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": convertResponse(mahasiswa),
	})
}

func (h *mahasiswaHandler) GetMahasiswaId(ctx *gin.Context) {
	idString := ctx.Param("id")
	id, _ := strconv.Atoi(idString)
	// fmt.Println(id)
	b, err := h.mahasiswaService.FindByID(int(id))

	mahasiswaResponse := convertResponse(b)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
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
