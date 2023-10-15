package models

import (
	"github.com/astaxie/beego/logs"
	"github.com/go-xorm/xorm"
	"math"
	"reflect"
)

type PageResp struct {
	Count       int `json:"count"`       // 总记录数
	CurrentPage int `json:"currentPage"` // 当前第几页
	PageSize    int `json:"pageSize"`    // 每页数量
	TotalPage   int `json:"totalPage"`   // 总页数
}

/*
interface数组分页
- [list] 原长度数组
- [currentPage] 当前页码 当小于等于0时，获取全部
- [pageSize] 页面容量 当小于等于0时，获取全部

- [pageList] 分页后数据
*/
func (pageResp *PageResp) NewListPageResp(list interface{}, currentPage int, pageSize int) (pageList []interface{}) {
	if reflect.TypeOf(list).Kind() != reflect.Array && reflect.TypeOf(list).Kind() != reflect.Slice && reflect.TypeOf(list).Kind() != reflect.Map {
		logs.Error("list 传入类型不为数组", reflect.TypeOf(list).Kind())
		return
	}
	var reflectList []interface{}

	if reflect.TypeOf(list).Kind() == reflect.Array || reflect.TypeOf(list).Kind() == reflect.Slice {
		val := reflect.ValueOf(list)
		for i := 0; i < val.Len(); i++ {
			reflectList = append(reflectList, val.Index(i).Interface())
		}
	} else {
		LMaps := list.(map[string]interface{})
		for _, _map := range LMaps {
			reflectList = append(reflectList, _map)
		}

	}
	if currentPage <= 0 || pageSize <= 0 {
		pageList = reflectList
		return
	}
	//总条数
	total := len(reflectList)
	pageResp.BuildCount(total, currentPage, pageSize)
	//截取index
	firstIndex := (currentPage - 1) * pageSize
	////实际容量
	//cap := pageSize
	//拦截index
	endIndex := firstIndex + pageSize

	if firstIndex > endIndex {
		return
	}
	if firstIndex < 0 {
		pageList = reflectList
		return
	}
	if endIndex < 0 {
		pageList = reflectList
		return
	}
	if endIndex > total {
		pageList = reflectList[firstIndex:]
		return
	}
	pageList = reflectList[firstIndex:endIndex]
	return
}

func (pageResp *PageResp) NewTotalPageResp(total int, currentPage int, pageSize int) (pageList []interface{}) {

	pageResp.BuildCount(total, currentPage, pageSize)
	//截取index
	firstIndex := (currentPage - 1) * pageSize
	////实际容量
	//cap := pageSize
	//拦截index
	endIndex := firstIndex + pageSize

	if firstIndex > endIndex {
		return
	}
	if firstIndex < 0 {
		return
	}
	if endIndex < 0 {
		return
	}
	if endIndex > total {
		return
	}
	return
}

/*
map数组分页
- [list] 原长度数组
- [currentPage] 当前页码 当小于等于0时，获取全部
- [pageSize] 页面容量 当小于等于0时，获取全部

- [pageList] 分页后数据
*/
func (pageResp *PageResp) NewMapListPageResp(list []map[string]interface{}, currentPage int, pageSize int) (pageList []map[string]interface{}) {
	if reflect.TypeOf(list).Kind() != reflect.Array && reflect.TypeOf(list).Kind() != reflect.Slice {
		logs.Error("list 传入类型不为数组", reflect.TypeOf(list).Kind())
		return
	}
	var reflectList []map[string]interface{}
	//val := reflect.ValueOf(list)
	for i := 0; i < len(list); i++ {
		reflectList = append(reflectList, list[i])
	}
	if currentPage <= 0 || pageSize <= 0 {
		pageList = reflectList
		return
	}
	//总条数
	total := len(reflectList)
	pageResp.BuildCount(total, currentPage, pageSize)
	//截取index
	firstIndex := (currentPage - 1) * pageSize
	////实际容量
	//cap := pageSize
	//拦截index
	endIndex := firstIndex + pageSize

	if firstIndex > endIndex {
		return
	}
	if firstIndex < 0 {
		pageList = reflectList
		return
	}
	if endIndex < 0 {
		pageList = reflectList
		return
	}
	if endIndex > total {
		pageList = reflectList[firstIndex:]
		return
	}
	pageList = reflectList[firstIndex:endIndex]
	return
}

/*计算页码*/
func (pageRes *PageResp) Build(page int, pageSize int) {
	pageRes.CurrentPage = page
	pageRes.PageSize = pageSize
	if pageSize == 0 {
		pageRes.TotalPage = 0
	} else {
		r := float64(pageRes.Count) / float64(pageSize)
		pageRes.TotalPage = int(math.Ceil(r))
	}
}

/*计算页码*/
func (pageRes *PageResp) BuildCount(count int, page int, pageSize int) {
	pageRes.Count = count
	pageRes.CurrentPage = page
	pageRes.PageSize = pageSize
	if pageSize == 0 {
		pageRes.TotalPage = 0
	} else if pageSize == -1 {
		pageRes.TotalPage = 1
	} else {
		r := float64(pageRes.Count) / float64(pageSize)
		pageRes.TotalPage = int(math.Ceil(r))
	}
}

/*计算pipeline分页*/
func (pageRes *PageResp) CollTotalCount(coll xorm.Session, page, pageSize int) (err error) {
	var total int64
	//获取最大条数
	total, err = coll.Count()
	pageRes.BuildCount(int(total), page, pageSize)

	return
}
