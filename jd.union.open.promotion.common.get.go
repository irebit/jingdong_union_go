package jingdong_union_go

import (
	"encoding/json"
	"errors"
	"log"
)

type JdUnionOpenPromotionCommonGetTopLevel struct {
	JdUnionOpenPromotionCommonGetResponse JdUnionOpenPromotionCommonGetResponse `json:"jd_union_open_promotion_common_get_responce"`
}

type JdUnionOpenPromotionCommonGetResponse struct {
	Result string `json:"getResult"`
	Code   string `json:"code"`
}

type JdUnionOpenPromotionCommonGetResult struct {
	Code      int64              `json:"code"`
	Data      PromotionCommonURL `json:"data"`
	Message   string             `json:"message"`
	RequestID string             `json:"requestId"`
}

type PromotionCommonURL struct {
	ClickURL string `json:"clickURL"`
}

func (app *App) JdUnionOpenPromotionCommonGet(params map[string]interface{}) (result *JdUnionOpenPromotionCommonGetResult, err error) {

	body, err := app.Request("jd.union.open.promotion.common.get", map[string]interface{}{"promotionCodeReq": params})
	resp := &JdUnionOpenPromotionCommonGetTopLevel{}
	if err != nil {
		log.Println(string(body))
		return
	}
	log.Printf("%v", string(body))
	if err = json.Unmarshal(body, resp); err != nil {
		return
	}
	if resp.JdUnionOpenPromotionCommonGetResponse.Result != "" {
		result = &JdUnionOpenPromotionCommonGetResult{}
		if err = json.Unmarshal([]byte(resp.JdUnionOpenPromotionCommonGetResponse.Result), result); err != nil {
			return
		}
	} else {
		err = errors.New("result is null")
	}
	return
}
