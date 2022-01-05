package env

import (
	"os"
)

var (
	PORT string
)

func Load() {
	PORT = os.Getenv("PORT")
}
