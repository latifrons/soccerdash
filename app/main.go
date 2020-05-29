package main

import (
	"fmt"
	"github.com/latifrons/soccerdash"
	"github.com/sirupsen/logrus"
	"time"
)

func main() {
	logrus.StandardLogger().Level = logrus.DebugLevel

	r := soccerdash.Reporter{
		Id:            "C1",
		TargetAddress: "172.28.152.101:32010",
	}

	for {
		fmt.Println("xxx")
		for i := 0; i < 10; i++ {
			go r.Report("time", time.Now().Second(), false)
		}

		time.Sleep(time.Second)
	}
}
