package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli/v2"
)

func GenerateApp() *cli.App {
	flags := []cli.Flag{&cli.StringFlag{
		Name:  "host",
		Usage: "Nome do host para encontrar o IP",
		Value: "localhost",
	}}
	app := &cli.App{
		Name:  "go-cli",
		Usage: "Uma CLI simples em Go",
	}

	app.Commands = []*cli.Command{
		{
			Name:  "hello",
			Usage: "Exibe uma mensagem de saudação",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "name",
					Usage: "Nome da pessoa a ser saudada",
					Value: "Mundo",
				},
			},
			Action: greeting,
		},
		{
			Name:   "ip",
			Usage:  "Encontra o endereço IP de um host",
			Flags:  flags,
			Action: findIps,
		},
		{
			Name:   "server",
			Usage:  "Encontra o nome do servidor de um host",
			Flags:  flags,
			Action: findServers,
		},
	}
	return app
}

func greeting(c *cli.Context) error {
	name := c.String("name")
	if name == "" {
		name = "Mundo"
	}
	println("Olá, " + name + "!")
	return nil
}

func findIps(c *cli.Context) error {
	host := c.String("host")
	if host == "" {
		host = "localhost"
	}
	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}
	for _, ip := range ips {
		fmt.Println("IP:", ip)
	}
	return nil
}

func findServers(c *cli.Context) error {
	host := c.String("host")
	if host == "" {
		host = "localhost"
	}
	servers, err := net.LookupNS(host)
	if err != nil {
		log.Fatal(err)
	}
	for _, server := range servers {
		fmt.Println("SERVER:", server.Host)
	}
	return nil
}
