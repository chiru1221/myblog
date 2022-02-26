package data

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*
	basic data structure
*/
type Profile struct {
	Content string `bson:"content"`
	Date    string `bson:"date"`
}

type Blog struct {
	Id      primitive.ObjectID `bson:"_id"`
	Article string             `bson:"article"`
	Open    bool               `bson:"open"`
	Tag     []string           `bson:"tag"`
	Title   string             `bson:"title"`
	Date    string             `bson:"date"`
}

/*
	related slack
*/
type DialogRequest struct {
	TriggerId string `json:"trigger_id"`
	Form      Dialog `json:"dialog"`
}

type Dialog struct {
	Callback     string    `json:"callback_id"`
	Title        string    `json:"title"`
	Label        string    `json:"submit_label"`
	NotifyCancel bool      `json:"notify_on_cancel"`
	Elements     []Element `json:"elements"`
}

type Element struct {
	Type     string   `json:"type"`
	Label    string   `json:"label"`
	Name     string   `json:"name"`
	Hint     string   `json:"hint"`
	Subtype  string   `json:"subtype"`
	Optional bool     `json:"optional"`
	Options  []Option `json:"options"`
}

type Option struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

type Payload struct {
	Body        SubmitBody `json:"submission"`
	CallbackId  string     `json:"callback_id"`
	ResponseURL string     `json:"response_url"`
}

type SubmitBody struct {
	Id      string `json:id`
	Title   string `json:"title"`
	Tag     string `json:"tag"`
	Article string `json:"article"`
	Open    string `json:"open"`
	Date    string `json:"date"`
}
