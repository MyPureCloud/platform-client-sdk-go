package platformclientv2
import (
	"encoding/json"
)

// Conversationsocialexpressioneventtopicjourneycustomer
type Conversationsocialexpressioneventtopicjourneycustomer struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// IdType
	IdType *string `json:"idType,omitempty"`

}

// String returns a JSON representation of the model
func (o *Conversationsocialexpressioneventtopicjourneycustomer) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
