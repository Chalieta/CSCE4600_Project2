//Sussie Manovapillai
//Description: Implementing the ls command.
//4-18-2023 CSCE4600

package builtins       

import (
	"fmt"
	"log"
	"os"
)

func ListFilesDirectory(args ...string) error{
	if len(args) == 0{
		files,err := os.ReadDir(".")
		if err != nil {
                      log.Fatal(err)//log?
		      return fmt.Errorf("No arguement given(ls)")	
		}
		for _, file := range files {
                      fmt.Println(file.Name())
                }
	}else{
		return fmt.Errorf("Expected only one arguement (ls)")
	}
	return fmt.Errorf("")
}
