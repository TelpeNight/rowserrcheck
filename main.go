package main

import (
	"github.com/TelpeNight/rowserrcheck/passes/rowserr"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() { singlechecker.Main(rowserr.NewAnalyzer()) }
