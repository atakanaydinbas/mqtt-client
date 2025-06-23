package cmd

import (
	"fmt"
	"os"
	"time"
	"github.com/spf13/cobra"
	"mqtt-client/pkg/mqtt"
)

var message string

func init() {
	publishCmd := publishCmd()
	publishCmd.Flags().StringVarP(&message, "message", "m", "", "Message to publish (overrides timestamp loop)")
	rootCmd.AddCommand(publishCmd)
}

func publishCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "publish [message] [flags]",
		Short: "Publish a message or the current timestamp every 10 seconds. Flags: --broker, --port, --topic, --message/-m",
		Args:  cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			client, err := mqtt.NewClient(broker, port)
			if err != nil {
				fmt.Println("Failed to connect:", err)
				os.Exit(1)
			}
			fmt.Println("Connected to MQTT broker")

			// Priority: positional arg > --message flag > timestamp loop
			var msg string
			if len(args) > 0 {
				msg = args[0]
			} else if message != "" {
				msg = message
			}

			if msg != "" {
				fmt.Printf("Publishing: %s\n", msg)
				err := client.Publish(topic, msg)
				if err != nil {
					fmt.Println("Publish error:", err)
				}
				return
			}

			for {
				now := time.Now()
				timestamp := now.Format("2006-01-02 15:04:05.") + fmt.Sprintf("%03d", now.Nanosecond()/1e6)
				fmt.Printf("Publishing: %s\n", timestamp)
				err := client.Publish(topic, timestamp)
				if err != nil {
					fmt.Println("Publish error:", err)
				}
				time.Sleep(10 * time.Second)
			}
		},
	}
} 