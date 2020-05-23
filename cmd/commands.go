package cmd

import (
	"net"
	"github.com/spf13/cobra"
	"log"
)

// LookUpDomain - Returns IP for domain
func LookUpDomain() *cobra.Command {
	return &cobra.Command{
		Use: 		"domain",
		ValidArgs:  []string{"www.google.com"},
		RunE: func(cmd *cobra.Command, args[] string) error {
			domain := args[0]
			res, err := net.LookupIP(domain)

			if len(res) == 0 { log.Println(err) }

			cmd.Println(args[0] + " : {")
			cmd.Println("IPv6 Address: ", res[0])
			cmd.Println("IPv4 Address: ", res[1])
			cmd.Println("}")

			return nil
		},
	}
}

// LookUpNameServices - Looks up the name services for a particular domain
func LookUpNameServices() *cobra.Command {
	return &cobra.Command{
		Use: "ns", 
		RunE: func(cmd *cobra.Command, args[] string) error {
			domain := args[0]
			res, _ := net.LookupNS(domain)

			if len(res) == 0 { log.Println("Failed to lookup name services for: ", domain) }

			for _, nameService := range  res {
				cmd.Printf("%s\n", nameService.Host)
			}
			return nil // Return nil as there were no errors.
		},
	}
}





