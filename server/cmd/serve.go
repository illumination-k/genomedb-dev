/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"genomedb/ent"
	genomedbv1 "genomedb/protogen/genomedb/v1"
	"genomedb/protogen/genomedb/v1/genomedbv1connect"
	"net/http"

	"log"

	"github.com/bufbuild/connect-go"
	"github.com/spf13/cobra"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type GenomeServer struct{}

func (s *GenomeServer) Get(
	ctx context.Context,
	req *connect.Request[genomedbv1.GenomeGetRequest],
) (*connect.Response[genomedbv1.GenomeGetResponse], error) {
	res := connect.NewResponse(&genomedbv1.GenomeGetResponse{Genome: &genomedbv1.Genome{Name: req.Msg.GenomeName}})
	res.Header().Set("Genomedb-Version", "v1")
	return res, nil
}

func (s *GenomeServer) ListGenomeNames(
	ctx context.Context,
	req *connect.Request[emptypb.Empty],
) (*connect.Response[genomedbv1.ListGenomeNamesResponse], error) {
	res := connect.NewResponse(&genomedbv1.ListGenomeNamesResponse{})
	return res, nil
}

var port int32
var host string

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve gRPC",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := ent.Open("sqlite3", databaseUri)

		if err != nil {
			return err
		}

		defer client.Close()

		if err != nil {
			log.Fatalf("failed listening: %s", err)
		}

		genomeServer := &GenomeServer{}
		mux := http.NewServeMux()
		path, handler := genomedbv1connect.NewGenomeServiceHandler(genomeServer)

		mux.Handle(path, handler)
		http.ListenAndServe(
			fmt.Sprintf("%s:%d", host, port),
			// Use h2c so we can serve HTTP/2 without TLS.
			h2c.NewHandler(mux, &http2.Server{}),
		)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	serveCmd.Flags().StringVarP(&host, "host", "", "localhost", "host")
	serveCmd.Flags().Int32VarP(&port, "port", "p", 50000, "port")
}
