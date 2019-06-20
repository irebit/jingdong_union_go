package jingdong_union_go

import (
	"encoding/json"
	"errors"
	"log"
)

type JdUnionOpenPromotionBysubunionidGetTopLevel struct {
	JdUnionOpenPromotionBysubunionidGetResponse JdUnionOpenPromotionBysubunionidGetResponse `json:"jd_union_open_promotion_bysubunionid_get_response"`
}

type JdUnionOpenPromotionBysubunionidGetResponse struct {
	Result string `json:"result"`
	Code   string `json:"code"`
}

type JdUnionOpenPromotionBysubunionidGetResult struct {
	Code      int64        `json:"code"`
	Data      PromotionUrl `json:"data"`
	Message   string       `json:"message"`
	RequestID string       `json:"requestId"`
}

type PromotionUrl struct {
	ClickURL string `json:"clickURL"`
	ShortURL string `json:"shortURL"`
}

func (app *App) JdUnionOpenPromotionBysubunionidGet(params map[string]interface{}) (result *JdUnionOpenPromotionBysubunionidGetResult, err error) {

	body, err := app.Request("jd.union.open.promotion.bysubunionid.get", map[string]interface{}{"promotionCodeReq": params})
	resp := &JdUnionOpenPromotionBysubunionidGetTopLevel{}
	if err != nil {
		log.Println(string(body))
		return
	}

	if err = json.Unmarshal(body, resp); err != nil {
		return
	}
	if resp.JdUnionOpenPromotionBysubunionidGetResponse.Result != "" {
		result = &JdUnionOpenPromotionBysubunionidGetResult{}
		if err = json.Unmarshal([]byte(resp.JdUnionOpenPromotionBysubunionidGetResponse.Result), result); err != nil {
			return
		}
	} else {
		err = errors.New("result is null")
	}
	return
}
