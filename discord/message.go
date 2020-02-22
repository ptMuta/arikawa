package discord

type Message struct {
	ID        Snowflake   `json:"id,string"`
	Type      MessageType `json:"type"`
	ChannelID Snowflake   `json:"channel_id,string"`
	GuildID   Snowflake   `json:"guild_id,string,omitempty"`

	// The author object follows the structure of the user object, but is only
	// a valid user in the case where the message is generated by a user or bot
	// user. If the message is generated by a webhook, the author object
	// corresponds to the webhook's id, username, and avatar. You can tell if a
	// message is generated by a webhook by checking for the webhook_id on the
	// message object.
	Author User `json:"author"`

	// The member object exists in MESSAGE_CREATE and MESSAGE_UPDATE
	// events from text-based guild channels.
	Member *Member `json:"member,omitempty"`

	Content string `json:"content"`

	Timestamp       Timestamp `json:"timestamp,omitempty"`
	EditedTimestamp Timestamp `json:"edited_timestamp,omitempty"`

	TTS    bool `json:"tts"`
	Pinned bool `json:"pinned"`

	// The user objects in the mentions array will only have the partial
	// member field present in MESSAGE_CREATE and MESSAGE_UPDATE events from
	// text-based guild channels.
	Mentions []GuildUser `json:"mentions"`

	MentionRoleIDs  []Snowflake `json:"mention_roles"`
	MentionEveryone bool        `json:"mention_everyone"`

	// Not all channel mentions in a message will appear in mention_channels.
	MentionChannels []ChannelMention `json:"mention_channels,omitempty"`

	Attachments []Attachment `json:"attachments"`
	Embeds      []Embed      `json:"embeds"`

	Reactions []Reaction `json:"reactions,omitempty"`

	// Used for validating a message was sent
	Nonce string `json:"nonce,omitempty"`

	WebhookID   Snowflake           `json:"webhook_id,string,omitempty"`
	Activity    *MessageActivity    `json:"activity,omitempty"`
	Application *MessageApplication `json:"application,omitempty"`
	Reference   *MessageReference   `json:"message_reference,omitempty"`
	Flags       MessageFlags        `json:"flags"`
}

// URL generates a Discord client URL to the message. If the message doesn't
// have a GuildID, it will generate a URL with the guild "@me".
func (m Message) URL() string {
	var head = "https://discordapp.com/channels/"
	var tail = "/" + m.ChannelID.String() + "/" + m.ID.String()

	if !m.GuildID.Valid() {
		return head + "@me" + tail
	}

	return head + m.GuildID.String() + tail
}

type MessageType uint8

const (
	DefaultMessage MessageType = iota
	RecipientAddMessage
	RecipientRemoveMessage
	CallMessage
	ChannelNameChangeMessage
	ChannelIconChangeMessage
	ChannelPinnedMessage
	GuildMemberJoinMessage
	NitroBoostMessage
	NitroTier1Message
	NitroTier2Message
	NitroTier3Message
	ChannelFollowAddMessage
	GuildDiscoveryDisqualifiedMessage
	GuildDiscoveryRequalifiedMessage
)

type MessageFlags uint8

const (
	CrosspostedMessage MessageFlags = 1 << iota
	MessageIsCrosspost
	SuppressEmbeds
	SourceMessageDeleted
	UrgentMessage
)

type ChannelMention struct {
	ChannelID   Snowflake   `json:"id,string"`
	GuildID     Snowflake   `json:"guild_id,string"`
	ChannelType ChannelType `json:"type"`
	ChannelName string      `json:"name"`
}

type GuildUser struct {
	User
	Member *Member `json:"member,omitempty"`
}

//

type MessageActivity struct {
	Type MessageActivityType `json:"type"`

	// From a Rich Presence event
	PartyID string `json:"party_id,omitempty"`
}

type MessageActivityType uint8

const (
	JoinMessage MessageActivityType = iota + 1
	SpectateMessage
	ListenMessage
	JoinRequestMessage
)

//

type MessageApplication struct {
	ID          Snowflake `json:"id,string"`
	CoverID     string    `json:"cover_image,omitempty"`
	Description string    `json:"description"`
	Icon        string    `json:"icon"`
	Name        string    `json:"name"`
}

//

type MessageReference struct {
	ChannelID Snowflake `json:"channel_id,string"`

	// Field might not be provided
	MessageID Snowflake `json:"message_id,string,omitempty"`
	GuildID   Snowflake `json:"guild_id,string,omitempty"`
}

//

type Attachment struct {
	ID       Snowflake `json:"id,string"`
	Filename string    `json:"filename"`
	Size     uint64    `json:"size"`

	URL   URL `json:"url"`
	Proxy URL `json:"proxy_url"`

	// Only if Image
	Height uint `json:"height,omitempty"`
	Width  uint `json:"width,omitempty"`
}

//

type Reaction struct {
	Count int   `json:"count"`
	Me    bool  `json:"me"` // for current user
	Emoji Emoji `json:"emoji"`
}
