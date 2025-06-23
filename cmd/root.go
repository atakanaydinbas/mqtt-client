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
} 