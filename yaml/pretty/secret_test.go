package pretty

import "testing"

func TestSecret(t *testing.T) {
	ok, err := diff("testdata/secret.yml")
	if err != nil {
		t.Error(err)
	} else if !ok {
		t.Errorf("Unepxected formatting")
	}
}

func TestExternalSecret(t *testing.T) {
	ok, err := diff("testdata/secret_extern.yml")
	if err != nil {
		t.Error(err)
	} else if !ok {
		t.Errorf("Unepxected formatting")
	}
}
