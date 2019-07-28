package eventing

import (
	"testing"
)

const (
	InvalidPath = "../testdata/invalid.yml"
	SnippetPath = "../testdata/snippet.yml"
	TopicsPath  = "../testdata/topics.yml"
	TotalTopics = 3
)

func TestTopics(t *testing.T) {
	topicsList := &TopicsList{
		TopicsList: []string{
			"notification-events",
			"request-events",
		},
	}

	topicNames := topicsList.Topics()

	if len(topicNames) != len(topicsList.Topics()) {
		t.Errorf("TestTopics Failed: Expected len() to be equal")
	}

	for i, topic := range topicsList.Topics() {
		if topicNames[i] != topic {
			t.Errorf("TestTopics Failed: Expected topic names to be equal")
		}
	}
}

func TestLoadTopics(t *testing.T) {
	topicsList, err := LoadTopics(TopicsPath)
	if err != nil {
		t.Errorf("TestLoadTopics Failed: Error loading topics - %v", err)
	}
	if len(topicsList.TopicsList) != TotalTopics {
		t.Errorf("TestLoadTopics Failed: Expected len() of %d but got %d", TotalTopics, len(topicsList.TopicsList))
	}

	topicsList, err = LoadTopics(SnippetPath)
	if len(topicsList.TopicsList) != 0 {
		t.Error("TestLoadTopics Failed: Expected error to not be nil")
	}

	_, err = LoadTopics(InvalidPath)
	if err == nil {
		t.Error("TestLoadTopics Failed: Expected error to not be nil")
	}
}
