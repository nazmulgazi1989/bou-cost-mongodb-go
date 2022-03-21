package models

type Cost struct {
	ID              string  `json:"id,omitempty" bson:"_id,omitempty"`
	Cost            float64 `json:"cost"`
	CostType        string  `json:"costType"`
	CostDate        string  `json:"costDate"`
	CostDescription string  `json:"costDescription"`
}