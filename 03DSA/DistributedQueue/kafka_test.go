package distributedqueue

import (
	"testing"
	"time"
)


func TestQueueSystemBasicFunctionality(t *testing.T) {
	queueSystem := NewQueueSystem()

	// Create topics
	queueSystem.CreateTopic("topic1")
	queueSystem.CreateTopic("topic2")

	// Create producers
	producer1 := &Producer{ID: "producer1"}
	producer2 := &Producer{ID: "producer2"}

	// Create consumers
	consumer1 := &Consumer{ID: "consumer1"}
	consumer2 := &Consumer{ID: "consumer2"}

	// Get topics
	topic1 := queueSystem.GetTopic("topic1")
	topic2 := queueSystem.GetTopic("topic2")

	// Consumers subscribe to topics
	consumer1.Subscribe(topic1)
	consumer1.Subscribe(topic2)
	consumer2.Subscribe(topic1)

	// Producers publish messages to topics
	producer1.Publish(topic1, "message 1 to topic1")
	producer1.Publish(topic2, "message 1 to topic2")
	producer2.Publish(topic1, "message 2 to topic1")

	// Give some time for messages to be processed
	time.Sleep(1 * time.Second)
}

func TestMultipleProducersConsumers(t *testing.T) {
	queueSystem := NewQueueSystem()

	// Create a topic
	queueSystem.CreateTopic("topic1")

	// Create multiple producers
	producer1 := &Producer{ID: "producer1"}
	producer2 := &Producer{ID: "producer2"}

	// Create multiple consumers
	consumer1 := &Consumer{ID: "consumer1"}
	consumer2 := &Consumer{ID: "consumer2"}
	consumer3 := &Consumer{ID: "consumer3"}

	// Get topic
	topic1 := queueSystem.GetTopic("topic1")

	// Consumers subscribe to the topic
	consumer1.Subscribe(topic1)
	consumer2.Subscribe(topic1)
	consumer3.Subscribe(topic1)

	// Producers publish messages concurrently
	go producer1.Publish(topic1, "message 1 from producer1")
	go producer2.Publish(topic1, "message 1 from producer2")
	go producer1.Publish(topic1, "message 2 from producer1")
	go producer2.Publish(topic1, "message 2 from producer2")

	// Give some time for messages to be processed
	time.Sleep(1 * time.Second)
}


func TestTopicCreation(t *testing.T) {
	queueSystem := NewQueueSystem()

	// Create a topic
	topicName := "topic1"
	queueSystem.CreateTopic(topicName)

	// Verify that the topic was created
	topic := queueSystem.GetTopic(topicName)
	if topic == nil {
		t.Errorf("Expected topic %s to be created, but it was not found", topicName)
	}

	// Attempt to create the same topic again and verify it is not duplicated
	queueSystem.CreateTopic(topicName)
	if len(queueSystem.Topics) != 1 {
		t.Errorf("Expected only 1 topic to exist, but found %d", len(queueSystem.Topics))
	}
}


func TestTopicPublishingAndSubscribing(t *testing.T) {
	queueSystem := NewQueueSystem()

	// Create a topic
	queueSystem.CreateTopic("topic1")

	// Create a producer and a consumer
	producer := &Producer{ID: "producer"}
	consumer := &Consumer{ID: "consumer"}

	// Get the topic and subscribe the consumer
	topic1 := queueSystem.GetTopic("topic1")
	consumer.Subscribe(topic1)

	// Publish a message to the topic
	producer.Publish(topic1, "test message")

	// Give some time for the message to be processed
	time.Sleep(1 * time.Second)
}

