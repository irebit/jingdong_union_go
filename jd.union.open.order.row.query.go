package jingdong_union_go

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/shopspring/decimal"
)

type JdUnionOpenOrderRowQueryResponseTopLevel struct {
	JdUnionOpenOrderRowQueryResponse JdUnionOpenOrderRowQueryResponse `json:"jd_union_open_order_row_query_responce"`
}

type JdUnionOpenOrderRowQueryResponse struct {
	Result string `json:"queryResult"`
	Code   string `json:"code"`
}

type JdUnionOpenOrderRowQueryResult struct {
	Code      int64       `json:"code"`
	Data      []*OrderRow `json:"data"`
	HasMore   bool        `json:"hasMore"`
	Message   string      `json:"message"`
	RequestID string      `json:"requestId"`
}

type OrderRow struct {
	ActualCosPrice      decimal.Decimal    `json:"actualCosPrice"`
	ActualFee           decimal.Decimal    `json:"actualFee"`
	BalanceExt          string             `json:"balanceExt"`
	Cid1                int                `json:"cid1"`
	Cid2                int                `json:"cid2"`
	Cid3                int                `json:"cid3"`
	CommissionRate      decimal.Decimal    `json:"commissionRate"`
	CpActID             int                `json:"cpActId"`
	EstimateCosPrice    decimal.Decimal    `json:"estimateCosPrice"`
	EstimateFee         decimal.Decimal    `json:"estimateFee"`
	ExpressStatus       int                `json:"expressStatus"`
	Ext1                string             `json:"ext1"`
	FinalRate           decimal.Decimal    `json:"finalRate"`
	FinishTime          string             `json:"finishTime"`
	GiftCouponKey       string             `json:"giftCouponKey"`
	GiftCouponOcsAmount decimal.Decimal    `json:"giftCouponOcsAmount"`
	GoodsInfo           *OrderRowGoodsInfo `json:"goodsInfo"`
	ID                  string             `json:"id"`
	ModifyTime          string             `json:"modifyTime"`
	OrderEmt            int                `json:"orderEmt"`
	OrderID             int                `json:"orderId"`
	OrderTime           string             `json:"orderTime"`
	ParentID            int                `json:"parentId"`
	PayMonth            int                `json:"payMonth"`
	Pid                 string             `json:"pid"`
	Plus                int                `json:"plus"`
	PopId               int                `json:"popId"`
	PositionId          int                `json:"positionId"`
	Price               decimal.Decimal    `json:"price"`
	ProPriceAmount      decimal.Decimal    `json:"proPriceAmount"`
	Rid                 int                `json:"rid"`
	SiteId              int                `json:"siteId"`
	SkuFrozenNum        int                `json:"skuFrozenNum"`
	SkuId               int                `json:"skuId"`
	SkuName             string             `json:"skuName"`
	SkuNum              int                `json:"skuNum"`
	SkuReturnNum        int                `json:"skuReturnNum"`
	SubSideRate         decimal.Decimal    `json:"subSideRate"`
	SubUnionID          string             `json:"subUnionId"`
	SubSidyRate         decimal.Decimal    `json:"subsidyRate"`
	TraceType           int                `json:"traceType"`
	UnionAlias          string             `json:"unionAlias"`
	UnionId             int                `json:"unionId"`
	UnionRole           int                `json:"unionRole"`
	UnionTag            string             `json:"unionTag"`
	ValidCode           int                `json:"validCode"`
}

type OrderRowGoodsInfo struct {
	ImageUrl  string `json:"imageUrl"`
	MainSkuId int    `json:"mainSkuId"`
	Owner     string `json:"owner"`
	ProductId int    `json:"productId"`
	ShopId    int    `json:"shopId"`
	ShopName  string `json:"shopName"`
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
