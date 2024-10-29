package main

import (
	"fmt"
	"os"

	"github.com/fighterlyt/dtm/client/dtmcli"
	"github.com/fighterlyt/dtm/dtmsvr"
	"github.com/fighterlyt/dtm/dtmsvr/config"
	"github.com/fighterlyt/dtm/dtmsvr/storage/registry"
	"github.com/fighterlyt/dtm/helper/bench/svr"
	"github.com/fighterlyt/dtm/test/busi"
	"github.com/fighterlyt/logger"
)

var usage = `bench is a bench test server for dtmf
usage:
    redis   prepare for redis bench test
    db      prepare for mysql|postgres bench test
		boltdb  prepare for boltdb bench test
`

func hintAndExit() {
	fmt.Print(usage)
	os.Exit(0)
}

var conf = &config.Config

func main() {
	if len(os.Args) <= 1 {
		hintAndExit()
	}
	logger.Infof("starting bench server")
	config.MustLoadConfig("")
	logger.InitLog(conf.LogLevel)
	registry.WaitStoreUp()
	dtmsvr.PopulateDB(false)
	if os.Args[1] == "db" {
		if busi.BusiConf.Driver == "mysql" {
			dtmcli.SetCurrentDBType(busi.BusiConf.Driver)
			svr.PrepareBenchDB()
		}
		busi.PopulateDB(false)
	} else if os.Args[1] == "redis" || os.Args[1] == "boltdb" {

	} else {
		hintAndExit()
	}
	dtmsvr.StartSvr()
	go dtmsvr.CronExpiredTrans(-1)
	svr.StartSvr()
	select {}
}
