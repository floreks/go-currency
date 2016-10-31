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

type convertedRates map[string]float64

// MarshalXML - marshals convertedRates map into XML
func (c convertedRates) MarshalXML(enc *xml.Encoder, startElem xml.StartElement) error {
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

type ConverterResponse struct {
	XMLName xml.Name `json:"-" xml:"ConverterResponse"`

	Amount float64 `json:"amount" xml:"amount"`

	Currency string `json:"currency" xml:"currency"`

	Converted convertedRates `json:"converted" xml:"converted"`
}

type ConverterProvider interface {
	Convert(float64, string) (*ConverterResponse, error)
	Name() string
}

func GetProviders() []ConverterProvider {
	return []ConverterProvider{
		NewFixerIOProvider(),
		LocalProvider{},
	}
}
