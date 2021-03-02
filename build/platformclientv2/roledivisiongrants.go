package platformclientv2
import (
	"encoding/json"
)

// Roledivisiongrants
type Roledivisiongrants struct { 
	// Grants - A list containing pairs of role and division IDs
	Grants *[]Roledivisionpair `json:"grants,omitempty"`

}

// String returns a JSON representation of the model
func (o *Roledivisiongrants) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
