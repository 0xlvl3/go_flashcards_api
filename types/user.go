package types

type User struct {
	ID             string `bson:"_id,omitempty" json:"id,omitempty"`
	Username       string `bson:"username" json:"username"`
	TotalCardCount int    `bson:"totalCardCount" json:"totalCardCount"`
}
