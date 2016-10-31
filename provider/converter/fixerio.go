package converter

import (
	"errors"
	"fmt"
	"log"

	"github.com/floreks/go-currency/common"
)

const apiUrl = "http://api.fixer.io/latest?base=%s"

type FixerAPIError string

type FixerAPIRates map[string]float64

type FixerAPIResponse struct {
	Error FixerAPIError `json:"error"`
	Base  string        `json:"base"`
	Date  string        `json:"date"`
	Rates FixerAPIRates `json:"rates"`
}

type FixerIOProvider struct {
	url string
}

func (f FixerIOProvider) Name() string {
	return FixerIO
}

func (f FixerIOProvider) Convert(amount float64, currency string) (*ConverterResponse, error) {
	log.Printf("FixerIO - converting %.2f %s", amount, currency)

	rates, err := f.getRates(currency)
	if err != nil {
		return nil, err
	}

	converted := f.convert(rates, amount)
	return &ConverterResponse{Amount: amount, Currency: currency, Converted: converted}, nil
}

func (f FixerIOProvider) getRates(currency string) (FixerAPIRates, error) {
	fixerAPIResponse := new(FixerAPIResponse)
	err := common.GetJson(fmt.Sprintf(f.url, currency), &fixerAPIResponse)
	if err != nil {
		log.Printf("Error during request to fixer.io: %s", err)
		return nil, err
	}

	if len(fixerAPIResponse.Error) != 0 {
		log.Printf("Fixer.io returned error: %s", fixerAPIResponse.Error)
		return nil, errors.New(string(fixerAPIResponse.Error))
	}

	return fixerAPIResponse.Rates, nil
}

func (f FixerIOProvider) convert(rates FixerAPIRates, amount float64) convertedRates {
	for cur, rate := range rates {
		rates[cur] = common.Round((rate * amount), 2)
	}

	return convertedRates(rates)
}

func NewFixerIOProvider() FixerIOProvider {
	return FixerIOProvider{url: apiUrl}
}
