package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationsocialexpressioneventtopicdisconnectreason
type Queueconversationsocialexpressioneventtopicdisconnectreason struct { 
	// VarType
	VarType *string `json:"type,omitempty"`


	// Code
	Code *int `json:"code,omitempty"`


	// Phrase
	Phrase *string `json:"phrase,omitempty"`

}

func (u *Queueconversationsocialexpressioneventtopicdisconnectreason) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueconversationsocialexpressioneventtopicdisconnectreason

	

	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		Code *int `json:"code,omitempty"`
		
		Phrase *string `json:"phrase,omitempty"`
		*Alias
	}{ 
		VarType: u.VarType,
		
		Code: u.Code,
		
		Phrase: u.Phrase,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Queueconversationsocialexpressioneventtopicdisconnectreason) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
