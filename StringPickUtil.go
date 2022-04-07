package main

import (
	"fmt"
	"regexp"
)

func PickHref(s string) [][]string {

	str := regexp.MustCompile(`/xh.*?html"`)

	if str == nil {
		fmt.Println("Error")
		return nil
	}

	return str.FindAllStringSubmatch(s, -1)

}

func PickLore(s string) [][]string {

	body, err := httpGet(s)

	if err != nil {
		fmt.Println(err)
	}

	str := regexp.MustCompile("[1-9]\u3001.*[。|”]")

	if str == nil {
		fmt.Println("Error")
		return nil
	}

	return str.FindAllStringSubmatch(body, -1)

}

func GetPageTitle(s string) string {

	body, err := httpGet(s)

	if err != nil {
		fmt.Println(err)
	}

	str := regexp.MustCompile("<title>.*?</title>")

	if str == nil {
		fmt.Println("Error")
		return ""
	}

	return str.FindAllStringSubmatch(body, -1)[0][0]

}
