package config

import(
	"os"
)

func GetNodeName() (string) {
	name := "miner" // default name
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	return name
}