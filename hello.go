package hello

import (
	"fmt"

	cli "github.com/urfave/cli/v2"
	"github.com/tejzpr/zetools/commands"
)

func init() {
	commands.RegisterCommand(HelloCommandName, &helloCommand{}, nil)
}

// HelloCommandName returns the name of the command
const HelloCommandName commands.CommandName = "hello"

// helloCommand is the base64 command
type helloCommand struct {
}

// Name returns the name of the command
func (b *helloCommand) Name() commands.CommandName {
	return HelloCommandName
}

// Aliases returns the aliases of the command
func (b *helloCommand) Aliases() []string {
	return []string{"hello"}
}

// Usage returns the usage of the command
func (b *helloCommand) Usage() string {
	return "Hello"
}

// Subcommands returns the subcommands of the command
func (b *helloCommand) Subcommands() []*cli.Command {
	return []*cli.Command{}
}

// Flags returns the flags of the command
func (b *helloCommand) Flags() []cli.Flag {
	return []cli.Flag{}
}

// Action returns the action of the command
func (b *helloCommand) Action(c *cli.Context) error {
	fmt.Print("Hello!\n")
	return nil
}
