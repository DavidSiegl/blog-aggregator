package main

import (
	"fmt"
	"log"
	"os"

	"blog-aggregator/internal/config"
)

type state struct {
	cfg *config.Config
}

type command struct {
	name string
	args []string
}

type commands struct {
	handlers map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.handlers[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	f, ok := c.handlers[cmd.name]
	if !ok {
		return fmt.Errorf("unknown command: %s", cmd.name)
	}
	return f(s, cmd)
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("login requires a username argument")
	}

	if err := s.cfg.SetUser(cmd.args[0]); err != nil {
		return err
	}
	fmt.Println("User has been set to", cmd.args[0])
	return nil
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	s := &state{cfg: &cfg}

	cmds := commands{handlers: make(map[string]func(*state, command) error)}
	cmds.register("login", handlerLogin)

	if len(os.Args) < 2 {
		log.Fatal("usage: gator <command> [args...]")
	}

	cmd := command{name: os.Args[1], args: os.Args[2:]}
	if err := cmds.run(s, cmd); err != nil {
		log.Fatal(err)
	}
}
