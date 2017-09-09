package main

import (
	"encoding/base64"
	"fmt"
	"hello-go/netprogram/util"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

const auth = "jannewmarch:mypassword"

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: ", os.Args[0], "http://proxy-host:port http://host:port/page")
		os.Exit(1)
	}
	proxy := os.Args[1]
	proxyUrl, err := url.Parse(proxy)
	util.CheckError(err)
	rawURL := os.Args[2]
	url, err := url.Parse(rawURL)
	util.CheckError(err)

	// encode the auth
	basic := "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))

	transport := &http.Transport{Proxy: http.ProxyURL(proxyUrl)}
	client := &http.Client{Transport: transport}

	request, err := http.NewRequest("GET", url.String(), nil)

	request.Header.Add("Proxy-Authorization", basic)
	dump, _ := httputil.DumpRequest(request, false)
	fmt.Println(string(dump))

	// send the request
	response, err := client.Do(request)

	util.CheckError(err)
	fmt.Println("Read ok")

	if response.Status != "200 OK" {
		fmt.Println(response.Status)
		os.Exit(2)
	}
	fmt.Println("Response ok")

	var buf [512]byte
	reader := response.Body
	for {
		n, err := reader.Read(buf[0:])
		if err != nil {
			os.Exit(0)
		}
		fmt.Print(string(buf[0:n]))
	}
	os.Exit(0)

}
