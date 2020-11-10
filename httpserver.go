package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"grpc-pro/helper"
	"grpc-pro/services"
	"log"
	"net/http"
)

func main(){
	gwmux := runtime.NewServeMux() //创建路由
	opt := []grpc.DialOption{grpc.WithTransportCredentials(helper.GetClientCreds())}	//指定客户端请求时用的证书
	err := services.RegisterProdServiceHandlerFromEndpoint(context.Background(), gwmux, "localhost:8081", opt)	//这个方法在Prod.pb.gw.go,endpoint指的是grpc服务端的地址
	if err != nil{
		log.Fatal(err)
	}
	httpServer := &http.Server{
		Addr: ":8080",
		Handler: gwmux,
	}
	httpServer.ListenAndServe()
}
