package converter

import (
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"

	"fmt"
	"github.com/emicklei/go-restful"
	"github.com/floreks/go-currency/provider/converter"
)

type ConverterQuery struct {
	// Amount represents amount of money that should be converted to another currency.
	Amount float64

	// Currency represents target currency to which conversion should be applied.
	Currency string

	// Provider is a optional parameter that represents provider that should be used for conversion.
	Provider converter.ConverterProvider
}

// ConverterService converts given amount of money in given currency to currencies supported by
// selected provider. By default FixerIO provider is used for conversion.
type ConverterService struct {
	providers []converter.ConverterProvider
}

func (c ConverterService) getProvider(providerName string) converter.ConverterProvider {
	for _, provider := range c.providers {
		if strings.Compare(provider.Name(), providerName) == 0 {
			return provider
		}
	}

	return nil
}

func (c ConverterService) getDefaultProvider() converter.ConverterProvider {
	return c.getProvider(converter.FixerIO)
}

func (c ConverterService) Handler() *restful.WebService {
	ws := new(restful.WebService)
	ws.
		Path("/convert").
		Consumes(restful.MIME_JSON, restful.MIME_XML).
		Produces(restful.MIME_JSON, restful.MIME_XML)

	ws.Route(ws.GET("/").To(c.convert).
		Doc("Converts currency from one to another").
		Writes(converter.ConverterResponse{}))

	return ws
}

func (c ConverterService) convert(request *restful.Request, response *restful.Response) {
	converterQuery, err := c.parseConverterParameters(request)
	if err != nil {
		response.WriteError(http.StatusBadRequest, err)
		return
	}

	converterResponse, err :=
		converterQuery.Provider.Convert(converterQuery.Amount, converterQuery.Currency)
	if err != nil {
		response.WriteError(http.StatusInternalServerError, err)
		return
	}

	response.WriteHeaderAndEntity(http.StatusOK, converterResponse)
}

func (c ConverterService) parseConverterParameters(
	request *restful.Request) (*ConverterQuery, error) {

	amountParam := request.QueryParameter("amount")
	amount, err := strconv.ParseFloat(amountParam, 64)
	if err != nil {
		log.Printf("Provided amount is invalid or empty: '%s'.", amountParam)
		return nil, errors.New(fmt.Sprintf("Provided amount is invalid or empty: '%s'.", amountParam))
	}

	currency := request.QueryParameter("currency")
	if currency == "" {
		log.Println("Currency parameter can not be empty.")
		return nil, errors.New("Currency parameter can not be empty.")
	}

	var provider converter.ConverterProvider
	providerName := request.QueryParameter("provider")
	provider = c.getProvider(providerName)
	if provider == nil {
		provider = c.getDefaultProvider()
		log.Printf("Provider is either empty or invalid. Falling back to default provider: %s",
			provider.Name())
	}

	return &ConverterQuery{Amount: amount, Currency: currency, Provider: provider}, nil
}

func NewConverterService() ConverterService {
	return ConverterService{providers: converter.GetProviders()}
}
