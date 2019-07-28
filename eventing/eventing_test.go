package eventing

import (
	"testing"
)

const (
	InvalidPath = "../testdata/invalid.yml"
	SnippetPath = "../testdata/snippet.yml"
	TopicsPath  = "../testdata/topics.yml"
	TotalTopics = 4
)

func TestTopics(t *testing.T) {
	topics := &TopicsList{
		TopicsList: []Topic{
			{Key: "INTERACTION", Name: "adp-interaction-events"},
			{Key: "USER_SESSION", Name: "adp-user-session-events"},
		},
	}

	topicNames := topics.Topics()

	if len(topicNames) != len(topics.TopicsList) {
		t.Errorf("TestTopics Failed: Expected len() to be equal")
	}

	for i, topic := range topics.TopicsList {
		if topicNames[i] != topic.Name {
			t.Errorf("TestTopics Failed: Expected topic names to be equal")
		}
	}
}

func TestLoadTopics(t *testing.T) {
	topics, err := LoadTopics(TopicsPath)
	if err != nil {
		t.Errorf("TestLoadTopics Failed: Error loading topics - %v", err)
	}
	if len(topics.TopicsList) != TotalTopics {
		t.Errorf("TestLoadTopics Failed: Expected len() of %d but got %d", TotalTopics, len(topics.TopicsList))
	}

	topics, err = LoadTopics(SnippetPath)
	if len(topics.TopicsList) != 0 {
		t.Error("TestLoadTopics Failed: Expected error to not be nil")
	}

	_, err = LoadTopics(InvalidPath)
	if err == nil {
		t.Error("TestLoadTopics Failed: Expected error to not be nil")
	}
}
