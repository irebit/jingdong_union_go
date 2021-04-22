package jingdong_union_go

import (
	"encoding/json"
	"errors"
	"log"
)

type JdUnionOpenActivityQueryTopLevel struct {
	JdUnionOpenActivityQueryResponse JdUnionOpenActivityQueryResponse `json:"jd_union_open_activity_query_responce"`
}

type JdUnionOpenActivityQueryResponse struct {
	Result string `json:"queryResult"`
	Code   string `json:"code"`
}

type JdUnionOpenActivityQueryResult struct {
	Code       int64                           `json:"code"`
	Data       []*JdUnionOpenActivityQueryData `json:"data"`
	Message    string                          `json:"message"`
	TotalCount int                             `json:"totalCount"`
	RequestID  string                          `json:"requestId"`
}

type JdUnionOpenActivityQueryData struct {
	ActStatus    int                              `json:"actStatus"`
	Advantage    string                           `json:"advantage"`
	Content      string                           `json:"content"`
	DownloadCode string                           `json:"downloadCode"`
	DownloadUrl  string                           `json:"downloadUrl"`
	EndTime      int                              `json:"endTime"`
	ID           int                              `json:"id"`
	PlatformType int                              `json:"platformType"`
	StartTime    int                              `json:"startTime"`
	Tag          string                           `json:"tag"`
	Title        string                           `json:"title"`
	UpdateTime   int                              `json:"updateTime"`
	UrlM         string                           `json:"urlM"`
	UrlPC        string                           `json:"urlPC"`
	CategoryList []*JdUnionOpenActivityQueryCate  `json:"categoryList"`
	ImgList      []*JdUnionOpenActivityQueryImage `json:"imgList"`
}

type JdUnionOpenActivityQueryCate struct {
	CategoryID int `json:"categoryId"`
	Type       int `json:"type"`
}

type JdUnionOpenActivityQueryImage struct {
	ImgName     string `json:"imgName"`
	ImgUrl      string `json:"imgUrl"`
	WidthHeight string `json:"widthHeight"`
}

func (app *App) JdUnionOpenActivityQuery(params map[string]interface{}) (result *JdUnionOpenActivityQueryResult, err error) {

	body, err := app.Request("jd.union.open.activity.query", map[string]interface{}{"activityReq": params})
	log.Println(string(body))
	resp := &JdUnionOpenActivityQueryTopLevel{}
	if err != nil {
		log.Println(string(body))
		return
	}

	if err = json.Unmarshal(body, resp); err != nil {
		return
	}
	if resp.JdUnionOpenActivityQueryResponse.Result != "" {
		result = &JdUnionOpenActivityQueryResult{}
		if err = json.Unmarshal([]byte(resp.JdUnionOpenActivityQueryResponse.Result), result); err != nil {
			return
		}
	} else {
		err = errors.New("result is null")
	}
	return
}
