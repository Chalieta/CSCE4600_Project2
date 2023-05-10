//Herbert Flowers IV 
//CSCE 4600.
package builtins

import (
	"fmt"
	"os"
)

func OpenFileDirectory(args ...string) error {
	switch len(args) {
	case 1:
    f, err := os.OpenFile(args[0], os.O_RDONLY, 0)
    if err != nil {
    	return fmt.Errorf("Open is unsuccessful")
	}
    if err := f.Close(); err != nil {
		return fmt.Errorf("Open is unsuccessful")
	}
	println("Open is successful")
	return err

	default:
		return fmt.Errorf("%w: expected one argument (file name)", ErrInvalidArgCount)
	}
}
