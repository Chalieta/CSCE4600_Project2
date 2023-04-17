// Chalieta Audreylia (ca0467)
// University of North Texas
package builtins

import (
	// "errors" // Not used
	"fmt"
	"os"
)

// var (
// 	ErrInvalidArgCount = errors.New("invalid argument count") // Already declared in cd
// )

func MakeDirectory(args ...string) error {
	switch len(args) {
	case 1:
    err := os.Mkdir(args[0], 0750)
    if err != nil && !os.IsExist(err) {
  		return fmt.Errorf("mkdir is unsuccessful")
	  } else {
      println("mkdir is successful")
      return err
    }
	default:
		return fmt.Errorf("%w: expected one argument (new directory name)", ErrInvalidArgCount)
	}
}
