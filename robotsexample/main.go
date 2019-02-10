package main

import (
	"net/http"
	"github.com/temoto/robotstxt"
)

func main() {
	resp, err := http.Get("https://www.packtpub.com/robots.txt")
	if err != nil {
		panic(err)
	}

	data, err := robotstxt.FromResponse(resp)
	if err != nil {
		panic(err)
	}

	grp := data.FindGroup("Go-http-client/1.1")
	if grp != nil {
		testUrls := []string{
			"/all",
			"/all?search=Go",
			"/bundles",
			"/contact/",
			"/search/",
			"/user/password/",
		}

		for _, url := range testUrls {
			print("checking " + url + "...")

			if grp.Test(url) == true {
				println("OK")
			} else {
				println("X")
			}
		}
	}
}
