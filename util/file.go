package util

import (
	"errors"

	model "github.com/cloudreve/Cloudreve/v3/models"
)

func GetFileByPath(path []string, user model.User) (*model.File, *model.Folder, uint, error) {
	folder, idx, err := GetFolderIDByPath(path, user)
	if err == nil {
		return nil, folder, idx, errors.New("here is a folder")
	}
	if err.Error() != "record not found" {
		return nil, nil, idx, err
	}
	if idx == uint(len(path)-1) {
		file, err := folder.GetChildFile(path[len(path)-1])
		return file, folder, idx, err
	}
	return nil, folder, idx, err
}
