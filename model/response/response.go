package response

import (
	model "HotApi/model"
	"encoding/json"
	"net/http"
)

const (
	ERROR   = 1001
	SUCCESS = 1000
)

type Response struct {
	Code    int                 `json:"code"`
	Data    model.HotSearchData `json:"data"`
	Message string              `json:"message"`
}

func OkWithData(w http.ResponseWriter, data model.HotSearchData) {
	_ = json.NewEncoder(w).Encode(Response{
		Code:    SUCCESS,
		Data:    data,
		Message: "success",
	})
}

func Failed(w http.ResponseWriter, err error) {
	_ = json.NewEncoder(w).Encode(Response{
		Code:    ERROR,
		Data:    model.HotSearchData{},
		Message: err.Error(),
	})
}
