//Leif Messinger
//Description: Implementing the cat command.
//5-10-2023 CSCE4600

package builtins

import (
	"bufio"	//scanner
	"fmt"
	"log"
	"os"
)

func ConcatenateFiles(args ...string) error{
	if len(args) == 0{
		//Redirect user input into user output until ctrlD
		//That means I gotta handle the EOF signal

		//Cat program done here https://stackoverflow.com/a/45914294/10141528
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			fmt.Println(scanner.Text()) // Println will add back the final '\n'
		}
		return scanner.Err()
	}else{
		//Concatenate all files in args
		for i := 0; i < len(args); i++ {
			data,err := os.ReadFile(args[i])
			//From documentation https://pkg.go.dev/os#ReadFile
			os.Stdout.Write(data)
			if err != nil {
				log.Fatal(err)
			}
		}
		return nil
	}
	return nil	//Shouldn't get here in theory
}
