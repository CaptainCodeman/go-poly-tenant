package main

import (
	"encoding/json"
	"html/template"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
)

type (
	tenantContext struct {
		Tenant *Tenant `json:"tenant"`
		API    string  `json:"api"`
	}

	manifest struct {
		Name            string         `json:"name"`
		ShortName       string         `json:"short_name"`
		Icons           []manifestIcon `json:"icons"`
		StartURL        string         `json:"start_url"`
		Display         string         `json:"standalone"`
		BackgroundColor string         `json:"background_color"`
		ThemeColor      string         `json:"theme_color"`
	}

	manifestIcon struct {
		Src   string `json:"src"`
		Sizes string `json:"sizes"`
		Type  string `json:"type"`
	}
)

var (
	templates = template.Must(template.ParseFiles("index.html"))
)

func init() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/manifest.json", manifestHandler)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	tenant, err := getTenantForDomain(c, r.Host)
	if err != nil {
		if err == datastore.ErrNoSuchEntity {
			log.Warningf(c, "unknown tenant %v", err)
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		} else {
			log.Errorf(c, "error loading tenant %v", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}

	x := &tenantContext{
		Tenant: tenant,
		API:    "https://api-dot-go-poly-tenant.appspot.com",
	}

	if err := templates.ExecuteTemplate(w, "index.html", x); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func manifestHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	tenant, err := getTenantForDomain(ctx, r.Host)
	if err != nil {
		if err == datastore.ErrNoSuchEntity {
			log.Warningf(ctx, "unknown tenant %v", err)
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		} else {
			log.Errorf(ctx, "error loading tenant %v", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}

	m := &manifest{
		Name:      tenant.Name,
		ShortName: tenant.ShortName,
		Icons: []manifestIcon{
			{
				Src:   "/images/manifest/icon-192x192.png",
				Sizes: "192x192",
				Type:  "image/png",
			},
			{
				Src:   "/images/manifest/icon-512x512.png",
				Sizes: "512x512",
				Type:  "image/png",
			},
		},
		StartURL:        "/?homescreen=1",
		Display:         "standalone",
		BackgroundColor: tenant.BackgroundColor,
		ThemeColor:      tenant.ThemeColor,
	}

	w.Header().Set("Content-Type", "application/json")

	enc := json.NewEncoder(w)
	enc.Encode(m)
}
