package main

import (
	"fmt"

	greetv1 "github.com/arianics/bazel-buf-demo/src/protos/greetapis/greet/v1"
)

func main() {
	fmt.Println("Hello greet service")

	fmt.Println(greetv1.SayHelloRequest{
		Name: "Greeter Name",
	})
}
