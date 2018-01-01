package fcm

// WebpushConfig represents the Webpush protocol outlined in https://tools.ietf.org/html/rfc8030
// https://firebase.google.com/docs/reference/fcm/rest/v1/projects.messages#WebpushConfig
type WebpushConfig struct {
	// HTTP headers defined in webpush protocol. Refer to Webpush
	// protocol for supported headers, e.g. "TTL": "15".

	// An object containing a list of "key": value pairs.
	// Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	Headers map[string]string `json:"headers,omitempty"`

	// Arbitrary key/value payload. If present, it will override
	// google.firebase.fcm.v1.Message.data.

	// An object containing a list of "key": value pairs.
	// Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	Data map[string]string `json:"data,omitempty"`

	// A web notification to send.
	Notification *WebpushNotification `json:"notification,omitempty"`
}

// WebpushNotification represents a web notification to send via webpush protocol.
// https://firebase.google.com/docs/reference/fcm/rest/v1/projects.messages#AndroidNotification
type WebpushNotification struct {
	// The notification's title. If present, it will override
	// google.firebase.fcm.v1.Notification.title.
	Title string `json:"title,omitempty"`

	// The notification's body text. If present, it will override
	// google.firebase.fcm.v1.Notification.body.
	Body string `json:"body,omitempty"`

	// The URL to use for the notification's icon.
	Icon string `json:"icon,omitempty"`
}
