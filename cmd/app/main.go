package main

import (
	"fmt"
	"net/http"

	"github.com/pwnsbd/footsteps/internal/router"
)

func main() {
	var router router.Router
	fmt.Println("Started the server on :3000...")
	http.ListenAndServe(":3000", router)
}
