package models

// Guest struct
type Guest struct {
	EventID          string `json:"event_id"`
	UserID           string `json:"user_id"`
	Confirm          bool   `json:"confirm"`
	Old              int    `json:"old"`
	OldModificated   int    `json:"old_modificated"`
	Child            int    `json:"child"`
	ChildModificated int    `json:"child_modificated"`
	Total            int    `json:"total"`
}
