package main

import (
	"fmt"
	rsp "github.com/tlboright/randomsp"
	"gopkg.in/urfave/cli.v1"
)

func StartCli(args []string) ([]string, error) {
	var resp []string
	app := cli.NewApp()
	app.Name = "randomsp-cli"
	app.Version = "1.1.0"
	app.Usage = "a cli tool to randomly query stocks from a few major indices"
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Tyler Boright",
			Email: "ru.lai.development@gmail.com",
		},
	}

	app.Action = func(c *cli.Context) error {
		resp = append(resp, "rsp: A stock index wasn't chosen\nFor help, type: rsp help\n")
		return nil
	}

	app.Commands = []cli.Command{
		{
			Name:    "random",
			Aliases: []string{"rs"},
			Usage:   "Return a random stock from any of the supported indices",
			Action: func(c *cli.Context) error {
				stock, err := rsp.GetRandomIndexStock()
				if err != nil {
					return fmt.Errorf("randomsp-cli failed with following error: %s. \nCheck your internet connection and try again.\n", err)
				}

				resp = append(resp, fmt.Sprintf("Your randomly selected stock symbol is:\n%s\n\n", stock.Symbol))
				resp = append(resp, fmt.Sprintf("%s is a stock on the %s stock index\n", stock.Symbol, stock.Index))

				return nil
			},
		},

		{
			Name:    "index",
			Aliases: []string{"idx"},
			Usage:   "Return a random stock from a specific index.  Defaults to S&P500.",
			Subcommands: []cli.Command{
				{
					Name:    "SP",
					Aliases: []string{"sp"},
					Usage:   "Returns a random stock from the Standard & Poors 500 index",
					Action: func(c *cli.Context) error {
						stock, err := rsp.GetRandomSPStock()
						if err != nil {
							return fmt.Errorf("randomsp-cli failed with following error: %s. \nCheck your internet connection and try again.\n", err)
						}

						resp = append(resp, fmt.Sprintf("Your randomly selected %s stock symbol is:\n%s\n", stock.Index, stock.Symbol))

						return nil
					},
				},
				{
					Name:    "DAX",
					Aliases: []string{"dax"},
					Usage:   "Returns a random stock from the DAX index",
					Action: func(c *cli.Context) error {
						stock, err := rsp.GetRandomDaxStock()
						if err != nil {
							return fmt.Errorf("randomsp-cli failed with following error: %s. \nCheck your internet connection and try again.\n", err)
						}

						resp = append(resp, fmt.Sprintf("Your randomly selected %s stock symbol is:\n%s\n", stock.Index, stock.Symbol))

						return nil
					},
				},
				{
					Name:    "FinancialTimes",
					Aliases: []string{"ft"},
					Usage:   "Returns a random stock from the Financial Times index",
					Action: func(c *cli.Context) error {
						stock, err := rsp.GetRandomFinancialTimesStock()
						if err != nil {
							return fmt.Errorf("randomsp-cli failed with following error: %s. \nCheck your internet connection and try again.\n", err)
						}

						resp = append(resp, fmt.Sprintf("Your randomly selected %s stock symbol is:\n%s\n", stock.Index, stock.Symbol))

						return nil
					},
				},
				{
					Name:    "Nasdaq",
					Aliases: []string{"nas", "nasdaq"},
					Usage:   "Returns a random stock from the Nasdaq index",
					Action: func(c *cli.Context) error {
						stock, err := rsp.GetRandomNasdaqStock()
						if err != nil {
							return fmt.Errorf("randomsp-cli failed with following error: %s. \nCheck your internet connection and try again.\n", err)
						}

						resp = append(resp, fmt.Sprintf("Your randomly selected %s stock symbol is:\n%s\n", stock.Index, stock.Symbol))

						return nil
					},
				},
				{
					Name:    "Nikkei",
					Aliases: []string{"nik", "nikkei"},
					Usage:   "Returns a random stock from the Nikkei index.\n    Note: Since the Nikkei 225 index uses numbers for symbols, the tool returns the name of the company and its number from the Nikkei 225 index",
					Action: func(c *cli.Context) error {
						stock, err := rsp.GetRandomNikkeiStock()
						if err != nil {
							return fmt.Errorf("randomsp-cli failed with following error: %s. \nCheck your internet connection and try again.\n", err)
						}

						resp = append(resp, fmt.Sprintf("Your randomly selected %s stock symbol is:\n%s\n", stock.Index, stock.Symbol))

						return nil
					},
				},
				{
					Name:    "ItalianFinancialTimes",
					Aliases: []string{"ift"},
					Usage:   "Returns a random stock from the Italian Financial Times index",
					Action: func(c *cli.Context) error {
						stock, err := rsp.GetRandomItalianFinancialTimesStock()
						if err != nil {
							return fmt.Errorf("randomsp-cli failed with following error: %s. \nCheck your internet connection and try again.\n", err)
						}

						resp = append(resp, fmt.Sprintf("Your randomly selected %s stock symbol is:\n%s\n", stock.Index, stock.Symbol))

						return nil
					},
				},
			},
		},
	}
	err := app.Run(args)
	return resp, err
}
