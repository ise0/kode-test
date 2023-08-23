package speller

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/ise0/kode-test/src/shared/logger"
)

type spellerError struct {
	Code int      `json:"code"`
	Pos  int      `json:"pos"`
	Col  int      `json:"col"`
	Row  int      `json:"row"`
	Len  int      `json:"len"`
	Word string   `json:"word"`
	S    []string `json:"s"`
}

func CheckText(text string) ([]spellerError, error) {
	body := new(bytes.Buffer)
	mp := multipart.NewWriter(body)
	mp.WriteField("text", text)

	res, err := http.Post(
		"https://speller.yandex.net/services/spellservice.json/checkText",
		mp.FormDataContentType(),
		body)

	if err != nil {
		logger.L.Errorw("Request failed", "error", err)
		return nil, err
	}

	b, err := io.ReadAll(res.Body)
	if err != nil {
		logger.L.Errorw("io body read failed", "error", err)
		return nil, err
	}

	if res.StatusCode != 200 {
		logger.L.Errorw("Request failed", "error", string(b))
	}

	var spellerRes []spellerError
	err = json.Unmarshal(b, &spellerRes)
	if err != nil {
		logger.L.Errorw("Failed to parse json", "error", err)
		return nil, err
	}

	return spellerRes, nil
}
