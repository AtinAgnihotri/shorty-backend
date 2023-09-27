package main

import (
	"fmt"
	"net/http"

	"github.com/AtinAgnihotri/shorty-backend/helpers"
	"github.com/go-chi/chi/v5"
)

const (
	SHORT_URL_PARAM = "shortUrlBlob"
)

func (cfg *ServerConf) RedirectHandler(w http.ResponseWriter, r *http.Request) {
	shortLinkBlob := chi.URLParam(r, SHORT_URL_PARAM)

	if len(shortLinkBlob) == 0 {
		helpers.RespondWithError(w, http.StatusBadRequest, "Need to have a shortend url blob at the end")
		// TODO handle with react error
		// http.Redirect(w, r, "/404", http.StatusMovedPermanently)
		return
	}

	data, err := cfg.DB.GetLinkByBlob(r.Context(), shortLinkBlob)
	if err != nil {
		helpers.RespondWithError(w, http.StatusNotFound, fmt.Sprintf("Unable to find corresponding link for %v", shortLinkBlob))
		// TODO handle with react 404
		// http.Redirect(w, r, "/404", http.StatusMovedPermanently)
		return
	}

	http.Redirect(w, r, data.LongLink, http.StatusMovedPermanently)
}
