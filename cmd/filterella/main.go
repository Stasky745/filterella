package main

import (
	"fmt"
	"strings"

	"github.com/alecthomas/kong"
)

var CLI struct {
	Job struct {
		Recursive bool `short:"r" help:"Recursively apply filter." env:"FIL_PROCESS_RECURSIVE"`

		Paths []string `short:"p" name:"path" help:"Paths to apply filters to separated by commas." type:"path" env:"FIL_PROCESS_PATHS"`
	} `cmd:"" help:"Run Filterella to apply filters on a folder."`

	Server struct {
		Webhook struct {
			Url string `optional:"" default:"localhost:8080" help:"URL of the webhook" env:"FIL_WEBHOOK_URL"`
		} `cmd:"" help:"Run the server using a webhook."`

		Polling struct {
			Paths    []string `short:"p" optional:"" name:"path" help:"Paths to list." type:"path" default:"inbox" env:"FIL_POLL_PATHS"`
			Interval int      `short:"i" optional:"" help:"Interval to poll in seconds" default:"3600" env:"FIL_POLL_INTERVAL"`
		} `cmd:"" help:"Run the server using polling."`
	} `cmd:"" help:"Start Filterella as a server." default:"withargs"`

	Config string `short:"c" help:"Path to configuration file." required:"" type:"existingfile" env:"FIL_CONFIG_FILE"`

	IMAPServer   string `help:"IMAP Server." required:"" env:"FIL_IMAP_SERVER"`
	IMAPPort     int    `help:"IMAP Port." required:"" env:"FIL_IMAP_PORT"`
	IMAPUsername string `help:"IMAP Username." required:"" env:"FIL_IMAP_USERNAME"`
	IMAPPassword string `help:"IMAP Password." required:"" env:"FIL_IMAP_PASSWORD"`

	SMTPServer   string `help:"SMTP Server." env:"FIL_SMTP_SERVER"`
	SMTPPort     int    `help:"SMTP Port." env:"FIL_SMTP_PORT"`
	SMTPUsername string `help:"SMTP Username." env:"FIL_SMTP_USERNAME"`
	SMTPPassword string `help:"SMTP Password." env:"FIL_SMTP_PASSWORD"`
}

func main() {
	ctx := kong.Parse(&CLI)
	switch strings.SplitN(ctx.Command(), " ", 2)[0] {
	case "job":
		fmt.Println("job")
	case "server":
		fmt.Println("server")
	default:
		panic(ctx.Command())
	}
}
