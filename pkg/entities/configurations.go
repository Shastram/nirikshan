package entities

type SiteConfigs struct {
	ID               string   `bson:"_id,omitempty" json:"_id"`
	SiteName         string   `bson:"site_name,omitempty" json:"site_name"`
	ForwardingURL    string   `bson:"forwarding_url,omitempty" json:"forwarding_url,omitempty"`
	BlockedUserAgent string   `bson:"blocked_user_agent,omitempty" json:"blocked_user_agent"`
	BlockedLocations string   `bson:"blocked_locations,omitempty" json:"blocked_locations"`
	BlockedIP        []string `bson:"blocked_ip,omitempty" json:"blocked_ip"`
	CreatedAt        string   `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt        string   `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}
