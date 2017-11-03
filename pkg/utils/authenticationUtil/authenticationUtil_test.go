package authenticationUtil

import "testing"

func TestCreateToken(t *testing.T) {
	fn := CreateToken

	if val, _ := fn(); len(val) != 32 {
		t.Fatalf("The token format is incorrect : %s", val)
	}
}
