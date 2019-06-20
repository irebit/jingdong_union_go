package jingdong_union_go

import (
	"fmt"
	"log"
	"testing"
)

func TestOpenGoodsQuery(t *testing.T) {
	app := &App{
		Key:    "5510523a691c199aa038c637e59acf03",
		Secret: "28e4b19760e14f9fa83f333d07929a20",
	}

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

//获取商品订单
func TestOpenOrderQuery(t *testing.T) {
	app := &App{
		Key:    "5510523a691c199aa038c637e59acf03",
		Secret: "28e4b19760e14f9fa83f333d07929a20",
	}

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
	log.Println(fmt.Sprintf("https://wqitem.jd.com/item/view?sku=%d", skuId))
	app := &App{
		Key:    "5510523a691c199aa038c637e59acf03",
		Secret: "28e4b19760e14f9fa83f333d07929a20",
	}
	res, err := app.JdUnionOpenPromotionBysubunionidGet(map[string]interface{}{
		"subUnionId": "{\"u\":\"1020\"}",
		"positionId": 1000,
		"chainType":  3,
		"materialId": fmt.Sprintf("https://wqitem.jd.com/item/view?sku=%d", skuId),
		"couponUrl":  "http://coupon.jd.com/ilink/couponSendFront/send_index.action?key=02d2b6ff587c42fda6d4cac7ff1c2d6a&roleId=20498843&to=mall.jd.com/index-821028.html",
	})
	log.Println(res, err)
}
