package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	model "github.com/cloudreve/Cloudreve/v3/models"
	"github.com/cloudreve/Cloudreve/v3/pkg/conf"
	"github.com/cloudreve/Cloudreve/v3/pkg/util"
	ci "github.com/yindaheng98/cloudreveimport/util"
)

var (
	userName string
	confPath string
	dataPath string
)

func init() {
	flag.StringVar(&confPath, "c", util.RelativePath("conf.ini"), "Path to the config file.")
	flag.StringVar(&userName, "u", "admin@cloudreve.org", "Email of the target user.")
	flag.StringVar(&dataPath, "m", "", "Map of files to be imported.")
	flag.Parse()
	conf.Init(confPath)
	model.Init()
}

func main() {
	user, err := model.GetUserByEmail(userName)
	if err != nil {
		return
	}
	util.Log().Info(fmt.Sprintf("User: %+v\n", user))
	file, err := os.Open(dataPath)
	if err != nil {
		util.Log().Error("%+v", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data := scanner.Bytes()
		v, err := ci.Unmarshal(data)
		if err != nil {
			util.Log().Error("%s %+v", string(data), err)
			continue
		}
		util.Log().Debug("%+v", v)
		err = ci.ImportFile(v.DstPath, v.SourceName, user)
		if err != nil {
			if err.Error() == "file exists" {
				util.Log().Debug("exists %+v", v)
			} else {
				util.Log().Info("error  %+v %+v", v, err)
			}
		} else {
			util.Log().Info("done   %+v", v)
		}
	}
}
