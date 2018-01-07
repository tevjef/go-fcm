package fcm

import (
	"encoding/json"
	"errors"
	"strings"
	"time"
)

var (
	// ErrInvalidMessage occurs if push notitication message is nil.
	ErrInvalidMessage = errors.New("message is invalid")

	// ErrInvalidTarget occurs if message topic is empty.
	ErrInvalidTarget = errors.New("target is invalid. topic, token or condition maybe used")

	// ErrInvalidTimeToLive occurs if TimeToLive more then 2419200.
	ErrInvalidTimeToLive = errors.New("messages time-to-live is invalid")

	// ErrInvalidApnsPriority occurs if the priority is not 5 or 10.
	ErrInvalidApnsPriority = errors.New("apns message priority is invalid")
)

// SendRequest has a flag for testing and the actual message to send.
type SendRequest struct {
	// Flag for testing the request without actually delivering the message.
	ValidateOnly bool `json:"validate_only,omitempty"`
	// Message to send.
	Message *Message `json:"message,omitempty"`
}

// Notification specifies the basic notification template to use across all platforms.
type Notification struct {
	// The notification's title.
	Title string `json:"title,omitempty"`

	// The notification's body text.
	Body string `json:"body,omitempty"`
}

// Message represents list of targets, options, and payload for HTTP JSON
// messages.
type Message struct {
	// The identifier of the message sent, in the format of projects/*/messages/{message_id}.
	Name string `json:"name,omitempty"`

	// Registration token to send a message to.
	Token string `json:"token,omitempty"`

	// Topic name to send a message to, e.g. "weather". Note: "/topics/" prefix should not be provided.
	Topic string `json:"topic,omitempty"`

	// Condition to send a message to, e.g. "'foo' in topics && 'bar' in topics".
	Condition string `json:"condition,omitempty"`

	// Apple Push Notification Service specific options.
	Apns *ApnsConfig `json:"apns,omitempty"`

	// Webpush protocol options.
	Webpush *WebpushConfig `json:"webpush,omitempty"`

	// Android specific options for messages sent through FCM connection server.
	Android *AndroidConfig `json:"android,omitempty"`

	// Basic notification template to use across all platforms.
	Notification *Notification `json:"notification,omitempty"`

	// Arbitrary key/value payload.
	// An object containing a list of "key": value pairs.
	// Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	Data map[string]string `json:"data,omitempty"`
}

// MessageID returns the message id the successful send request.
func (msg Message) MessageID() string {
	lastIndex := strings.LastIndex(msg.Name, "/")
	if len(msg.Name) > lastIndex {
		return msg.Name[lastIndex+1:]
	}

	return ""
}

// Validate returns an error if the message is not well-formed.
func (msg *Message) Validate() error {
	if msg == nil {
		return ErrInvalidMessage
	}

	var targets = 0
	// validate target: `topic` or `condition`, or `token`
	if msg.Topic != "" {
		targets = targets + 1
	}

	if msg.Condition != "" {
		targets = targets + 1
	}

	opCnt := strings.Count(msg.Condition, "&&") + strings.Count(msg.Condition, "||")
	if opCnt > 2 {
		return ErrInvalidTarget
	}

	if msg.Token != "" {
		targets = targets + 1
	}

	if targets == 0 || targets > 1 {
		return ErrInvalidTarget
	}

	if msg.Android != nil && msg.Android.TTL != "" {
		if _, err := time.ParseDuration(msg.Android.TTL); err != nil {
			return ErrInvalidTimeToLive
		}
	}

	if msg.Apns != nil {
		b, err := json.Marshal(msg.Apns.Payload)
		if err != nil {
			return err
		}

		var payload ApnsPayload
		err = json.Unmarshal(b, &payload)
		if err != nil {
			return err
		}

		if msg.Apns.Headers != nil {
			if payload.Aps.ContentAvailable == int(ApnsContentAvailable) &&
				msg.Apns.Headers.Priority == string(ApnsHighPriority) {
				return ErrInvalidApnsPriority
			}
		}
	}

	return nil
}
