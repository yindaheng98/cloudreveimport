package util

import (
	model "github.com/cloudreve/Cloudreve/v3/models"
)

func GetFolderIDByPath(path []string, user model.User) (*model.Folder, error) {
	root, err := user.Root()
	if err != nil {
		return nil, err
	}
	folder := root
	for _, dirname := range path {
		folder, err = folder.GetChild(dirname)
		if err != nil {
			return nil, err
		}
	}
	return folder, nil
}
