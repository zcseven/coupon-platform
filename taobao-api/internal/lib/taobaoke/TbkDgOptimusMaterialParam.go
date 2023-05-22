package taobaoke

import "strconv"

//TbkDgOptimusMaterialParam 通用物料搜索API（导购）
//https://open.taobao.com/api.htm?docId=35896&docType=2
//通用物料搜索API（导购）
type TbkDgOptimusMaterialParam struct {
	AdzoneId   int64 `json:"adzone_id"`
	MaterialId int64 `json:"material_id"`

	PageNo        int64  `json:"page_no"`
	PageSize      int64  `json:"page_size"`
	DeviceValue   string `json:"device_value"`
	DeviceType    string `json:"device_type"`
	DeviceEncrypt string `json:"device_encrypt"`
	ContentId     string `json:"content_id"`
	ContentSource string `json:"content_source"`
	ItemId        string `json:"item_id"`
	FavoritesId   string `json:"favorites_id"`
}

//APIName API名称
func (m TbkDgOptimusMaterialParam) APIName() string {
	return "taobao.tbk.dg.optimus.material"
}

//Params 参数
func (m TbkDgOptimusMaterialParam) Params() map[string]string {
	var params = make(map[string]string)
	params["adzone_id"] = strconv.FormatInt(m.AdzoneId, 10)
	params["material_id"] = strconv.FormatInt(m.MaterialId, 10)

	return params
}

//ExtJSONParamName JSON类型参数的名称
func (m TbkDgOptimusMaterialParam) ExtJSONParamName() string {
	return "" //请求参数的名称
}

//ExtJSONParamValue JSON类型参数的值
func (m TbkDgOptimusMaterialParam) ExtJSONParamValue() string {
	return marshal(m)
}

//AddParam 增加参数.暂时未用到
func (m *TbkDgOptimusMaterialParam) AddParam(key string, value interface{}) {

}
