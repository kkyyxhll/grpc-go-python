# 操作流程  

## git操作流程  

### 首次提交  

1.touch README.md  
2.git init  // 初始化  
3.git add . // .代表添加文件夹中所有文件,从工作区到暂存区  
4.git commit -m "first commit"  //从暂存区到版本库，即.git中的Git版本库，Repository  
5.git remote add origin `git@github.com`:kkyyxhll/grpc-go-python.git //建立远程连接  
6.git push -u origin master //上传  

### 后续提交  

1.git add  
2.git commit -m "next commit"  
3.git push -u origin master

## grpc初次创建流程  

`https://blog.csdn.net/m0_73537205/article/details/140258402`
`https://grpc.org.cn/docs/languages/go/quickstart/`  

### golang版本更新

1.echo $GOROOT  
2.rm -rf $GOROOT  
3.golang官方下载压缩包  
4.tar -C /usr/local/ -zxvf 压缩包路径  
5.export GOROOT=/usr/local/go  
6.export PATH=$PATH:$GOROOT/bin  
7.source /etc/profile  
8.cp -f $GOROOT/bin/go* /usr/bin/  

1.创建main.go
2.go mod init main  
3.go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28  
4.go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2  
