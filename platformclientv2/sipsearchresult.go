package platformclientv2
import (
	"encoding/json"
)

// Sipsearchresult
type Sipsearchresult struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Status - Status of the search request
	Status *int `json:"status,omitempty"`


	// Sid - Session id associated to the search request
	Sid *string `json:"sid,omitempty"`


	// Auth - Auth token used for this search request
	Auth *string `json:"auth,omitempty"`


	// Message - Any messages returned from homer as part of the response
	Message *string `json:"message,omitempty"`


	// Data - Homer search data that is returned
	Data *[]Homerrecord `json:"data,omitempty"`


	// Count - Number of records returned
	Count *int `json:"count,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Sipsearchresult) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
