package platformclientv2
import (
	"encoding/json"
)

// Bulkresponseresultexternalcontactexternalcontact
type Bulkresponseresultexternalcontactexternalcontact struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Success
	Success *bool `json:"success,omitempty"`


	// Entity
	Entity *Externalcontact `json:"entity,omitempty"`


	// VarError
	VarError *Bulkerrorexternalcontact `json:"error,omitempty"`

}

// String returns a JSON representation of the model
func (o *Bulkresponseresultexternalcontactexternalcontact) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
