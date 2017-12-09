package hlp

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/rubencaro/cardo/lib/hlp"
)

// AliveLoop takes a folder path, a version string, and a step, and then
// loops forever waiting `step` milliseconds each time.
// On each iteration it replaces the content of an 'alive' file
// inside given folder with the given version string.
// Given folder must exist.
// If the 'alive' file does not exist, it will be created.
//
// This is intended to let an outside force know this program is running,
// and which version is the current one.
//
// This may be called from a new goroutine, as it starts an infinite loop.
//
func AliveLoop(path string, version string, step time.Duration) {
	for {
		_, err := os.Stat(path)
		if os.IsNotExist(err) {
			log.Fatal("Path does not exist: ", path)
		}

		time.Sleep(step * time.Millisecond)

		cmd := fmt.Sprintf("echo \"%s\" > \"%s/alive\"", version, path)
		err2 := hlp.Run(cmd)
		if err2 != nil {
			log.Fatal("Error running: ", cmd, "\nThe error was: ", err2)
		}
	}
}
