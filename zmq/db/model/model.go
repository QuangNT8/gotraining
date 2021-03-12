package model

import "time"

type User struct {
	/* ID will get from keystone*/
	ID string `sql:"size:64;unique;index" json:"id" gorm:"primary_key"`
	/* Equal with "name" of keystone auth request*/
	Name         string         `sql:"size:256;unique;index" json:"name"`
	UserName     string         `json:"user_name"`
	Role         string         `json:"role"`
	Email        string         `json:"email"`
	Phone        string         `json:"phone"`
	Enabled      bool           `json:"enabled"`
	LastActiveAt *time.Time     `json:"last_active_at"`
	LastLoginAt  *time.Time     `json:"last_login_at"`
	Option       UserOption     `json:"option"`
	Whilelist    WhileList      `json:"whilelist"`
	Deleted      bool           `json:"deleted" gorm:"type:boolean;default:false"`
	AuditIDs     string         `json:"audit_ids"`
	Retention    *DataRetention `json:"retention"`
}

type UserOption struct {
	UserID   string `json:"user_id" gorm:"primary_key"`
	Lang     string `json:"lang"`
	Timezone string `json:"timezone"`
	Alert    string `json:"alert"`
	// for alert
	Email string `json:"email"`
	Freq  int64  `gorm:"default:5" json:"freq"`
	Iface string `json:"iface"`
	Pps   bool   `gorm:"default:false" json:"pps"`
}

type WhileList struct {
	UserID string `json:"user_id" gorm:"primary_key"`
	/* list ip will store with format 192.168.1.2,10.1.2.3 */
	Ips string `json:"ips"`
}

type DataRetention struct {
	UserID           string `json:"user_id" gorm:"primary_key"`
	RetentionPerDays int    `gorm:"type:int;default:15" json:"retention_per_days"`
}

type UserUI struct {
	ID       uint   `gorm:"primary_key"`
	UserID   string `json:"user_id"`
	ScreenID string `json:"screen_id"`
	H        uint32 `json:"h"`
	W        uint32 `json:"w"`
	Y        uint32 `json:"y"`
	X        uint32 `json:"x"`
	/*chart_name will store with format: xs_chart_name or lg_chart_name or xxs_chart_name or sm_chart_name*/
	Info string `sql:"size:256" json:"info"`
}
