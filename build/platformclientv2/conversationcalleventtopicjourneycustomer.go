package platformclientv2
import (
	"encoding/json"
)

// Conversationcalleventtopicjourneycustomer
type Conversationcalleventtopicjourneycustomer struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// IdType
	IdType *string `json:"idType,omitempty"`

}

// String returns a JSON representation of the model
func (o *Conversationcalleventtopicjourneycustomer) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
