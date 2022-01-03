package main

import (
	"gpm/day09/k8s_visitor"
	"gpm/day09/pipeline_visitor"
	"gpm/day09/simple_visitor"
)

func main() {
	simple_visitor.Example()
	pipeline_visitor.Example()
	k8s_visitor.Example()
}
