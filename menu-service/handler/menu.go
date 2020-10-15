package handler

import (
	"net/http"

	"github.com/Bayu200720/microservices/utils"
)

func AddMenu(w http.ResponseWriter, r *http.Request) {
	utils.WrapAPISuccess(w, r, "success", http.StatusOK)

}
