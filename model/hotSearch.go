package model

type HotItem struct {
	Id          int    `json:"id"`          // 排名
	Title       string `json:"title"`       // 标题
	Description string `json:"description"` // 描述
	Picture       string `json:"picture"`       // 图片
	Popularity  string `json:"popularity"`  // 热度
	URL         string `json:"url"`         // 链接
}

type HotSearchData struct {
	Source     string    `json:"source"`      // 数据源
	UpdateTime string    `json:"update_time"` // 更新时间
	HotList    []HotItem `json:"hot_list"`    // 热搜列表
}
