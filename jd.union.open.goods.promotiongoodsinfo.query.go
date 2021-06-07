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
	SkuID             int64   `json:"skuId"`             //	商品ID
	GoodsName         string  `json:"goodsName"`         //商品名称
	UnitPrice         float64 `json:"unitPrice"`         //商品单价即京东价
	MaterialURL       string  `json:"materialUrl"`       //商品落地页
	WlUnitPrice       float64 `json:"wlUnitPrice"`       //商品无线京东价（单价为-1表示未查询到该商品单价）
	CommisionRatioWl  float64 `json:"commisionRatioWl"`  //无线佣金比例
	CommisionRatioPC  float64 `json:"commisionRatioPc"`  //PC佣金比例
	ImgURL            string  `json:"imgUrl"`            //图片地址
	Cid               int64   `json:"cid"`               //一级类目ID
	CidName           string  `json:"cidName"`           //一级类目名称
	Cid2              int64   `json:"cid2"`              //二级类目ID
	Cid2Name          string  `json:"cid2Name"`          //二级类目名称
	Cid3              int64   `json:"cid3"`              //三级类目ID
	Cid3Name          string  `json:"cid3Name"`          //三级类目名称
	IsFreeFreightRisk int64   `json:"isFreeFreightRisk"` //是否支持运费险(1:是,0:否)
	IsSeckill         int64   `json:"isSeckill"`         //是否秒杀(1:是,0:否)
	IsFreeShipping    int64   `json:"isFreeShipping"`    //是否包邮(1:是,0:否,2:自营商品遵从主站包邮规则)
	InOrderCount      int64   `json:"inOrderCount"`      //30天引单数量
	Vid               int64   `json:"vid"`               //商家ID
	ShopID            int64   `json:"shopId"`            //店铺ID
	IsJdSale          int64   `json:"isJdSale"`          //是否自营(1:是,0:否)
	StartDate         int64   `json:"startDate"`         //推广开始日期（时间戳，毫秒）
	EndDate           int64   `json:"endDate"`           //推广结束日期(时间戳，毫秒)
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
