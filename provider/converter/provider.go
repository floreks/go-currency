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
	"encoding/xml"
	"fmt"
)

// Supported providers
const (
	FixerIO = "fixerio"
	Local   = "local"
)

// ConvertedRates converted rates that will be presented to the user.
type ConvertedRates map[string]float64

// MarshalXML - marshals convertedRates map into XML
func (c ConvertedRates) MarshalXML(enc *xml.Encoder, startElem xml.StartElement) error {
	tokens := []xml.Token{startElem}

	for key, value := range c {
		t := xml.StartElement{Name: xml.Name{"", key}}
		tokens = append(tokens, t, xml.CharData(fmt.Sprintf("%v", value)), xml.EndElement{t.Name})
	}

	tokens = append(tokens, xml.EndElement{startElem.Name})

	for _, t := range tokens {
		err := enc.EncodeToken(t)
		if err != nil {
			return err
		}
	}

	err := enc.Flush()
	if err != nil {
		return err
	}

	return nil
}

// ConverterResponse is a structure returned by converter providers.
type ConverterResponse struct {
	// XMLName needed for correct xml response
	XMLName xml.Name `json:"-" xml:"ConverterResponse"`

	// Amount of money which is a base for exchange rate calculation
	Amount float64 `json:"amount" xml:"amount"`

	// Currency for which we should calculate exchange rates
	Currency string `json:"currency" xml:"currency"`

	// Converted rates based on given amount and currency
	Converted ConvertedRates `json:"converted" xml:"converted"`
}

// ConverterProvider is an abstract interface in order to allow providing multiple conversion
// providers.
type ConverterProvider interface {
	Convert(float64, string) (*ConverterResponse, error)
	Name() string
}

// GetProviders returns list of supported providers.
func GetProviders() []ConverterProvider {
	return []ConverterProvider{
		NewFixerIOProvider(),
		LocalProvider{},
	}
}
