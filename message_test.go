package fcm

import (
	"testing"
)

func TestValidate(t *testing.T) {
	t.Run("extract message id", func(t *testing.T) {
		msg := &Message{
			Name: "projects/*/messages/{message_id}",
		}
		expected := "{message_id}"
		result := msg.MessageID()
		if expected != result {
			t.Fatalf("expected: %v got: %v", expected, result)
		}
	})

	t.Run("valid with token", func(t *testing.T) {
		msg := &Message{
			Topic: "test",
		}
		err := msg.Validate()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	t.Run("invalid message", func(t *testing.T) {
		var msg *Message
		err := msg.Validate()
		if err == nil {
			t.Fatalf("expected <%v> error, but got <nil>", ErrInvalidMessage)
		}
	})

	t.Run("invalid target", func(t *testing.T) {
		msg := &Message{
			Data: map[string]string{
				"message": "This is a Firebase Cloud Messaging Topic Message!",
			},
		}
		err := msg.Validate()
		if err == nil {
			t.Fatalf("expected <%v> error, but got nil", ErrInvalidTarget)
		}
	})

	t.Run("invalid TTL", func(t *testing.T) {
		timeToLive := "5"
		msg := &Message{
			Topic: "test",
			Android: &AndroidConfig{
				TTL: timeToLive,
			},
		}
		err := msg.Validate()
		if err == nil {
			t.Fatalf("expected <%v> error, but got nil", ErrInvalidTimeToLive)
		}
	})

	t.Run("valid target with condition", func(t *testing.T) {
		msg := &Message{
			Condition: "'dogs' in topics || 'cats' in topics",
		}
		err := msg.Validate()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	t.Run("valid target with topic", func(t *testing.T) {
		msg := &Message{
			Topic: "cats",
		}
		err := msg.Validate()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	t.Run("valid target with token", func(t *testing.T) {
		msg := &Message{
			Token: "12345678",
		}
		err := msg.Validate()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	t.Run("invalid condition", func(t *testing.T) {
		msg := &Message{
			Condition: "'TopicA' in topics && ('TopicB' in topics || 'TopicC' in topics || 'TopicD' in topics )",
		}
		err := msg.Validate()
		if err == nil {
			t.Fatalf("expected <%v> error, but got nil", ErrInvalidTarget)
		}
	})

	t.Run("invalid union with topic and condition", func(t *testing.T) {
		msg := &Message{
			Topic:     "cats",
			Condition: "'TopicA' in topics && ('TopicB' in topics || 'TopicC' in topics || 'TopicD' in topics )",
		}
		err := msg.Validate()
		if err == nil {
			t.Fatalf("expected <%v> error, but got nil", ErrInvalidTarget)
		}
	})

	t.Run("invalid union with token and condition", func(t *testing.T) {
		msg := &Message{
			Token:     "12345678",
			Condition: "'TopicA' in topics && ('TopicB' in topics || 'TopicC' in topics || 'TopicD' in topics )",
		}
		err := msg.Validate()
		if err == nil {
			t.Fatalf("expected <%v> error, but got nil", ErrInvalidTarget)
		}
	})

	t.Run("invalid union with token and topic", func(t *testing.T) {
		msg := &Message{
			Token: "12345678",
			Topic: "cats",
		}
		err := msg.Validate()
		if err == nil {
			t.Fatalf("expected <%v> error, but got nil", ErrInvalidTarget)
		}
	})

	t.Run("invalid apns priority", func(t *testing.T) {
		payload := &ApnsPayload{
			Aps: &ApsDictionary{
				ContentAvailable: int(ApnsContentAvailable),
			},
		}

		payloadMap, err := payload.ToMap()
		if err != nil {
			t.Error(err)
		}

		msg := &Message{
			Token: "12345678",
			Apns: &ApnsConfig{
				Headers: &ApnsHeaders{
					Priority: string(ApnsHighPriority),
				},
				Payload: payloadMap,
			},
		}

		err = msg.Validate()
		if err == nil {
			t.Fatalf("expected <%v> error, but got nil", ErrInvalidApnsPriority)
		}
	})
}
