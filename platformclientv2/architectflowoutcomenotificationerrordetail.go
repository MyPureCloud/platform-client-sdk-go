package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Architectflowoutcomenotificationerrordetail
type Architectflowoutcomenotificationerrordetail struct { 
	// ErrorCode
	ErrorCode *string `json:"errorCode,omitempty"`


	// EntityId
	EntityId *string `json:"entityId,omitempty"`


	// EntityName
	EntityName *string `json:"entityName,omitempty"`


	// FieldName
	FieldName *string `json:"fieldName,omitempty"`

}

// String returns a JSON representation of the model
func (o *Architectflowoutcomenotificationerrordetail) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
