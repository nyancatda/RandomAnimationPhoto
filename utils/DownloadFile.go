package utils

import (
    "os"
    "fmt"
    "net/http"
    "io"
    "path"
    "bufio"
    "time"
)

//下载文件
func GetFile(imgUrl string,imgPath string) {
    fileName := path.Base(imgUrl)
	//设置30秒超时
    client := http.Client{
    	Timeout: 30 * time.Second,
    }
    res, err := client.Get(imgUrl)
    if err != nil {
    	fmt.Println("A error occurred!")
    	return
    }
    defer res.Body.Close()
    reader := bufio.NewReaderSize(res.Body, 32 * 1024)
    file, err := os.Create(imgPath + fileName)
    if err != nil {
    	panic(err)
    }
    writer := bufio.NewWriter(file)
    io.Copy(writer, reader)
    //written, _ := io.Copy(writer, reader)
    //fmt.Printf("Total length: %d", written)
}