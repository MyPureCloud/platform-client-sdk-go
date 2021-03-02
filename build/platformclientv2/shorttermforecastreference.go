package platformclientv2
import (
	"encoding/json"
)

// Shorttermforecastreference - A pointer to a short term forecast
type Shorttermforecastreference struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`


	// WeekDate - The weekDate of the short term forecast in yyyy-MM-dd format
	WeekDate *string `json:"weekDate,omitempty"`


	// Description - The description of the short term forecast
	Description *string `json:"description,omitempty"`

}

// String returns a JSON representation of the model
func (o *Shorttermforecastreference) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
