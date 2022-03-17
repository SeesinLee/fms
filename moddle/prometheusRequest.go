package moddle

type PrometheusMetrics struct {
	Status string `json:"status"`
	Data   data   `json:"data"`
}

type data struct {
	ResultType *string   `json:"resultType"`
	Result     *[]Result `json:"result"`
}
type Result struct {
	Metric *metric  `json:"metric"`
	Value  []interface{} `json:"value"`
}

type metric struct {
	Name                string `json:"__name__"`
	App                 string `json:"app"`
	Author              string `json:"author"`
	Container           string `json:"container"`
	Endpoint            string `json:"endpoint"`
	Instance            string `json:"instance"`
	Job                 string `json:"job"`
	KubernetesName      string `json:"kubernetes_name"`
	KubernetesNamespace string `json:"kubernetes_namespace"`
	Namespace           string `json:"namespace"`
	Service             string `json:"service"`
}

type InstanceData struct {
	Instances []string	`json:"instances"`
}
