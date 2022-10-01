package models

import (
	"encoding/json"
	"io"
	"net/http"
	ce "new-alunos/custom_errors"
	pb "new-alunos/protoimp"

	"github.com/paemuri/brdoc"
)

type Address struct {
	Cep         string `json:"cep,omitempty"`
	Logradouro  string `json:"logradouro,omitempty"`
	Bairro      string `json:"bairro,omitempty"`
	Localidade  string `json:"localidade,omitempty"`
	Uf          string `json:"uf,omitempty"`
	Complemento string `json:"complemento,omitempty"`
}

func GetData(url string) ([]byte, error) {

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	jsonByte, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return jsonByte, nil

}

func isCepValid(cep string) bool {
	return brdoc.IsCEP(cep)
}

func ToAdress(cep string) (*Address, *pb.Address, error) {
	endereco := &Address{}
	proto_endereco := &pb.Address{}

	if cep == "" {
		return nil, nil, ce.ErrCEPnaoInformado
	}
	if !isCepValid(cep) {
		return nil, nil, ce.ErrCEPInvalido
	}
	data, errData := GetData("https://viacep.com.br/ws/" + cep + "/json/")
	if errData != nil {
		return nil, nil, ce.ErrCEPnaoRecuperado
	}

	err := json.Unmarshal(data, &endereco)
	if err != nil {
		return nil, nil, ce.ErrCEPInvalido
	}

	err = json.Unmarshal(data, &proto_endereco)
	if err != nil {
		return nil, nil, ce.ErrCEPInvalido
	}

	return endereco, proto_endereco, nil

}
