package cmd

import (
	"fss/pkg/http"
	"fss/version"

	"github.com/spf13/cobra"
)

var (
	dsn  string
	addr string
)

func NewRunCMD() *cobra.Command {
	serverCMD := &cobra.Command{
		Use:   "run",
		Short: "run the fss server",
		Run: func(cmd *cobra.Command, args []string) {
			http.InitServer(dsn).Run(addr)
		},
	}
	serverCMD.Flags().StringVarP(&dsn, "dsn", "d", "mongodb://localhost:27017", "the dsn of dbbase")
	serverCMD.Flags().StringVarP(&addr, "addr", "a", ":8888", "the address which server listen on")
	return serverCMD
}

func NewVersionCMD() *cobra.Command {

	versionCMD := &cobra.Command{
		Use:   "version",
		Short: "show the version of binary",
		Run: func(cmd *cobra.Command, args []string) {
			version.Print()
		},
	}
	return versionCMD
}
