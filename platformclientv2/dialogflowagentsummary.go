package platformclientv2
import (
	"encoding/json"
)

// Dialogflowagentsummary
type Dialogflowagentsummary struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Project - The project this Dialogflow agent belongs to
	Project *Dialogflowproject `json:"project,omitempty"`


	// Description - A description of the Dialogflow agent
	Description *string `json:"description,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Dialogflowagentsummary) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
