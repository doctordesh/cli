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
	Arguments   []Argument
}

type Argument struct {
	Name        string
	Description string
}

func (self *SubCommand) AddArgument(name, desc string) {
	self.Arguments = append(self.Arguments, Argument{name, desc})
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
			parts = append(parts, fmt.Sprintf("<%s>", v.Name))
		}
		argumentsString = strings.Join(parts, " ")
	}

	fmt.Printf("%s\n", self.Description)
	fmt.Printf("Usage: %s %s%s\n", self.Name, optionsString, argumentsString)

	if len(self.Arguments) > 0 {
		fmt.Println("")
		fmt.Println("Arguments:")
		for _, v := range self.Arguments {
			fmt.Printf("  %s\n", v.Name)
			fmt.Printf("\t%s\n", v.Description)
		}
	}

	if numberOfOptions > 0 {
		fmt.Println("")
		fmt.Println("Options:")
		self.FlagSet.PrintDefaults()
	}
}

func (self *SubCommand) argumentNames() []string {
	names := []string{}
	for _, v := range self.Arguments {
		names = append(names, v.Name)
	}
	return names
}
