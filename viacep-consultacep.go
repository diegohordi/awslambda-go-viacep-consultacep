/**
 * ViaCEP JSON GO client
 */
package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	ViaCepUrl = "https://viacep.com.br/ws/%s/json"
	ConsultaIndisponivel = "CONSULTA_INDISPONIVEL"
	CepNaoEncontrado = "CEP_NAO_ENCONTRADO"
	ErroInesperado = "ERRO_INESPERADO"
)

type Cep struct {
	Cep string `json:"cep"`
	Logradouro string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro string `json:"bairro"`
	Localidade string `json:"localidade"`
	Uf string `json:"uf"`
	Unidade string `json:"unidade"`
	Ibge string `json:"ibge"`
	Erro bool `json:"erro"`
}

func logInfo(message string){
	fmt.Printf("[INFO] %s %s \n", time.Now().Format("2006-01-02 15:04:05"), message)
}

func logError(err error){
	fmt.Printf("[ERROR] %s %s \n", time.Now().Format("2006-01-02 15:04:05"), err.Error())
}

func parse(response * http.Response)(resultado Cep, err error){
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logError(err)
		err = errors.New(ErroInesperado)
		return
	}
	err = json.Unmarshal(body, &resultado)
	if err != nil {
		logError(err)
		err = errors.New(ErroInesperado)
		return
	}
	if resultado.Erro {
		err = errors.New(CepNaoEncontrado)
		logError(err)
		return
	}
	return
}

func Consultar(cep string)(resultado Cep, err error){
	url := fmt.Sprintf(ViaCepUrl, cep)
	logInfo(fmt.Sprintf("Pesquisando %s", url))
	response, err := http.Get(url)
	if err != nil {
		logError(err)
		err = errors.New(ConsultaIndisponivel)
		return
	}
	defer response.Body.Close()
	if response.StatusCode != 200{
		err = errors.New(CepNaoEncontrado)
		logError(err)
		return
	}
	resultado, err = parse(response)
	return
}

func HandleRequest(ctx context.Context, cep Cep) (Cep, error) {
	return Consultar(cep.Cep)
}

func main() {
	lambda.Start(HandleRequest)
}
