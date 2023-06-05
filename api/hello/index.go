package handler

import (
	"fmt"
	"net/http"

	"github.com/Crade47/vercel-api-testing/middleware"
	"github.com/Crade47/vercel-api-testing/utils"
)

func HandleHello(w http.ResponseWriter, r *http.Request) {
	middleware.EnsureValidToken()(http.HandlerFunc((func(innerW http.ResponseWriter, innerR *http.Request) {
		fmt.Println("Inside the edge function")
		helloResponse(w, r)
	})))
}

func helloResponse(w http.ResponseWriter, r *http.Request) {
	if err := utils.WriteJSON(w, http.StatusOK, map[string]string{"message": "hello!!"}); err != nil {
		fmt.Errorf("%s", err.Error())
	}
}
