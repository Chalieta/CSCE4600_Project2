// Sakchham Sangroula (SS1960)
// University of North Texas
package builtins

import (
	// importing package for printing the errors and the output
	"fmt"
)



func Echo(args ...string) error {
	switch len(args) {
		//if no argument given, error is thrown
	case 0:
		return fmt.Errorf("Expected one argument")
	default:
		// Loop over all the arguments and print them
		for i := 0; i < len(args); i++ {
			fmt.Print(string(args[i]))
			// Add a space after each one of the word except for the last word
			if (i < len(args)-1){
				fmt.Print(string(" "))
			}
		}
		return fmt.Errorf("")
	}
}
