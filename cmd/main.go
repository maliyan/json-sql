// Package main 程序入口.
package main

import (
	"flag"

	"github.com/maliyan/json-sql/internal/helper/log"
)

var (
	query = flag.String("q", "", "query")
)

func main() {
	flag.Parse()
	if *query == "" {
		log.ConsoleErr("empty query statement")
		return
	}

}
