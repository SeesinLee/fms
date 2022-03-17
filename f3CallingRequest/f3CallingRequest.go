package f3CallingRequest

type F3Calling struct {
	Status string	`json:"status"`
	Data data	`json:"data"`
}

type data struct {
	ResultType	string	`json:"resultType"`
	Result	*[]Result	`json:"result"`
}

type Result struct {
	Metric	*metric	`json:"metric"`
	Value []interface{}	`json:"value"`
}

type metric struct {
	MName string 	`json:"__name__"`
	ExportedInstance string `json:"exported_instance"`
	ExportedJob	string	`json:"exported_job"`
	Instance string	`json:"instance"`
	Job string	`json:"job"`
	Name string	`json:"name"`
}