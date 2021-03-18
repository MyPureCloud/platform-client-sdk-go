package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationemaileventtopicjourneycustomer
type Queueconversationemaileventtopicjourneycustomer struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// IdType
	IdType *string `json:"idType,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queueconversationemaileventtopicjourneycustomer) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
