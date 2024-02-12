package util

import (
	model "github.com/cloudreve/Cloudreve/v3/models"
	"github.com/cloudreve/Cloudreve/v3/pkg/util"
)

type Invoker struct {
	User model.User
}

func (i *Invoker) Invoke(v Command) {
	switch v.Command {
	case ImportFileCommand:
		i.invokeImportFile(v)
	case UpdateFileTimeCommand:
		i.invokeUpdateFolderTime(v)
	case UpdateFolderTimeCommand:
		i.invokeUpdateFileTime(v)
	default:
		util.Log().Error("Unrecogenized command: %s", v.Command)
	}
}

func (i *Invoker) invokeImportFile(v Command) {
	err := ImportFile(v.DstPath, v.SourceName, i.User)
	if err != nil {
		if err.Error() == "file exists" {
			util.Log().Debug("exists %+v", v)
		} else {
			util.Log().Info("error  %+v %+v", v, err)
		}
	} else {
		util.Log().Info("new   %+v", v)
	}
}

func (i *Invoker) invokeUpdateFolderTime(v Command) {
	folder, _, err := GetFolderByPath(v.DstPath, i.User)
	if err != nil {
		util.Log().Error("not exists %s", v.DstPath)
		return
	}
	ctime := folder.CreatedAt
	mtime := folder.UpdatedAt
	if v.CreatedAt.Unix() > 0 {
		ctime = v.CreatedAt
	}
	if v.UpdatedAt.Unix() > 0 {
		mtime = v.UpdatedAt
	}
	err = UpdateFolderTime(folder, ctime, mtime, nil)
	if err != nil {
		util.Log().Error("%+v %+v", err, v)
		return
	}
}

func (i *Invoker) invokeUpdateFileTime(v Command) {
	file, _, _, err := GetFileByPath(v.DstPath, i.User)
	if err != nil {
		util.Log().Error("not exists %s", v.DstPath)
		return
	}
	ctime := file.CreatedAt
	mtime := file.UpdatedAt
	if v.CreatedAt.Unix() > 0 {
		ctime = v.CreatedAt
	}
	if v.UpdatedAt.Unix() > 0 {
		mtime = v.UpdatedAt
	}
	err = UpdateFileTime(file, ctime, mtime, nil)
	if err != nil {
		util.Log().Error("%+v %+v", err, v)
		return
	}
}
