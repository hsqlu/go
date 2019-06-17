package main

import (
	"flag"
	"learning/example"
)

func hello(val string) string { return "hello " + val }

func main() {
	var method string
	flag.StringVar(&method, "name", "", "Selected running exeample id")
	// web.Get("/(.*)", hello)
	// web.Run("0.0.0.0:9999")

	// var methodName string
	// if len(method) > 0 {
	// 	if !strings.HasPrefix(method, "Example") {
	// 		methodName = "Example" + method
	// 	}
	// 	reflect.
	// } else {
	// 	flag.PrintDefaults()
	// }
	// fmt.Println(os.Args)
	example.Example03()
}
