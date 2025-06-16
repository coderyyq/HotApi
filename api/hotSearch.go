package api

import (
	"HotApi/api/hot"
	"HotApi/model/response"
	"errors"
	"net/http"
	"strings"
)

func GetHotListHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response.Failed(w, errors.New("仅允许GET请求~"))
		return
	}

	w.Header().Set("Content-Type", "application/json")

	path := r.URL.Path
	segments := strings.Split(path, "/")
	if len(segments) != 4 {
		response.Failed(w, errors.New("URL参数错误~"))
		return
	}
	source := hot.NewSource(segments[3])
	if source == nil {
		response.Failed(w, errors.New("无法获取此源数据~"))
		return
	}
	hotSearchData, err := source.GetHotSearchData(30)
	if err != nil {
		response.Failed(w, errors.New(err.Error()))
		return
	}
	response.OkWithData(w, hotSearchData)
}
