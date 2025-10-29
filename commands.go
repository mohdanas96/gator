package main

type command struct {
	name string
	args []string
}

type commandHandler func(*state, command) error

type commands struct {
	commandRegistry map[string]commandHandler
}

func (c *commands) run(s *state, cmd command) error {
	fn := c.commandRegistry[cmd.name]
	err := fn(s, cmd)
	if err != nil {
		return err
	}
	return nil
}

func (c *commands) register(name string, f commandHandler) error {
	_, ok := c.commandRegistry[name]
	if !ok {
		c.commandRegistry[name] = f
		return nil
	}
	return nil
}
