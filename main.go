package pkg

import (
        "errors"
        "fmt"
        "io"
        "net/http"
        "os"

	"github/achiarato/sr-app/pkg"
)

func main() {
	http.HandleFunc("/", pkg.GetSRuSID)

	err := http.ListenAndServe(":3333", nil)
if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
