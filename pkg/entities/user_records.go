package entities

import "time"

type UserRecords struct {
	ID       string    `json:"id,omitempty" bson:"id"`
	SiteID   string    `json:"site_id,omitempty" bson:"site_id"`
	SiteName string    `json:"site_name,omitempty" bson:"site_name"`
	Device   string    `json:"device,omitempty" bson:"device"`
	Os       string    `json:"os,omitempty" bson:"os"`
	Browser  string    `json:"browser,omitempty" bson:"browser"`
	IP       string    `json:"ip,omitempty" bson:"ip"`
	Time     time.Time `json:"time,omitempty" bson:"time"`
}
