package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
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
type VideoSearchResult struct {
	Page         int32   `json:"page"`
	PerPage      int32   `json:"per_page"`
	TotalResults int32   `json:"total_results"`
	NextPage     string  `json:"next_page"`
	Videos       []Video `json:"videos"`
}
type Video struct {
}
type PopularVideo struct {
}
type VideoFiles struct {
}
type VideoPictures struct {
}

func (c *Client) SearchPhotos(query string, perPage, page int) (*SearchResult, error) {
	url := fmt.Sprintf(PhotoApi+"/search?query=%s&per_page=%d&page=%d", query, perPage, page)
	resp, err := c.requestDoWithAuth("GET", url)
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result SearchResult
	err = json.Unmarshal(data, &result)
	return &result, err
}
func (c *Client) requestDoWithAuth(method, url string) (*http.Response, err) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", c.Token)
	resp, err := c.hc.Do(req)
	if err != nil {
		return resp, err
	}
	times, err := strconv.Atoi(resp.Header.Get("X-Ratelimit-Remaining"))
	if err != nil {
		return resp, nil
	} else {
		c.RemainingTimes = int32(times)
	}
	return resp, nil

}
func (c *Client) GetPhoto(id int32) (*Photo, error) {
	url := fmt.Sprintf(PhotoApi+"/photos/%d", id)
	resp, err := c.requestDoWithAuth("GET", url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	var result Photo
	err = json.Unmarshal(data, &result)
	return &result, err

}

func (c *Client) GetRandomPhoto() (*Photo, error) {
	rand.Seed(time.Now().Unix())
	randNum := rand.Intn(1001)
	result, err := c.CuratedPhotos(1, randNum)
	if err != nil && len(result.Photos) == 1 {
		return &result.Photos[0], nil
	}
	return nil, err

}

func (c *Client) SearchVideo(query, perPage, page int) (*VideoSearchReasult, error) {

}
func (c *Client) PopularVideo(perPage, page int) (*PopularVideos, error) {

}
func (c *Client) GetRandomVideo() (*Video, error) {

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
