package jingdong_union_go

import (
	"encoding/json"
	"errors"
	"log"
)

type JdUnionOpenCategoryGoodsGetTopLevel struct {
	JdUnionOpenCategoryGoodsGetResponse JdUnionOpenCategoryGoodsGetResponse `json:"jd_union_open_category_goods_get_responce"`
}

type JdUnionOpenCategoryGoodsGetResponse struct {
	Result string `json:"queryResult"`
	Code   string `json:"code"`
}

type JdUnionOpenCategoryGoodsGetResult struct {
	Code      int64  `json:"code"`
	Data      []Cate `json:"data"`
	Message   string `json:"message"`
	RequestID string `json:"requestId"`
}

type Cate struct {
	Grade    int64  `json:"grade"`
	Name     string `json:"name"`
	ID       int64  `json:"id"`
	ParentID int64  `json:"parentId"`
}

func (app *App) JdUnionOpenCategoryGoodsGet(params map[string]interface{}) (result *JdUnionOpenCategoryGoodsGetResult, err error) {

	body, err := app.Request("jd.union.open.category.goods.get", map[string]interface{}{"req": params})
	resp := &JdUnionOpenCategoryGoodsGetTopLevel{}
	if err != nil {
		log.Println(string(body))
		return
	}

	if err = json.Unmarshal(body, resp); err != nil {
		return
	}
	if resp.JdUnionOpenCategoryGoodsGetResponse.Result != "" {
		result = &JdUnionOpenCategoryGoodsGetResult{}
		if err = json.Unmarshal([]byte(resp.JdUnionOpenCategoryGoodsGetResponse.Result), result); err != nil {
			return
		}
	} else {
		err = errors.New("result is null")
	}
	return
}
