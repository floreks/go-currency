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
	"strings"
	"encoding/json"
	"github.com/floreks/go-currency/common"
)

const (
	BASE_PLN = `{
    	"base":"PLN",
    	"date":"2016-10-31",
    	"rates":{"AUD":0.33266,"BGN":0.45192,"BRL":0.80494,"CAD":0.33886,"CHF":0.25001,"CNY":1.7135,"CZK":6.2443,"DKK":1.719,"GBP":0.20807,"HKD":1.9614,"HRK":1.7351,"HUF":71.269,"IDR":3298.2,"ILS":0.97366,"INR":16.889,"JPY":26.565,"KRW":289.96,"MXN":4.7865,"MYR":1.0622,"NOK":2.0876,"NZD":0.35383,"PHP":12.261,"RON":1.0412,"RUB":16.001,"SEK":2.2794,"SGD":0.3524,"THB":8.856,"TRY":0.78481,"USD":0.25292,"ZAR":3.4309,"EUR":0.23106}
	}`
	BASE_USD = `{
    	"base":"USD",
    	"date":"2016-10-31",
    	"rates":{"AUD":1.3153,"BGN":1.7868,"BRL":3.1825,"CAD":1.3398,"CHF":0.98849,"CNY":6.7747,"CZK":24.688,"DKK":6.7964,"GBP":0.82267,"HKD":7.7551,"HRK":6.8603,"HUF":281.78,"IDR":13040.0,"ILS":3.8496,"INR":66.777,"JPY":105.03,"KRW":1146.4,"MXN":18.925,"MYR":4.1997,"NOK":8.2537,"NZD":1.399,"PHP":48.477,"PLN":3.9538,"RON":4.1166,"RUB":63.265,"SEK":9.0124,"SGD":1.3933,"THB":35.015,"TRY":3.103,"ZAR":13.565,"EUR":0.91358}
	}`
	BASE_EUR = `{
    	"base":"EUR",
    	"date":"2016-10-31",
    	"rates":{"AUD":1.4397,"BGN":1.9558,"BRL":3.4836,"CAD":1.4665,"CHF":1.082,"CNY":7.4156,"CZK":27.024,"DKK":7.4393,"GBP":0.9005,"HKD":8.4887,"HRK":7.5093,"HUF":308.44,"IDR":14273.82,"ILS":4.2138,"INR":73.094,"JPY":114.97,"KRW":1254.89,"MXN":20.715,"MYR":4.597,"NOK":9.0345,"NZD":1.5313,"PHP":53.063,"PLN":4.3278,"RON":4.506,"RUB":69.2498,"SEK":9.865,"SGD":1.5251,"THB":38.327,"TRY":3.3965,"USD":1.0946,"ZAR":14.8482}
	}`

	CURRENCY_PLN = "PLN"
	CURRENCY_USD = "USD"
	CURRENCY_EUR = "EUR"
)

// Rates is a local rates map. Created for code clarity.
type Rates map[string]float64

// LocalBaseRates is a structure used for local conversion. Similar to FixerAPIResponse.
type LocalBaseRates struct {
	// Base - base currency
	Base  string
	// Date - date on which local base rates were saved
	Date  string
	// Rates - exchange rates
	Rates Rates
}

// LocalProvider represents localprovider used to convert exchange rates. Supports only 3 base
// currencies: PLN, EUR, USD
type LocalProvider struct{}

// Name returns name of this provider
func (l LocalProvider) Name() string {
	return Local
}

// Convert - takes the amount in one currency and converts it to other currencies
func (l LocalProvider) Convert(amount float64, currency string) (*ConverterResponse, error) {
	log.Printf("Local provider - converting %.2f %s", amount, currency)

	baseRates, err := l.getBase(currency)
	if err != nil {
		return nil, err
	}

	converted := l.convert(baseRates.Rates, amount)
	return &ConverterResponse{Amount: amount, Currency: currency, Converted: converted}, nil
}

// Returns base exchange rate structure that is used for further conversion or error if given
// currency is not supported.
func (l LocalProvider) getBase(currency string) (*LocalBaseRates, error) {
	var baseJsonString string
	result := new(LocalBaseRates)

	switch strings.ToUpper(currency) {
	case CURRENCY_EUR:
		baseJsonString = BASE_EUR
	case CURRENCY_PLN:
		baseJsonString = BASE_PLN
	case CURRENCY_USD:
		baseJsonString = BASE_USD
	default:
		log.Printf("Currency %s not supported by local provider.", currency)
		return nil,
			errors.New(fmt.Sprintf("Currency %s not supported by local provider.", currency))
	}

	if err := json.Unmarshal([]byte(baseJsonString), &result); err != nil {
		return nil, err
	}

	return result, nil
}

// Does the actual conversion based on rates map and amount of currency
func (LocalProvider) convert(rates Rates, amount float64) ConvertedRates {
	for cur, rate := range rates {
		rates[cur] = common.Round((rate * amount), 2)
	}

	return ConvertedRates(rates)
}
