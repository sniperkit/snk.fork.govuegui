package govuegui

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRouterWithRice(t *testing.T) {
	useRice = true
	routerTest(t)
	useRice = false
	routerTest(t)
}

func routerTest(t *testing.T) {
	gui := NewGui()
	ts := httptest.NewServer(NewRouter(gui))
	defer ts.Close()
	testUrls := []string{
		"lib/vue.min.js",
		"lib/vue-resource.min.js",
		"lib/vue-router.min.js",
		"lib/pure.min.css",
		"lib/app.js",
		"",
	}
	for _, tURL := range testUrls {
		res, err := http.Get(ts.URL + PathPrefix + "/" + tURL)
		if err != nil {
			t.Error(err)
		}
		if res.StatusCode != 200 {
			t.Errorf("Did not find %s", tURL)
		}
	}
}
