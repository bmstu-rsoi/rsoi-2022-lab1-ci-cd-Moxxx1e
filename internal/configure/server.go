package configure

import (
	"fmt"
	"os"
)

func GetConnString() string {
	return fmt.Sprintf(":%s", os.Getenv("PORT"))
}
