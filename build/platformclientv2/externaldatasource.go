package platformclientv2
import (
	"encoding/json"
)

// Externaldatasource - Describes a link to a record in an external system that contributed data to a Relate record
type Externaldatasource struct { 
	// Platform - The platform that was the source of the data.  Example: a CRM like SALESFORCE.
	Platform *string `json:"platform,omitempty"`


	// Url - An URL that links to the source record that contributed data to the associated entity.
	Url *string `json:"url,omitempty"`

}

// String returns a JSON representation of the model
func (o *Externaldatasource) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
