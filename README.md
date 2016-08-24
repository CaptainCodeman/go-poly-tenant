# Go Polymer MultiTenancy

An example of using server-side Go templating to serve a multi-tenanted
Polymer client using AppEngine. Each client could be served from it's own
custom domain or using a CNAME from a root service (e.g. tenant.service.com).

Examples:

* https://red-dot-go-poly-tenant.appspot.com/
* https://green-dot-go-poly-tenant.appspot.com/
* https://blue-dot-go-poly-tenant.appspot.com/

The dummy tenants are generated via the `/app/warmup.go` file.

Yes, those are some funky URLs, the '-dot-' part is todo with how AppEngine
provides HTTPS for appspot.com domains. In reality you'd use a custom domain
so you would have tenant URLs like red.myservice.com or www.red.com.

The index.html page and manifest.json file is generated dynamically for the
domain it's being served from. All other static files are served by the
AppEngine frontend edge cache (like a CDN).

Creating the index.html page from a template saves having to download the app
and then make a separate request to get the runtime config for the tenant. We
want fast and zippy and no waiting to render. This could also be used to add
additional server-side meta-data for SEO, open-graph etc... by adding a little
more server-side routing to the mix.

The app config and tenant data (for use within the app) is set to the MyApp
global object. It can contain additional config settings such as api, auth or
image endpoints.

## Dependencies

You need to have

* Go SDK for AppEngine
* Polymer CLI

## Setup

The `/app` folder contains the AppEngine application and also serves the polymer
frontend SPA from a static folder which uses a symlink to point to the regular
Polymer-CLI build folder. Create it in the `/app` folder using:

    $ ln -s ../build/unbundled static

## Run Locally

Build the polymer app using `./build.sh`. This runs `polymer build` with the
parameters to include the additonal files required (images and the webcomponent
polyfill).

Start the app locally by running `goapp serve` within the `/app` folder

Go to one of the tenant sites:

* http://red.127.0.0.1.xip.io:8080/
* http://green.127.0.0.1.xip.io:8080/
* http://blue.127.0.0.1.xip.io:8080/

xip.io is a wildcard DNS service that can be used when developing locally.

## Develop Locally

I normally have the root index.html page set to use a dev tenant configuration.