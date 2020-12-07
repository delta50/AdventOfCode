package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func Download(day int) string {
	sessionId, err := ioutil.ReadFile("session")
	if err != nil {
		log.Fatalln(err)
	}

	client := &http.Client{}
	cookie := http.Cookie{
		Name:     "session",
		Value:    strings.TrimRight(string(sessionId), "\n"),
		MaxAge:   0,
		Secure:   true,
		HttpOnly: true}
	request, err := http.NewRequest("GET", fmt.Sprintf("https://adventofcode.com/2020/day/%d/input", day), nil)
	request.AddCookie(&cookie)
	response, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
	}
	defer response.Body.Close()

	input, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return string(input)
}
