package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Messagingtemplaterequest
type Messagingtemplaterequest struct { 
	// ResponseId - A Response Management response identifier for a messaging template defined response
	ResponseId *string `json:"responseId,omitempty"`


	// Parameters - A list of Response Management response substitutions for the response's messaging template
	Parameters *[]Templateparameter `json:"parameters,omitempty"`

}

func (u *Messagingtemplaterequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Messagingtemplaterequest

	

	return json.Marshal(&struct { 
		ResponseId *string `json:"responseId,omitempty"`
		
		Parameters *[]Templateparameter `json:"parameters,omitempty"`
		*Alias
	}{ 
		ResponseId: u.ResponseId,
		
		Parameters: u.Parameters,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Messagingtemplaterequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
