package service

import (
	"fmt"
	"sort"
	"strings"

	"github.com/urfave/cli/v2"
)

func Command() *cli.Command {
	return &cli.Command{
		Name:    "service",
		Aliases: []string{"s"},
		Usage:   "query stark service",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "registry",
				Usage:   "registry name [mdns,etcd,consul]",
				Value:   "mdns",
				EnvVars: []string{"STARK_CTL_REGISTRY"},
			},
			&cli.StringFlag{
				Name:    "registry_addr",
				Usage:   "registry host:port",
				EnvVars: []string{"STARK_CTL_REGISTRY_ADDR"},
			},
			&cli.StringFlag{
				Name:     "name",
				Usage:    "service name",
				EnvVars:  []string{"STARK_CTL_NAME"},
				Required: true,
			},
		},
		Action: serviceAction,
	}
}

func serviceAction(c *cli.Context) error {
	registry := c.String("registry")
	registryAddr := c.String("registry_addr")
	name := c.String("name")

	rg, err := newRegistry(registry, registryAddr)
	if err != nil {
		return fmt.Errorf("registry error: %v", err)
	}

	service, err := rg.GetService(name)
	if err != nil {
		return fmt.Errorf("service error: %v", err)
	}

	for _, s := range service {
		fmt.Printf("%s %s\n", s.Name, s.Version)

		for _, node := range s.Nodes {
			fmt.Printf("\t%s %s %s\n", node.Id, node.Address, mapToString(node.Metadata))
		}
	}
	return nil
}

func mapToString(kv map[string]string) string {
	var ret []string
	for k, v := range kv {
		ret = append(ret, fmt.Sprintf("%s:%s", k, v))
	}

	sort.Strings(ret)

	return strings.Join(ret, ",")
}
