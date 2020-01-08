package strategy

import "fmt"

//ConsolSquare not need inner field ,it will always print the word Square to the console
type ConsolSquare struct {
}

func (c *ConsolSquare) Print() error {

	fmt.Println("Square")
	return nil
}
