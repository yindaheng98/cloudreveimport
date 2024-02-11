package util

import (
	model "github.com/cloudreve/Cloudreve/v3/models"
)

func GetFolderIDByPath(path []string, user model.User) (*model.Folder, uint, error) {
	root, err := user.Root()
	if err != nil {
		return nil, 0, err
	}
	parent := root
	for i, dirname := range path {
		folder, err := parent.GetChild(dirname)
		if err != nil {
			folder.Name = dirname
			folder.ParentID = &parent.ID
			return folder, uint(i), err
		}
		parent = folder
	}
	return parent, uint(len(path)), nil
}

func CreateFolderByPath(path []string, user model.User) {
}
