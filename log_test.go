package rings

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestLog(t *testing.T) {
	var actual string
	route := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		R.UpdateLog(r, "test0", "test")
		R.UpdateLog(r, "test1", "test")
		R.UpdateLog(r, "test2", "test")

		entry := R.GetLog(r)
		entry.Level = logrus.InfoLevel
		actual, _ = entry.String()
	})

	server := httptest.NewServer(route)
	defer server.Close()

	http.Get(server.URL)

	expected := "time=\"0001-01-01T00:00:00Z\" level=info test0=test test1=test test2=test\n"
	if actual != expected {
		t.Errorf("\nexpected:\n\t%q\ngot\n\t%q\n", expected, actual)
	}
}
