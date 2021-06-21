package model

type TbkDgMaterialOptionalRequest struct {
	StartDsr          int64            `json:"start_dsr,omitempty"`
	PageSize          int64            `json:"page_size,omitempty"`
	PageNO            int64            `json:"page_no,omitempty"`
	Platform          int64            `json:"platform,omitempty"`
	EndTkRate         int64            `json:"end_tk_rate,omitempty"`
	StartTkRate       int64            `json:"start_tk_rate,omitempty"`
	EndPrice          int64            `json:"end_price,omitempty"`
	StartPrice        int64            `json:"start_price,omitempty"`
	IsOverseas        bool             `json:"is_overseas"`
	IsTmall           bool             `json:"is_tmall"`
	Sort              string           `json:"sort,omitempty"`
	Itemloc           string           `json:"itemloc,omitempty"`
	Cat               string           `json:"cat,omitempty"`
	Q                 string           `json:"q,omitempty"`
	MaterialID        int64            `json:"material_id,omitempty"`
	HasCoupon         bool             `json:"has_coupon"`
	Ip                string           `json:"ip,omitempty"`
	AdzoneID          int64            `json:"adzone_id,omitempty"`
	NeedFreeShipment  bool             `json:"need_free_shipment,omitempty"`
	NeedPrepay        bool             `json:"need_prepay,omitempty"`
	IncludePayRate30  bool             `json:"include_pay_rate_30,omitempty"`
	IncludeGoodRate   bool             `json:"include_good_rate,omitempty"`
	IncludeRfdRate    bool             `json:"include_rfd_rate,omitempty"`
	NpxLevel          int              `json:"npx_level,omitempty"`
	EndKaTkRate       int64            `json:"end_ka_tk_rate,omitempty"`
	StartKaTkRate     int64            `json:"start_ka_tk_rate,omitempty"`
	DeviceEncrypt     string           `json:"device_encrypt,omitempty"`
	DeviceValue       string           `json:"device_value,omitempty"`
	DeviceType        string           `json:"device_type,omitempty"`
	LockRateEndTime   int64            `json:"lock_rate_end_time,omitempty"`
	LockRateStartTime int64            `json:"lock_rate_start_time,omitempty"`
	Longitude         string           `json:"longitude,omitempty"`
	Latitude          string           `json:"latitude,omitempty"`
	CityCode          string           `json:"city_code,omitempty"`
	SellerIDs         string           `json:"seller_ids,omitempty"`
	SpecialID         string           `json:"special_id,omitempty"`
	RelationID        string           `json:"relation_id,omitempty"`
	PageResultKey     string           `json:"page_result_key,omitempty"`
	UcrowdID          int64            `json:"ucrowd_id,omitempty"`
	GetTopnRate       int64            `json:"get_topn_rate,omitempty"`
	UcrowdRankItems   []UcrowdRankItem `json:"ucrowd_rank_items,omitempty"`
}

type UcrowdRankItem struct {
	Commirate int64  `json:"commirate,omitempty"`
	Price     string `json:"price,omitempty"`
	ItemID    string `json:"item_id,omitempty"`
}

type TbkDgMaterialOptionalResponse struct {
	BaseResponse
	TotalResults                  int64                  `json:"total_results" xml:"total_results"`
	PageResultKey                 string                 `json:"page_result_key" xml:"page_result_key"`
	TbkDgMaterialOptionalResponse map[string]interface{} `json:"tbk_dg_material_optional_response,omitempty" xml:"tbk_dg_material_optional_response"`
}
