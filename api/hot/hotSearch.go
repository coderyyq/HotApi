package hot

import "HotApi/model"

type Source interface {
	GetHotSearchData(maxNum int) (HotSearchData model.HotSearchData, err error)
}

func NewSource(source string) Source {
	switch source {
	case "baidu":
		return &Baidu{}
	case "bilibili":
		return &Bilibili{}
	case "douyin":
		return &Douyin{}
	case "kuaishou":
		return &Kuaishou{}
	case "pengpai": // 20
		return &Pengpai{}
	case "qqnews":
		return &QQnews{}
	case "sina":
		return &Sina{}
	case "sougou":
		return &Sougou{}
	case "sspai":
		return &Sspai{}
	case "tieba":
		return &Tieba{}
	case "toutiao":
		return &Toutiao{}
	case "weibo":
		return &Weibo{}
	default:
		return nil
	}
}
