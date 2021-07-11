package main

import (
	"bufio"
	"fmt"
	"github.com/sr-2020/gateway/app/adapters/config"
	"github.com/sr-2020/gateway/app/domain"
	"github.com/sr-2020/gateway/app/usecases"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	cfg := config.LoadConfigFromFile("./config.yaml")

	jswUsecase := usecases.GenerateJwt{
		Secret: cfg.JwtSecret,
	}

	app := &cli.App{
		Name: "jwt",
		Usage: "Generate jwt",
		Flags: []cli.Flag {
			&cli.StringFlag{
				Name: "host",
				Value: "gateway.evarun.ru",
				Usage: "Host",
			},
			&cli.StringFlag{
				Name: "method",
				Value: "get",
				Usage: "Method",
			},
			&cli.StringFlag{
				Name: "scopes",
				Value: "auth",
				Usage: "Scopes",
			},
		},
		Action: func(c *cli.Context) error {
			scanner := bufio.NewScanner(os.Stdin)

			tokens := make([]domain.Payload, 0)
			for scanner.Scan() {
				t := strings.Split(scanner.Text(), ",")

				modelId, err := strconv.Atoi(t[0])
				if err != nil {
					return err
				}

				tokens = append(tokens, domain.Payload{
					Auth: t[1],
					ModelId: modelId,
					Exp: 1668149976,
				})
			}

			if err := scanner.Err(); err != nil {
				return err
			}

			request := usecases.GenerateJwtRequest{
				Host: c.String("host"),
				Method: c.String("method"),
				Scopes: strings.Split(c.String("scopes"), ","),
				Tokens: tokens,
			}

			resp, err := jswUsecase.Execute(request)
			if err != nil {
				return err
			}

			fmt.Println(resp.Result)
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
