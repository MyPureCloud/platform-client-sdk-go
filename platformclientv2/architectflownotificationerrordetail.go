package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Architectflownotificationerrordetail
type Architectflownotificationerrordetail struct { 
	// ErrorCode
	ErrorCode *string `json:"errorCode,omitempty"`


	// EntityId
	EntityId *string `json:"entityId,omitempty"`


	// EntityName
	EntityName *string `json:"entityName,omitempty"`


	// FieldName
	FieldName *string `json:"fieldName,omitempty"`

}

func (u *Architectflownotificationerrordetail) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Architectflownotificationerrordetail

	

	return json.Marshal(&struct { 
		ErrorCode *string `json:"errorCode,omitempty"`
		
		EntityId *string `json:"entityId,omitempty"`
		
		EntityName *string `json:"entityName,omitempty"`
		
		FieldName *string `json:"fieldName,omitempty"`
		*Alias
	}{ 
		ErrorCode: u.ErrorCode,
		
		EntityId: u.EntityId,
		
		EntityName: u.EntityName,
		
		FieldName: u.FieldName,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Architectflownotificationerrordetail) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
