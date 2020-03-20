package platformclientv2
import (
	"encoding/json"
)

// Qualityaudit
type Qualityaudit struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// User
	User *User `json:"user,omitempty"`


	// JobId
	JobId *string `json:"jobId,omitempty"`


	// Level
	Level *string `json:"level,omitempty"`


	// Entity
	Entity *Auditentity `json:"entity,omitempty"`


	// Timestamp
	Timestamp *string `json:"timestamp,omitempty"`


	// Action
	Action *string `json:"action,omitempty"`


	// Status
	Status *string `json:"status,omitempty"`


	// Changes
	Changes *[]Change `json:"changes,omitempty"`


	// EntityType
	EntityType *string `json:"entityType,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Qualityaudit) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
