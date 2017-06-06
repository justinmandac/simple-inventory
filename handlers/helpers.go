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

func createJSONResponse(err []error, data interface{}) (res models.Response) {
	var errsCopy []error
	// TODO: eliminate double loops
	for i := 0; i < len(err); i++ {
		if err[i] != nil {
			errsCopy = append(errsCopy, err[i])
		}
	}

	if len(errsCopy) > 0 {
		var stringErrs []string

		for i := 0; i < len(err); i++ {
			stringErrs = append(stringErrs, err[i].Error())
		}

		res = models.Response{Err: 1, Message: stringErrs, Data: nil}
	} else {
		res = models.Response{Err: 0, Message: []string{}, Data: data}
	}

	return res
}
