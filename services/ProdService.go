package services

import (
	"context"
)

type ProdService struct {	//创建实现proto中GetProdStock的结构体

}

func (this *ProdService) GetProdStock(context.Context, *ProdRequest) (*ProdResponse, error){	//方法的具体参数和返回直接复制pb.go中prodServiceClient实现的GetProdStock
	return &ProdResponse{ProdStock:20}, nil		//自定义处理函数，返回的ProdResponse和proto中定义的ProdResponse一样
}