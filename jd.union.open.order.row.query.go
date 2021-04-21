package jingdong_union_go

import (
	"encoding/json"
	"errors"
	"log"
)

type JdUnionOpenOrderRowQueryResponseTopLevel struct {
	JdUnionOpenOrderRowQueryResponse JdUnionOpenOrderRowQueryResponse `json:"jd_union_open_order_row_query_response"`
}

type JdUnionOpenOrderRowQueryResponse struct {
	Result string `json:"result"`
	Code   string `json:"code"`
}

type JdUnionOpenOrderRowQueryResult struct {
	Code      int64         `json:"code"`
	Data      []interface{} `json:"data"`
	HasMore   bool          `json:"hasMore"`
	Message   string        `json:"message"`
	RequestID string        `json:"requestId"`
}

func (app *App) JdUnionOpenOrderRowQuery(params map[string]interface{}) (result *JdUnionOpenOrderRowQueryResult, err error) {

	body, err := app.Request("jd.union.open.order.row.query", map[string]interface{}{"orderReq": params})

	resp := &JdUnionOpenOrderRowQueryResponseTopLevel{}
	if err != nil {
		log.Println(string(body))
		return
	}

	if err = json.Unmarshal(body, resp); err != nil {
		return
	}
	if resp.JdUnionOpenOrderRowQueryResponse.Result != "" {
		result = &JdUnionOpenOrderRowQueryResult{}
		if err = json.Unmarshal([]byte(resp.JdUnionOpenOrderRowQueryResponse.Result), result); err != nil {
			return
		}
	} else {
		err = errors.New("result is null")
	}
	return
}
