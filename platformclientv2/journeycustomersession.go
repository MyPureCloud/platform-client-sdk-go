package platformclientv2
import (
	"encoding/json"
)

// Journeycustomersession
type Journeycustomersession struct { 
	// Id - An ID of a Customer/User's session within the Journey System at a point-in-time
	Id *string `json:"id,omitempty"`


	// VarType - The type of the Customer/User's session within the Journey System (e.g. web, app)
	VarType *string `json:"type,omitempty"`

}

// String returns a JSON representation of the model
func (o *Journeycustomersession) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
