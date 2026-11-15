package ids

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

// NewUniqueString creates new unique string
func NewUniqueString(
	moduleName string,
) string {
	id := uuid.NewString()
	return fmt.Sprintf(
		"%s:%s:%d",
		moduleName,
		id,
		time.Now().UnixNano(),
	)
}
