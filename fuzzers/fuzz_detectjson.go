package myfuzz

import "github.com/bytedance/godlp"

func Fuzz(data []byte) int {
	jsonBody := string(data)
	eng, err := dlp.NewEngine("replace.your.psm")
	if err != nil {
		return 1
	}

	err = eng.ApplyConfigDefault()
	if err != nil {
		return 1
	}
	_, err = eng.DetectJSON(jsonBody)
	if err != nil {
		return 1
	}
	return 0

}
