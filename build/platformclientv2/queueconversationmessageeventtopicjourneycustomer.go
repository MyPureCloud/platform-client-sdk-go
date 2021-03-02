package platformclientv2
import (
	"encoding/json"
)

// Queueconversationmessageeventtopicjourneycustomer
type Queueconversationmessageeventtopicjourneycustomer struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// IdType
	IdType *string `json:"idType,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queueconversationmessageeventtopicjourneycustomer) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
