package main

import (
	comm "github.com/6ixbit/Lookup-CLI/cmd"
	"github.com/spf13/cobra"
	"os"
	"log"
)

func main() {
	// Parent command for program
	cmd := &cobra.Command{
		Use:	 "./Lookup-CLI",
		Short: 	 "A CLI network lookup tool",
		Long: 	 `A command line network tool written in Go that can be used 
		to resolve information. `,
	}

	// Attach subcommands to main parent command
	cmd.AddCommand(comm.LookUpDomain())
	cmd.AddCommand(comm.LookUpNameServices())

	if err := cmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}

}