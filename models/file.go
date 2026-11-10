package model

type CSVResponse struct {
	Filename string              `json:"filename"`
	Total    int                 `json:"total"`
	Data     []map[string]string `json:"data"`
}
