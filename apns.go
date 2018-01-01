package fcm

import (
	"encoding/json"
	"log"
)

// ApnsConfig represents Apple Push Notification Service specific options.
type ApnsConfig struct {
	Headers *ApnsHeaders           `json:"headers,omitempty"`
	Payload map[string]interface{} `json:"payload,omitempty"`
}

// ApnsPayload defines an APNS notification.
type ApnsPayload struct {
	Aps *ApsDictionary `json:"aps,omitempty"`
}

// ToMap converts a ApnsPayload struct to a map[string]interface{}.
// It also returns an error if the operation fails.
func (payload *ApnsPayload) ToMap() (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if payload == nil {
		return m, nil
	}

	b, err := json.Marshal(payload)
	if err != nil {
		return m, err
	}

	err = json.Unmarshal(b, &m)
	return m, err
}

// MustToMap converts a ApnsPayload struct to a map[string]interface{}.
// It exits the program is this operation fails.
func (payload *ApnsPayload) MustToMap() map[string]interface{} {
	m, err := payload.ToMap()
	if err != nil {
		log.Fatal(err.Error())
	}

	return m
}

// ApsDictionary defines an APNS notification.
type ApsDictionary struct {
	// Include this key when you want the system to display a standard
	// alert or a banner. The notification settings for your app on
	// the user’s device determine whether an alert or banner is displayed.
	Alert *ApnsAlert `json:"alert,omitempty"`

	// Include this key when you want the system to modify the badge of
	// your app icon. If this key is not included in the dictionary,
	// the badge is not changed. To remove the badge, set the value of this key to 0.
	Badge int `json:"badge,omitempty"`

	// Include this key when you want the system to play a sound. The value of
	// this key is the name of a sound file in your app’s main bundle or in the
	// Library/Sounds folder of your app’s data container. If the sound file
	// cannot be found, or if you specify defaultfor the value, the system plays
	// the default alert sound.
	Sound string `json:"sound,omitempty"`

	// Provide this key with a string value that represents the notification’s type.
	// This value corresponds to the value in the "identifier" property of one of
	// your app’s registered categories.
	Category string `json:"category,omitempty"`

	// Provide this key with a string value that represents the app-specific identifier
	// for grouping notifications. If you provide a Notification Content app extension,
	// you can use this value to group your notifications together.
	ThreadID string `json:"thread-id,omitempty"`

	// Include this key with a value of 1 to configure a background update notification.
	// When this key is present, the system wakes up your app in the background and
	// delivers the notification to its app delegate.
	ContentAvailable int `json:"content-available,omitempty"`
}

// ApnsAlert represents a APNS alert
type ApnsAlert struct {
	// A short string describing the purpose of the notification.
	// Apple Watch displays this string as part of the notification interface.
	// This string is displayed only briefly and should be crafted so that
	// it can be understood quickly. This key was added in iOS 8.2.
	Title string `json:"title,omitempty"`

	// The text of the alert message.
	Body string `json:"body,omitempty"`

	// The key to a title string in the Localizable.strings file for the current
	// localization. The key string can be formatted with %@ and %n$@ specifiers
	// to take the variables specified in the title-loc-args array.
	TitleLocKey string `json:"title-loc-key,omitempty"`

	// Variable string values to appear in place of the format specifiers in title-loc-key.
	TitleLocArgs []string `json:"title-loc-args,omitempty"`
	// If a string is specified, the system displays an alert that includes
	// the Close and View buttons. The string is used as a key to get a
	// localized string in the current localization to use for the right button’s
	// title instead of “View”.
	ActionLocKey string `json:"action-loc-key,omitempty"`

	// A key to an alert-message string in a Localizable.strings file for the current
	// localization (which is set by the user’s language preference). The key string
	// can be formatted with %@ and %n$@ specifiers to take the variables specified in
	// the loc-args array.
	LocKey string `json:"loc-key,omitempty"`

	// Variable string values to appear in place of the format specifiers in loc-key.
	LocArgs []string `json:"loc-args,omitempty"`

	// The filename of an image file in the app bundle, with or without the filename
	// extension. The image is used as the launch image when users tap the action button
	// or move the action slider. If this property is not specified, the system either uses
	// the previous snapshot, uses the image identified by the UILaunchImageFile key in the
	//  app’s Info.plist file, or falls back to Default.png.
	LaunchImage string `json:"launch-image,omitempty"`
}

// ApnsHeaders represents a collection of APNS headers
type ApnsHeaders struct {
	// A UNIX epoch date expressed in seconds (UTC). This header identifies the date when the
	// notification is no longer valid and can be discarded.

	// If this value is nonzero, APNs stores the notification and tries to deliver it at least
	// once, repeating the attempt as needed if it is unable to deliver the notification the
	// first time. If the value is 0, APNs treats the notification as if it expires immediately
	// and does not store the notification or attempt to redeliver it.
	Expiration string `json:"apns-expiration,omitempty"`

	// The priority of the notification. Specify one of the following values:

	// 10–Send the push message immediately. Notifications with this priority
	// must trigger an alert, sound, or badge on the target device. It is an error
	// to use this priority for a push notification that contains only the
	// content-available key.

	// 5—Send the push message at a time that takes into account power considerations
	// for the device. Notifications with this priority might be grouped and delivered
	// in bursts. They are throttled, and in some cases are not delivered.

	// If you omit this header, the APNs server sets the priority to 10.
	Priority string `json:"apns-priority,omitempty"`

	// The topic of the remote notification, which is typically the bundle ID for your app.
	// The certificate you create in your developer account must include the capability for
	// this topic.

	// If your certificate includes multiple topics, you must specify a value for this header.

	// If you omit this request header and your APNs certificate does not specify multiple topics,
	// the APNs server uses the certificate’s Subject as the default topic.

	// If you are using a provider token instead of a certificate, you must specify a value for
	// this request header. The topic you provide should be provisioned for the your team named
	// in your developer account.
	Topic string `json:"apns-topic,omitempty"`

	// Multiple notifications with the same collapse identifier are displayed to the user as
	//  a single notification. The value of this key must not exceed 64 bytes.
	CollapseID string `json:"apns-collapse-id,omitempty"`
}

// ApnsMessagePriority represents the priority of the notification. Specify one of the following values:
type ApnsMessagePriority string

var (
	// ApnsNormalPriority sends the push message at a time that takes into account power considerations
	// for the device. Notifications with this priority might be grouped and delivered
	// in bursts. They are throttled, and in some cases are not delivered.
	ApnsNormalPriority ApnsMessagePriority = "5"

	// ApnsHighPriority sends the push message immediately. Notifications with this priority
	// must trigger an alert, sound, or badge on the target device. It is an error
	// to use this priority for a push notification that contains only the
	// content-available key.
	ApnsHighPriority ApnsMessagePriority = "10"
)

// ApnsContentAvailability represents the content-available key in a APNS notification
// which may wither be 1 or 0.
type ApnsContentAvailability int

const (
	// ApnsContentUnavailable flags the notification to be delievered directly to the user
	// without first waking your app up.
	ApnsContentUnavailable ApnsContentAvailability = 0

	// ApnsContentAvailable flags the notification to be delivered to the user’s device in the background.
	// iOS wakes up your app in the background and gives it up to 30 seconds to run.
	ApnsContentAvailable ApnsContentAvailability = 1
)
