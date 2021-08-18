package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Schedulegenerationmessage
type Schedulegenerationmessage struct { 
	// VarType - The type of the message
	VarType *string `json:"type,omitempty"`


	// Arguments - The arguments describing the message
	Arguments *[]Schedulermessageargument `json:"arguments,omitempty"`

}

func (u *Schedulegenerationmessage) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Schedulegenerationmessage

	

	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		Arguments *[]Schedulermessageargument `json:"arguments,omitempty"`
		*Alias
	}{ 
		VarType: u.VarType,
		
		Arguments: u.Arguments,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Schedulegenerationmessage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
