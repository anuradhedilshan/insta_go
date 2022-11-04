package main

import (
	"math/rand"
)

var agents = [...]string{"Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)",
	"Googlebot/2.1 (+http://www.googlebot.com/bot.html)",
	"Googlebot/2.1 (+http://www.google.com/bot.html)"}

func RandomUserAgent() string {
	n := rand.Int() % len(agents)
	return agents[n]
}
