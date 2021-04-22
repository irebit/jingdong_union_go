package jingdong_union_go

import (
	"encoding/json"
	"errors"
	"log"
)

type JdUnionOpenCouponGiftStopTopLevel struct {
	JdUnionOpenCouponGiftStopResponse JdUnionOpenCouponGiftStopResponse `json:"jd_union_open_coupon_gift_stop_responce"`
}

type JdUnionOpenCouponGiftStopResponse struct {
	Result string `json:"queryResult"`
	Code   string `json:"code"`
}

type JdUnionOpenCouponGiftStopResult struct {
	Code      int64  `json:"code"`
	Message   string `json:"message"`
	RequestID string `json:"requestId"`
}

func (app *App) JdUnionOpenCouponGiftStop(params map[string]interface{}) (result *JdUnionOpenCouponGiftStopResult, err error) {

	body, err := app.Request("jd.union.open.coupon.gift.stop", map[string]interface{}{"couponReq": params})
	log.Println(string(body))
	resp := &JdUnionOpenCouponGiftStopTopLevel{}
	if err != nil {
		log.Println(string(body))
		return
	}

	if err = json.Unmarshal(body, resp); err != nil {
		return
	}
	if resp.JdUnionOpenCouponGiftStopResponse.Result != "" {
		result = &JdUnionOpenCouponGiftStopResult{}
		if err = json.Unmarshal([]byte(resp.JdUnionOpenCouponGiftStopResponse.Result), result); err != nil {
			return
		}
	} else {
		err = errors.New("result is null")
	}
	return
}
