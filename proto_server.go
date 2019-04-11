package main

import (
	"github.com/goweb/hello"
	"github.com/gogo/protobuf/proto"
	"fmt"
)

func main() {

	test2 := &hello.Test{
		Label:"Hi",
		Type:25,
		Reps:21,
	}

	data, err := proto.Marshal(test2)
	if err != nil {
		fmt.Println(err)
	}

	newTest := hello.Test{}

			proto.Unmarshal(data,&newTest)
	fmt.Println(newTest.Reps,newTest.Type,newTest.Label)
}
