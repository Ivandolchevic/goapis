package basicAuthenticationHandler

import (
	"testing"
	"net/http"
)

func TestGetAll(t *testing.T) {
	fn := GetAll

	var w := ResponseWriter()
	var r := Request()

	if val, _ := fn(w,r); len(val) != 32 {
		t.Fatalf("The token format is incorrect : %s", val)
	}
}
