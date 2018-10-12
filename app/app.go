package main
import (
	"fmt"
	"log"
	"net/http"
	"os"
)
var version string
var color string
func mainHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL.Path, r.RemoteAddr, r.UserAgent())
	if r.URL.Path == "/color" {
		fmt.Fprintf(w, `<html><body style="background-color: %s">Version: %s<br>Hostname: %s</body></html>`, color, version, os.Getenv("HOSTNAME"))
	} else {
		w.Header().Set("Content-Type","text/plain")
		fmt.Fprintf(w, "Version: %s\nHostname: %s\n", version, os.Getenv("HOSTNAME"))
	}
}
func healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL.Path, r.RemoteAddr, r.UserAgent())
	w.Header().Set("Content-Type","text/plain")
	fmt.Fprintf(w, "OK\n")
}
func main() {
	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/healthcheck", healthcheckHandler)
	http.ListenAndServe(":8080", nil)
}
