package model

import (
	"github.com/crazy-max/diun/v3/pkg/registry"
)

// NotifEntry represents a notification entry
type NotifEntry struct {
	Status   ImageStatus       `json:"status,omitempty"`
	Provider string            `json:"provider,omitempty"`
	Image    registry.Image    `json:"image,omitempty"`
	Manifest registry.Manifest `json:"manifest,omitempty"`
}

// Notif holds data necessary for notification configuration
type Notif struct {
	Amqp       *NotifAmqp       `yaml:"amqp,omitempty"`
	Gotify     *NotifGotify     `yaml:"gotify,omitempty"`
	Mail       *NotifMail       `yaml:"mail,omitempty"`
	RocketChat *NotifRocketChat `yaml:"rocketchat,omitempty"`
	Script     *NotifScript     `yaml:"script,omitempty"`
	Slack      *NotifSlack      `yaml:"slack,omitempty"`
	Teams      *NotifTeams      `yaml:"teams,omitempty"`
	Telegram   *NotifTelegram   `yaml:"telegram,omitempty"`
	Webhook    *NotifWebhook    `yaml:"webhook,omitempty"`
}

// NotifAmqp holds amqp notification configuration details
type NotifAmqp struct {
	Username     string `yaml:"username,omitempty"`
	UsernameFile string `yaml:"username_file,omitempty"`
	Password     string `yaml:"password,omitempty"`
	PasswordFile string `yaml:"password_file,omitempty"`
	Host         string `yaml:"host,omitempty"`
	Port         int    `yaml:"port,omitempty"`
	Queue        string `yaml:"queue,omitempty"`
	Exchange     string `yaml:"exchange,omitempty"`
}

// NotifGotify holds gotify notification configuration details
type NotifGotify struct {
	Endpoint string `yaml:"endpoint,omitempty"`
	Token    string `yaml:"token,omitempty"`
	Priority int    `yaml:"priority,omitempty"`
	Timeout  int    `yaml:"timeout,omitempty"`
}

// NotifMail holds mail notification configuration details
type NotifMail struct {
	Host               string `yaml:"host,omitempty"`
	Port               int    `yaml:"port,omitempty"`
	SSL                *bool  `yaml:"ssl,omitempty"`
	InsecureSkipVerify *bool  `yaml:"insecure_skip_verify,omitempty"`
	Username           string `yaml:"username,omitempty"`
	UsernameFile       string `yaml:"username_file,omitempty"`
	Password           string `yaml:"password,omitempty"`
	PasswordFile       string `yaml:"password_file,omitempty"`
	From               string `yaml:"from,omitempty"`
	To                 string `yaml:"to,omitempty"`
}

// NotifRocketChat holds Rocket.Chat notification configuration details
type NotifRocketChat struct {
	Endpoint string `yaml:"endpoint,omitempty"`
	Channel  string `yaml:"channel,omitempty"`
	UserID   string `yaml:"user_id,omitempty"`
	Token    string `yaml:"token,omitempty"`
	Timeout  int    `yaml:"timeout,omitempty"`
}

// NotifScript holds script notification configuration details
type NotifScript struct {
	Cmd  string   `yaml:"cmd,omitempty"`
	Args []string `yaml:"args,omitempty"`
	Dir  string   `yaml:"dir,omitempty"`
}

// NotifSlack holds slack notification configuration details
type NotifSlack struct {
	WebhookURL string `yaml:"webhook_url,omitempty"`
}

// NotifTeams holds Teams notification configuration details
type NotifTeams struct {
	WebhookURL string `yaml:"webhook_url,omitempty"`
}

// NotifTelegram holds Telegram notification configuration details
type NotifTelegram struct {
	BotToken string  `yaml:"token,omitempty"`
	ChatIDs  []int64 `yaml:"chat_ids,omitempty"`
}

// NotifWebhook holds webhook notification configuration details
type NotifWebhook struct {
	Endpoint string            `yaml:"endpoint,omitempty"`
	Method   string            `yaml:"method,omitempty"`
	Headers  map[string]string `yaml:"headers,omitempty"`
	Timeout  int               `yaml:"timeout,omitempty"`
}
