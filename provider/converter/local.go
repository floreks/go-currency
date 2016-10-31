package converter

import (
	"log"
)

type LocalProvider struct{}

func (l LocalProvider) Name() string {
	return Local
}

func (l LocalProvider) Convert(amount float64, currency string) (*ConverterResponse, error) {
	log.Printf("Local - converting %.2f %s", amount, currency)
	return &ConverterResponse{}, nil
}
