package mqtt

import (
	"fmt"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	MQTT "github.com/eclipse/paho.mqtt.golang"
)

type Client struct {
	client MQTT.Client
}

func NewClient(broker string, port int, username, password string, tlsEnabled bool, caCert, clientCert, clientKey string, insecureSkipVerify bool) (*Client, error) {
	var brokerURL string
	if tlsEnabled {
		brokerURL = fmt.Sprintf("wss://%s:%d/mqtt", broker, port)
	} else {
		brokerURL = fmt.Sprintf("ws://%s:%d/mqtt", broker, port)
	}
	opts := MQTT.NewClientOptions().AddBroker(brokerURL)
	if username != "" {
		opts.SetUsername(username)
	}
	if password != "" {
		opts.SetPassword(password)
	}
	if tlsEnabled {
		tlsConfig := &tls.Config{InsecureSkipVerify: insecureSkipVerify}
		if caCert != "" {
			caCertData, err := ioutil.ReadFile(caCert)
			if err != nil {
				return nil, fmt.Errorf("failed to read CA cert: %w", err)
			}
			caPool := x509.NewCertPool()
			if !caPool.AppendCertsFromPEM(caCertData) {
				return nil, fmt.Errorf("failed to append CA cert")
			}
			tlsConfig.RootCAs = caPool
		}
		if clientCert != "" && clientKey != "" {
			cert, err := tls.LoadX509KeyPair(clientCert, clientKey)
			if err != nil {
				return nil, fmt.Errorf("failed to load client cert/key: %w", err)
			}
			tlsConfig.Certificates = []tls.Certificate{cert}
		}
		opts.SetTLSConfig(tlsConfig)
	}
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