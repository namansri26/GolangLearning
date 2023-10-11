package model

type NetFlix struct {
	ID        string `json:"_Id, omitempty" bson : "_id, omitempty"`
	Movie     string `json:"movie,omitempty"`
	isWatched bool   `json: "isWatched,omitempty"`
}
