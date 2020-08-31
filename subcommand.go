package cli

import (
	"flag"
	"fmt"
	"strings"
)

type SubCommand struct {
	Name        string
	Description string
	FlagSet     *flag.FlagSet
	Action      func([]string) error
	Arguments   []string
}

func (self *SubCommand) usage() {
	numberOfOptions := 0
	self.FlagSet.VisitAll(func(fn *flag.Flag) {
		numberOfOptions += 1
	})

	optionsString := ""
	if numberOfOptions > 0 {
		optionsString = "[options] "
	}

	argumentsString := ""
	if len(self.Arguments) > 0 {
		parts := []string{}
		for _, v := range self.Arguments {
			parts = append(parts, fmt.Sprintf("<%s>", v))
		}
		argumentsString = strings.Join(parts, " ")
	}

	fmt.Printf("%s: %s\n", self.Name, self.Description)
	fmt.Printf("Usage: %s %s%s\n", self.Name, optionsString, argumentsString)

	if numberOfOptions > 0 {
		fmt.Println("")
		fmt.Println("Options:")
		self.FlagSet.PrintDefaults()
	}
}
