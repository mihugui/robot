package joke

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	appId     = "sqnmugmnspqditkh"
	appSecret = "akw0SkhCOVoweWhWT3FQQ3dJNHgxUT09"
)

type Joke struct {
	Content string `json:"content"`
}

type Response struct {
	Code int64  `json:"code"`
	Data []Joke `json:"data"`
}

func GetJoke() string {

	// 获取城市ID
	res, err := http.Get("https://www.mxnzp.com/api/jokes/list/random?app_id=" + appId + "&app_secret=" + appSecret)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	var response Response
	// string 转json
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	if response.Code == 1 {
		return response.Data[0].Content
	} else {
		return ""
	}
}
