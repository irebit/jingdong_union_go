package jingdong_union_go

import (
	"encoding/json"
	"errors"
	"log"
)

type JdUnionOpenCouponGiftGetTopLevel struct {
	JdUnionOpenCouponGiftGetResponse JdUnionOpenCouponGiftGetResponse `json:"jd_union_open_coupon_gift_get_responce"`
}

type JdUnionOpenCouponGiftGetResponse struct {
	Result string `json:"getResult"`
	Code   string `json:"code"`
}

type JdUnionOpenCouponGiftGetResult struct {
	Code      int64       `json:"code"`
	Data      *CouponGift `json:"data"`
	Message   string      `json:"message"`
	RequestID string      `json:"requestId"`
}

type CouponGift struct {
	GiftCouponKey string `json:"giftCouponKey"`
}

func (app *App) JdUnionOpenCouponGiftGet(params map[string]interface{}) (result *JdUnionOpenCouponGiftGetResult, err error) {

	body, err := app.Request("jd.union.open.coupon.gift.get", map[string]interface{}{"couponReq": params})
	log.Println(string(body))
	resp := &JdUnionOpenCouponGiftGetTopLevel{}
	if err != nil {
		log.Println(string(body))
		return
	}

	if err = json.Unmarshal(body, resp); err != nil {
		return
	}
	if resp.JdUnionOpenCouponGiftGetResponse.Result != "" {
		result = &JdUnionOpenCouponGiftGetResult{}
		if err = json.Unmarshal([]byte(resp.JdUnionOpenCouponGiftGetResponse.Result), result); err != nil {
			return
		}
	} else {
		err = errors.New("result is null")
	}
	return
}
