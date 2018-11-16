package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"regexp"
	"strings"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "nslookupper"
	app.Usage = "github.com/akm/nslookupper"
	app.UsageText = app.Name + " [GLOBAL OPTIONS] HOST_NAME"
	app.Version = Version

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "name-server",
			Value: "8.8.8.8",
			Usage: "Name server to look up",
		},
	}
	app.Action = process
	app.Run(os.Args)
}

var PatternInclude = regexp.MustCompile(`\Ainclude:`)
var PatternIp4 = regexp.MustCompile(`\Aip4:`)

func digDomain(ctx context.Context, resolver *net.Resolver, domain string) ([]string, error) {
	records, err := resolver.LookupTXT(ctx, domain)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to LookupTXT %q because of %v\n", domain, err)
		return nil, err
	}

	result := []string{}
	for _, r := range records {
		parts := strings.Split(r, " ")
		for _, part := range parts {
			if PatternInclude.MatchString(part) {
				name := PatternInclude.ReplaceAllString(part, "")
				ips, err := digDomain(ctx, resolver, name)
				if err != nil {
					return nil, err
				}
				result = append(result, ips...)
			} else if PatternIp4.MatchString(part) {
				result = append(result, PatternIp4.ReplaceAllString(part, ""))
			}
		}
	}
	return result, nil
}

func process(c *cli.Context) error {
	if c.NArg() < 1 {
		cli.ShowAppHelp(c)
		os.Exit(1)
		return nil
	}

	resolver := &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			d := net.Dialer{}
			return d.DialContext(ctx, "udp", net.JoinHostPort(c.String("name-server"), "53"))
		},
	}

	for _, arg := range c.Args() {
		ips, err := digDomain(context.Background(), resolver, arg)
		if err != nil {
			return err
		}
		for _, ip := range ips {
			fmt.Fprintln(os.Stdout, ip)
		}
	}
	return nil
}
