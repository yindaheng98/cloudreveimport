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

func CreateFolderByPath(path []string, user model.User) (*model.Folder, error) {
	parent, idx, err := GetFolderIDByPath(path, user)
	if err != nil {
		if err.Error() != "record not found" {
			return nil, err
		}
		for i := int(idx); i < len(path); i++ {
			folder := &model.Folder{
				Name:     path[i],
				ParentID: &parent.ID,
				OwnerID:  user.ID,
			}
			folder.ID, err = folder.Create()
			if err != nil {
				return nil, err
			}
			parent = folder
		}
	}
	return parent, nil
}
