package scheduler

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/clients/types"
	"github.com/edgexfoundry/go-mod-core-contracts/clients/unmarshaler"
	"github.com/edgexfoundry/go-mod-core-contracts/models"
)

var testID1 = "one"
var testID2 = "two"
var testInterval1 = models.Interval{
	Timestamps: models.Timestamps{
		Created:  123,
		Modified: 123,
		Origin:   123,
	},
	ID:        testID1,
	Name:      "testName",
	Start:     "20060102T150405",
	End:       "20070102T150405",
	Frequency: "24h",
	Cron:      "1",
	RunOnce:   false,
}
var testInterval2 = models.Interval{
	Timestamps: models.Timestamps{
		Created:  321,
		Modified: 321,
		Origin:   321,
	},
	ID:        testID2,
	Name:      "testNombre",
	Start:     "20080102T150405",
	End:       "20090102T150405",
	Frequency: "48h",
	Cron:      "10",
	RunOnce:   false,
}

var UnmarshalToError = unmarshaler.NewCallback(func(_ []byte, _ interface{}) error {
	return errors.New("expected error")
})

type MockEndpoint struct {
}

func (e MockEndpoint) Monitor(params types.EndpointParams) chan string {
	return make(chan string, 1)
}

// testHttpServer instantiates a test HTTP Server to be used for conveniently verifying a client's invocation
func testHttpServer(t *testing.T, matchingRequestMethod string, matchingRequestUri string) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		if r.Method == matchingRequestMethod && r.RequestURI == matchingRequestUri {
			w.Write([]byte("Ok"))
		} else if r.Method != matchingRequestMethod {
			t.Fatalf("expected method %s to be invoked by client, %s invoked", matchingRequestMethod, r.Method)
		} else if r.RequestURI == matchingRequestUri {
			t.Fatalf("expected endpoint %s to be invoked by client, %s invoked", matchingRequestUri, r.RequestURI)
		}
	}))
}
