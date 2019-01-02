package common

type FormOption struct{
	Key interface{} `json:"key"`
	Value string `json:"value"`
}

type AccLists map[string]map[string]struct{}