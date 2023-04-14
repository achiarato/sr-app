package pkg

import (
	"time"
	"bufio"
	"strings"
	"fmt"
	"net/http"
	"io"
	"encoding/json"
	"os"
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
	var uSID string
	var input string

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
		//inputs just to pause the automatic process
                fmt.Println("BREAKPOINT. Press RETURN to continue the process")
                reader := bufio.NewReader(os.Stdin)
                input, err = reader.ReadString('\n')
		_ = err
                if err != nil {
	                fmt.Println("An error occured while reading input. Please try again", err)
                }
                input = strings.TrimSuffix(input, "\n")
                //finisce input
		fmt.Println("Querying ArangoDB")
		time.Sleep(2*time.Second)
		err, uSID = ArangoDBQuery(ctx, src, dst, "Shortest Path")
	        if err != nil {
        	        fmt.Printf("Arango DB Query Failed: %s\n", err)
        	}

		if uSID == "" {
			fmt.Printf("No uSID created. Please double check the query type")
		}
		fmt.Printf("Query to ArangoDB worked! Got the uSID: %s\n", uSID)
		time.Sleep(1*time.Second)
		srdata := SRdata {
				  Src: src,
				  Dst: dst,
				  USid: uSID,
				  Query: "Shortest Path",
				}
		fmt.Printf("Encoding JSON response: %s\n", srdata)
		json.NewEncoder(w).Encode(srdata)
	}
}
