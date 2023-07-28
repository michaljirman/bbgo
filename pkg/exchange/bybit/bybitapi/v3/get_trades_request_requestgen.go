// Code generated by "requestgen -method GET -responseType .APIResponse -responseDataField Result -url /spot/v3/private/my-trades -type GetTradesRequest -responseDataType .TradesResponse"; DO NOT EDIT.

package v3

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/c9s/bbgo/pkg/exchange/bybit/bybitapi"
	"net/url"
	"reflect"
	"regexp"
	"strconv"
	"time"
)

func (g *GetTradesRequest) Symbol(symbol string) *GetTradesRequest {
	g.symbol = &symbol
	return g
}

func (g *GetTradesRequest) OrderId(orderId string) *GetTradesRequest {
	g.orderId = &orderId
	return g
}

func (g *GetTradesRequest) Limit(limit uint64) *GetTradesRequest {
	g.limit = &limit
	return g
}

func (g *GetTradesRequest) StartTime(startTime time.Time) *GetTradesRequest {
	g.startTime = &startTime
	return g
}

func (g *GetTradesRequest) EndTime(endTime time.Time) *GetTradesRequest {
	g.endTime = &endTime
	return g
}

func (g *GetTradesRequest) FromTradeId(fromTradeId string) *GetTradesRequest {
	g.fromTradeId = &fromTradeId
	return g
}

func (g *GetTradesRequest) ToTradeId(toTradeId string) *GetTradesRequest {
	g.toTradeId = &toTradeId
	return g
}

// GetQueryParameters builds and checks the query parameters and returns url.Values
func (g *GetTradesRequest) GetQueryParameters() (url.Values, error) {
	var params = map[string]interface{}{}
	// check symbol field -> json key symbol
	if g.symbol != nil {
		symbol := *g.symbol

		// assign parameter of symbol
		params["symbol"] = symbol
	} else {
	}
	// check orderId field -> json key orderId
	if g.orderId != nil {
		orderId := *g.orderId

		// assign parameter of orderId
		params["orderId"] = orderId
	} else {
	}
	// check limit field -> json key limit
	if g.limit != nil {
		limit := *g.limit

		// assign parameter of limit
		params["limit"] = limit
	} else {
	}
	// check startTime field -> json key startTime
	if g.startTime != nil {
		startTime := *g.startTime

		// assign parameter of startTime
		// convert time.Time to milliseconds time stamp
		params["startTime"] = strconv.FormatInt(startTime.UnixNano()/int64(time.Millisecond), 10)
	} else {
	}
	// check endTime field -> json key endTime
	if g.endTime != nil {
		endTime := *g.endTime

		// assign parameter of endTime
		// convert time.Time to milliseconds time stamp
		params["endTime"] = strconv.FormatInt(endTime.UnixNano()/int64(time.Millisecond), 10)
	} else {
	}
	// check fromTradeId field -> json key fromTradeId
	if g.fromTradeId != nil {
		fromTradeId := *g.fromTradeId

		// assign parameter of fromTradeId
		params["fromTradeId"] = fromTradeId
	} else {
	}
	// check toTradeId field -> json key toTradeId
	if g.toTradeId != nil {
		toTradeId := *g.toTradeId

		// assign parameter of toTradeId
		params["toTradeId"] = toTradeId
	} else {
	}

	query := url.Values{}
	for _k, _v := range params {
		query.Add(_k, fmt.Sprintf("%v", _v))
	}

	return query, nil
}

// GetParameters builds and checks the parameters and return the result in a map object
func (g *GetTradesRequest) GetParameters() (map[string]interface{}, error) {
	var params = map[string]interface{}{}

	return params, nil
}

// GetParametersQuery converts the parameters from GetParameters into the url.Values format
func (g *GetTradesRequest) GetParametersQuery() (url.Values, error) {
	query := url.Values{}

	params, err := g.GetParameters()
	if err != nil {
		return query, err
	}

	for _k, _v := range params {
		if g.isVarSlice(_v) {
			g.iterateSlice(_v, func(it interface{}) {
				query.Add(_k+"[]", fmt.Sprintf("%v", it))
			})
		} else {
			query.Add(_k, fmt.Sprintf("%v", _v))
		}
	}

	return query, nil
}

// GetParametersJSON converts the parameters from GetParameters into the JSON format
func (g *GetTradesRequest) GetParametersJSON() ([]byte, error) {
	params, err := g.GetParameters()
	if err != nil {
		return nil, err
	}

	return json.Marshal(params)
}

// GetSlugParameters builds and checks the slug parameters and return the result in a map object
func (g *GetTradesRequest) GetSlugParameters() (map[string]interface{}, error) {
	var params = map[string]interface{}{}

	return params, nil
}

func (g *GetTradesRequest) applySlugsToUrl(url string, slugs map[string]string) string {
	for _k, _v := range slugs {
		needleRE := regexp.MustCompile(":" + _k + "\\b")
		url = needleRE.ReplaceAllString(url, _v)
	}

	return url
}

func (g *GetTradesRequest) iterateSlice(slice interface{}, _f func(it interface{})) {
	sliceValue := reflect.ValueOf(slice)
	for _i := 0; _i < sliceValue.Len(); _i++ {
		it := sliceValue.Index(_i).Interface()
		_f(it)
	}
}

func (g *GetTradesRequest) isVarSlice(_v interface{}) bool {
	rt := reflect.TypeOf(_v)
	switch rt.Kind() {
	case reflect.Slice:
		return true
	}
	return false
}

func (g *GetTradesRequest) GetSlugsMap() (map[string]string, error) {
	slugs := map[string]string{}
	params, err := g.GetSlugParameters()
	if err != nil {
		return slugs, nil
	}

	for _k, _v := range params {
		slugs[_k] = fmt.Sprintf("%v", _v)
	}

	return slugs, nil
}

func (g *GetTradesRequest) Do(ctx context.Context) (*TradesResponse, error) {

	// no body params
	var params interface{}
	query, err := g.GetQueryParameters()
	if err != nil {
		return nil, err
	}

	apiURL := "/spot/v3/private/my-trades"

	req, err := g.client.NewAuthenticatedRequest(ctx, "GET", apiURL, query, params)
	if err != nil {
		return nil, err
	}

	response, err := g.client.SendRequest(req)
	if err != nil {
		return nil, err
	}

	var apiResponse bybitapi.APIResponse
	if err := response.DecodeJSON(&apiResponse); err != nil {
		return nil, err
	}
	var data TradesResponse
	if err := json.Unmarshal(apiResponse.Result, &data); err != nil {
		return nil, err
	}
	return &data, nil
}
