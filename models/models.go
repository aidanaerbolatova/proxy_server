package models

type Request struct {
	Method  string            `json:"method"`
	URL     string            `json:"URL"`
	Headers map[string]string `json:"headers"`
}

type Response struct {
	Id      string              `json:"id"`
	Status  string              `json:"status"`
	Headers map[string][]string `json:"headers"`
	Length  int                 `json:"length"`
}
