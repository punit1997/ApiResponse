package eecom

type JobOrderResponse struct {
	Raw string
	Awbnumber int
	Ordernumber string
	Reason string
	Status string
	Success bool
	Error string
	External_trace_id string
	MozartId string
	MozartSuccess string
}