package main

import (
    "time"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
    "net/url"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parseUrl(s string) string{
    u, _ := url.Parse(s)
    return u.Path[1:]
}

func download(_url string, filename string) {
	img, _ := os.Create(filename)
    resp, _ := http.Get(_url)
    io.Copy(img, resp.Body)
	defer img.Close()
	defer resp.Body.Close()
}

func ns() {
    now := time.Now()
    fmt.Println(now)
}

func main() {
	content, _ := ioutil.ReadFile("./urls.txt")
	lines := strings.Split(string(content), "\n")

	for _, v := range lines {
		if v != "" {
            download(v, parseUrl(v))
            // println(parseUrl(v))
            // download(v)
		}
	}
}
