package handlers

import "net/http"

//RootHandlers handel the routing of root
func RootHandlers(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("not Found\n"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Running API V1\n"))
}
