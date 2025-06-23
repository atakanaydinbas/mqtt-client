package mqtt

import (
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
)

type Client struct {
	client MQTT.Client
}

func NewClient(broker string, port int) (*Client, error) {
	opts := MQTT.NewClientOptions().AddBroker(fmt.Sprintf("ws://%s:%d/mqtt", broker, port))
	client := MQTT.NewClient(opts)
	token := client.Connect()
	token.Wait()
	if err := token.Error(); err != nil {
		return nil, err
	}
	return &Client{client: client}, nil
}

func (c *Client) Subscribe(topic string, handler func(topic, payload string)) error {
	token := c.client.Subscribe(topic, 0, func(_ MQTT.Client, msg MQTT.Message) {
		handler(msg.Topic(), string(msg.Payload()))
	})
	token.Wait()
	return token.Error()
}

func (c *Client) Publish(topic, payload string) error {
	token := c.client.Publish(topic, 0, false, payload)
	token.Wait()
	return token.Error()
} 