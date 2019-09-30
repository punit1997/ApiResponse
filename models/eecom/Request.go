package eecom

type JobOrderRequest struct {
	Order_number string
	Product string
	Consignee string
	Consignee_adr1 string
	collectable_value string
	desig_city string
	pincode string
	state string
	mobile string
	telephone string
	item_desc string
	drop_vendor_code string
	drop_name string
	drop_adr1 string
	drop_pin string
	drop_mobile string
	actcode string `json: "act.code"`
	date string
	time_slot string
}
