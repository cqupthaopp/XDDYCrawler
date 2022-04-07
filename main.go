package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const webAddress = "http://xiaodiaodaya.cn/xh/sh/"
const webIP = "http://xiaodiaodaya.cn"

func httpGet(ip string) (string, error) {

	resp, err := http.Get(ip)

	if err != nil {
		return "", errors.New("GetError")
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	return string(body), nil

}

func main() {

	body, err := httpGet(webAddress)

	if err != nil {
		fmt.Println(err)
	}

	for _, v := range PickHref(body) {

		address := webIP + strings.Replace(v[0], "\"", "", -1)

		title := GetPageTitle(address)
		title = strings.Replace(title, "<title>", "", -1)
		title = strings.Replace(title, "</title>", "", -1)

		os.Mkdir("./Lore", 077)

		f, err := os.OpenFile("./Lore/"+title, os.O_RDWR|os.O_CREATE, 0766)

		if err != nil {
			fmt.Println(err)
		}

		f.Write([]byte("Title: " + title + "\n\n\n"))

		for _, lore := range PickLore(address) {
			f.Write([]byte(lore[0] + "\n"))
		}

		f.Close()

	}

}
