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
	case DeleteFileCommand:
		i.invokeDeleteFile(v)
	case UpdateFileStatCommand:
		i.invokeUpdateFileStat(v)
	case UpdateFileMetaCommand:
		i.invokeUpdateFileMeta(v)
	case UpdateFolderTimeCommand:
		i.invokeUpdateFolderTime(v)
	default:
		util.Log().Error("Unrecogenized command: %+v", v)
	}
}

func (i *Invoker) invokeImportFile(v Command) {
	err := ImportFile(v.DstPath, v.SourceName, v.Size, i.User)
	if err != nil {
		if err.Error() == "file exists" {
			util.Log().Debug("file exists %+v", v)
		} else {
			util.Log().Error("error  %+v %+v", v, err)
		}
	} else {
		util.Log().Info("new file %+v %+v %+v", v.DstPath, v.SourceName, v.Size)
	}
}

func (i *Invoker) invokeDeleteFile(v Command) {
	file, _, _, err := GetFileByPath(v.DstPath, i.User)
	if err != nil {
		if err.Error() == "record not found" {
			util.Log().Error("file not exists %+v", v)
		} else {
			util.Log().Error("error  %+v %+v", v, err)
		}
		return
	}
	err = DeleteFile(file)
	if err != nil {
		util.Log().Error("error  %+v %+v", v, err)
	} else {
		util.Log().Info("deleted file %+v", v.DstPath)
	}
}

func (i *Invoker) invokeUpdateFolderTime(v Command) {
	folder, _, err := GetFolderByPath(v.DstPath, i.User)
	if err != nil {
		if err.Error() == "record not found" {
			util.Log().Error("folder not exists %+v", v)
		} else {
			util.Log().Error("error  %+v %+v", v, err)
		}
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
	} else {
		util.Log().Info("folder time updated %+v %+v %+v", v.DstPath, ctime, mtime)
	}
}

func (i *Invoker) invokeUpdateFileStat(v Command) {
	file, _, _, err := GetFileByPath(v.DstPath, i.User)
	if err != nil {
		if err.Error() == "record not found" {
			util.Log().Error("file not exists %+v", v)
		} else {
			util.Log().Error("error  %+v %+v", v, err)
		}
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
	size := file.Size
	if v.Size > 0 {
		size = v.Size
	}
	err = UpdateFileStat(file, ctime, mtime, nil, size)
	if err != nil {
		util.Log().Error("%+v %+v", err, v)
	} else {
		util.Log().Info("file stat updated %+v %+v %+v %+v", v.DstPath, ctime, mtime, size)
	}
}

func (i *Invoker) invokeUpdateFileMeta(v Command) {
	file, _, _, err := GetFileByPath(v.DstPath, i.User)
	if err != nil {
		if err.Error() == "record not found" {
			util.Log().Error("file not exists %+v", v)
		} else {
			util.Log().Error("error  %+v %+v", v, err)
		}
		return
	}
	err = UpdateFileMeta(file, v.Metadata)
	if err != nil {
		util.Log().Error("%+v %+v", err, v)
	} else {
		util.Log().Info("file meta updated %+v %+v", v.DstPath, v.Metadata)
	}
}
