package cli

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func New(name string) *cli {
	c := cli{
		Name:        name,
		subcommands: SubCommands{},
	}
	flag.Usage = c.usage
	return &c
}

type cli struct {
	Name        string
	subcommands SubCommands
}

func (self *cli) usage() {
	lengthOfLongestName := 0
	for _, subcommand := range self.subcommands {
		if len(subcommand.Name) > lengthOfLongestName {
			lengthOfLongestName = len(subcommand.Name)
		}
	}

	fmt.Printf("Usage: %s <sub-command>\n", self.Name)
	fmt.Println("Sub commands:")
	for _, subcommand := range self.subcommands {
		fmt.Printf("  %-*s - %s\n", lengthOfLongestName, subcommand.Name, subcommand.Description)
	}
}

func (self *cli) AddSubCommand(name string, description string) *SubCommand {
	s := &SubCommand{
		Name:        name,
		Description: description,
		FlagSet:     flag.NewFlagSet(name, flag.ExitOnError),
		Arguments:   []Argument{},
	}
	s.FlagSet.Usage = s.usage
	self.subcommands = append(self.subcommands, s)
	return s
}

func (self *cli) Run() int {
	// Make sure there's a subcommand
	if len(os.Args) < 2 {
		fmt.Println("error: valid subcommand expected")
		flag.Usage()
		return 2
	}

	flag.Parse()

	// Get subcommand
	command, found := self.subcommands.Get(os.Args[1])
	if found == false {
		fmt.Println("error: valid subcommand expected")
		flag.Usage()
		return 2
	}

	// Parse flags
	err := command.FlagSet.Parse(os.Args[2:])
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return 2
	}

	// Run the action
	args := command.FlagSet.Args()
	if len(args) != len(command.Arguments) {
		if len(args) > len(command.Arguments) {
			extraArgs := args[len(command.Arguments):]
			fmt.Printf("error: unknown argument(s): %v\n", strings.Join(extraArgs, ", "))
		} else {
			missingArgs := command.argumentNames()[len(args):]
			fmt.Printf("error: missing argument(s): %v\n", strings.Join(missingArgs, ", "))
		}
		command.FlagSet.Usage()
		return 2
	}

	err = command.Action(args)
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return 1
	}

	return 0
}
