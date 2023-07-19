package main

import (
	"flag"
	"github.com/sirupsen/logrus"
	"minimalism/common/config"
	"minimalism/common/db"
	"minimalism/common/log"
)

var (
	cfg = flag.String("c", "./config/conf.yaml", "config file.")
)

func main() {
	flag.Parse()

	err := config.Parse(*cfg)
	if err != nil {
		panic(err)
	}

	log.InitLog()

	err = db.InitDB()
	if err != nil {
		logrus.Panic(err)
	}

}
