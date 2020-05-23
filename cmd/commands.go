package cmd

import (
	"net"
	"github.com/spf13/cobra"
	"log"
)

// LookUpIP - Returns info about an IP address
func LookUpDomain() *cobra.Command {
	return &cobra.Command{
		Use: "Domain", 
		RunE: func(cmd *cobra.Command, args[] string) error {
			res, err := net.LookupIP(args[0])

			if err != nil {
				log.Println(err)
			}

			cmd.Println(args[0] + " : {")
			cmd.Println("IPv6 Address: ", res[0])
			cmd.Println("IPv4 Address: ", res[1])
			cmd.Println("}")

			return nil
		},
	}
}