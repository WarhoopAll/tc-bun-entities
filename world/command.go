package model

import "github.com/uptrace/bun"

type Command struct {
	Name string `json:"name,omitempty"`
	Help string `json:"help,omitempty"`
}

type CommandSlice []Command

type DBCommand struct {
	bun.BaseModel `bun:"table:command,alias:command"`
	Name string `bun:"name"`
	Help string `bun:"help"`
}

type DBCommandSlice []DBCommand

func (entry *Command) ToDB() *DBCommand {
	if entry == nil {
		return nil
	}
	return &DBCommand{
		Name: entry.Name,
		Help: entry.Help,
	}
}

func (entry *DBCommand) ToWeb() *Command {
	if entry == nil {
		return nil
	}
	return &Command{
		Name: entry.Name,
		Help: entry.Help,
	}
}

func (data CommandSlice) ToDB() DBCommandSlice {
	result := make(DBCommandSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToDB())
	}
	return result
}

func (data DBCommandSlice) ToWeb() CommandSlice {
	result := make(CommandSlice, 0, len(data))
	for _, d := range data {
		result = append(result, *d.ToWeb())
	}
	return result
}
