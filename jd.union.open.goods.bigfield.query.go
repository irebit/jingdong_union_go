package jingdong_union_go

import (
	"encoding/json"
	"errors"
	"log"
)

type JdUnionOpenGoodsBigfieldQueryTopLevel struct {
	JdUnionOpenGoodsBigfieldQueryResponse JdUnionOpenGoodsBigfieldQueryResponse `json:"jd_union_open_goods_bigfield_query_response"`
}

type JdUnionOpenGoodsBigfieldQueryResponse struct {
	Result string `json:"result"`
	Code   string `json:"code"`
}

type JdUnionOpenGoodsBigfieldQueryResult struct {
	Code      int64         `json:"code"`
	Data      []interface{} `json:"data"`
	Message   string        `json:"message"`
	RequestID string        `json:"requestId"`
}

func (app *App) JdUnionOpenGoodsBigfieldQuery(params map[string]interface{}) (result *JdUnionOpenGoodsBigfieldQueryResult, err error) {

	body, err := app.Request("jd.union.open.goods.bigfield.query", map[string]interface{}{"goodsReq": params})
	log.Println(string(body))
	resp := &JdUnionOpenGoodsBigfieldQueryTopLevel{}
	if err != nil {
		log.Println(string(body))
		return
	}

	if err = json.Unmarshal(body, resp); err != nil {
		return
	}
	if resp.JdUnionOpenGoodsBigfieldQueryResponse.Result != "" {
		result = &JdUnionOpenGoodsBigfieldQueryResult{}
		if err = json.Unmarshal([]byte(resp.JdUnionOpenGoodsBigfieldQueryResponse.Result), result); err != nil {
			return
		}
	} else {
		err = errors.New("result is null")
	}
	return
}
