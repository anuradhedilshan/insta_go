package main

func main() {

	a := ProxySwicher{}
	a.Init()

}

// func handler(url string) {
// 	f, er := os.Create("./test_.txt")
// 	if er != nil {
// 		fmt.Println(er)
// 		return
// 	}

// 	client := &http.Client{}

// 	req, err := http.NewRequest("GET", url, nil)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	req.Header.Set("User-Agent", RandomUserAgent())
// 	fmt.Print("before")
// 	resp, err := client.Do(req)
// 	fmt.Print("Afteany
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	defer resp.Body.Close()
// 	body resp.Body
// 	res := bytes.Index(body, []byte(`{"config"`))

// 	f.Write(body[res:])

// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// }
