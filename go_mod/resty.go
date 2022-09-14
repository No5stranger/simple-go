package go_mod

import (
	"fmt"
	"github.com/go-resty/resty/v2"
)

type GithubUser struct {
	ID  int    `json:"id"`
	URL string `json:"url"`
}

func Hello() {
	fmt.Println("hello world")
}

func RestyRequest() {
	client := resty.New()
	//client.JSONMarshal = json.Marshal
	//client.JSONUnmarshal = json.Unmarshal
	gu := new(GithubUser)
	resp, err := client.R().SetResult(gu).EnableTrace().Get("https://api.github.com/users/No5stranger")
	if err != nil {
		fmt.Println("Err: ", err)
	}
	fmt.Println(resp.Proto())
	fmt.Println("Result: ", gu)
	fmt.Println("Status: ", resp.Status())
}
