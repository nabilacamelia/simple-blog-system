package setup

import (
	"log"
	"simple-blog-system/config"
	"simple-blog-system/config/db"

	"github.com/gin-contrib/cors" // Tambahkan ini (setelah go get)
	"github.com/gin-gonic/gin"
)

// BaseURL base url of api
const BaseURL = "/v1/api"

// CloseDB close connection to db
var CloseDB func() error

type SetupData struct {
	ConfigData  config.Config
	InternalApp InternalAppStruct
	Router      *gin.Engine // Tambahkan ini agar main.go bisa akses router
}

// Fungsi Init Utama
func Init() SetupData {
	configData := config.GetConfig()

	// 1. DB INIT
	dbConn, err := db.Init(configData.DB.DSN)
	if err != nil {
		log.Println("database error")
	}

	CloseDB = func() error {
		if err := dbConn.CloseConnection(); err != nil {
			return err
		}
		return nil
	}

	// 2. GIN INIT + CORS (IZIN UNTUK REACT)
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"POST", "GET", "OPTIONS", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	// 3. INTERNAL APP INIT
	internalAppVar := initInternalApp(dbConn.GormDB)

	return SetupData{
		ConfigData:  configData,
		InternalApp: internalAppVar,
		Router:      r, // Kembalikan router yang sudah dipasang CORS
	}
}

func initInternalApp(gormDB *db.GormDB) InternalAppStruct {
	var internalAppVar InternalAppStruct

	initAppRepo(gormDB, &internalAppVar)
	initAppService(&internalAppVar)
	initAppHandler(&internalAppVar)

	return internalAppVar
}