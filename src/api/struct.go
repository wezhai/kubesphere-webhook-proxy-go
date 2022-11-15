package api

type Alert struct {
	StartsAt    string `json:"startsAt"`
	Annotations struct {
		AliasName string `json:"aliasName"`
		Kind      string `json:"kind"`
		Summary   string `json:"summary"`
		Resources string `json:"resources"`
	} `json:"annotations"`
	Labels struct {
		Deployment string `json:"deployment"`
		Node       string `json:"node"`
		Workload   string `json:"workload"`
	} `json:"labels"`
}

type Hook struct {
	CommonLables struct {
		Severity  string `json:"severity"`
		Alertname string `json:"alertname"`
		Namespace string `json:"namespace"`
	} `json:"commonLabels"`
	Alerts            []Alert `json:"alerts"`
	CommonAnnotations struct {
		AliasName string `json:"aliasName"`
	} `json:"commonAnnotations"`
}
