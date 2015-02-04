package models

type Commodity struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	TypeId   int    `json:"category_id"`
	AvgPrice int    `json:"average_price"`
}

type CommodityType struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

// Listing is a specific commodity buy or sell listing at a station.
type Listing struct {
	Id          uint64 `json:"id"`
	CommodityId int    `json:"commodity_id"`
	Supply      int    `json:"supply"`
	BuyPrice    int    `json:"buy_price"`
	Demand      int    `json:"demand"`
	SellPrice   int    `json:"sell_price"`
	Timestamp   int    `json:"collected_at"`
}
