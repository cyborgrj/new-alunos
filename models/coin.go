package models

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	ce "new-alunos/custom_errors"
	pb "new-alunos/protoimp"
	"time"
)

type CoinResponse struct {
	// Name é o nome completo (por extenso) da moeda, em caps
	Name string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	// Symbol é o nome reduzido (BTC, ETH) da moeda, em caps
	Symbol string `protobuf:"bytes,2,opt,name=Symbol,proto3" json:"Symbol,omitempty"`
	// MarketCapUSD é o valor corrente da moeda no mercado
	MarketCapUSD float64 `protobuf:"fixed64,3,opt,name=MarketCapUSD,proto3" json:"MarketCapUSD,omitempty"`
	// Price é o valor corrente da moeda em USD
	Price float64 `protobuf:"fixed64,4,opt,name=Price,proto3" json:"Price,omitempty"`
	// CirculatingSupply é a quantidade da mueda que está disponível/circulando no mercado
	CirculatingSupply float64 `protobuf:"fixed64,5,opt,name=CirculatingSupply,proto3" json:"CirculatingSupply,omitempty"`
	// Volume é o valor todal em dólar que foi negociado nas últimas 24 horas
	Volume float64 `protobuf:"fixed64,6,opt,name=Volume,proto3" json:"Volume,omitempty"`
	// Change é o valor percentual que indica as variações no valor corrente,
	// na última hora ou dia
	ChangeHour string `protobuf:"bytes,8,opt,name=ChangeHour,proto3" json:"ChangeHour,omitempty"`
	ChangeDay  string `protobuf:"bytes,9,opt,name=ChangeDay,proto3" json:"ChangeDay,omitempty"`
}

type Coin struct {
	Status struct {
		Timestamp time.Time `json:"timestamp"`
	} `json:"status"`
	Data struct {
		Num1 struct {
			ID        int       `json:"id"`
			Name      string    `json:"name"`
			Symbol    string    `json:"symbol"`
			Slug      string    `json:"slug"`
			DateAdded time.Time `json:"date_added"`
			Quote     struct {
				Usd struct {
					Price           float64   `json:"price"`
					Volume24H       float64   `json:"volume_24h"`
					VolumeChange24H float64   `json:"volume_change_24h"`
					LastUpdated     time.Time `json:"last_updated"`
				} `json:"USD"`
			} `json:"quote"`
		} `json:"1"`
	} `json:"data"`
}

func GetCoinData(urlCoinMarket string, coinSymbol string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", urlCoinMarket, nil)
	if err != nil {
		return nil, err
	}

	q := url.Values{}
	q.Add("symbol", coinSymbol)
	q.Add("convert", "USD")

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", "bac5c65d-4865-41ec-86c1-70cd328bdedc")
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		log.Println("Erro enviando request ao servidor")
		return nil, err
	}

	jsonByte, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return jsonByte, nil

}

func ToCoin(coin string) (*CoinResponse, *pb.GetCoinResponse, error) {
	apiCoin := Coin{}
	fiberCoin := &CoinResponse{}
	protoCoin := &pb.GetCoinResponse{}
	urlCoinMarket := "https://pro-api.coinmarketcap.com/v2/cryptocurrency/quotes/latest"

	coinData, err := GetCoinData(urlCoinMarket, "BTC")
	if err != nil {
		return nil, nil, ce.ErrRecuperarCoinInformada
	}

	err = json.Unmarshal(coinData, &apiCoin)
	if err != nil {
		return nil, nil, ce.ErrUnmarshalCoinInformada
	}

	protoCoin.Name = apiCoin.Data.Num1.Name
	protoCoin.Price = apiCoin.Data.Num1.Quote.Usd.Price
	protoCoin.Symbol = apiCoin.Data.Num1.Symbol
	protoCoin.Volume = apiCoin.Data.Num1.Quote.Usd.Volume24H

	fiberCoin.Name = apiCoin.Data.Num1.Name
	fiberCoin.Price = apiCoin.Data.Num1.Quote.Usd.Price
	fiberCoin.Symbol = apiCoin.Data.Num1.Symbol
	fiberCoin.Volume = apiCoin.Data.Num1.Quote.Usd.Volume24H

	return fiberCoin, protoCoin, nil
}