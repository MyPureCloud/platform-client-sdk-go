package platformclientv2
import (
	"time"
	"encoding/json"
)

// Dialerrulesetconfigchangeruleset
type Dialerrulesetconfigchangeruleset struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// DateCreated
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateModified
	DateModified *time.Time `json:"dateModified,omitempty"`


	// Version
	Version *int32 `json:"version,omitempty"`


	// ContactList
	ContactList *Dialerrulesetconfigchangeurireference `json:"contactList,omitempty"`


	// Queue
	Queue *Dialerrulesetconfigchangeurireference `json:"queue,omitempty"`


	// Rules
	Rules *[]Dialerrulesetconfigchangerule `json:"rules,omitempty"`


	// AdditionalProperties
	AdditionalProperties *map[string]interface{} `json:"additionalProperties,omitempty"`

}

// String returns a JSON representation of the model
func (o *Dialerrulesetconfigchangeruleset) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
