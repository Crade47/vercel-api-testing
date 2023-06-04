package handler

import (
	"net/http"

	"github.com/Crade47/vercel-api-testing/utils"
)

func HandleHello(w http.ResponseWriter, r *http.Request) error {
	return utils.WriteJSON(w, http.StatusOK, map[string]string{"message": "hello!!"})
}
