package main

import (
	comm "github.com/6ixbit/Lookup-CLI/cmd"
	"github.com/spf13/cobra"
	"os"
	"log"
)

func main() {
	cmd := &cobra.Command{
		Use:	 "lookup",
		Short: 	 "A CLI network lookup tool",
		Long: 	 `A command line network tool written in Go that can be used 
		to resolve information. `,
	}

	cmd.AddCommand(comm.LookUpDomain())

	if err := cmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}

}