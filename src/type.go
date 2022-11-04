package main

type Proxylist struct {
	Data []Proxy `json:"data"`
}

type Proxy struct {
	Id        string   `json:"_id"`
	Ip        string   `json:"ip"`
	Port      string   `json:"port"`
	Speed     int16    `json:"speed"`
	Google    bool     `json:"google"`
	Protocols []string `json:"protocols"`
}
