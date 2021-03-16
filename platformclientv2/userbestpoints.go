package platformclientv2
import (
	"encoding/json"
)

// Userbestpoints
type Userbestpoints struct { 
	// User - The requested user for the best points
	User *Userreference `json:"user,omitempty"`


	// BestPoints - List of best point for the requested user
	BestPoints *[]Userbestpointsitem `json:"bestPoints,omitempty"`

}

// String returns a JSON representation of the model
func (o *Userbestpoints) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
