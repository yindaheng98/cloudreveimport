package util

import (
	"testing"

	model "github.com/cloudreve/Cloudreve/v3/models"
	"github.com/cloudreve/Cloudreve/v3/pkg/conf"
)

func Test_GetFolderIDByPath(t *testing.T) {
	conf.Init("D:\\Documents\\MyPrograms\\cloudreveimport\\conf.ini")
	model.Init()
	user, err := model.GetUserByEmail("yindaheng98@gmail.com")
	if err != nil {
		t.Error(err)
		return
	}
	folder, err := GetFolderIDByPath([]string{"www", "qqq"}, user)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(folder)
	folder, err = GetFolderIDByPath([]string{"www", "www"}, user)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(folder)
}
