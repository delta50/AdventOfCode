package main

import (
	"log"
	"os"
	"reflect"
	"regexp"
	"strconv"
)

type T struct{}

func main() {
	method := os.Args[1]
	day_regex := regexp.MustCompile("(\\d+)")
	day, err := strconv.Atoi(day_regex.FindString(method))
	if err != nil {
		log.Fatal(err)
	}
	input := Download(day)
	var t T
	reflect.ValueOf(&t).MethodByName(method).Call([]reflect.Value{reflect.ValueOf(input)})
}
