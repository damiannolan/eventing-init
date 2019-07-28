package config

import (
	"os"
	"strings"

	"github.com/Shopify/sarama"
)

// Configuration Fallbacks
const (
	KafkaHost    = "kafka"
	KafkaPort    = "9092"
	KafkaVersion = "2.3.0"
	TopicsPath   = "/etc/config/topics.yml"
)

var (
	// KafkaSettings - Contains the Kafka server host and port values
	KafkaSettings *kafkaSettings
	// SaramaConfig - The Sarama Kafka Client configuration
	SaramaConfig *sarama.Config
)

func init() {
	KafkaSettings = newKafkaSettings()
	SaramaConfig = newSaramaConfig()
}

func newSaramaConfig() *sarama.Config {
	version, err := sarama.ParseKafkaVersion(KafkaSettings.Version())
	if err != nil {
		panic("Unsupported Kafka Version")
	}

	config := sarama.NewConfig()
	config.ClientID = "eventing-init"
	config.Version = version

	return config
}

type kafkaSettings struct {
	host    string
	port    string
	version string
}

func newKafkaSettings() *kafkaSettings {
	return &kafkaSettings{
		host:    getStringEnvironmentVariable("KAFKA_INSTANCE", KafkaHost),
		port:    getStringEnvironmentVariable("KAFKA_PORT", KafkaPort),
		version: getStringEnvironmentVariable("KAFKA_VERSION", KafkaVersion),
	}
}

func (k *kafkaSettings) Host() string {
	return k.host
}

func (k *kafkaSettings) Port() string {
	return k.port
}

func (k *kafkaSettings) Version() string {
	return k.version
}

func (k *kafkaSettings) Broker() string {
	return strings.Join([]string{k.host, k.port}, ":")
}

func getStringEnvironmentVariable(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
