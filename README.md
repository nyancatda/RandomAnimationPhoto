# RandomAnimationPhoto
随机下载指定数量的二次元图片  
Golang入门写的练手项目，能run

### 说明:  
随机从API获取指定数量的二次元图片链接，并下载至同级img目录(运行时自动创建)

## 运行
```
go run ./main.go
```
如果遇到错误`build command-line-arguments: cannot find module for path xxxxxxxxxx`，请先运行
```
go env -w GO111MODULE=auto
```

## 构建
```
go build ./main.go
```