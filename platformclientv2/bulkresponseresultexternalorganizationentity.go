package platformclientv2
import (
	"encoding/json"
)

// Bulkresponseresultexternalorganizationentity
type Bulkresponseresultexternalorganizationentity struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Success
	Success *bool `json:"success,omitempty"`


	// Entity
	Entity *Externalorganization `json:"entity,omitempty"`


	// VarError
	VarError *Bulkerrorentity `json:"error,omitempty"`

}

// String returns a JSON representation of the model
func (o *Bulkresponseresultexternalorganizationentity) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
