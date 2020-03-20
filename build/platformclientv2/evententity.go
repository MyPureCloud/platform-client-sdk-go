package platformclientv2
import (
	"encoding/json"
)

// Evententity
type Evententity struct { 
	// EntityType - Type of entity the event pertains to. e.g. integration
	EntityType *string `json:"entityType,omitempty"`


	// Id - ID of the entity the event pertains to.
	Id *string `json:"id,omitempty"`

}

// String returns a JSON representation of the model
func (o *Evententity) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
