package entities

type SiteConfigs struct {
	ID               string   `bson:"_id,omitempty" json:"_id"`
	SiteName         string   `bson:"site_name,omitempty" json:"site_name"`
	ForwardingURL    string   `bson:"forwarding_url,omitempty" json:"forwarding_url,omitempty"`
	BlockedOS        string   `bson:"blocked_os" json:"blocked_os"`
	BlockedBrowser   string   `bson:"blocked_browser" json:"blocked_browser"`
	BlockedDevice    string   `bson:"blocked_device" json:"blocked_device"`
	BlockedOSVersion string   `bson:"blocked_os_version" json:"blocked_os_version"`
	BlockedLocations string   `bson:"blocked_locations,omitempty" json:"blocked_locations"`
	BlockedIP        []string `bson:"blocked_ip,omitempty" json:"blocked_ip"`
	CreatedAt        string   `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt        string   `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}
