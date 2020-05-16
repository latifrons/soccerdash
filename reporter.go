package soccerdash

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net"
	"time"
)

type Reporter struct {
	Name          string
	TargetAddress string
	conn          net.Conn
}

type Message struct {
	Name   string      `json:"name"`
	Global bool        `json:"global"`
	Key    string      `json:"key"`
	Value  interface{} `json:"value"`
}

func (r *Reporter) ensureConnection() {
	var err error
	for {
		if r.conn != nil {
			break
		}
		logrus.WithField("server", r.TargetAddress).Debug("connecting to ws server")
		r.conn, err = net.DialTimeout("tcp", r.TargetAddress, time.Millisecond*5000)
		if err == nil {
			logrus.Debug("connected to ws server")
			break
		}
		time.Sleep(time.Second * 3)
	}
}

func (r *Reporter) Report(key string, value interface{}, global bool) {
	r.ensureConnection()
	b, err := json.Marshal(Message{
		Global: global,
		Key:    key,
		Value:  value,
		Name:   r.Name,
	})
	if err != nil {
		panic(err)
	}
	logrus.WithField("content", string(b)).Debug("sending content")
	r.conn.Write(b)
	r.conn.Write([]byte("\n"))
}
