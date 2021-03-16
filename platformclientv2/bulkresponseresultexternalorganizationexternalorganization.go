package platformclientv2
import (
	"encoding/json"
)

// Bulkresponseresultexternalorganizationexternalorganization
type Bulkresponseresultexternalorganizationexternalorganization struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Success
	Success *bool `json:"success,omitempty"`


	// Entity
	Entity *Externalorganization `json:"entity,omitempty"`


	// VarError
	VarError *Bulkerrorexternalorganization `json:"error,omitempty"`

}

// String returns a JSON representation of the model
func (o *Bulkresponseresultexternalorganizationexternalorganization) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
