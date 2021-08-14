package utils

import (
    "os"
    "fmt"
    "net/http"
    "io/ioutil"
    "encoding/json"
)

type UrlJson struct {
    Code string `json:"code"`
    Message string `json:"message"`
    Time int `json:"time"`
    Version string `json:"version"`
    Data DataJson `json:"data"`
}

type DataJson struct {
    Num string `json:"num"`
    Function string `json:"function"`
    Urls []string `json:"url"`
}

//获取色图链接
func GetSexUrl()(string){
    url := "https://api.nyan.xyz/httpapi/sexphoto/"
    res, err := http.Get(url)
    if err != nil {
        fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
        os.Exit(1)
    }
    body, err := ioutil.ReadAll(res.Body)
    res.Body.Close()
    if err != nil {
        fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
        os.Exit(1)
    }
    var config UrlJson
	json.Unmarshal([]byte(body), &config)
    PhotoUrl := config.Data.Urls[0]
    return PhotoUrl
}