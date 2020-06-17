package models

import (
	"errors"
	"io"
	"strconv"
	"time"

	"github.com/99designs/gqlgen/graphql"
)

func MarshalTimestamp(t time.Time) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, strconv.FormatInt(t.Unix(), 10))
	})
}

func UnmarshalTimestamp(v interface{}) (time.Time, error) {
	if timestampStr, ok := v.(int64); ok {
		return time.Unix(timestampStr, 0), nil
	}
	return time.Time{}, errors.New("time should be a unix timestamp")
}
