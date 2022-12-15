/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"genomedb/ent"
	"genomedb/ent/proto/entpb"
	"log"
	"net"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var port int32

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

		svc := entpb.NewTranscriptService(client)
		server := grpc.NewServer()

		entpb.RegisterTranscriptServiceServer(server, svc)

		lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
		reflection.Register(server)

		if err != nil {
			log.Fatalf("failed listening: %s", err)
		}

		// トラフィックを無期限にリッスン
		if err := server.Serve(lis); err != nil {
			log.Fatalf("server ended: %s", err)
		}
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

	serveCmd.Flags().Int32VarP(&port, "port", "p", 50000, "port")
}
