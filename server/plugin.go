package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"sync"

	"github.com/mattermost/mattermost-server/v5/plugin"
)

const (
	helloFileLocation  = "hello_default.html"
	errorFileLocation  = "error_default.html"
	thanksFileLocation = "thanks_default.html"
)

// Plugin implements the interface expected by the Mattermost server to communicate between the server and plugin processes.
type Plugin struct {
	plugin.MattermostPlugin

	// configurationLock synchronizes access to the configuration.
	configurationLock sync.RWMutex

	// configuration is the active plugin configuration. Consult getConfiguration and
	// setConfiguration for usage.
	configuration *configuration
}

// ServeHTTP accepts HTTP requests to the plugin and routes them accordingly
func (p *Plugin) ServeHTTP(c *plugin.Context, w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		p.handleGet(w, r)
	case http.MethodPost:
		http.Error(w, "You didn't say the magic word.", 400)
	}
}

// See https://developers.mattermost.com/extend/plugins/server/reference/

func (p *Plugin) handleGet(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/hello":
		p.handleLanding(w, r)
	case "/error":
		p.handleError(w, r)
	case "/thanks":
		p.handleThanks(w, r)
	default:
		http.NotFound(w, r)
	}
}

func (p *Plugin) handlePost(w http.ResponseWriter, r *http.Request) {

}

func (p *Plugin) handleLanding(w http.ResponseWriter, r *http.Request) {
	p.handleHTML(helloFileLocation, w, r)
}

func (p *Plugin) handleError(w http.ResponseWriter, r *http.Request) {
	p.handleHTML(errorFileLocation, w, r)
}

func (p *Plugin) handleThanks(w http.ResponseWriter, r *http.Request) {
	p.handleHTML(thanksFileLocation, w, r)
}

func (p *Plugin) handleHTML(filename string, w http.ResponseWriter, r *http.Request) {
	bundlePath, errGetBundlePath := p.API.GetBundlePath()
	if errGetBundlePath != nil {
		http.Error(w, "Error getting plugin bundle path", 500)
		p.API.LogError(errGetBundlePath.Error())
	}

	landingHTML, errReadFile := ioutil.ReadFile(filepath.Join(bundlePath, "assets", filename))

	if errReadFile != nil {
		http.Error(w, fmt.Sprintf("Error reading %s", filename), 500)
		p.API.LogError(errReadFile.Error())
	}

	w.Header().Add("Content-Type", "text/html; charset=utf-8")
	w.Header().Add("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, string(landingHTML))
}
