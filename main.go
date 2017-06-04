package main

import (
    "bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parseUrl(s string) string {
	u, _ := url.Parse(s)
	return u.Path[1:]
}

func download(_url string, filename string) {
	img, _ := os.Create(filename)
	resp, _ := http.Get(_url)

    defer img.Close()
	defer resp.Body.Close()

    io.Copy(img, resp.Body)
}

func readStdin() {
    scanner := bufio.NewScanner(os.Stdin)

    for scanner.Scan() {
        line := scanner.Text()
        download(line, parseUrl(line))
    }
}

func downloadFile() {
	content, _ := ioutil.ReadFile("./urls.txt")
	lines := strings.Split(string(content), "\n")

    for _, v := range lines {
		if v != "" {
			download(v, parseUrl(v))
		}
	}
}

func main() {
	readStdin()
}
