package util

import (
	"errors"
	"time"

	model "github.com/cloudreve/Cloudreve/v3/models"
)

func GetFileByPath(path []string, user model.User) (*model.File, *model.Folder, uint, error) {
	folder, idx, err := GetFolderByPath(path, user)
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

func ImportFile(path []string, source string, user model.User) error {
	file, folder, idx, err := GetFileByPath(path, user)
	if err != nil {
		if err.Error() == "here is a folder" {
			return err
		}
		if err.Error() == "record not found" {
			if idx < uint(len(path)-1) {
				folder, err = CreateSubFolders(folder, path[idx:len(path)-1], user)
				if err != nil {
					return err
				}
			}
			file = &model.File{
				Name:       path[len(path)-1],
				SourceName: source,
				UserID:     user.ID,
				FolderID:   folder.ID,
				PolicyID:   user.Policy.ID,
			}
			return file.Create()
		}
	}
	return errors.New("file exists")
}

func UpdateFileTime(file *model.File, ctime, mtime time.Time, dtime *time.Time, size uint64) error {
	var updates = map[string]interface{}{}
	if file.CreatedAt != ctime {
		updates["created_at"] = ctime
	}
	if file.UpdatedAt != mtime {
		updates["updated_at"] = mtime
	}
	if file.DeletedAt != dtime {
		updates["deleted_at"] = dtime
	}
	if file.Size != size {
		updates["size"] = size
	}
	return model.DB.Model(&model.File{}).Where("id = ?", file.ID).UpdateColumns(updates).Error
}
