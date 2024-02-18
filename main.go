package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

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
	flag.StringVar(&dataPath, "m", "", "A folder to be imported.")
	flag.StringVar(&dataPath, "t", "", "Import the folder to which folder in cloudreve.")
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
	invoker := ci.Invoker{User: user}
	if dataPath == "-" {
		import_from_stdin(invoker)
	} else if info, err := os.Stat(dataPath); err == nil {
		if info.IsDir() {
			import_from_folder(invoker)
		} else {
			util.Log().Error("%s is not a folder", dataPath)
		}
	} else {
		util.Log().Error("%s %+v", dataPath, err)
	}
}

func import_from_stdin(invoker ci.Invoker) {
	var reader io.Reader = bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		data := scanner.Bytes()
		v, err := ci.Unmarshal(data)
		if err != nil {
			util.Log().Error("%s %+v", string(data), err)
			continue
		}
		util.Log().Debug("%+v", v)
		invoker.Invoke(*v)
	}
}

func import_from_folder(invoker ci.Invoker) {
	err := filepath.Walk(dataPath, func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		rel, err := filepath.Rel(dataPath, path)
		if err != nil {
			util.Log().Error("%s %+v", path, err)
		}
		v := ci.Command{
			SourceName:         path,
			DstPath:            strings.Split(rel, string(os.PathSeparator)),
			Size:               uint64(info.Size()),
			UpdatedAt:          info.ModTime(),
			UpdatedAtTimestamp: info.ModTime().Unix(),
		}
		v.Command = ci.ImportFileCommand
		invoker.Invoke(v)
		v.Command = ci.UpdateFileStatCommand
		invoker.Invoke(v)
		return err
	})
	if err != nil {
		util.Log().Error("%s %+v", dataPath, err)
	}
}
