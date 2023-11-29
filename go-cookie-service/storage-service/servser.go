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
		maintenance.GetEnv("DB_URI", "mongodb://root:password@localhost:27017"),
		maintenance.GetEnv("DB_NAME", "test"))
	if err != nil {
		maintenance.LogData("Error: " + err.Error())
	}

	storage, err := storage.NewMongoDBFileStorage(db)
	if err != nil {
		maintenance.LogData("Error: " + err.Error())
	}

	mdb = *storage

	r := gin.Default()
	r.GET("/files", GetFiles)
	r.GET("/files/:id", GetFile)
	r.GET("/files/:id/info", GetFileInfo)
	r.POST("/files", UploadFile)
	r.PUT("/files/:id", UpdateFile)
	r.DELETE("/files/:id", DeleteFile)

	if err := r.Run(":" + maintenance.GetEnv("PORT", "8080")); err != nil {
		maintenance.LogData(err.Error())
		panic(err)
	}

}

func GetFiles(c *gin.Context) {
	files, err := mdb.GetFilesNames()
	if err != nil {
		c.JSON(404, "Что-то пошло не так"+err.Error())
		maintenance.LogData("Что-то пошло не так: " + err.Error())
		return
	}
	c.JSON(200, files)
}

func GetFile(c *gin.Context) {
	id := c.Param("id")
	file, err := mdb.GetFile(id)
	if err != nil {
		c.JSON(404, "Что-то пошло не так"+err.Error())
		maintenance.LogData("Что-то пошло не так: " + err.Error())
		return
	}
	content, err := io.ReadAll(file)
	if err != nil {
		c.JSON(404, "Ошибка при чтении файла"+err.Error())
		maintenance.LogData("Что-то пошло не так: " + err.Error())
		return
	}
	c.Data(http.StatusOK, "application/octet-stream", content)
}

func GetFileInfo(c *gin.Context) {
	id := c.Param("id")
	info, err := mdb.GetFileInfo(id)
	if err != nil {
		c.JSON(404, "Что-то пошло не так"+err.Error())
		maintenance.LogData("Что-то пошло не так: " + err.Error())
		return
	}
	c.JSON(http.StatusOK, info)
}

func UploadFile(c *gin.Context) {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		c.JSON(400, "Не удалось получить файл"+err.Error())
		maintenance.LogData("Что-то пошло не так: " + err.Error())
		return
	}
	file, err := fileHeader.Open()
	if err != nil {
		c.JSON(500, "Не удалось открыть файл"+err.Error())
		maintenance.LogData("Что-то пошло не так: " + err.Error())
		return
	}
	id, err := mdb.SaveFile(file, fileHeader)
	if err != nil {
		c.JSON(500, "Не удалось сохранить файл"+err.Error())
		maintenance.LogData("Что-то пошло не так: " + err.Error())
		return
	}
	c.JSON(200, id)
}

func UpdateFile(c *gin.Context) {
	id := c.Param("id")
	fileHeader, err := c.FormFile("file")
	if err != nil {
		c.JSON(400, "Не удалось получить файл"+err.Error())
		maintenance.LogData("Что-то пошло не так: " + err.Error())
		return
	}
	file, err := fileHeader.Open()
	if err != nil {
		c.JSON(500, "Не удалось открыть файл"+err.Error())
		maintenance.LogData("Что-то пошло не так: " + err.Error())
		return
	}
	err = mdb.UpdateFile(id, file, fileHeader)
	if err != nil {
		c.JSON(500, "Ошибка при обновлении файла"+err.Error())
		maintenance.LogData("Что-то пошло не так: " + err.Error())
	}
	c.JSON(200, "")
}

func DeleteFile(c *gin.Context) {
	id := c.Param("id")
	err := mdb.DeleteFile(id)
	if err != nil {
		c.JSON(500, "Ошибка при удалении файла")
		maintenance.LogData("Что-то пошло не так: " + err.Error())
		return
	}
	c.JSON(200, "")
}
