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
	"reflect"
	"testing"
)

func TestConvertFixerIO(t *testing.T) {
	provider := new(FixerIOProvider)
	cases := []struct {
		rates    FixerAPIRates
		amount   float64
		expected ConvertedRates
	}{
		{
			FixerAPIRates{
				"USD": 1.234,
				"PLN": 1.01,
				"EUR": 0.05,
				"SEK": 0.00055,
			},
			100,
			ConvertedRates{
				"USD": 123.4,
				"PLN": 101,
				"EUR": 5,
				"SEK": 0.06,
			},
		},
	}

	for _, c := range cases {
		actual := provider.convert(c.rates, c.amount)

		if !reflect.DeepEqual(actual, c.expected) {
			t.Errorf("FixerIOProvider.convert(%v, %v) == \ngot: %v, \nexpected %v",
				c.rates, c.amount, actual, c.expected)
		}
	}
}

func TestNewFixerIOProvider(t *testing.T) {
	cases := []struct {
		expected FixerIOProvider
	}{
		{FixerIOProvider{url: FixerIOApiUrl}},
	}

	for _, c := range cases {
		actual := NewFixerIOProvider()

		if !reflect.DeepEqual(actual, c.expected) {
			t.Errorf("NewFixerIOProvider() == \ngot: %v, \nexpected %v", actual, c.expected)
		}
	}
}
