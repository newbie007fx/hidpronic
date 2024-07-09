package mqtt

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"hidroponic/internal/platform/configuration"
	"log"
	"math/rand"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"golang.org/x/exp/slog"
)

type MqttService struct {
	mqqtClient mqtt.Client
	cs         *configuration.ConfigService
}

type CallbackFunc func(payload []byte)

func New(cs *configuration.ConfigService) *MqttService {
	return &MqttService{
		cs: cs,
	}
}

func (m *MqttService) Setup() (err error) {
	config := m.cs.GetConfig().MqttConfig
	connectAddress := fmt.Sprintf("%s://%s:%d", config.Protocol, config.Broker, config.Port)
	clientID := fmt.Sprintf("go-client-%d", rand.Int())

	log.Println("connect address: ", connectAddress)

	opts := mqtt.NewClientOptions()
	opts.AddBroker(connectAddress)
	opts.SetUsername(config.Username)
	opts.SetPassword(config.Password)
	opts.SetClientID(clientID)
	opts.SetKeepAlive(time.Second * 60)
	opts.SetConnectRetry(true)
	opts.SetConnectRetryInterval(time.Second * 60)
	opts.SetConnectTimeout(time.Second * 10)
	opts.SetAutoReconnect(true)

	opts.SetTLSConfig(loadTLSConfig("./emqxsl-ca.crt"))

	mqqtClient := mqtt.NewClient(opts)
	token := mqqtClient.Connect()
	if token.WaitTimeout(3*time.Second) && token.Error() != nil {
		err = token.Error()
		return
	}

	m.mqqtClient = mqqtClient

	return
}

func (m MqttService) Publish(topic string, data any) {
	payload, _ := json.Marshal(data)
	qos := 0
	if token := m.mqqtClient.Publish(topic, byte(qos), false, payload); token.Wait() && token.Error() != nil {
		slog.Error(fmt.Sprintf("publish failed, topic: %s, payload: %s with message: %s\n", topic, payload, token.Error().Error()))
	} else {
		slog.Debug(fmt.Sprintf("publish success, topic: %s, payload: %s\n", topic, payload))
	}
}

func (m MqttService) Subscribe(topic string, callback CallbackFunc) {
	qos := 0
	m.mqqtClient.Subscribe(topic, byte(qos), func(client mqtt.Client, msg mqtt.Message) {
		slog.Debug(fmt.Sprintf("Received `%s` from `%s` topic\n", msg.Payload(), msg.Topic()))
		callback(msg.Payload())
		msg.Ack()
	})
}

func loadTLSConfig(caFile string) *tls.Config {
	var tlsConfig tls.Config
	tlsConfig.InsecureSkipVerify = false
	if caFile != "" {
		certpool := x509.NewCertPool()
		ca, err := os.ReadFile(caFile)
		if err != nil {
			slog.Error(err.Error())
		}
		certpool.AppendCertsFromPEM(ca)
		tlsConfig.RootCAs = certpool
	}
	return &tlsConfig
}

func (mqtt MqttService) Shutdown() {
	mqtt.mqqtClient.Disconnect(1)
}
