package models

import (
	"new-alunos/custom_errors"
	pb "new-alunos/protoimp"

	gecko "github.com/superoo7/go-gecko/v3"
)

type CoinResponse struct {
	// Name é o nome da moeda em letras minúsculas ex.: bitcoin
	Name string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	// Price é o valor corrente da crypto em "Currency"
	Price float64 `protobuf:"fixed64,2,opt,name=Price,proto3" json:"Price,omitempty"`
	// Symbol é a moeda a ser utilizada ex.: USD
	Currency string `protobuf:"bytes,3,opt,name=Currency,proto3" json:"Currency,omitempty"`
}

func ToCoin(c string) (*CoinResponse, *pb.GetCoinResponse, error) {
	fiberCoin := &CoinResponse{}
	protoCoin := &pb.GetCoinResponse{}

	cg := gecko.NewClient(nil)

	price, err := cg.SimpleSinglePrice(c, "usd")
	if err != nil {
		return nil, nil, custom_errors.ErrRecuperarCoinInformada
	}

	protoCoin.Name = price.ID
	protoCoin.Price = float64(price.MarketPrice)
	protoCoin.Currency = price.Currency

	fiberCoin.Name = price.ID
	fiberCoin.Price = float64(price.MarketPrice)
	fiberCoin.Currency = price.Currency

	return fiberCoin, protoCoin, nil
}
