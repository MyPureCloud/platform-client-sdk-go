package platformclientv2
import (
	"encoding/json"
)

// Bulkresponseresultexternalcontactentity
type Bulkresponseresultexternalcontactentity struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Success
	Success *bool `json:"success,omitempty"`


	// Entity
	Entity *Externalcontact `json:"entity,omitempty"`


	// VarError
	VarError *Bulkerrorentity `json:"error,omitempty"`

}

// String returns a JSON representation of the model
func (o *Bulkresponseresultexternalcontactentity) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
