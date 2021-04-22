package jingdong_union_go

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/shopspring/decimal"
)

type JdUnionOpenStatisticsGiftcouponQueryTopLevel struct {
	JdUnionOpenStatisticsGiftcouponQueryResponse JdUnionOpenStatisticsGiftcouponQueryResponse `json:"jd_union_open_statistics_giftcoupon_query_response"`
}

type JdUnionOpenStatisticsGiftcouponQueryResponse struct {
	Result string `json:"result"`
	Code   string `json:"code"`
}

type JdUnionOpenStatisticsGiftcouponQueryResult struct {
	Code      int64               `json:"code"`
	Message   string              `json:"message"`
	Data      []*CouponEffectData `json:"data"`
	RequestID string              `json:"requestId"`
}

type CouponEffectDataResp struct {
	GiftCouponEffectDataResp *CouponEffectData `json:"giftCouponEffectDataResp"`
}

type CouponEffectData struct {
	Amount           decimal.Decimal `json:"amount"`
	ContentMatch     int             `json:"contentMatch"`
	GiftCouponKey    string          `json:"giftCouponKey"`
	CostNum          int             `json:"costNum"`
	ExpireType       int             `json:"expireType"`
	UseStartTime     string          `json:"useStartTime"`
	YgCommission     decimal.Decimal `json:"ygCommission"`
	ReceiveStartTime string          `json:"receiveStartTime"`
	CostAmount       decimal.Decimal `json:"costAmount"`
	ShowStatus       int             `json:"showStatus"`
	Type             int             `json:"type"`
	Denomination     decimal.Decimal `json:"denomination"`
	UseEndTime       string          `json:"useEndTime"`
	ReceiveEndTime   string          `json:"receiveEndTime"`
	EffectiveDays    int             `json:"effectiveDays"`
	Share            int             `json:"share"`
	ReceiveNum       int             `json:"receiveNum"`
	SkuIdList        []int           `json:"skuIdList"`
	Status           int             `json:"status"`
	CouponTitle      string          `json:"couponTitle"`
}

func (app *App) JdUnionOpenStatisticsGiftcouponQuery(params map[string]interface{}) (result *JdUnionOpenStatisticsGiftcouponQueryResult, err error) {

	body, err := app.Request("jd.union.open.statistics.giftcoupon.query", map[string]interface{}{"effectDataReq": params})
	log.Println(string(body))
	resp := &JdUnionOpenStatisticsGiftcouponQueryTopLevel{}
	if err != nil {
		log.Println(string(body))
		return
	}

	if err = json.Unmarshal(body, resp); err != nil {
		return
	}
	if resp.JdUnionOpenStatisticsGiftcouponQueryResponse.Result != "" {
		result = &JdUnionOpenStatisticsGiftcouponQueryResult{}
		if err = json.Unmarshal([]byte(resp.JdUnionOpenStatisticsGiftcouponQueryResponse.Result), result); err != nil {
			return
		}
	} else {
		err = errors.New("result is null")
	}
	return
}
