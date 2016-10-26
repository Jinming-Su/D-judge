package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	log "github.com/Sirupsen/logrus"
	"github.com/VOID001/D-judge/config"
	"github.com/VOID001/D-judge/judge-controller"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

const (
	INFO    = 0
	WARN    = 1
	DEBUG   = 2
	DirPerm = 0744
)

const (
	ErrNoDir  = "no such file or directory"
	ErrNoFile = "no such file or directory"
)

var path string
var debuglv int64
var GlobalConfig config.SystemConfig

func init() {
	flag.StringVar(&path, "c", "config.toml", "select configuration file")
	flag.Int64Var(&debuglv, "d", 0, "debug mode enabled")
	flag.Parse()
	_, err := toml.DecodeFile(path, &config.GlobalConfig)
	GlobalConfig = config.GlobalConfig
	if err != nil {
		err = errors.Wrap(err, "Processing config file error")
		log.Fatal(err)
	}
	cwd, err := os.Getwd()
	if err != nil {
		err = errors.Wrap(err, "Get current directory error")
	}
	if GlobalConfig.JudgeRoot[0] != '/' {
		GlobalConfig.JudgeRoot = filepath.Join(cwd, GlobalConfig.JudgeRoot)
	}
	if GlobalConfig.CacheRoot[0] != '/' {
		GlobalConfig.CacheRoot = filepath.Join(cwd, GlobalConfig.CacheRoot)
	}
	if debuglv == INFO {
		log.SetLevel(log.InfoLevel)
	}
	if debuglv == WARN {
		log.SetLevel(log.WarnLevel)
	}
	if debuglv == DEBUG {
		log.SetLevel(log.DebugLevel)
	}
}

func main() {
	log.Infof("Settings %+v", GlobalConfig)
	// Perform Sanity Check
	log.Infof("sanity check start")
	err := sanity_check_dir(GlobalConfig.JudgeRoot)
	if err != nil {
		err = errors.Wrap(err, "sanity check dir judgeroot error")
		log.Fatal(err)
	}
	err = sanity_check_dir(GlobalConfig.CacheRoot)
	if err != nil {
		err = errors.Wrap(err, "sanity check dir cacheroot error")
		log.Fatal(err)
	}
	err = sanity_check_connection(GlobalConfig.EndpointURL)
	if err != nil {
		err = errors.Wrap(err, "sanity check connection error")
		log.Fatal(err)
	}
	err = sanity_check_docker()
	if err != nil {
		err = errors.Wrap(err, "sanity check docker error")
		log.Fatal(err)
	}
	log.Infof("sanity check success")

	// PerformRequest Lifcycle

	// Request For Judge

	// Start Compile

	// Start Run

	// Start Compare

	// NextTestcase
}

func sanity_check_dir(dir string) (err error) {
	_, err = ioutil.ReadDir(dir)
	if err != nil && os.IsNotExist(err) {
		err = os.Mkdir(dir, DirPerm)
		if err != nil {
			err = errors.Wrap(err, fmt.Sprintf("cannot make %s", dir))
			return
		}
		log.Infof("created dir %s with mode %04o", dir, DirPerm)
	}
	info, err := os.Stat(dir)
	log.Infof("dir %s mode bits %s", dir, info.Mode())
	return
}

func sanity_check_connection(endpoint string) (err error) {
	req, err := http.NewRequest(http.MethodPost, endpoint, nil)
	if err != nil {
		err = errors.Wrap(err, fmt.Sprintf("cannot create request", endpoint))
		return
	}
	cli := http.Client{}
	resp, err := cli.Do(req)
	if err != nil {
		err = errors.Wrap(err, fmt.Sprintf("cannot connect to %s", endpoint))
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 300 || resp.StatusCode < 200 {
		err = errors.New(fmt.Sprintf("server response error %s", resp.Status))
		return
	}
	return
}

func sanity_check_docker() (err error) {
	err = controller.Ping(context.Background())
	if err != nil {
		err = errors.Wrap(err, "docker Ping error")
	}
	return
}
