package jingdong_union_go

import (
	"encoding/json"
	"errors"
	"log"
)

type JdUnionOpenGoodsJingfenQueryTopLevel struct {
	JdUnionOpenGoodsJingfenQueryResponse JdUnionOpenGoodsJingfenQueryResponse `json:"jd_union_open_goods_jingfen_query_responce"`
}

type JdUnionOpenGoodsJingfenQueryResponse struct {
	Result string `json:"queryResult"`
	Code   string `json:"code"`
}

func (app *App) JdUnionOpenGoodsJingfenQuery(params map[string]interface{}) (result *JdUnionOpenGoodsQueryResult, err error) {

	body, err := app.Request("jd.union.open.goods.jingfen.query", map[string]interface{}{"goodsReq": params})
	resp := &JdUnionOpenGoodsJingfenQueryTopLevel{}
	if err != nil {
		log.Println(string(body))
		return
	}
	if err = json.Unmarshal(body, resp); err != nil {
		return
	}
	if resp.JdUnionOpenGoodsJingfenQueryResponse.Result != "" {
		result = &JdUnionOpenGoodsQueryResult{}
		if err = json.Unmarshal([]byte(resp.JdUnionOpenGoodsJingfenQueryResponse.Result), result); err != nil {
			return
		}
	} else {
		err = errors.New("result is null")
	}
	return
}
