package handlers

import (
	"net/http"

	"github.com/bongerka/microservice/api/data"
)

func (p *Products) Update(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "application/json")

	prod := r.Context().Value(KeyProduct{}).(data.Product)
	p.l.Println("Updating record id", prod.ID)

	rw.WriteHeader(http.StatusNoContent)
}
