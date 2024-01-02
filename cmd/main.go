package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/emejotaw/multithread-fullcycle-challenge/pkg/client/httpclient"
	"github.com/emejotaw/multithread-fullcycle-challenge/pkg/dto"
)

func main() {

	zipCode := os.Args[1]
	viaCepChannel := make(chan *dto.ViaCepAddressResponseDTO)
	brasilApiChannel := make(chan *dto.BrasilApiAddressResponseDTO)
	encoder := json.NewEncoder(os.Stdout)

	go GetBrasilApiAddress(zipCode, brasilApiChannel)
	go GetViaCepAddress(zipCode, viaCepChannel)

	select {
	case responseBody := <-viaCepChannel:
		log.Println("Got address from viacep")
		encoder.Encode(responseBody)
	case responseBody := <-brasilApiChannel:
		log.Println("Got address from brasilapi")
		encoder.Encode(responseBody)
	case <-time.After(time.Second):
		fmt.Println("timeout")
	}
}

func GetViaCepAddress(zipCode string, channel chan<- *dto.ViaCepAddressResponseDTO) error {

	url := fmt.Sprintf("http://viacep.com.br/ws/%s/json/", zipCode)

	client := &httpclient.HttpClient{}
	responseBody, err := client.GetAddress(url)

	if err != nil {
		return err
	}

	responseDTO := &dto.ViaCepAddressResponseDTO{}
	err = json.Unmarshal(responseBody, responseDTO)

	if err != nil {
		return err
	}
	channel <- responseDTO
	return nil
}

func GetBrasilApiAddress(zipCode string, channel chan *dto.BrasilApiAddressResponseDTO) error {

	url := fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", zipCode)

	client := &httpclient.HttpClient{}
	responseBody, err := client.GetAddress(url)

	if err != nil {
		return err
	}

	responseDTO := &dto.BrasilApiAddressResponseDTO{}
	err = json.Unmarshal(responseBody, responseDTO)

	if err != nil {
		return err
	}

	channel <- responseDTO
	return nil
}
