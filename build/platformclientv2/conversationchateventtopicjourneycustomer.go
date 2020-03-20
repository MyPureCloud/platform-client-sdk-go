package platformclientv2
import (
	"encoding/json"
)

// Conversationchateventtopicjourneycustomer
type Conversationchateventtopicjourneycustomer struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// IdType
	IdType *string `json:"idType,omitempty"`

}

// String returns a JSON representation of the model
func (o *Conversationchateventtopicjourneycustomer) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
