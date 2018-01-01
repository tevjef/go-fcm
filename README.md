# go-fcm for FCM HTTP v1 API

[![GoDoc](https://godoc.org/github.com/tevjef/go-fcm?status.svg)](https://godoc.org/github.com/tevjef/go-fcm)
[![Build Status](https://travis-ci.org/tevjef/go-fcm.svg?branch=master)](https://travis-ci.org/tevjef/go-fcm)
[![Go Report Card](https://goreportcard.com/badge/github.com/tevjef/go-fcm)](https://goreportcard.com/report/github.com/tevjef/go-fcm)

Golang client library for [Firebase Cloud Messaging](https://firebase.google.com/docs/cloud-messaging/) v1 API.

The [Legacy FCM HTTP Protocol](https://firebase.google.com/docs/cloud-messaging/http-server-ref) has no construct to make a distinction between Android, iOS and Web notifications. The new [HTTP v1 API](https://firebase.google.com/docs/reference/fcm/rest/v1/projects.messages) does. 

With this library, you:
* Cannot send notifications to multiple registration ids or devices in one request. Use the [Legacy HTTP Protocol](https://firebase.google.com/docs/cloud-messaging/http-server-ref_)

* Cannot receive messages from devices. Use the [Legacy XMPP Protocol](https://firebase.google.com/docs/cloud-messaging/xmpp-server-ref) 

## Getting Started

To install fcm, use `go get`:


## Usage

```bash
go get github.com/tevjef/go-fcm

```

### Example

```go
import (
	"encoding/json"
	"log"

	"github.com/tevjef/go-fcm"
)

func main() {
	// Create the message to be sent.
	msg := &fcm.SendRequest{
		ValidateOnly: true,
		Message: &fcm.Message{
			Token: "bk3RNwTe3H0:CI2k_HHwgIpoDKCIZvvDMExUdFQ3P1...",
			Notification: &fcm.Notification{
				Title: "FCM Message",
				Body:  "This is a Firebase Cloud Messaging Topic Message!",
			},
			Apns: &fcm.ApnsConfig{
				Payload: &fcm.ApnsPayload{
					Aps: &fcm.ApsDictionary{
						Alert: &fcm.ApnsAlert{
							LaunchImage: "UILaunchImageFileKey",
						},
						Badge:            1,
						Category:         "NEW_MESSAGE_CATEGORY",
						ContentAvailable: int(fcm.ApnsContentAvailable),
					},
				},
			},
			Android: &fcm.AndroidConfig{
				Priority: string(fcm.AndroidHighPriority),
				TTL:      "84000s",
				Notification: &fcm.AndroidNotification{
					Icon:        "ic_notification",
					Color:       "#rrggbb",
					ClickAction: "MainActivity",
				},
			},
		},
	}

	// Create a FCM client to send the message.
	client, err := fcm.NewClient("projectID", "sa.json")
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
```

### Example JSON sent to FCM HTTP v1 API

```json
{
   "validate_only": true,
   "message": {
     "token": "bk3RNwTe3H0:CI2k_HHwgIpoDKCIZvvDMExUdFQ3P1...",
     or
     "topic": "cats", 
     or
     "condition": "'dogs' in topics || 'cats' in topics", 
     "apns": {
       "headers": {
         "apns-expiration": "14567890",
         "apns-priority": "10",
         "apns-topic": "my-topic",
         "apns-collapse-id": "my-collapse-id"
       },
       "payload": {
         "acme1": "bar",
         "acme2": [
           "bang",
           "whiz"
         ],
         "aps": {
           "alert": {
             "action-loc-key": "PLAY",
             "body": "Acme message received from Johnny Appleseed",
             "launch-image": "UILaunchImageFileKey",
             "loc-args": [
               "Jenna",
               "Frank"
             ],
             "loc-key": "GAME_PLAY_REQUEST_FORMAT",
             "title": "Acme title",
             "title-loc-args": [
               "Jenna",
               "Frank"
             ],
             "title-loc-key": "GAME_PLAY_REQUEST_FORMAT"
           },
           "badge": 1,
           "category": "NEW_MESSAGE_CATEGORY",
           "content-available": 1,
           "sound": "chime.aiff",
           "thread-id": "my-thread-id"
         }
       }
     },
     "android": {
       "collapse_key": "my-collapse-key",
       "priority": "HIGH",
       "ttl": "84000s",
       "restricted_package_name": "com.github.go-fcm",
       "data": {
         "acme1": "bar",
         "acme2": [
           "bang",
           "whiz"
         ]
       },
       "notification": {
         "title": "FCM Message",
         "body": "This is a Firebase Cloud Messaging Topic Message!",
         "icon": "ic_notification",
         "sound": "res_raw_notification_sound.mp3",
         "tag": "my-notification-tag",
         "color": "#rrggbb",
         "click_action": "MainActivity",
         "body_loc_key": "notification_body",
         "body_loc_args": [
           "Jenna",
           "Frank"
         ],
         "title_loc_key": "notification_title",
         "title_loc_args": [
           "Jenna",
           "Frank"
         ]
       }
     },
     "notification": {
       "title": "FCM Message",
       "body": "This is a Firebase Cloud Messaging Topic Message!"
     }
   }
 }
```
