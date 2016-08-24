package main

import (
	"strings"

	"github.com/qedus/nds"
	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
)

type (
	// Tenant is a tenant of the system
	Tenant struct {
		ID              string `json:"id"               datastore:"-"`
		Name            string `json:"name"             datastore:"name,noindex"`
		ShortName       string `json:"short_name"       datastore:"short_name,noindex"`
		Domain          string `json:"domain"           datastore:"domain"`
		ThemeColor      string `json:"theme_color"      datastore:"theme_color,noindex"`
		BackgroundColor string `json:"background_color" datastore:"background_color,noindex"`
	}
)

var rootHost string

func init() {
	if appengine.IsDevAppServer() {
		rootHost = ".127.0.0.1.xip.io:8080"
	} else {
		rootHost = "-dot-go-poly-tenant.appspot.com"
	}
}

func (t *Tenant) save(ctx context.Context) error {
	key := datastore.NewKey(ctx, "tenant", t.ID, 0, nil)
	if _, err := nds.Put(ctx, key, t); err != nil {
		return err
	}
	return nil
}

func getTenantForDomain(ctx context.Context, host string) (*Tenant, error) {
	var key *datastore.Key
	log.Infof(ctx, "host %s", host)
	if strings.HasSuffix(host, rootHost) {
		id := host[:len(host)-len(rootHost)]
		log.Infof(ctx, "tenant id:%s", id)
		key = datastore.NewKey(ctx, "tenant", id, 0, nil)
	} else {
		q := datastore.NewQuery("tenant").Filter("domain =", host).Limit(1).KeysOnly()
		keys, err := q.GetAll(ctx, nil)
		if err != nil {
			return nil, err
		}
		if len(keys) == 0 {
			return nil, datastore.ErrNoSuchEntity
		}
		key = keys[0]
	}

	tenant := new(Tenant)
	if err := nds.Get(ctx, key, tenant); err != nil {
		return nil, err
	}
	tenant.ID = key.StringID()
	return tenant, nil
}
