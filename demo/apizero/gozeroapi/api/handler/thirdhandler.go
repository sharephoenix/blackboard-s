package handler

import (
	"github.com/tal-tech/go-zero/rest/httpx"
	"net/http"
)

type Request struct {
	User string `form:"user"`
}

func Handle(w http.ResponseWriter, r *http.Request) {
	var req Request
	err := httpx.Parse(r, &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	httpx.OkJson(w, "helllo, " + req.User)
}