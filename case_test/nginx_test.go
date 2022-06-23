package case_test

import (
	"github.com/edisonwong520/gonginx/parser"
	"log"
	"testing"
)

func TestNgxParse(t *testing.T) {
	err := parse()
	if err != nil {
		log.Printf("err:%v", err.Error())
	}
}

func parse() (err error) {
	p , err := parser.NewParser("/Users/edison/GoProjects/src/process/workspace/nginx-public/sites/api.conf")
	if err != nil {
		log.Printf("parse err:%v" , err.Error())
		return
	}
	log.Printf("start to parse...")
	cfg, err := p.Parse()
	if err != nil {
		return
	}
	log.Printf(cfg.FilePath)
	return
}
