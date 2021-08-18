package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Bulkerrordetail
type Bulkerrordetail struct { 
	// FieldName
	FieldName *string `json:"fieldName,omitempty"`


	// Value
	Value *string `json:"value,omitempty"`


	// Message
	Message *string `json:"message,omitempty"`

}

func (u *Bulkerrordetail) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Bulkerrordetail

	

	return json.Marshal(&struct { 
		FieldName *string `json:"fieldName,omitempty"`
		
		Value *string `json:"value,omitempty"`
		
		Message *string `json:"message,omitempty"`
		*Alias
	}{ 
		FieldName: u.FieldName,
		
		Value: u.Value,
		
		Message: u.Message,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Bulkerrordetail) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
