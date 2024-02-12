package main

import (
	"flag"
	"fmt"

	model "github.com/cloudreve/Cloudreve/v3/models"
	"github.com/cloudreve/Cloudreve/v3/pkg/conf"
	"github.com/cloudreve/Cloudreve/v3/pkg/util"
)

var (
	userName string
	confPath string
)

func init() {
	flag.StringVar(&confPath, "c", util.RelativePath("conf.ini"), "Path to the config file.")
	flag.StringVar(&userName, "u", "admin@cloudreve.org", "Email of the target user.")
	flag.Parse()
	conf.Init(confPath)
	model.Init()
}

func main() {
	user, err := model.GetUserByEmail(userName)
	if err != nil {
		return
	}
	util.Log().Info(fmt.Sprintf("%v\n", user))
}
