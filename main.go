package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

var html string = `
<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Crab Life !</title>
</head>

<body>
    <div id="crab" style="margin: auto; text-align: center; padding: 5em;">
        <img src="%s" alt="crab" width="300" />
    </div>
    <div id="crab_story" style="margin: auto; text-align: center; font-size: 2em;">
        <div style="color:chocolate; font-weight: bold;">Hi ! I was born on %s</div>
        <div style="color: cadetblue;">My motherboard's name is <strong>%s</strong></div>
		<div style="color: coral;font-size: 0.5em;">(This page will be kept in cache for %d seconds)</div>
    </div>
</body>

</html>
`

var crabURL *string
var hostname string
var ttl *int

func crab(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Cache-Control", fmt.Sprintf("public, max-age=%d", *ttl))
	fmt.Fprintf(w, html, *crabURL, time.Now().Format(time.Stamp), hostname, *ttl)
}

func healthz(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "i'm alive, this is my crab life !")
}

func main() {
	var err error

	// Get TTL value for cache configuration
	ttl = flag.Int("cache-ttl", 0, "time in second to cache HTTP response")
	crabURL = flag.String("crab-url", "https://i.gifer.com/3QZn.gif", "url to crab gif")
	flag.Parse()

	// Get Hostname of the server running this app
	hostname, err = os.Hostname()
	if err != nil {
		hostname = "unknown"
	}

	http.HandleFunc("/", crab)
	http.HandleFunc("/healthz", healthz)

	fmt.Println("listening to 0.0.0.0:8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
