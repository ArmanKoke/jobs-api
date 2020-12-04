package responses

// Create struct
// swagger:parameters CreateResponse
type Create struct {
	Saved bool
}

// Show struct
// swagger:parameters ShowResponse
type Show struct {
	ID      int64   `db:"id" json:"id"`
	Value   string  `db:"val" json:"value"`
	Comment *string `db:"comment" json:"comment"`
}
