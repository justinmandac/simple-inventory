package handlers

import (
	"encoding/json"
	"net/http"
	"simple-inventory/models"
)

func writeJSON(w http.ResponseWriter, response models.Response) (err error) {
	res, err := json.Marshal(response)

	w.Header().Set("content-type", "application/json")
	if err != nil {
		data := models.Response{Err: 1, Message: []string{err.Error()}, Data: nil}
		errRes, _ := json.Marshal(data)
		w.Write(errRes)
		return err
	}

	w.Write(res)
	return
}
