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
	case UpdateFolderTimeCommand:
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
