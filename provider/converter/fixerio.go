// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
