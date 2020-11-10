package helper

import (
	"crypto/tls"
	"crypto/x509"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
)

//获取服务端证书
func GetServerCreds() credentials.TransportCredentials{
	cert, _ := tls.LoadX509KeyPair("cert/server.pem","cert/server.key")
	certPool := x509.NewCertPool()
	ca, _ := ioutil.ReadFile("cert/ca.pem")
	certPool.AppendCertsFromPEM(ca)
	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},	//服务端证书
		ClientAuth: tls.RequireAndVerifyClientCert,	//需要验证客户端证书
		ClientCAs: certPool,	//指定客户端的ca
	})
	return  creds
}

//获取客户端证书,http请求过来后他会主动请求grpc服务器，所以需要客户端证书
func GetClientCreds() credentials.TransportCredentials{
	cert, _ := tls.LoadX509KeyPair("cert/client.pem","cert/client.key")
	certPool := x509.NewCertPool()
	ca, _ := ioutil.ReadFile("cert/ca.pem")
	certPool.AppendCertsFromPEM(ca)
	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},	//客户端证书
		ServerName: "localhost",	//域名
		RootCAs: certPool,
	})
	return creds
}