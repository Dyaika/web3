package storage_service

import (
	"github.com/gin-gonic/gin"
	"go-cookie-service/maintenance"
	"go-cookie-service/storage"
	"io"
	"net/http"
)

var mdb storage.MongoDBFileStorage

func Start() {
	db, err := storage.InitConnection(
		maintenance.GetEnv("CONNECTION_STRING", "mongodb://root:password@localhost:27017"),
		maintenance.GetEnv("DATABASE", "test"))
	if err != nil {
		maintenance.LogData("Error: " + err.Error())
	}

	fileStorage, err := storage.NewMongoDBFileStorage(db)
	if err != nil {
		maintenance.LogData("Error: " + err.Error())
	}

	mdb = *fileStorage

	r := gin.Default()
	r.GET("/files", GetFiles)
	r.GET("/files/:id", GetFile)
	r.GET("/files/:id/info", GetFileInfo)
	r.POST("/files", AddFile)
	r.PUT("/files/:id", UpdateFile)
	r.DELETE("/files/:id", RemoveFile)

	if err := r.Run(":" + maintenance.GetEnv("PORT", "8080")); err != nil {
		maintenance.LogData(err.Error())
		panic(err)
	}

}

func GetFiles(c *gin.Context) {
	files, err := mdb.GetFilesNames()
	if err != nil {
		c.JSON(500, err.Error())
		maintenance.LogData("Error: " + err.Error())
		return
	}
	c.JSON(200, files)
}

func GetFile(c *gin.Context) {
	id := c.Param("id")
	file, err := mdb.GetFile(id)
	if err != nil {
		c.JSON(500, err.Error())
		maintenance.LogData("Error: " + err.Error())
		return
	}
	content, err := io.ReadAll(file)
	if err != nil {
		c.JSON(500, err.Error())
		maintenance.LogData("Error while reading: " + err.Error())
		return
	}
	c.Data(http.StatusOK, "application/octet-stream", content)
}

func GetFileInfo(c *gin.Context) {
	id := c.Param("id")
	info, err := mdb.GetFileInfo(id)
	if err != nil {
		c.JSON(500, err.Error())
		maintenance.LogData("Error: " + err.Error())
		return
	}
	c.JSON(http.StatusOK, info)
}

func AddFile(c *gin.Context) {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		c.JSON(400, err.Error())
		maintenance.LogData("File error: " + err.Error())
		return
	}
	file, err := fileHeader.Open()
	if err != nil {
		c.JSON(500, err.Error())
		maintenance.LogData("Can't open file: " + err.Error())
		return
	}
	id, err := mdb.SaveFile(file, fileHeader)
	if err != nil {
		c.JSON(500, err.Error())
		maintenance.LogData("Can't save file: " + err.Error())
		return
	}
	c.JSON(200, id)
}

func UpdateFile(c *gin.Context) {
	id := c.Param("id")
	fileHeader, err := c.FormFile("file")
	if err != nil {
		c.JSON(400, err.Error())
		maintenance.LogData("File error: " + err.Error())
		return
	}
	file, err := fileHeader.Open()
	if err != nil {
		c.JSON(500, err.Error())
		maintenance.LogData("Can't open file: " + err.Error())
		return
	}
	err = mdb.UpdateFile(id, file, fileHeader)
	if err != nil {
		c.JSON(500, err.Error())
		maintenance.LogData("Can't update file: " + err.Error())
	}
	c.JSON(200, "")
}

func RemoveFile(c *gin.Context) {
	id := c.Param("id")
	err := mdb.DeleteFile(id)
	if err != nil {
		c.JSON(400, err.Error())
		maintenance.LogData("Error while deleting file: " + err.Error())
		return
	}
	c.JSON(200, "")
}
