package main

import (
	"log"
	"os"

	_ "github.com/liumc/goinaction/chapter02/sample/matchers"
	"github.com/liumc/goinaction/chapter02/sample/search"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("president")
}
