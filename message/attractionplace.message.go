package message

type AttractionPlace struct {
	ID             string        `json:"id"`
	TitleTh        string        `json:"title_th"`
	AddressTh      Address       `json:"address_th"`
	Opentime       Opentime      `json:"opentime"`
	Recommened     Recommened    `json:"recommened"`
	Source         string        `json:"source"`
	AddressFull    string        `json:"address_full"`
	Verified       bool          `json:"verified"`
	Tag            []interface{} `json:"tag"`
	AboutTh        string        `json:"about_th"`
	Type           string        `json:"type"`
	Rating         float64       `json:"rating"`
	Iframe         string        `json:"iframe"`
	GoogleMapURL   string        `json:"googleMap_url"`
	GoogleMapShare string        `json:"googleMap_share"`
	Lat            float64       `json:"lat"`
	Long           float64       `json:"long"`
	Contact        []interface{} `json:"contact"`
	Reviews        Reviews       `json:"reviews"`
	AddressEn      Address       `json:"address_en"`
	TitleEn        string        `json:"title_en"`
}

type Address struct {
	SubDistrict string `json:"sub_district"`
	District    string `json:"district"`
	Province    string `json:"province"`
	PostCode    string `json:"post_code"`
}

type Opentime struct {
	Monday    Day `json:"monday"`
	Tuesday   Day `json:"tuesday"`
	Wednesday Day `json:"wednesday"`
	Thursday  Day `json:"thursday"`
	Friday    Day `json:"friday"`
}

type Day struct {
	O string `json:"o"`
	C string `json:"c"`
}

type Recommened struct {
}

type Reviews struct {
	AmountAll int     `json:"amount_all"`
	Excellent Average `json:"excellent"`
	VeryGood  Average `json:"very_good"`
	Average   Average `json:"average"`
	Poor      Average `json:"poor"`
	Terrible  Average `json:"terrible"`
}

type Average struct {
	Amount  float32 `json:"amount"`
	LazyURL string  `json:"lazy_url"`
}
