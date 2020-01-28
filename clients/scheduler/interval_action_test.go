package scheduler

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/clients"
	"github.com/edgexfoundry/go-mod-core-contracts/clients/interfaces"
	"github.com/edgexfoundry/go-mod-core-contracts/clients/types"
	"github.com/edgexfoundry/go-mod-core-contracts/models"
)

var testIntervalAction1 = models.IntervalAction{
	ID:         testID1,
	Created:    123,
	Modified:   123,
	Origin:     123,
	Name:       "testName",
	Interval:   "123",
	Parameters: "123",
	Target:     "testNombre",
	Protocol:   "123",
	HTTPMethod: "get",
	Address:    "localhost",
	Port:       2700,
	Path:       "123",
	Publisher:  "123",
	User:       "123",
	Password:   "123",
	Topic:      "123",
}

var testIntervalAction2 = models.IntervalAction{
	ID:         testID2,
	Created:    321,
	Modified:   321,
	Origin:     321,
	Name:       "testNombre",
	Interval:   "321",
	Parameters: "321",
	Target:     "testName",
	Protocol:   "321",
	HTTPMethod: "post",
	Address:    "127.0.0.1",
	Port:       3000,
	Path:       "321",
	Publisher:  "321",
	User:       "321",
	Password:   "321",
	Topic:      "321",
}

func TestIntervalActionRestClient_IntervalAction(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		if r.Method != http.MethodGet {
			t.Fatalf("expected http method is GET, active http method is : %s", r.Method)
		}

		expectedURL := clients.ApiIntervalActionRoute + "/" + testID1
		if r.URL.EscapedPath() != expectedURL {
			t.Fatalf("expected uri path is %s, actual uri path is %s", expectedURL, r.URL.EscapedPath())
		}

		data, err := json.Marshal(testIntervalAction1)
		if err != nil {
			t.Fatalf("marshaling error: %s", err.Error())
		}
		w.Write(data)
	}))

	defer ts.Close()

	var tests = []struct {
		name             string
		IntervalActionID string
		unmarshaler      interfaces.Unmarshaler
		ic               IntervalActionClient
		expectedError    bool
	}{
		{
			"happy path",
			testIntervalAction1.ID,
			nil,
			NewIntervalActionClient(types.EndpointParams{
				ServiceKey:  clients.SupportSchedulerServiceKey,
				Path:        clients.ApiIntervalActionRoute,
				UseRegistry: false,
				Url:         ts.URL + clients.ApiIntervalActionRoute,
				Interval:    clients.ClientMonitorDefault,
			}, MockEndpoint{}),
			false,
		},
		{
			"nil client",
			testIntervalAction1.ID,
			nil,
			NewIntervalActionClient(types.EndpointParams{}, nil),
			true,
		},
		{"bad marshal",
			testIntervalAction1.ID,
			nil,
			NewIntervalActionClientWithUnmarshaler(types.EndpointParams{
				ServiceKey:  clients.SupportSchedulerServiceKey,
				Path:        clients.ApiIntervalActionRoute,
				UseRegistry: false,
				Url:         ts.URL + clients.ApiIntervalActionRoute,
				Interval:    clients.ClientMonitorDefault,
			},
				MockEndpoint{},
				&UnmarshalToError,
			),
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := tt.ic.IntervalAction(tt.IntervalActionID, context.Background())

			emptyIntervalAction := models.IntervalAction{}

			if !tt.expectedError && res == emptyIntervalAction {
				t.Error("unexpected empty response")
			} else if tt.expectedError && res != emptyIntervalAction {
				t.Errorf("expected empty response, was %s", res)
			}

			if !tt.expectedError && err != nil {
				t.Errorf("unexpected error %s", err.Error())
			} else if tt.expectedError && err == nil {
				t.Error("expected error")
			}
		})
	}
}
