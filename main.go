package main

import (
	"log"
	"mahasiswa/handler"
	"mahasiswa/model"
	"mahasiswa/repository"
	"mahasiswa/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=postgres dbname=mahasiswa port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("db connection error")
	}
	db.AutoMigrate(&model.Mahasiswa{})
	mahasiswaRepository := repository.NewRepository(db)
	mahasiswaServices := services.NewService(mahasiswaRepository)
	mahasiswaHandler := handler.NewMahasiswaHandler(mahasiswaServices)

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders: []string{"Content-Type,access-control-allow-origin, access-control-allow-headers"},
	}))

	router.GET("/mahasiswa", mahasiswaHandler.GetAllMahasiswa)
	router.POST("/mahasiswa", mahasiswaHandler.CreateMahasiswa)
	router.GET("/mahasiswa/:id", mahasiswaHandler.GetMahasiswaId)
	router.PUT("/mahasiswa/:id", mahasiswaHandler.UpdateMahasiswa)
	router.DELETE("/mahasiswa/:id", mahasiswaHandler.DeleteMahasiswa)
	log.Println("Starting app at :8888")
	router.Run(":8888")

}
