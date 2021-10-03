package sunucu

import (
	"fmt"
	"net/http"
)

//start

func (server *Server) handler(w http.ResponseWriter, r *http.Request) {

	if !server.isKlasor {
		http.ServeFile(w, r, server.kokDizin)
		return
	}

	urlPath := r.URL.Path[1:]
	if urlPath != "" {
		if server.KlasorMu(urlPath) {
			server.KlasorListe(urlPath)
		} else {
			http.ServeFile(w, r, server.kokDizin+urlPath)
			return
		}
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<pre>")
	for _, item := range server.files {

		fmt.Fprintf(w, "<a href=\"%s\">%s</a>\n", item, item)
	}
	fmt.Fprintf(w, "</pre><p>Eray :) | 2021</p>")

}

func (server *Server) HTTPServer() {
	//sunucu HTTP olarak başlatılıyor
	http.HandleFunc("/", server.handler)
	http.Serve(server.listener, nil)
}
