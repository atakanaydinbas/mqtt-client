package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"mqtt-client/pkg/mqtt"
)

func init() {
	rootCmd.AddCommand(subscribeCmd())
}

func subscribeCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "subscribe",
		Short: "Subscribe to a topic",
		Run: func(cmd *cobra.Command, args []string) {
			client, err := mqtt.NewClient(broker, port, username, password, tlsEnabled, caCert, clientCert, clientKey, insecureSkipVerify)
			if err != nil {
				fmt.Println("Failed to connect:", err)
				os.Exit(1)
			}
			fmt.Println("Connected to MQTT broker")
			err = client.Subscribe(topic, func(topic, payload string) {
				fmt.Printf("Received message on %s: %s\n", topic, payload)
			})
			if err != nil {
				fmt.Println("Subscribe error:", err)
				os.Exit(1)
			}
			fmt.Printf("Subscribed to topic: %s\n", topic)
			select {}
		},
	}
} 