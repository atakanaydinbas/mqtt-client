package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
)

var (
	broker string
	port int
	topic string
	username string
	password string

	tlsEnabled bool
	caCert string
	clientCert string
	clientKey string
	insecureSkipVerify bool
	rootCmd = &cobra.Command{
		Use:   "mqtt-client",
		Short: "MQTT WebSocket Client CLI",
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&broker, "broker", "b", "localhost", "MQTT broker address (default: localhost)")
	rootCmd.PersistentFlags().IntVarP(&port, "port", "p", 9001, "MQTT WebSocket port (default: 9001)")
	rootCmd.PersistentFlags().StringVarP(&topic, "topic", "t", "test/topic", "MQTT topic (default: test/topic)")
	rootCmd.PersistentFlags().StringVar(&username, "username", "", "Username for broker authentication (optional)")
	rootCmd.PersistentFlags().StringVar(&password, "password", "", "Password for broker authentication (optional)")
	rootCmd.PersistentFlags().BoolVar(&tlsEnabled, "tls", false, "Enable TLS (wss://) connection")
	rootCmd.PersistentFlags().StringVar(&caCert, "ca-cert", "", "Path to CA certificate file (optional, for self-signed brokers)")
	rootCmd.PersistentFlags().StringVar(&clientCert, "client-cert", "", "Path to client certificate file (optional, for mTLS)")
	rootCmd.PersistentFlags().StringVar(&clientKey, "client-key", "", "Path to client private key file (optional, for mTLS)")
	rootCmd.PersistentFlags().BoolVar(&insecureSkipVerify, "insecure", false, "Skip server certificate verification (not recommended)")
} 