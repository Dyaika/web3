package storage

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"io"
	"mime/multipart"
)

type MongoDBFileStorage struct {
	db *mongo.Database
	fs *gridfs.Bucket
}

type FileInfo struct {
	Name string `bson:"filename"`
	Size int64  `bson:"length"`
}

func NewMongoDBFileStorage(db *mongo.Database) (*MongoDBFileStorage, error) {
	fs, err := gridfs.NewBucket(db)
	if err != nil {
		return nil, err
	}

	return &MongoDBFileStorage{
		db: db,
		fs: fs,
	}, nil
}

func (m *MongoDBFileStorage) SaveFile(file multipart.File, header *multipart.FileHeader) (string, error) {
	fileID, err := m.fs.UploadFromStream(header.Filename, io.Reader(file))
	if err != nil {
		panic(err)
	}
	return fileID.String(), nil
}

func (m *MongoDBFileStorage) GetFile(id string) (io.Reader, error) {
	fileID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	file, err := m.fs.OpenDownloadStream(fileID)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (m *MongoDBFileStorage) GetFileInfo(id string) (FileInfo, error) {
	fileID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return FileInfo{}, err
	}

	cursor, err := m.fs.Find(bson.M{"_id": fileID})
	if err != nil {
		return FileInfo{}, err
	}

	var found []FileInfo
	if err = cursor.All(context.TODO(), &found); err != nil {
		panic(err)
	}
	return found[0], nil
}

func (m *MongoDBFileStorage) UpdateFile(id string, file multipart.File, header *multipart.FileHeader) error {
	err := m.DeleteFile(id)
	if err != nil {
		return err
	}
	fileID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	err = m.fs.UploadFromStreamWithID(fileID, header.Filename, io.Reader(file))
	if err != nil {
		panic(err)
	}

	return nil
}

func (m *MongoDBFileStorage) DeleteFile(id string) error {
	fileID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	return m.fs.Delete(fileID)
}

func (m *MongoDBFileStorage) GetFilesNames() ([]string, error) {
	cursor, err := m.fs.Find(bson.M{})
	if err != nil {
		return []string{}, err
	}
	var fileNames []string

	// Итерируем по результатам запроса
	for cursor.Next(context.Background()) {
		var file FileInfo
		err := cursor.Decode(&file)
		if err != nil {
			return nil, fmt.Errorf("failed to decode file info: %v", err)
		}
		fileNames = append(fileNames, file.Name)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %v", err)
	}

	return fileNames, nil
}
