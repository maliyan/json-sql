// Package main 程序入口.
package main

import (
	"flag"

	"github.com/maliyan/json-sql/internal/helper/log"
)

var (
	query    = flag.String("q", "", "query")
	fromFile = flag.String("f", "", "specify file to query")
)

func main() {
	flag.Parse()
	if *query == "" {
		log.ConsoleErr("empty query statement")
		return
	}

	// 从文件读取
	if *fromFile != "" {

	}
}
