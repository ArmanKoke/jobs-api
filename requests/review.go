package requests

import (
	"regexp"
)

// Review input
// swagger:parameters CreateRequest
type Review struct {
	ID      string  `json:"id"`
	Value   string  `json:"value"`
	Comment *string `json:"comment"`
}

// Validate request
func (r *Review) Validate() {
	re := regexp.MustCompile("(?i)^null\\b$") //match null in any case
	if re.MatchString(*r.Comment) || *r.Comment == "" {
		r.Comment = nil
	}
}
