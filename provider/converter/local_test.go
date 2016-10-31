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
	"reflect"
	"testing"
)

func TestConvertLocal(t *testing.T) {
	provider := new(LocalProvider)
	cases := []struct {
		rates    Rates
		amount   float64
		expected ConvertedRates
	}{
		{
			Rates{
				"USD": 1,
				"PLN": 1.5,
				"EUR": 0.5,
			},
			10,
			ConvertedRates{
				"USD": 10,
				"PLN": 15,
				"EUR": 5,
			},
		},
	}

	for _, c := range cases {
		actual := provider.convert(c.rates, c.amount)

		if !reflect.DeepEqual(actual, c.expected) {
			t.Errorf("LocalProvider.convert(%v, %v) == \ngot: %v, \nexpected %v",
				c.rates, c.amount, actual, c.expected)
		}
	}
}

func TestConvert(t *testing.T) {
	provider := new(LocalProvider)
	cases := []struct {
		currency      string
		amount        float64
		expected      *ConverterResponse
		expectedError error
	}{
		{
			currencyPLN, 10,
			&ConverterResponse{
				Amount:   10,
				Currency: currencyPLN,
				Converted: ConvertedRates{
					"IDR": 32982, "JPY": 265.65, "ZAR": 34.31, "BRL": 8.05, "HRK": 17.35,
					"MXN": 47.87, "MYR": 10.62, "NOK": 20.88, "USD": 2.53, "CNY": 17.14,
					"HKD": 19.61, "PHP": 122.61, "RUB": 160.01, "CHF": 2.5, "NZD": 3.54,
					"SEK": 22.79, "EUR": 2.31, "ILS": 9.74, "GBP": 2.08, "KRW": 2899.6,
					"BGN": 4.52, "CAD": 3.39, "CZK": 62.44, "DKK": 17.19, "HUF": 712.69,
					"INR": 168.89, "AUD": 3.33, "TRY": 7.85, "SGD": 3.52, "THB": 88.56,
					"RON": 10.41,
				}}, nil,
		},
		{
			"ERR_CURRENCY", 10,
			nil,
			errors.New("Currency ERR_CURRENCY not supported by local provider."),
		},
	}

	for _, c := range cases {
		actual, err := provider.Convert(c.amount, c.currency)

		if !reflect.DeepEqual(err, c.expectedError) {
			t.Errorf("LocalProvider.Convert(%f, %s) == \ngot: %s, \nexpected: %s",
				c.amount, c.currency, err, c.expectedError)
		}

		if !reflect.DeepEqual(actual, c.expected) {
			t.Errorf("LocalProvider.Convert(%f, %s) == \ngot: %v, \nexpected: %v",
				c.amount, c.currency, actual, c.expected)
		}
	}
}
