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

## goProject初次创建流程  

1.创建main.go
2.go mod init main  
