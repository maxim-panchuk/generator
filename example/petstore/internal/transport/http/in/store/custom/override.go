package custom

import (
	"net/http"
)

func (c *ControllerImpl) GetOrderById(
	w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("override testing"))
	return
}
