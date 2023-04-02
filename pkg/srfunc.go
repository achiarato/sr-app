package pkg

import (
	"fmt"
	"net/http"
	"io"
	"encoding/json"
)

const KeyServerAddr = "serverAddr"

type SRdata struct {
     Src   string  `json:"src"`
     Dst   string  `json:"dst"`
     USid  string  `json:"uSid"`
     Query string  `json:"Query"`
}



func GetShortestPathSRuSID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var err error

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
		w.Header().Set("Content-Type", "application/json")
	        err = Newclient()
	        if err != nil {
        	        fmt.Printf("New DB Client creation failed: %s\n", err)
        	}

		srdata := SRdata {
				  Src: src,
				  Dst: dst,
				  USid: "2001",
				  Query: "Shortest Path",
				}
		json.NewEncoder(w).Encode(srdata)
	}
}
