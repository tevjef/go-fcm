package main

import (
	"log"

	"github.com/tevjef/go-fcm"
)

func main() {
	// Create the message to be sent.
	apnsPayload := &fcm.ApnsPayload{
		Aps: &fcm.ApsDictionary{
			Alert: &fcm.ApnsAlert{
				Title:        "Acme title",
				Body:         "Acme message received from Johnny Appleseed",
				LocKey:       "GAME_PLAY_REQUEST_FORMAT",
				LocArgs:      []string{"Jenna", "Frank"},
				TitleLocKey:  "GAME_PLAY_REQUEST_FORMAT",
				TitleLocArgs: []string{"Jenna", "Frank"},
				ActionLocKey: "PLAY",
				LaunchImage:  "UILaunchImageFileKey",
			},
			Badge:            1,
			Sound:            "chime.aiff",
			Category:         "NEW_MESSAGE_CATEGORY",
			ThreadID:         "my-thread-id",
			ContentAvailable: int(fcm.ApnsContentAvailable),
		},
	}

	payload, err := apnsPayload.ToMap()
	if err != nil {
		log.Fatal(err)
	}

	// optionall additional data in APNS message.
	payload["acme1"] = "bar"
	payload["acme2"] = []string{"bang", "whiz"}

	msg := &fcm.SendRequest{
		ValidateOnly: true,
		Message: &fcm.Message{
			Topic:     "cats",
			Token:     "bk3RNwTe3H0:CI2k_HHwgIpoDKCIZvvDMExUdFQ3P1...",
			Condition: "'dogs' in topics || 'cats' in topics",
			Notification: &fcm.Notification{
				Title: "FCM Message",
				Body:  "This is a Firebase Cloud Messaging Topic Message!",
			},
			Apns: &fcm.ApnsConfig{
				Headers: &fcm.ApnsHeaders{
					Expiration: "14567890",
					Priority:   string(fcm.ApnsHighPriority),
					Topic:      "my-topic",
					CollapseID: "my-collapse-id",
				},
				Payload: payload,
			},
			Android: &fcm.AndroidConfig{
				CollapseKey: "my-collapse-key",
				Priority:    string(fcm.AndroidHighPriority),
				TTL:         "84000s",
				Data: map[string]string{
					"acme1": "bar",
				},
				RestrictedPackageName: "com.github.go-fcm",
				Notification: &fcm.AndroidNotification{
					Title:        "FCM Message",
					Body:         "This is a Firebase Cloud Messaging Topic Message!",
					Icon:         "ic_notification",
					Color:        "#rrggbb",
					Sound:        "res_raw_notification_sound.mp3",
					Tag:          "my-notification-tag",
					ClickAction:  "MainActivity",
					BodyLocKey:   "notification_body", // R.string.notification_body
					BodyLocArgs:  []string{"Jenna", "Frank"},
					TitleLocKey:  "notification_title", // R.string.notification_title
					TitleLocArgs: []string{"Jenna", "Frank"},
				},
			},
		},
	}

	// Create a FCM client to send the message.
	client, err := fcm.NewClient("projectID", "secrets/sa.json")
	if err != nil {
		log.Fatalln(err)
	}

	// Send the message and receive the response without retries.
	response, err := client.Send(msg)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("%#v\n", response)
}
