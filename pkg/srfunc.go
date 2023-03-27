package pkg

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func GetSRuSID(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "It works!\n")
}
