package eventing

import (
	"io/ioutil"
	"log"
	"time"

	"github.com/Shopify/sarama"
	"github.com/jpillora/backoff"
	yaml "gopkg.in/yaml.v2"
)

// Topic - Represents a KafkaTopic where Key is the data type and Name is the topic name
type Topic struct {
	Key  string `yaml:"key"`
	Name string `yaml:"name"`
}

// TopicsList - Wrapper Struct to allow parsing of yaml to a slice of type Topic
type TopicsList struct {
	TopicsList []Topic `yaml:"topics"`
}

// Topics - Returns string slice containing the TopicsList's topic names
func (t TopicsList) Topics() []string {
	topics := []string{}

	for _, topic := range t.TopicsList {
		topics = append(topics, topic.Name)
	}

	return topics
}

// LoadTopics - Attempts to read the yaml file at the provided path and unmarshal to a *TopicsList
func LoadTopics(path string) (*TopicsList, error) {
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	topicsList := new(TopicsList)
	err = yaml.Unmarshal(yamlFile, topicsList)
	if err != nil {
		return nil, err
	}

	return topicsList, nil
}

// WaitForKafka - Attempts to connect to the provided addrs and create a sarama.ClusterAdmin
func WaitForKafka(backoff *backoff.Backoff, addrs []string, conf *sarama.Config) sarama.ClusterAdmin {
	log.Printf("Attempting to connect to Kafka on - %s", addrs)

WaitForKafka:
	for {
		clusterAdmin, err := sarama.NewClusterAdmin(addrs, conf)
		if err != nil {
			log.Printf("Error: Failed to connect to Kafka - %v", err)

			d := backoff.Duration()
			log.Printf("Retrying in %v...", d)
			time.Sleep(d)
			continue WaitForKafka
		}

		log.Printf("Successfully connected to Kafka on - %s", addrs)
		return clusterAdmin
	}
}

// WaitForTopics - Uses the provided sarama.ClusterAdmin to list cluster topics and waits until the required *TopicsList has been satisfied
func WaitForTopics(backoff *backoff.Backoff, clusterAdmin sarama.ClusterAdmin, requiredTopics *TopicsList) {
	log.Printf("Waiting for Required Topics - %s", requiredTopics.Topics())

WaitForTopics:
	for {
		topics, err := clusterAdmin.ListTopics()
		if err != nil {
			log.Printf("Error: Failed to retrieve topics from Kafka - %v", err)

			d := backoff.Duration()
			log.Printf("Retrying in %v...", d)
			time.Sleep(d)
			continue WaitForTopics
		}

		for _, t := range requiredTopics.Topics() {
			topicInfo, ok := topics[t]
			if !ok {
				log.Printf("Failed to find Topic - %s", t)

				d := backoff.Duration()
				log.Printf("Retrying in %v...", d)
				time.Sleep(d)
				continue WaitForTopics
			}
			log.Printf("Successfully found Topic: %s Metadata: [Partitions: %d ReplicationFactor: %d]", t, topicInfo.NumPartitions, topicInfo.ReplicationFactor)
		}

		return
	}
}
