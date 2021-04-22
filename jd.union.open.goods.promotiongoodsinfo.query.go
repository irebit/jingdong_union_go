package jingdong_union_go

import (
	"encoding/json"
	"errors"
	"log"
)

type JdUnionOpenGoodsPromotiongoodsinfoQueryTopLevel struct {
	JdUnionOpenGoodsPromotiongoodsinfoQueryResponse JdUnionOpenGoodsPromotiongoodsinfoQueryResponse `json:"jd_union_open_goods_promotiongoodsinfo_query_responce"`
}

type JdUnionOpenGoodsPromotiongoodsinfoQueryResponse struct {
	Result string `json:"queryResult"`
	Code   string `json:"code"`
}

type JdUnionOpenGoodsPromotiongoodsinfoQueryResult struct {
	Code      int64       `json:"code"`
	Data      []GoodsInfo `json:"data"`
	Message   string      `json:"message"`
	RequestID string      `json:"requestId"`
}

type GoodsInfo struct {
	UnitPrice         float64 `json:"unitPrice"`
	MaterialURL       string  `json:"materialUrl"`
	EndDate           int64   `json:"endDate"`
	IsFreeFreightRisk int64   `json:"isFreeFreightRisk"`
	IsFreeShipping    int64   `json:"isFreeShipping"`
	CommisionRatioWl  float64 `json:"commisionRatioWl"`
	CommisionRatioPC  float64 `json:"commisionRatioPc"`
	ImgURL            string  `json:"imgUrl"`
	Vid               int64   `json:"vid"`
	CidName           string  `json:"cidName"`
	WlUnitPrice       float64 `json:"wlUnitPrice"`
	Cid2Name          string  `json:"cid2Name"`
	IsSeckill         int64   `json:"isSeckill"`
	Cid2              int64   `json:"cid2"`
	Cid3Name          string  `json:"cid3Name"`
	InOrderCount      int64   `json:"inOrderCount"`
	Cid3              int64   `json:"cid3"`
	ShopID            int64   `json:"shopId"`
	IsJdSale          int64   `json:"isJdSale"`
	GoodsName         string  `json:"goodsName"`
	SkuID             int64   `json:"skuId"`
	StartDate         int64   `json:"startDate"`
	Cid               int64   `json:"cid"`
}

func (app *App) JdUnionOpenGoodsPromotiongoodsinfoQuery(params map[string]interface{}) (result *JdUnionOpenGoodsPromotiongoodsinfoQueryResult, err error) {

	body, err := app.Request("jd.union.open.goods.promotiongoodsinfo.query", params)
	resp := &JdUnionOpenGoodsPromotiongoodsinfoQueryTopLevel{}
	if err != nil {
		log.Println(string(body))
		return
	}

	if err = json.Unmarshal(body, resp); err != nil {
		return
	}
	if resp.JdUnionOpenGoodsPromotiongoodsinfoQueryResponse.Result != "" {
		result = &JdUnionOpenGoodsPromotiongoodsinfoQueryResult{}
		if err = json.Unmarshal([]byte(resp.JdUnionOpenGoodsPromotiongoodsinfoQueryResponse.Result), result); err != nil {
			return
		}
	} else {
		err = errors.New("result is null")
	}
	return
}
