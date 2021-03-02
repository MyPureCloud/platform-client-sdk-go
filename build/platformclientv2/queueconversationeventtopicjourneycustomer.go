package platformclientv2
import (
	"encoding/json"
)

// Queueconversationeventtopicjourneycustomer
type Queueconversationeventtopicjourneycustomer struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// IdType
	IdType *string `json:"idType,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queueconversationeventtopicjourneycustomer) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
