package user

import (
	"fmt"
)

// ToRegisterToken converts raw token to register token
func ToRegisterToken(
	rawToken string,
) string {
	return fmt.Sprintf(
		"register:%s",
		rawToken,
	)
}
