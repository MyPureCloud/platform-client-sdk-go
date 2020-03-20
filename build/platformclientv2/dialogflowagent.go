package platformclientv2
import (
	"encoding/json"
)

// Dialogflowagent
type Dialogflowagent struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Project - The project this Dialogflow agent belongs to
	Project *Dialogflowproject `json:"project,omitempty"`


	// Languages - The target languages of the Dialogflow agent
	Languages *[]string `json:"languages,omitempty"`


	// Intents - An array of Intents associated with this bot alias
	Intents *[]Dialogflowintent `json:"intents,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Dialogflowagent) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
