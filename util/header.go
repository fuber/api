package util

import (
	"fmt"
	"time"
)

const TimeFormat = "Mon, 02 Jan 2006 15:04:05 GMT"
const HTTPHeader = `HTTP/1.1 200 OK
Date: %s
Content-Type: text/plain; charset=UTF-8
Transfer-Encoding: chunked

`

func ChunkedHeader() string {
	return fmt.Sprintf(HTTPHeader, time.Now().UTC().Format(TimeFormat))
}
