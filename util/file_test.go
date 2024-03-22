package util

import (
	"testing"
	"time"

	model "github.com/cloudreve/Cloudreve/v3/models"
	"github.com/cloudreve/Cloudreve/v3/pkg/conf"
)

func Test_GetFileByPath(t *testing.T) {
	conf.Init("D:\\Documents\\MyPrograms\\cloudreveimport\\test\\conf.ini")
	model.Init()
	user, err := model.GetUserByEmail("yindaheng98@gmail.com")
	if err != nil {
		t.Error(err)
		return
	}
	file, folder, idx, err := GetFileByPath([]string{"www", "qqq"}, user)
	t.Log(file)
	t.Log(folder)
	t.Log(idx)
	t.Log(err)

	file, folder, idx, err = GetFileByPath([]string{"www", "qqq", "fff.png"}, user)
	t.Log(file)
	t.Log(folder)
	t.Log(idx)
	t.Log(err)
}

func Test_ImportFile(t *testing.T) {
	conf.Init("D:\\Documents\\MyPrograms\\cloudreveimport\\test\\conf.ini")
	model.Init()
	user, err := model.GetUserByEmail("yindaheng98@gmail.com")
	if err != nil {
		t.Error(err)
		return
	}
	err = ImportFile([]string{"www", "qqq"}, "/gallery/data/twitter/eumi_114/1426380913575743488_1.jpg", 222, user)
	t.Log(err)

	err = ImportFile([]string{"www", "qqq", "fff.png"}, "/gallery/data/twitter/eumi_114/1426380913575743488_1.jpg", 222, user)
	t.Log(err)

	err = ImportFile([]string{"www", "qqq", "iii.png"}, "/gallery/data/twitter/eumi_114/1426380913575743488_1.jpg", 333, user)
	t.Log(err)

	err = ImportFile([]string{"www", "ggg", "ggg", "iii.png"}, "/gallery/data/twitter/eumi_114/1426380913575743488_1.jpg", 444, user)
	t.Log(err)
}

func Test_UpdateFileStat(t *testing.T) {
	conf.Init("D:\\Documents\\MyPrograms\\cloudreveimport\\test\\conf.ini")
	model.Init()
	user, err := model.GetUserByEmail("yindaheng98@gmail.com")
	if err != nil {
		t.Error(err)
		return
	}
	file, _, _, err := GetFileByPath([]string{"www", "qqq", "fff.png"}, user)
	if err != nil {
		t.Log(file)
		t.Log(err)
	}
	err = UpdateFileStat(file, file.CreatedAt, time.Now(), nil, 12)
	t.Log(err)
}

func Test_UpdateFileMeta(t *testing.T) {
	conf.Init("D:\\Documents\\MyPrograms\\cloudreveimport\\test\\conf.ini")
	model.Init()
	user, err := model.GetUserByEmail("yindaheng98@gmail.com")
	if err != nil {
		t.Error(err)
		return
	}
	file, _, _, err := GetFileByPath([]string{"www", "qqq", "fff.png"}, user)
	if err != nil {
		t.Log(file)
		t.Log(err)
	}
	err = UpdateFileMeta(file, map[string]string{"thumb_sidecar": "true", "thumb_status": "exist"})
	t.Log(err)
}

func Test_DeleteFile(t *testing.T) {
	conf.Init("D:\\Documents\\MyPrograms\\cloudreveimport\\test\\conf.ini")
	model.Init()
	user, err := model.GetUserByEmail("yindaheng98@gmail.com")
	if err != nil {
		t.Error(err)
		return
	}
	err = ImportFile([]string{"www", "qqq", "fff.png"}, "/gallery/data/twitter/eumi_114/1426380913575743488_1.jpg", 111, user)
	t.Log(err)
	file, _, _, err := GetFileByPath([]string{"www", "qqq", "fff.png"}, user)
	if err != nil {
		t.Log(file)
		t.Log(err)
	}
	err = DeleteFile(file)
	t.Log(err)
	err = DeleteFile(file)
	t.Log(err)
}
