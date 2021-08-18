package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Schedulermessageargument
type Schedulermessageargument struct { 
	// VarType - The type of this message parameter
	VarType *string `json:"type,omitempty"`


	// Value - The value of this message parameter
	Value *string `json:"value,omitempty"`

}

func (u *Schedulermessageargument) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Schedulermessageargument

	

	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		Value *string `json:"value,omitempty"`
		*Alias
	}{ 
		VarType: u.VarType,
		
		Value: u.Value,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Schedulermessageargument) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
