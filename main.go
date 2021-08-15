package main

import (
    "os"
    "fmt"
    "sync"
)

import "./utils"

var num int

//创建目录
//判断目录是否存在,不存在则创建
func mkdirFolder(path string){
    _, err := os.Stat(path)
    if os.IsNotExist(err) {
        os.Mkdir(path, 0777)
    }
}

//使用多线程模式下载图片
func Multithreading(){
    var wg sync.WaitGroup
    for i:=1;i<=num;i++ {
        sexUrl := utils.GetSexUrl()
        fmt.Println("已获取到图片链接:"+sexUrl+" ,开始下载")
        //创建一个新线程下载图片
        wg.Add(1)
        go func(i int) {
            utils.GetFile(sexUrl,"./img/",240)
            defer wg.Done()
        }(i)
    }
    fmt.Println("请等待所有下载线程执行结束")
    wg.Wait()
}

func main() {
    var MultithreadingMode int =  1

    fmt.Println("是否启用多线程模式(1):\n1.启用\n2.不启用")
    fmt.Scanln(&MultithreadingMode)
    fmt.Println("请问需要下载几张图片:")
    fmt.Scanln(&num)

    mkdirFolder("./img")

    switch {
        case MultithreadingMode==2 :
            for i:=1;i<=num;i++ {
                sexUrl := utils.GetSexUrl()
                fmt.Println("已获取到图片链接:"+sexUrl+" ,正在下载")
                utils.GetFile(sexUrl,"./img/",60)
            }
        default:
            Multithreading()
    }
}