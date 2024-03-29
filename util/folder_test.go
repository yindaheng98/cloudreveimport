package util

import (
	"testing"
	"time"

	model "github.com/cloudreve/Cloudreve/v3/models"
	"github.com/cloudreve/Cloudreve/v3/pkg/conf"
)

func Test_GetFolderByPath(t *testing.T) {
	conf.Init("D:\\Documents\\MyPrograms\\cloudreveimport\\test\\conf.ini")
	model.Init()
	user, err := model.GetUserByEmail("yindaheng98@gmail.com")
	if err != nil {
		t.Error(err)
		return
	}
	folder, idx, err := GetFolderByPath([]string{"www", "qqq"}, user)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(idx)
	t.Log(folder)

	folder, idx, err = GetFolderByPath([]string{"qqq", "www"}, user)
	if err != nil {
		t.Log(folder)
		t.Log(err)
	}
	t.Log(idx)
	t.Log(folder)

	folder, idx, err = GetFolderByPath([]string{"www", "www"}, user)
	if err != nil {
		t.Log(folder)
		t.Log(err)
	}
	t.Log(idx)
	t.Log(folder)
}

func Test_CreateFolderByPath(t *testing.T) {
	conf.Init("D:\\Documents\\MyPrograms\\cloudreveimport\\test\\conf.ini")
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

	folder, err = CreateFolderByPath([]string{"www", "www"}, user)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(folder)

	folder, err = CreateFolderByPath([]string{"www", "www", "qqq", "qqq", "qqq"}, user)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(folder)
}

func Test_UpdateFolderTime(t *testing.T) {
	conf.Init("D:\\Documents\\MyPrograms\\cloudreveimport\\test\\conf.ini")
	model.Init()
	user, err := model.GetUserByEmail("yindaheng98@gmail.com")
	if err != nil {
		t.Error(err)
		return
	}
	folder, _, err := GetFolderByPath([]string{"qqq", "www"}, user)
	if err != nil {
		t.Log(folder)
		t.Log(err)
	}
	err = UpdateFolderTime(folder, folder.CreatedAt, time.Now(), nil)
	t.Log(err)
}
