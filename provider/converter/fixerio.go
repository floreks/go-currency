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

// Points to a fixer.io api
const fixerIOApiUrl = "http://api.fixer.io/latest?base=%s"

// FixerAPIError is an error string returned by fixer api
type FixerAPIError string

// FixerAPIRates is a map of current exchange rates returned by fixer api
type FixerAPIRates map[string]float64

// FixerAPIResponse is a structure returned by fixer api (it's either error or base,date,rates)
type FixerAPIResponse struct {
	// Error - error string returned by fixer api
	Error FixerAPIError `json:"error"`
	// Base - base currency string returned by fixer api
	Base string `json:"base"`
	// Date - date on which request for exchange rates was made
	Date string `json:"date"`
	// Rates - current exchange rates returned by fixer api
	Rates FixerAPIRates `json:"rates"`
}

// FixerIOProvider represents provider used to convert exchange rates based on Fixer.io service.
// Implements ConverterProvider interface.
type FixerIOProvider struct {
	url string
}

// Name returns name of this provider
func (f FixerIOProvider) Name() string {
	return FixerIO
}

// Convert - takes the amount in one currency and converts it to other currencies
func (f FixerIOProvider) Convert(amount float64, currency string) (*ConverterResponse, error) {
	log.Printf("FixerIO provider - converting %.2f %s", amount, currency)

	rates, err := f.getRates(currency)
	if err != nil {
		return nil, err
	}

	converted := f.convert(rates, amount)
	return &ConverterResponse{Amount: amount, Currency: currency, Converted: converted}, nil
}

// Queries Fixer.io api and returns current exchange rates for currencies
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

// Does the actual conversion based on rates map and amount of currency
func (f FixerIOProvider) convert(rates FixerAPIRates, amount float64) ConvertedRates {
	for cur, rate := range rates {
		rates[cur] = common.Round((rate * amount), 2)
	}

	return ConvertedRates(rates)
}

// NewFixerIOProvider returns initialized fixer io provider object
func NewFixerIOProvider() FixerIOProvider {
	return FixerIOProvider{url: fixerIOApiUrl}
}
