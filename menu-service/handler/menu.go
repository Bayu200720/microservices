package handler

import (
	"net/http"

	"github.com/Bayu200720/microservices/utils"
	//"github.com/wskurniawan/digitalent-microservice/utils"
)

func AddMenu(http.ResponseWriter, *http.Request) {
	utils.WrapAPISuccess(w, r, "success", http.StatusOK)

}
