package util

import (
	"time"

	model "github.com/cloudreve/Cloudreve/v3/models"
)

func GetFolderByPath(path []string, user model.User) (*model.Folder, uint, error) {
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
	parent, idx, err := GetFolderByPath(path, user)
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

func UpdateFolderTime(folder *model.Folder, ctime, mtime time.Time, dtime *time.Time) error {
	var updates = map[string]interface{}{}
	if folder.CreatedAt != ctime {
		updates["created_at"] = ctime
	}
	if folder.UpdatedAt != mtime {
		updates["updated_at"] = mtime
	}
	if folder.DeletedAt != dtime {
		updates["deleted_at"] = dtime
	}
	return model.DB.Model(&model.Folder{}).Where("id = ?", folder.ID).Update(updates).Error
}
