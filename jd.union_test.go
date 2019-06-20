package jingdong_union_go

import (
	"fmt"
	"log"
	"testing"
)

var app = &App{
	ID:     "1828155095",
	Key:    "7a043997b34c1d1e385f37570b7b3f57",
	Secret: "9b5e8e39e7a440538f10acd8023fe1a3",
}

//获取商品类目
func TestOpenCategoryGoodsGet(t *testing.T) {
	res, err := app.JdUnionOpenCategoryGoodsGet(map[string]interface{}{
		"parentId": 0,
		"grade":    0,
	})
	log.Println(res, err)
}

func TestOpenGoodsJingfenQuery(t *testing.T) {
	res, err := app.JdUnionOpenGoodsJingfenQuery(map[string]interface{}{
		"eliteId":   1,
		"sortName":  "price",
		"sort":      "asc",
		"pageIndex": 1,
		"pageSize":  10,
	})
	log.Println(res, err)
}

func TestOpenGoodsQuery(t *testing.T) {

	//单品查询
	res, err := app.JdUnionOpenGoodsQuery(map[string]interface{}{
		"skuIds":   []int{30881878056},
		"isCoupon": "1",
	})
	log.Println(res, err)

	//列表查询
	res, err = app.JdUnionOpenGoodsQuery(map[string]interface{}{
		"sort":                 "asc",   // asc desc
		"sortName":             "price", //price：单价, commissionShare：佣金比例, commission：佣金， inOrderCount30Days：30天引单量， inOrderComm30Days：30天支出佣金
		"isCoupon":             1,
		"commissionShareStart": 20, //佣金比例开始区间
		"pageIndex":            1,
		"pageSize":             10,
		"cid1":                 1315,
	})
	log.Println(res, err)
}

//获取单品信息
func TestOpenGoodsPromotionGoodsInfoQuery(t *testing.T) {
	res, err := app.JdUnionOpenGoodsPromotiongoodsinfoQuery(map[string]interface{}{
		"skuIds": "30881878056",
	})
	log.Println(res, err)
}

//获取通用推广链接
func TestOpenPromotionCommonGet(t *testing.T) {
	res, err := app.JdUnionOpenPromotionCommonGet(map[string]interface{}{
		"subUnionId": "test_subunionid",
		"ext1":       "test_ext",
		"siteId":     app.ID,
		"materialId": "https://wqitem.jd.com/item/view?sku=43415523405",
		"positionId": 1000,
		"couponUrl":  "http://coupon.jd.com/ilink/couponSendFront/send_index.action?key=02d2b6ff587c42fda6d4cac7ff1c2d6a&roleId=20498843&to=mall.jd.com/index-821028.html",
	})
	log.Println(res, err)
}

//获取商品订单
func TestOpenOrderQuery(t *testing.T) {

	//单品查询
	res, err := app.JdUnionOpenOrderQuery(map[string]interface{}{
		"type":     "1", //1 下单时间  2 完成时间 3 更新时间
		"time":     "201906141811",
		"pageNo":   1,
		"pagesize": 500,
	})

	log.Println(res, err)
}

// 通过subUnionid获取推广链接
//https://wqitem.jd.com/item/view?sku=
func TestOpenPromotionBySubUnionIdGet(t *testing.T) {
	skuId := 43415523405
	res, err := app.JdUnionOpenPromotionBysubunionidGet(map[string]interface{}{
		"subUnionId": "{\"u\":\"1020\"}",
		"positionId": 1000,
		"chainType":  3,
		"materialId": fmt.Sprintf("https://wqitem.jd.com/item/view?sku=%d", skuId),
		"couponUrl":  "http://coupon.jd.com/ilink/couponSendFront/send_index.action?key=02d2b6ff587c42fda6d4cac7ff1c2d6a&roleId=20498843&to=mall.jd.com/index-821028.html",
	})
	log.Println(res, err)
}
