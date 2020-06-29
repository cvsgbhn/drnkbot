package main

import (
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

// whatthedrink отдает рандомную подстроку из тех, что находятся внутри в var opts
func whatthedrink() string {
	opts := strings.Split("чй 🍵,кф ☕", ",")
	rand.Seed(time.Now().UnixNano())
	return opts[rand.Intn(len(opts))]
}

func getweather() string {
	var bodyString string
	resp, err := http.Get("http://wttr.in/SVO?format=1")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		bodyString = string(bodyBytes)
	}
	return bodyString
}

func createmessage(drink string, weather string) string {
	if len(drink) == 0 {
		return "сам выбирай UwU 🤍  в мск " + weather
	}
	if len(weather) == 0 {
		return drink + "🤍 у природы нет плохой погоды 🤍"
	}
	return drink + "     🤍  в мск " + weather
}
