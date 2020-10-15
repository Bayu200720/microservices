package handler

import (
	"net/http"

	"github.com/Bayu200720/microservices/utils"
)

func AddMenu(http.ResponseWriter, *http.Request) {
	utils.WrapAPISuccess(w, r, "success", http.StatusOK)

}
