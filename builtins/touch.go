// Chalieta Audreylia (ca0467)
// University of North Texas
package builtins

import (
	// "errors" // Not used
	"time"
	"fmt"
	"os"
)

// var (
// 	ErrInvalidArgCount = errors.New("invalid argument count") // Already declared in cd
// )

func TouchFile(args ...string) error {
	switch len(args) {
	case 1:
		fileName := args[0]
		_, err := os.Stat(fileName)
		if os.IsNotExist(err) {
			file, err := os.Create(fileName)
			if err != nil {
				return fmt.Errorf("Touch is unsuccessful")
			}
			defer file.Close()
			println("File created by touch")
		} else {
			currentTime := time.Now().Local()
			err = os.Chtimes(fileName, currentTime, currentTime)
			if err != nil {
				return fmt.Errorf("Touch is unsuccessful")
			}
			println("Timestamp updated by touch")
		}
		return err
	default:
		return fmt.Errorf("%w: expected one argument (file name)", ErrInvalidArgCount)
	}
}
