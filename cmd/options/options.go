package options

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/emicklei/go-restful/v3"

	"gagent/pkg/utils/log"
)

const (
	Listen = 80
	LogDir = "/var/log/Gagent"
)

type Options struct {
	Listen int    `json:"listen"`
	LogDir string `json:"log_dir"`

	WsContainer *restful.Container
}

func NewOptions() *Options {
	o := &Options{
		Listen: Listen,
		LogDir: LogDir,
	}
	return o
}

func (o *Options) Container() *Options {
	o.WsContainer = restful.NewContainer()
	return o
}

func (o *Options) LogClient() *Options {
	log.NewLogger(o.LogDir)
	return o
}

func (o *Options) Run() {
	fmt.Println("starting at go-restful service")
	go func() {
		if err := http.ListenAndServe(fmt.Sprintf(":%d", o.Listen), o.WsContainer); err != nil {
			panic(err)
		}
	}()
}

func (o *Options) AddFlags() *Options {
	flag.IntVar(&o.Listen, "listen", o.Listen, "please input service listen port")
	flag.StringVar(&o.LogDir, "log", o.LogDir, "please input log dir")
	flag.Parse()
	return o
}
