package main

import (
    "os"
    "fmt"
)

import "./utils"

//创建目录
//判断目录是否存在,不存在则创建
func mkdirFolder(path string){
    _, err := os.Stat(path)
    if err == nil {}
    if os.IsNotExist(err) {
        os.Mkdir(path, 777)
    }
}

func main() {
    var num int

    mkdirFolder("./img")

    fmt.Println("请问需要下载几张图片:")
    fmt.Scanf("%v", &num)

    for i:=1;i<=num;i++ {
        sexUrl := utils.GetSexUrl()
        fmt.Println("已获取到图片链接:"+sexUrl+" ,正在下载")
        utils.GetFile(sexUrl,"./img/")
    }
}