package main

import (
	"net/http"

	"google.golang.org/appengine"
)

func init() {
	http.HandleFunc("/_ah/warmup", warmupHandler)
}

// create some dummy data at startup
func warmupHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	red := &Tenant{
		ID:              "red",
		Name:            "Red Tenant",
		ShortName:       "Red",
		Domain:          "www.red.com",
		ThemeColor:      "#b53f51",
		BackgroundColor: "#b53f51",
	}
	red.save(ctx)

	green := &Tenant{
		ID:              "green",
		Name:            "Green Tenant",
		ShortName:       "Green",
		Domain:          "www.green.com",
		ThemeColor:      "#3fb551",
		BackgroundColor: "#3fb551",
	}
	green.save(ctx)

	blue := &Tenant{
		ID:              "blue",
		Name:            "Blue Tenant",
		ShortName:       "Blue",
		Domain:          "www.blue.com",
		ThemeColor:      "#3f51b5",
		BackgroundColor: "#3f51b5",
	}
	blue.save(ctx)
}
