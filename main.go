package main

import (
    "os"
    "fmt"
    "sync"
)

import "./utils"

//创建目录
//判断目录是否存在,不存在则创建
func mkdirFolder(path string){
    _, err := os.Stat(path)
    if os.IsNotExist(err) {
        os.Mkdir(path, 0777)
    }
}

func main() {
    var wg sync.WaitGroup
    var num int

    fmt.Println("请问需要下载几张图片:")
    fmt.Scanf("%v", &num)

    mkdirFolder("./img")

    for i:=1;i<=num;i++ {
        sexUrl := utils.GetSexUrl()
        fmt.Println("已获取到图片链接:"+sexUrl+" ,开始下载")
        wg.Add(1)
        //创建一个新线程下载图片
        go func(i int) {
            utils.GetFile(sexUrl,"./img/")
            defer wg.Done()
        }(i)
    }

    fmt.Println("请等待所有下载线程执行结束")
    wg.Wait()
}