package jingdong_union_go

import (
	"encoding/json"
	"errors"
	"log"
)

type JdUnionOpenGoodsQueryResponseTopLevel struct {
	JdUnionOpenGoodsQueryResponse JdUnionOpenGoodsQueryResponse `json:"jd_union_open_goods_query_response"`
}

type JdUnionOpenGoodsQueryResponse struct {
	Result string `json:"result"`
	Code   string `json:"code"`
}

type JdUnionOpenGoodsQueryResult struct {
	Code       int64   `json:"code"`
	Data       []Goods `json:"data"`
	Message    string  `json:"message"`
	RequestID  string  `json:"requestId"`
	TotalCount int64   `json:"totalCount"`
}

type Goods struct {
	BrandCode          string         `json:"brandCode"`
	BrandName          string         `json:"brandName"`
	CategoryInfo       CategoryInfo   `json:"categoryInfo"`
	Comments           int64          `json:"comments"`
	CommissionInfo     CommissionInfo `json:"commissionInfo"`
	CouponInfo         CouponInfo     `json:"couponInfo"`
	GoodCommentsShare  float64        `json:"goodCommentsShare"`
	ImageInfo          ImageInfo      `json:"imageInfo"`
	InOrderComm30Days  float64        `json:"inOrderComm30Days"`
	InOrderCount30Days int64          `json:"inOrderCount30Days"`
	IsHot              int64          `json:"isHot"`
	IsJdSale           int64          `json:"isJdSale"`
	MaterialURL        string         `json:"materialUrl"`
	Owner              string         `json:"owner"`
	PinGouInfo         PinGouInfo     `json:"pinGouInfo"`
	PingGouInfo        PingGouInfo    `json:"pingGouInfo"`
	PriceInfo          PriceInfo      `json:"priceInfo"`
	ShopInfo           ShopInfo       `json:"shopInfo"`
	SkuID              int64          `json:"skuId"`
	SkuName            string         `json:"skuName"`
	Spuid              int64          `json:"spuid"`
}

type CategoryInfo struct {
	Cid1     int64  `json:"cid1"`
	Cid1Name string `json:"cid1Name"`
	Cid2     int64  `json:"cid2"`
	Cid2Name string `json:"cid2Name"`
	Cid3     int64  `json:"cid3"`
	Cid3Name string `json:"cid3Name"`
}

type CommissionInfo struct {
	Commission      float64 `json:"commission"`
	CommissionShare float64 `json:"commissionShare"`
}

type CouponInfo struct {
	CouponList []Coupon `json:"couponList"`
}

type Coupon struct {
	BindType     int     `json:"bindType"`
	Discount     float64 `json:"discount"`
	Link         string  `json:"link"`
	PlatformType int     `json:"platformType"`
	Quota        float64 `json:"quota"`
	GetStartTime int     `json:"getStartTime"`
	GetEndTime   int     `json:"getEndTime"`
	UseStartTime int     `json:"useStartTime"`
	UseEndTime   int     `json:"useEndTime"`
}

type ImageInfo struct {
	ImageList []ImageList `json:"imageList"`
}

type ImageList struct {
	URL string `json:"url"`
}

type PinGouInfo struct {
	PingouEndTime   int64   `json:"pingouEndTime"`
	PingouPrice     float64 `json:"pingouPrice"`
	PingouStartTime int64   `json:"pingouStartTime"`
	PingouTmCount   int64   `json:"pingouTmCount"`
	PingouURL       string  `json:"pingouUrl"`
}

type PingGouInfo struct {
	PingouPrice   float64 `json:"pingouPrice"`
	PingouTmCount int64   `json:"pingouTmCount"`
	PingouURL     string  `json:"pingouUrl"`
}

type PriceInfo struct {
	Price float64 `json:"price"`
}

type ShopInfo struct {
	ShopID   int64  `json:"shopId"`
	ShopName string `json:"shopName"`
}

func (app *App) JdUnionOpenGoodsQuery(params map[string]interface{}) (result *JdUnionOpenGoodsQueryResult, err error) {

	body, err := app.Request("jd.union.open.goods.query", map[string]interface{}{"goodsReqDTO": params})
	resp := &JdUnionOpenGoodsQueryResponseTopLevel{}
	if err != nil {
		log.Println(string(body))
		return
	}

	if err = json.Unmarshal(body, resp); err != nil {
		return
	}
	if resp.JdUnionOpenGoodsQueryResponse.Result != "" {
		result = &JdUnionOpenGoodsQueryResult{}
		if err = json.Unmarshal([]byte(resp.JdUnionOpenGoodsQueryResponse.Result), result); err != nil {
			return
		}
	} else {
		err = errors.New("result is null")
	}
	return
}
