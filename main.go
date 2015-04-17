package main

import (
    "os"
    "fmt"
    "log"

    "github.com/amitkgupta/library/mymath"
    "github.com/codegangsta/cli"
)

func main() {
    app := cli.NewApp()
    app.Name = "MyMath CLI"
    app.Usage = "Do some funky math stuff!"

    app.Commands = []cli.Command{
        {
            Name:      "pi",
            Usage:     "calculate π",
            Flags: []cli.Flag{
                cli.IntFlag{
                    Name: "n",
                    Value: -1,
                    Usage: "number of terms in infinite series to calculate",
                },
            },
            Action: func(c *cli.Context) {
                n := c.Int("n")
                if n < 1 {
                    log.Fatalf("Invalid n '%d'", n)
                }

                fmt.Fprintf(os.Stdout, "%.9f\n", mymath.Pi(n))
            },
        },
    }

    app.Run(os.Args)
}
