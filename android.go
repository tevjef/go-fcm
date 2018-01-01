package fcm

// AndroidNotification represents a notification to send to android devices.
type AndroidNotification struct {
	// The notification's title. If present, it will override
	// google.firebase.fcm.v1.Notification.title
	Title string `json:"title,omitempty"`

	// The notification's body text. If present, it will override
	// google.firebase.fcm.v1.Notification.body.
	Body string `json:"body,omitempty"`

	// The notification's icon. Sets the notification icon to myicon
	// for drawable resource myicon. If you don't send this key in
	// the request, FCM displays the launcher icon specified in your
	// app manifest.
	Icon string `json:"icon,omitempty"`

	// The sound to play when the device receives the notification.
	// Supports "default" or the filename of a sound resource bundled
	// in the app. Sound files must reside in /res/raw/.
	Sound string `json:"sound,omitempty"`

	// Identifier used to replace existing notifications in the
	// notification drawer. If not specified, each request creates
	// a new notification. If specified and a notification with the
	// same tag is already being shown, the new notification
	// replaces the existing one in the notification drawer.
	Tag string `json:"tag,omitempty"`

	// The notification's icon color, expressed in #rrggbb format.
	Color string `json:"color,omitempty"`

	// The action associated with a user click on the notification.
	// If specified, an activity with a matching intent filter is
	// launched when a user clicks on the notification.
	ClickAction string `json:"click_action,omitempty"`

	// The key to the body string in the app's string resources to
	// use to localize the body text to the user's current localization.
	BodyLocKey string `json:"body_loc_key,omitempty"`

	// Variable string values to be used in place of the format specifiers
	// in body_loc_key to use to localize the body text to the user's
	// current localization.
	BodyLocArgs []string `json:"body_loc_args,omitempty"`

	// The key to the title string in the app's string resources to use to
	// localize the title text to the user's current localization.
	TitleLocKey string `json:"title_loc_key,omitempty"`

	// Variable string values to be used in place of the format specifiers
	// in title_loc_key to use to localize the title text to the user's
	// current localization.
	TitleLocArgs []string `json:"title_loc_args,omitempty"`
}

// AndroidConfig represents android specific options for messages sent through FCM connection server.
type AndroidConfig struct {
	// An identifier of a group of messages that can be collapsed, so that
	// only the last message gets sent when delivery can be resumed.
	// A maximum of 4 different collapse keys is allowed at any given time.
	CollapseKey string `json:"collapse_key,omitempty"`

	// Message priority. Can take "normal" and "high" values.
	Priority string `json:"priority,omitempty"`

	// How long (in seconds) the message should be kept in FCM storage if
	// the device is offline. The maximum time to live supported is 4 weeks,
	// and the default value is 4 weeks if not set. Set it to 0 if want to
	// send the message immediately. In JSON format, the Duration type is
	// encoded as a string rather than an object, where the string ends in
	// the suffix "s" (indicating seconds) and is preceded by the number of
	// seconds, with nanoseconds expressed as fractional seconds. For example,
	// 3 seconds with 0 nanoseconds should be encoded in JSON format as "3s",
	// while 3 seconds and 1 nanosecond should be expressed in JSON format as "3.000000001s".
	// The ttl will be rounded down to the nearest second.

	// A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
	TTL string `json:"ttl,omitempty"`

	// Package name of the application where the registration tokens must match in order to receive the message.
	RestrictedPackageName string `json:"restricted_package_name,omitempty"`

	// Arbitrary key/value payload. If present, it will override google.firebase.fcm.v1.Message.data.
	// An object containing a list of "key": value pairs.
	// Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	Data map[string]string `json:"data,omitempty"`
	// Notification to send to android devices.
	Notification *AndroidNotification `json:"notification,omitempty"`
}

// AndroidMessagePriority represents the priority of a message to send to Android devices.
type AndroidMessagePriority string

var (
	// AndroidNormalPriority is the default priority for data messages. Normal priority messages won't open network connections on
	// a sleeping device, and their delivery may be delayed to conserve the battery. For less
	// time-sensitive messages, such as notifications of new email or other data to sync, choose normal
	// delivery priority.
	AndroidNormalPriority AndroidMessagePriority = "normal"

	// AndroidHighPriority is the default priority for notification messages. FCM attempts to deliver high priority
	// messages immediately, allowing the FCM service to wake a sleeping device when possible and open a
	// network connection to your app server. Apps with instant messaging, chat, or voice call alerts,
	// for example, generally need to open a network connection and make sure FCM delivers the message
	// to the device without delay. Set high priority if the message is time-critical and requires the
	// user's immediate interaction, but beware that setting your messages to high priority contributes
	// more to battery drain compared with normal priority messages.
	AndroidHighPriority AndroidMessagePriority = "high"
)
