package util

import (
	"fmt"
	"testing"
)

func TestHttpPostBody(t *testing.T) {
	s, e := HttpPostBody("http://192.168.1.95:8280/api/admin/sys/datacollecttask/test", "{}", "text/plain")
	fmt.Println(s)
	fmt.Println(e)
}