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
	c := http.Client{}
	return &Client{Token: token, hc: c}

}

type SearchResult struct {
	Page         int32   `json:"page"`
	PerPage      int32   `json:"per_page"`
	TotalResults int32   `json:"total_Results"`
	NextPage     string  `json:"next_page"`
	Photos       []Photo `json:"photos"`
}
type Photo struct {
	Id              int32       `json:"id"`
	Width           int32       `json:"width"`
	Height          int32       `json:"height"`
	Url             string      `json:"url"`
	Photographer    string      `json:"photographer"`
	PhotographerUrl string      `json:"photographer_url"`
	Src             PhotoSource `json:"src"`
}

type PhotoSource struct {
	Original  string `json:"original"`
	Large     string `json:"large"`
	Large2x   string `json:"large2x"`
	Medium    string `json:"medium"`
	Potrait   string `json:"potrait"`
	Squire    string `json:"squire"`
	Landscape string `json:"landscape"`
	Tiny      string `json:"tiny"`
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
