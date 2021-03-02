package platformclientv2
import (
	"encoding/json"
)

// Queueconversationcalleventtopicjourneycustomer
type Queueconversationcalleventtopicjourneycustomer struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// IdType
	IdType *string `json:"idType,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queueconversationcalleventtopicjourneycustomer) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
