package platformclientv2
import (
	"encoding/json"
)

// Querydivision
type Querydivision struct { }

// String returns a JSON representation of the model
func (o *Querydivision) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
