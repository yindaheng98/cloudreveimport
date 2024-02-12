package util

import (
	"testing"

	model "github.com/cloudreve/Cloudreve/v3/models"
	"github.com/cloudreve/Cloudreve/v3/pkg/conf"
)

func Test_GetFileByPath(t *testing.T) {
	conf.Init("D:\\Documents\\MyPrograms\\cloudreveimport\\conf.ini")
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
