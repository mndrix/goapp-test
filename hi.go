package hello

import (
	"fmt"
	"net/http"
)

func SayHello(w http.ResponseWriter, whom string) {
	fmt.Fprintf(w, "Hello, %s", whom)
}
