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
	folder, idx, err := GetFolderIDByPath([]string{"www", "qqq"}, user)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(idx)
	t.Log(folder)

	folder, idx, err = GetFolderIDByPath([]string{"qqq", "www"}, user)
	if err != nil {
		t.Log(folder)
		t.Log(err)
	}
	t.Log(idx)
	t.Log(folder)

	folder, idx, err = GetFolderIDByPath([]string{"www", "www"}, user)
	if err != nil {
		t.Log(folder)
		t.Log(err)
	}
	t.Log(idx)
	t.Log(folder)
}

func Test_CreateFolderByPath(t *testing.T) {
	conf.Init("D:\\Documents\\MyPrograms\\cloudreveimport\\conf.ini")
	model.Init()
	user, err := model.GetUserByEmail("yindaheng98@gmail.com")
	if err != nil {
		t.Error(err)
		return
	}
	folder, err := CreateFolderByPath([]string{"www", "qqq"}, user)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(folder)

	folder, err = CreateFolderByPath([]string{"qqq", "www"}, user)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(folder)
}
