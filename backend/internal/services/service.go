package services

import (
	"errors"
	"media_batch_converter/backend/internal/models"
)

var mockFolders = []models.Folder{
	{ID: "1", Name: "Documents", Path: "/folders/documents"},
	{ID: "2", Name: "Images", Path: "/folders/images"},
}

func GetFolderById(id string) (*models.Folder, error) {
	for _, folder := range mockFolders {
		if folder.ID == id {
			return &folder, nil
		}
	}
	return nil, errors.New("folder not found")
}

func GetAllFolders() ([]models.Folder, error) {
	return mockFolders, nil
}
