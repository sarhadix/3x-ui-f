package model

// ClientDevice stores per-client device identifiers (HWID) to enforce device limits.
// It is populated when a subscription request includes an x-hwid header.
type ClientDevice struct {
	Id              int    `json:"id" gorm:"primaryKey;autoIncrement"`
	ClientEmail     string `json:"clientEmail" gorm:"index:idx_email_hwid,priority:1"`
	HWID            string `json:"hwid" gorm:"index:idx_email_hwid,priority:2"`
	DeviceOS        string `json:"deviceOs"`
	DeviceOSVersion string `json:"deviceOsVersion"`
	DeviceModel     string `json:"deviceModel"`
	UserAgent       string `json:"userAgent"`
	LastSeen        int64  `json:"lastSeen"`  // Unix ms of last subscription request
	CreatedAt       int64  `json:"createdAt"` // Unix ms of first sighting
}
