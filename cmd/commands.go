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
		Short: 		"Looks up IPv4 and IPv6 addresses of a given host",
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
		Short: "Looks up the name services for a particular host",
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

// ReverseLookUpAddress - Returns the CNAME for a domain by using 
func ReverseLookUpAddress() *cobra.Command {
	return &cobra.Command {
		Use: 	"rv", 
		Short: "Performs a Reverse Lookup for a given host",
		RunE: func(cmd *cobra.Command, args[] string) error {
			domain := args[0]
			res, err := net.LookupCNAME(domain)

			if err != nil { log.Println(err) }
			cmd.Println(res)

			return nil
		},
	}
}




