package main

import (
	"crypto/tls"
	"crypto/x509"
	"google.golang.org/grpc"
	"grpc-pro/helper"
	"grpc-pro/services"
	"net/http"
)

func main() {
	//cresd, err := credentials.NewServerTLSFromFile("keys/server.crt","keys/server_no_passwd.key")	//创建服务端证书
	//if err != nil{
	//	log.Fatal(err)
	//}
	creds := helper.GetServerCreds()

	rpcServer := grpc.NewServer(grpc.Creds(creds))	//使用证书
	services.RegisterProdServiceServer(rpcServer, new(services.ProdService))	//注册服务端，Register定义在pb.go
	//
	//lis, _ := net.Listen("tcp", ":8081")
	//rpcServer.Serve(lis)

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		rpcServer.ServeHTTP(writer, request)
	})
	httpServer := &http.Server{
		Addr: ":8081",
		Handler: mux,
	}
	httpServer.ListenAndServeTLS("keys/server.crt","keys/server_no_passwd.key")
}
