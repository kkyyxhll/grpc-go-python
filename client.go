// client.go
package main

import (
	"GRPC-GO-PYTHON/proto"
	"context"
	"encoding/base64"
	"fmt"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/grpclog"
)

func clientTest(addr string) {
	// 创建rpc连接
	cc, err := grpc.NewClient(addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		panic(err)
	}
	defer cc.Close()

	// 创建rpc调用客户端
	c := proto.NewImageServiceClient(cc)

	image_path := "./image/image.png"
	image_base64, err := image_2_base64(image_path)
	if err != nil {
		fmt.Print("hahah")
	}
	fmt.Println(image_base64)
	req := &proto.ImageInput{ImageInput: image_base64}

	res, err := c.ProcessImage(context.Background(), req)

	if err != nil {
		fmt.Println(res)
		grpclog.Fatalln("报错拉报错拉")
	}
	output_path := "output_image/1.png"
	base64_2_image(res.ImageOutput, output_path)

}

// base64_2_image将Base64编码数据转换为图片并保存到指定路径
func base64_2_image(base64Str string, outputPath string) error {
	// 解码Base64字符串为字节数据
	imageData, err := base64.StdEncoding.DecodeString(base64Str)
	if err != nil {
		return err
	}

	// 将字节数据写入到指定路径的文件中，创建图片文件
	err = os.WriteFile(outputPath, imageData, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

// image_2_base64将指定路径的图片转换为Base64编码
func image_2_base64(imagePath string) (string, error) {
	// 读取图片文件
	imageData, err := os.ReadFile(imagePath)
	if err != nil {
		return "", err
	}

	// 将图片数据转换为Base64编码
	base64Str := base64.StdEncoding.EncodeToString(imageData)

	return base64Str, nil
}
