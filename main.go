package main

import (
        "errors"
	"context"
        "fmt"
        "net/http"
        "os"
	"net"

	"github.com/achiarato/sr-app/pkg"
)

const keyServerAddr = "serverAddr"

func main() {
	fmt.Println("SR-App Server start listening on :3333")
	http.HandleFunc("/shortestpath", pkg.GetShortestPathSRuSID)

	ctx := context.Background()

	Server := &http.Server{
		Addr:    ":3333",
		BaseContext: func(l net.Listener) context.Context {
			ctx = context.WithValue(ctx, keyServerAddr, l.Addr().String())
			return ctx
		},
	}

	err := Server.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
