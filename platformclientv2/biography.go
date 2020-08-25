package platformclientv2
import (
	"encoding/json"
)

// Biography
type Biography struct { 
	// Biography - Personal detailed description
	Biography *string `json:"biography,omitempty"`


	// Interests
	Interests *[]string `json:"interests,omitempty"`


	// Hobbies
	Hobbies *[]string `json:"hobbies,omitempty"`


	// Spouse
	Spouse *string `json:"spouse,omitempty"`


	// Education - User education details
	Education *[]Education `json:"education,omitempty"`

}

// String returns a JSON representation of the model
func (o *Biography) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
