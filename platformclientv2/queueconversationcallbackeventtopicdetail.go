package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationcallbackeventtopicdetail
type Queueconversationcallbackeventtopicdetail struct { 
	// ErrorCode
	ErrorCode *string `json:"errorCode,omitempty"`


	// FieldName
	FieldName *string `json:"fieldName,omitempty"`


	// EntityId
	EntityId *string `json:"entityId,omitempty"`


	// EntityName
	EntityName *string `json:"entityName,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queueconversationcallbackeventtopicdetail) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
