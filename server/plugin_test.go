package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"os"
	"path/filepath"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/mattermost/mattermost-server/v5/plugin/plugintest"
	
)

func TestServeHTTP(t *testing.T) {
	htmlHeaders := http.Header{"Content-Type": []string{"text/html; charset=utf-8"}, "X-Content-Type-Options": []string{"nosniff"}}

	for name, test := range map[string]struct {
		RequestURL         string
		ExpectedStatusCode int
		ExpectedHeader     http.Header
		HTMLFile		   string
	}{
		"LandingPage": {
			RequestURL:         "/hello",
			ExpectedStatusCode: http.StatusOK,
			ExpectedHeader:     htmlHeaders,
			HTMLFile: helloFileLocation,
		},
		"ErrorPage": {
			RequestURL:         "/error",
			ExpectedStatusCode: http.StatusOK,
			ExpectedHeader:     htmlHeaders,
			HTMLFile: errorFileLocation,
		},
		"ThanksPage": {
			RequestURL:         "/thanks",
			ExpectedStatusCode: http.StatusOK,
			ExpectedHeader:     htmlHeaders,
			HTMLFile: thanksFileLocation,
		},

	} {
		t.Run(name, func(t *testing.T) {
			assert := assert.New(t)

			plugin := &Plugin{}

			currentDir, err := os.Getwd()

			if err != nil {
				t.Fatal("Unable to get current working directory")
			}

			bundlePath := filepath.Dir(currentDir)

			expectedBodyBytes, err := ioutil.ReadFile(filepath.Join(bundlePath, "assets", test.HTMLFile))
			if err != nil {
				t.Fatal(err.Error())
			}
			ExpectedBodyString := string(expectedBodyBytes)

			api := &plugintest.API{}
			api.On("GetBundlePath").Return(bundlePath, nil)
			defer api.AssertExpectations(t)
			plugin.SetAPI(api)

			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", test.RequestURL, nil)
			plugin.ServeHTTP(nil, w, r)

			result := w.Result()
			require.NotNil(t, result)

			bodyBytes, err := ioutil.ReadAll(result.Body)
			require.Nil(t, err)
			bodyString := string(bodyBytes)

			assert.Equal(ExpectedBodyString, bodyString)
			assert.Equal(test.ExpectedStatusCode, result.StatusCode)
			assert.Equal(test.ExpectedHeader, result.Header)
		})
	}
}