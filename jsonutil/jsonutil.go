package jsonutil

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, data any, status int, headers http.Header) error {
	out, err := json.Marshal(data)
	if err != nil {
		return err
	}
	out = append(out, '\n')
	for k, v := range headers {
		w.Header()[k] = v
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(out)
	return nil
}

func ReadJSON(w http.ResponseWriter, r *http.Request, dest any) error {
	const MAX_BYTES = 1_048_576
	r.Body = http.MaxBytesReader(w, r.Body, int64(MAX_BYTES))
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	err := dec.Decode(dest)
	if err != nil {
		return err
	}
	err = dec.Decode(&struct{}{})
	if !errors.Is(err, io.EOF) {
		return errors.New("body must contain a single JSON value")
	}
	return nil
}
