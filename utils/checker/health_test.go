package health

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealth(t *testing.T) {
	healthController := Create()

	mux := http.NewServeMux()
	mux.HandleFunc(`/health`, healthController.Get)

	srv := httptest.NewServer(mux)
	defer srv.Close()
	cli := srv.Client()

	req, err := http.NewRequest("GET", srv.URL+"/health", nil)
	if err != nil {
		t.Error(err)
	}

	resp, err := cli.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	// Ignore the possible error; Body should never be nil at this point.
	//noinspection GoUnhandledErrorResult
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}

	t.Run("Check output", func(t *testing.T) {
		var response Response
		err = json.Unmarshal(bytes, &response)
		if err != nil {
			t.Error(err)
		}

		if response.Status != true {
			t.Error("Unsuccessful response")
			t.Logf("%#v", response.Data)
		}
	})
}
