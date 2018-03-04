package handler

import (
	"fmt"
	"net/http"
)

// IndexHandle func(res http.ResponseWriter, req *http.Request)
func IndexHandle(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Route Index : %v\n", req.URL)
}
