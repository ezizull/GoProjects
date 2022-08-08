package model

// Employee exported
type Employee struct {
	ID		string		`json:"id,omitempty" bson:"_id,omitempty"`
	Name	string		`json:"name,omitempty"`
	Salary	float64		`json:"salary,omitempty"`
	Age		float64		`json:"age,omitempty"`
}
