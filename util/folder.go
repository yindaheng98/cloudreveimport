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
			return parent, uint(i), err
		}
		parent = folder
	}
	return parent, uint(len(path)), nil
}

func CreateFolderByPath(path []string, user model.User) (*model.Folder, error) {
	parent, idx, err := GetFolderIDByPath(path, user)
	if err != nil {
		if err.Error() != "record not found" {
			return nil, err
		}
		parent, err = CreateSubFolders(parent, path[idx:], user)
		if err != nil {
			return nil, err
		}
	}
	return parent, nil
}

func CreateSubFolders(parent *model.Folder, path []string, user model.User) (*model.Folder, error) {
	var err error
	for i := 0; i < len(path); i++ {
		folder := &model.Folder{
			Name:     path[i],
			OwnerID:  user.ID,
			ParentID: &parent.ID,
		}
		folder.Name = path[i]
		folder.ID, err = folder.Create()
		if err != nil {
			return nil, err
		}
		parent = folder
	}
	return parent, nil
}
