package model

type TbkDgOptimusMaterialRequest struct {
	PageSize      int64  `json:"page_size,omitempty"`
	AdzoneID      int64  `json:"adzone_id,omitempty"`
	PageNO        int64  `json:"page_no,omitempty"`
	MaterialID    int64  `json:"material_id,omitempty"`
	SiteID        int64  `json:"site_id,omitempty"`
	DeviceType    string `json:"device_type,omitempty"`
	DeviceEncrypt string `json:"device_encrypt,omitempty"`
	DeviceValue   string `json:"device_value,omitempty"`
	ContentID     int64  `json:"content_id,omitempty"`
	ContentSource string `json:"content_source,omitempty"`
	ItemID        int64  `json:"item_id,omitempty"`
}

type TbkDgOptimusMaterialResponse struct {
	BaseResponse
	Data map[string]interface{} `json:"tbk_dg_optimus_material_response" xml:"tbk_dg_optimus_material_response"`
}