package request

type MasterAddressData struct {
	Name          string `json:"name"`
	Address       string `json:"address"`
	LocationType  int    `json:"location_type"`
	DetailAddress string `json:"detail_address"`
	PhoneNumber   string `json:"phone_number"`
	Code          string `json:"code"`
	Longitude     string `json:"longitude"`
	Latitude      string `json:"latitude"`
}

type MasterGoodsData struct {
	GoodsId      int     `json:"goods_id"`
	GoodsName    string  `json:"goods_name,omitempty"`
	CommodityId  int     `json:"commodity_id,omitempty"`
	Unit         string  `json:"unit,omitempty"`
	Code         string  `json:"code,omitempty"`
	Qty          float64 `json:"qty"`
	UnitWeightKg float64 `json:"unit_weight_kg,omitempty"`
	TotalKg      float64 `json:"total_weight_kg,omitempty"`
}
