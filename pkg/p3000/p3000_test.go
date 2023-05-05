package p3000_test

import (
	"filetool/api/config"
	"flag"
	"testing"

	"github.com/zeromicro/go-zero/core/conf"
)

var configFile = flag.String("f", "../../etc/mwp3000-api.yaml", "the config file")

func TestOrdersPostToServer(t *testing.T) {
	LoadConfig()

	// var data []interface{}
	// data = append(data, "testing")

	// p3000.OrdersPostToServer(data)
}

func LoadConfig() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	// setting config
	// config.SetConfig(c)
}
