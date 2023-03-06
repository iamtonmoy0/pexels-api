package main

import (
	"fmt"
	"net/http"
	"os"
)

// api link
const (
	PhotoApi = "https://api.pexels.com/v1"
	VideoApi = "https://api.pexels.com/videos"
)

type Client struct {
	Token          string
	hc             http.Client
	RemainingTimes int32
}

func NewClient(token string) *Client {

}

func main() {
	os.Setenv("PexelsToken", "OLY1UXu7nWNqhhiV5XXXTcU8SHJPaMUEWzotNouYLKhqNuTyLsnXjgxS")
	var TOKEN = os.Getenv("PexelsToken")
	var c = NewClient(TOKEN)
	result, err := c.SearchPhotos("wave")
	if err != nil {
		fmt.Errorf("Search error:%v", err)
	}
	if result.Page == 0 {
		fmt.Errorf("wrong search result")
	}
	fmt.Println(result)
}
