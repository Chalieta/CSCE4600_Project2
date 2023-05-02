//Herbert Flowers IV 
//CSCE 4600
package builtins

import (
	"log"
	"fmt"
	"os"
)

func OpenFile(args ...string) error {
	switch len(args) {
	case 1:
    err := os.Open("notes.txt", os.O_RDWR|os.O_CREATE, 0755)
    if err != nil {
		log.Fatal(err)
      return fmt.Errorf("open is unsuccessful")
	}
    if err := f.Close(); err != nil {
		log.Fatal(err)
	}
    else {
      println("open is successful")
      return err
    }
	default:
		return fmt.Errorf("%w: expected one argument (openfile name)", ErrInvalidArgCount)
	}
}
