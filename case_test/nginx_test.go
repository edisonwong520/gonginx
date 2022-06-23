package case_test

import (
	"encoding/json"
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
	p , err := parser.NewParser("/Users/edison/gitdownload/keepgit/nginx_config/sites/api.conf")
	if err != nil {
		log.Printf("parse err:%v" , err.Error())
		return
	}
	log.Printf("start to parse...")
	cfg, err := p.Parse()
	if err != nil {
		return
	}
	log.Printf(PrintJson(cfg.GetDirectives()))
	return
}

func PrintJson(it interface{})(s string){
	b,_:=json.Marshal(it)
	return string(b)
}