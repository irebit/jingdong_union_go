# jingdong_union_go

[京东联盟](https://union.jd.com/helpcenter/12188-12384-46301)SDK

部分接口为为高级接口，需要联盟企业账号发送邮件申请

### API列表

#### 通用API
- [x] [订单查询接口 jd.union.open.order.query](https://union.jd.com/openplatform/api/650)
- [x] [京粉精选商品查询接口 jd.union.open.goods.jingfen.query](https://union.jd.com/openplatform/api/739)
- [ ] [获取推广商品信息接口 jd.union.open.goods.promotiongoodsinfo.query]()
- [ ] [获取通用推广链接 jd.union.open.promotion.common.get]()
- [x] [商品类目查询 jd.union.open.category.goods.get](https://union.jd.com/openplatform/api/693)

#### 高级API
- [x] [关键词商品查询接口【申请】jd.union.open.goods.query](https://union.jd.com/openplatform/api/628)
- [ ] [秒杀商品查询接口【申请】 jd.union.open.goods.seckill.query]()

- [ ] [学生价商品查询接口【申请】 jd.union.open.goods.stuprice.query]()

- [ ] [通过unionId获取推广链接【申请】 jd.union.open.promotion.byunionid.get]()

- [x] [通过subUnionId获取推广链接【申请】 jd.union.open.promotion.bysubunionid.get](https://union.jd.com/openplatform/api/634)

- [ ] [优惠券领取情况查询接口【申请】 jd.union.open.coupon.query]()

- [ ] [创建推广位【申请】 jd.union.open.position.create]()

- [ ] [查询推广位【申请】 jd.union.open.position.query]()

- [ ] [获取PID jd.union.open.user.pid.get]()

###  运行测试用例

需要在jd.union_test.go中填入自己的appKey,appSecret

``` golang
go test -v *.go
```

### 接口调用

``` golang
    var app = &App{
        Key:    "xxxxxxx",
        Secret: "xxxxxxx",
    }
    res, err := app.JdUnionOpenCategoryGoodsGet(map[string]interface{}{
        "parentId": 0,
        "grade":    0,
    })
    log.Println(res, err)
```