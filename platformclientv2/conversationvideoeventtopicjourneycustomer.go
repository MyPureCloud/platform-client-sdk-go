package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationvideoeventtopicjourneycustomer
type Conversationvideoeventtopicjourneycustomer struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// IdType
	IdType *string `json:"idType,omitempty"`

}

// String returns a JSON representation of the model
func (o *Conversationvideoeventtopicjourneycustomer) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
