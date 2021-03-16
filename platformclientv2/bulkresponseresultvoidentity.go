package platformclientv2
import (
	"encoding/json"
)

// Bulkresponseresultvoidentity
type Bulkresponseresultvoidentity struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Success
	Success *bool `json:"success,omitempty"`


	// Entity
	Entity *Void `json:"entity,omitempty"`


	// VarError
	VarError *Bulkerrorentity `json:"error,omitempty"`

}

// String returns a JSON representation of the model
func (o *Bulkresponseresultvoidentity) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
