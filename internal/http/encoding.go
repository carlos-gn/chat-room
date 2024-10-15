package http

import (
	"encoding/base64"
	"errors"
	"fmt"
	"strings"
	"time"
)

func EncodeCursor(t time.Time, id string) string {
	key := fmt.Sprintf("%s,%s", t.Format(time.RFC3339Nano), id)
	return base64.StdEncoding.EncodeToString([]byte(key))
}

func DecodeCursor(cursor string) (time.Time, string, error) {
	bytes, err := base64.StdEncoding.DecodeString(cursor)
	if err != nil {
		return time.Time{}, "", err
	}

	arrStr := strings.Split(string(bytes), ",")
	if len(arrStr) != 2 {
		err = errors.New("invalid cursor")
		return time.Time{}, "", err
	}

	res, err := time.Parse(time.RFC3339Nano, arrStr[0])
	if err != nil {
		return time.Time{}, "", err
	}

	id := arrStr[1]

	return res, id, nil
}
