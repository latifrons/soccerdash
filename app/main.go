package main

import (
	"github.com/latifrons/soccerdash"
	"github.com/sirupsen/logrus"
	"time"
)

func main() {
	logrus.StandardLogger().Level = logrus.DebugLevel

	r := soccerdash.Reporter{
		Name:          "C1",
		TargetAddress: "127.0.0.1:1088",
	}

	for {
		r.Report("time", time.Now().Second(), false)
		time.Sleep(time.Second)
	}
}
