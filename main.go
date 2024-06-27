package main

import (
	"fmt"
	"github.com/octoboy233/mypodtrace/pkg/config"
)

func main() {
	fmt.Println(config.Config.Exporter.Jaeger.Endpoint)
	//core.WatchPod()
}
