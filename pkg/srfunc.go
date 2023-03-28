package pkg

import (
	"fmt"
	"net/http"
	"io"

)

const KeyServerAddr = "serverAddr"

func GetShortestPathSRuSID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	src := r.URL.Query().Get("src")
	dst := r.URL.Query().Get("dst")

	if r.URL.Query().Has("src") == false || src == ""  {
		fmt.Printf("%s: got WRONG /shortestpath request missing source node's address\n", ctx.Value(KeyServerAddr))
		w.Header().Set("x-missing-field", "src")
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "Bad Request missing source node address\n")
	} else if r.URL.Query().Has("dst") == false || dst == "" {
		fmt.Printf("%s: got WRONG /shortestpath request missing destination node's address\n", ctx.Value(KeyServerAddr))
		w.Header().Set("x-missing-field", "dst")
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "Bad Request missing destination node's address\n")
	} else {
		fmt.Printf("%s: got CORRECT /shortestpath request for src '%s' and dst '%s'\n", ctx.Value(KeyServerAddr), src, dst)
		srt := fmt.Sprintf("uSID for the shortest path between src '%s' and dst '%s' is being calculated\n", src, dst)
		io.WriteString(w, srt)
	}
}
