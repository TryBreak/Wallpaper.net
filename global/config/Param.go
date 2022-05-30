package config

type BingImgType struct {
	StartTime     int64  `json:"StartTime"`
	EndTime       int64  `json:"EndTime"`
	Path          string `json:"Path"`
	Url           string `json:"Url"`
	Copyright     string `json:"Copyright"`
	CopyrightLink string `json:"CopyrightLink"`
	Title         string `json:"Title"`
}

var BingImg []BingImgType
