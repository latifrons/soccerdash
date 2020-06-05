package soccerdash

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net"
	"sync"
	"time"
)

type Reporter struct {
	Id              string
	TargetAddress   string
	BufferSize      int
	Logger          *logrus.Logger
	conn            net.Conn
	inited          bool
	outgoingChannel chan *Message
	quit            chan bool
	lock            sync.Mutex
}

type Message struct {
	Name   string      `json:"name"`
	Global bool        `json:"global"`
	Key    string      `json:"key"`
	Value  interface{} `json:"value"`
}

func (r *Reporter) Name() string {
	return "reporter"
}

func (r *Reporter) ensureConnection() {
	var err error
	for {
		if r.conn != nil {
			break
		}
		if r.Logger != nil {
			r.Logger.WithField("server", r.TargetAddress).Trace("connecting to ws server")
		}

		r.conn, err = net.DialTimeout("tcp", r.TargetAddress, time.Millisecond*5000)
		if err == nil {
			if r.Logger != nil {
				logrus.Trace("connected to ws server")
			}
			break
		}
		time.Sleep(time.Second * 3)
	}
}

func (r *Reporter) Report(key string, value interface{}, global bool) {
	if !r.inited {
		r.lock.Lock()
		if !r.inited {
			r.init()
		}
		r.lock.Unlock()
	}
	msg := &Message{
		Global: global,
		Key:    key,
		Value:  value,
		Name:   r.Id,
	}

	if len(r.outgoingChannel) > r.BufferSize {
		if r.Logger != nil {
			r.Logger.Warn("soccer dash message drop because of too many messages.")
		}

	} else {
		r.outgoingChannel <- msg
	}
}

func (r *Reporter) init() {
	r.inited = true
	if r.BufferSize == 0 {
		r.BufferSize = 20
	}
	r.outgoingChannel = make(chan *Message, r.BufferSize+10)
	r.quit = make(chan bool)

	go r.Start()
}

func (r *Reporter) Start() {
loop:
	for {
		select {
		case <-r.quit:
			break loop
		case msg := <-r.outgoingChannel:
			r.ensureConnection()
			b, err := json.Marshal(msg)
			if err != nil {
				if r.Logger != nil {
					r.Logger.WithError(err).Warn("bad soccerdash format")
				}
			}

			if r.Logger != nil {
				logrus.WithField("content", string(b)).Trace("sending content")
			}

			_ = r.conn.SetWriteDeadline(time.Now().Add(5 * time.Second))

			_, err = r.conn.Write(b)
			if err != nil {
				r.conn = nil
				if r.Logger != nil {
					logrus.WithError(err).Trace("soccerdash server lost")
				}

				break // break select
			}
			_ = r.conn.SetWriteDeadline(time.Now().Add(5 * time.Second))
			_, err = r.conn.Write([]byte("\n"))
			if err != nil {
				r.conn = nil
				if r.Logger != nil {
					logrus.WithError(err).Trace("soccerdash server lost")
				}

				break // break select
			}
		}
	}

}

func (r *Reporter) Stop() {
	r.quit <- true
}
