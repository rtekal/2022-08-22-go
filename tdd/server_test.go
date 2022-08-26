/* Alta3 Research | RZFeeser
   TDD - Building tests to drive feature development.
   This test should be written BEFORE the feature is implimented.

   Functions need testing. Before writing the function (PlayerServer), we'll define
   a test.

   Here we expect to get a plyers score (Pepper), which we expect to be 20.

   A narrow test such as this is "OK". */

package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// test the function PlayerServer
func TestGETPlayers(t *testing.T) {
	//
	t.Run("returns Pepper's score", func(t *testing.T) {

		// create a request           method,         path,              body
		request, _ := http.NewRequest(http.MethodGet, "/players/Pepper", nil)
		response := httptest.NewRecorder() // from the testing package, this "spy" lets us inspect what is in a response

		PlayerServer(response, request) // this is the function we are trying to test

		got := response.Body.String() // what did we get back?
		want := "20"                  // this is what we expect to get back

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("returns Floyd's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/Floyd", nil)
		response := httptest.NewRecorder()

		PlayerServer(response, request)

		got := response.Body.String()
		want := "10"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

}
