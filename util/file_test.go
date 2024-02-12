package util

import (
	"testing"

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
	err = ImportFile([]string{"www", "qqq"}, "/gallery/data/twitter/eumi_114/1426380913575743488_1.jpg", user)
	t.Log(err)

	err = ImportFile([]string{"www", "qqq", "fff.png"}, "/gallery/data/twitter/eumi_114/1426380913575743488_1.jpg", user)
	t.Log(err)

	err = ImportFile([]string{"www", "qqq", "iii.png"}, "/gallery/data/twitter/eumi_114/1426380913575743488_1.jpg", user)
	t.Log(err)

	err = ImportFile([]string{"www", "ggg", "ggg", "iii.png"}, "/gallery/data/twitter/eumi_114/1426380913575743488_1.jpg", user)
	t.Log(err)
}
