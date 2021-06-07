package jingdong_union_go

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/shopspring/decimal"
)

type JdUnionOpenGoodsQueryResponseTopLevel struct {
	JdUnionOpenGoodsQueryResponse JdUnionOpenGoodsQueryResponse `json:"jd_union_open_goods_query_responce"`
}

type JdUnionOpenGoodsQueryResponse struct {
	Result string `json:"queryResult"`
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
	CategoryInfo           CategoryInfo         `json:"categoryInfo"`           //类目信息
	Comments               int64                `json:"comments"`               //评论数
	CommissionInfo         CommissionInfo       `json:"commissionInfo"`         //佣金信息
	CouponInfo             CouponInfo           `json:"couponInfo"`             //优惠券信息
	GoodCommentsShare      decimal.Decimal      `json:"goodCommentsShare"`      //商品好评率
	ImageInfo              ImageInfo            `json:"imageInfo"`              //图片信息
	InOrderCount30Days     int64                `json:"inOrderCount30Days"`     //30天引单数量
	MaterialURL            string               `json:"materialUrl"`            //商品落地页
	PriceInfo              PriceInfo            `json:"priceInfo"`              //价格信息
	ShopInfo               ShopInfo             `json:"shopInfo"`               //店铺信息
	SkuID                  int64                `json:"skuId"`                  //商品ID
	SkuName                string               `json:"skuName"`                //商品名称
	IsHot                  int64                `json:"isHot"`                  //已废弃，请勿使用
	Spuid                  int64                `json:"spuid"`                  //spuid，其值为同款商品的主skuid
	BrandCode              string               `json:"brandCode"`              //品牌code
	BrandName              string               `json:"brandName"`              //品牌名
	Owner                  string               `json:"owner"`                  //g=自营，p=pop
	PinGouInfo             PinGouInfo           `json:"pinGouInfo"`             //拼购信息
	ResourceInfo           ResourceInfo         `json:"resourceInfo"`           //资源信息
	InOrderCount30DaysSku  int64                `json:"inOrderCount30DaysSku"`  //30天引单数量(sku维度)
	SeckillInfo            SeckillInfo          `json:"seckillInfo"`            //秒杀信息
	JxFlags                []int64              `json:"jxFlags"`                //京喜商品类型，1京喜、2京喜工厂直供、3京喜优选（包含3时可在京东APP购买）
	VideoInfo              VideoInfo            `json:"videoInfo"`              //视频信息
	DocumentInfo           DocumentInfo         `json:"documentInfo"`           //段子信息
	BookInfo               BookInfo             `json:"bookInfo"`               //图书信息
	ForbidTypes            []int64              `json:"forbidTypes"`            // 0普通商品，10微信京东购物小程序禁售，11微信京喜小程序禁售
	DeliveryType           int64                `json:"deliveryType"`           //京东配送 1：是，0：不是
	SkuLabelInfo           SkuLabelInfo         `json:"skuLabelInfo"`           //商品标签
	PromotionLabelInfoList []PromotionLabelInfo `json:"promotionLabelInfoList"` //商品促销标签集
	SecondPriceInfoList    []SecondPriceInfo    `json:"secondPriceInfoList"`    //双价格
	// IsJdSale               int64                  `json:"isJdSale"`
	// PingGouInfo            PingGouInfo            `json:"pingGouInfo"`
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
	Commission          decimal.Decimal `json:"commission"`          //佣金
	CommissionShare     decimal.Decimal `json:"commissionShare"`     //佣金比例
	CouponCommission    decimal.Decimal `json:"couponCommission"`    //券后佣金，（促销价-优惠券面额）*佣金比例
	PlusCommissionShare decimal.Decimal `json:"plusCommissionShare"` //plus佣金比例，plus用户购买推广者能获取到的佣金比例
	IsLock              int64           `json:"isLock"`              //是否锁定佣金比例：1是，0否
	StartTime           int64           `json:"startTime"`           //计划开始时间（时间戳，毫秒）
	EndTime             int64           `json:"endTime"`             //计划结束时间（时间戳，毫秒）
}

type CouponInfo struct {
	CouponList []Coupon `json:"couponList"`
}

type Coupon struct {
	BindType     int64           `json:"bindType"`     //券种类 (优惠券种类：0 - 全品类，1 - 限品类（自营商品），2 - 限店铺，3 - 店铺限商品券)
	Discount     decimal.Decimal `json:"discount"`     //券面额
	Link         string          `json:"link"`         //券链接
	PlatformType int64           `json:"platformType"` //券使用平台 (平台类型：0 - 全平台券，1 - 限平台券)
	Quota        decimal.Decimal `json:"quota"`        //券消费限额
	GetStartTime int64           `json:"getStartTime"` //领取开始时间(时间戳，毫秒)
	GetEndTime   int64           `json:"getEndTime"`   //券领取结束时间(时间戳，毫秒)
	UseStartTime int64           `json:"useStartTime"` //券有效使用开始时间(时间戳，毫秒)
	UseEndTime   int64           `json:"useEndTime"`   //券有效使用结束时间(时间戳，毫秒)
	IsBest       int64           `json:"isBest"`       //最优优惠券，1：是；0：否，购买一件商品可使用的面额最大优惠券
	HotValue     int64           `json:"hotValue"`     //券热度，值越大热度越高，区间:[0,10]
}

type ImageInfo struct {
	ImageList []ImageList `json:"imageList"`
}

type ImageList struct {
	URL string `json:"url"`
}

type PinGouInfo struct {
	PingouPrice     decimal.Decimal `json:"pingouPrice"`     //拼购价格
	PingouTmCount   int64           `json:"pingouTmCount"`   //拼购成团所需人数
	PingouURL       string          `json:"pingouUrl"`       //拼购落地页url
	PingouEndTime   int64           `json:"pingouEndTime"`   //拼购结束时间(时间戳，毫秒)
	PingouStartTime int64           `json:"pingouStartTime"` //拼购开始时间(时间戳，毫秒)
}

type PingGouInfo struct {
	PingouPrice   decimal.Decimal `json:"pingouPrice"`
	PingouTmCount int64           `json:"pingouTmCount"`
	PingouURL     string          `json:"pingouUrl"`
}

type ResourceInfo struct {
	EliteId   int64  `json:"eliteId"`   //频道id
	EliteName string `json:"eliteName"` //频道名称
}

type SeckillInfo struct {
	SeckillOriPrice  decimal.Decimal `json:"seckillOriPrice"`  //秒杀价原价
	SeckillPrice     decimal.Decimal `json:"seckillPrice"`     //秒杀价
	SeckillStartTime int64           `json:"seckillStartTime"` //秒杀开始时间(时间戳，毫秒)
	SeckillEndTime   int64           `json:"seckillEndTime"`   //秒杀结束时间(时间戳，毫秒)
}

type VideoInfo struct {
	VideoList []Video `json:"videoList"`
}

type Video struct {
	Width     int64  `json:"width"`     // 宽
	High      int64  `json:"high"`      // 高
	ImageUrl  string `json:"imageUrl"`  //视频图片地址
	VideoType int64  `json:"videoType"` // 1:主图，2：商详
	PlayType  string `json:"playType"`  // low：标清，high：高清
	Duration  int64  `json:"duration"`  // 时长(单位:s)
	PlayUrl   string `json:"playUrl"`   //播放地址
}

type DocumentInfo struct {
	Document string `json:"document"` //描述文案
	Discount string `json:"discount"` //优惠力度文案
}

type BookInfo struct {
	Isbn string `json:"isbn"`
}

type PriceInfo struct {
	Price             decimal.Decimal `json:"price"`             //商品价格
	LowestPrice       decimal.Decimal `json:"lowestPrice"`       //促销价
	LowestPriceType   int64           `json:"lowestPriceType"`   //促销价类型，1：无线价格；2：拼购价格； 3：秒杀价格
	LowestCouponPrice decimal.Decimal `json:"lowestCouponPrice"` //券后价（有无券都返回此字段）
	HistoryPriceDay   int64           `json:"historyPriceDay"`   //历史最低价天数（例：当前券后价最近180天最低）
}

type SkuLabelInfo struct {
	Is7ToReturn    int64                       `json:"is7ToReturn"`    //0：不支持； 1或null：支持7天无理由退货； 2：支持90天无理由退货； 4：支持15天无理由退货； 6：支持30天无理由退货；
	Fxg            int64                       `json:"fxg"`            //1：放心购商品
	FxgServiceList []CharacteristicServiceInfo `json:"fxgServiceList"` //放心购商品子标签集合
}

type CharacteristicServiceInfo struct {
	ServiceName string `json:"serviceName"` //服务名称
}

type PromotionLabelInfo struct {
	PromotionLabel   string `json:"promotionLabel"`   //商品促销文案
	LabelName        string `json:"labelName"`        //促销标签名称
	StartTime        int64  `json:"startTime"`        //促销开始时间
	EndTime          int64  `json:"endTime"`          //促销结束时间
	PromotionLabelId int64  `json:"promotionLabelId"` //促销ID
}

type SecondPriceInfo struct {
	SecondPriceType string          `json:"secondPriceType"` //双价格类型：18新人价
	SecondPrice     decimal.Decimal `json:"secondPrice"`     //价格
}
type ShopInfo struct {
	ShopID                        int64           `json:"shopId"`                        //店铺Id
	ShopName                      string          `json:"shopName"`                      //店铺名称（或供应商名称）
	ShopLevel                     decimal.Decimal `json:"shopLevel"`                     //店铺评分
	ShopLabel                     string          `json:"shopLabel"`                     // 1：京东好店 https://img12.360buyimg.com/schoolbt/jfs/t1/80828/19/2993/908/5d14277aEbb134d76/889d5265315e11ed.png
	UserEvaluateScore             string          `json:"userEvaluateScore"`             // 用户评价评分（仅pop店铺有值）
	CommentFactorScoreRankGrade   string          `json:"commentFactorScoreRankGrade"`   // 用户评价评级（仅pop店铺有值）
	LogisticsLvyueScore           string          `json:"logisticsLvyueScore"`           // 物流履约评分（仅pop店铺有值）
	LogisticsFactorScoreRankGrade string          `json:"logisticsFactorScoreRankGrade"` // 物流履约评级（仅pop店铺有值）
	AfterServiceScore             string          `json:"logisticsFactorScoreRankGrade"` // 售后服务评分（仅pop店铺有值）
	AfsFactorScoreRankGrade       string          `json:"afsFactorScoreRankGrade"`       // 售后服务评级（仅pop店铺有值）
	ScoreRankRate                 string          `json:"scoreRankRate"`                 // 店铺风向标（仅pop店铺有值）
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
