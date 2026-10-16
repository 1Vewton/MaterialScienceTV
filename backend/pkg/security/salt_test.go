package security

import (
	"testing"
)

// TestGenerateSalt tests that the salt generated is unique
func TestGenerateSalt(t *testing.T) {
	t.Parallel()
	salt1, err := GenerateSalt()
	if err != nil {
		t.Error(err)
	}
	if len(salt1) < 1 {
		t.Errorf(
			"salt1 got length of %d",
			len(salt1),
		)
	}
	salt2, err := GenerateSalt()
	if err != nil {
		t.Error(err)
	}
	if len(salt2) < 1 {
		t.Errorf(
			"salt1 got length of %d",
			len(salt2),
		)
	}
	if salt1 == salt2 {
		t.Errorf(
			"salt1 %s and salt2 %s should not be the same",
			salt1,
			salt2,
		)
	}
}

// TestPasswordGeneration tests the generation of passwords
func TestPasswordGeneration(t *testing.T) {
	t.Parallel()
	const password string = "abc114514"
	salt, err := GenerateSalt()
	if err != nil {
		t.Error(err)
	}
	hashVal := ToHash(
		password,
		salt,
	)
	hash2Val := ToHash(
		password,
		salt,
	)
	if hashVal != hash2Val {
		t.Errorf(
			"%s is not equals to %s",
			hashVal,
			hash2Val,
		)
	}
}
