package conf

import (
	"flag"
	"os"
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
	"time"
)

var (
	Conf Config
)

type Config struct {
	Db         *Db `toml:"db"`
	DbAct      *Db `toml:"dbact"`
	DbReport   *Db `toml:"dbreport"`
	Log 	   *Log `toml:"xlog"`
	Redis	   *Redis `toml:"redis"`
	HttpClient *HttpClient `toml:"httpClient"`
}

// Redis
type Redis struct {
	Proto string `toml:"proto"`
	Host  string `toml:"host"`
	Port  string `toml:"port"`
	Auth  string `toml:"auth"`
}

// httpclient
type HttpClient struct {
	Timeout	  time.Duration	`toml:"timeout"`
	KeepAlive time.Duration `toml:"keepalive"`
}

// Log
type Log struct {
	Dir	string	`toml:"dir"`
}

// DB
type Db struct {
	Addr     string `toml:"addr"`
	Port     string `toml:"port"`
	Username string `toml:"username"`
	Password string `toml:"password"`
	Db_name  string `toml:"db_name"`
}

func ParseConfig() (err error){
	_, err = toml.DecodeFile("./config.example.toml", &Conf)
	return
}

func Logger(dir string) (err error) {
	flag.Parse()
	outfile, err := os.OpenFile(dir, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(*outfile, "open failed")
		os.Exit(1)
	}
	log.SetOutput(outfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	return
}

