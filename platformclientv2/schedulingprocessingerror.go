package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Schedulingprocessingerror
type Schedulingprocessingerror struct { 
	// InternalErrorCode - An internal code representing the type of error. BadJson for 'Unable to parse json.' NotFound for 'Resource not found.' Fail for 'An unexpected server error occured.'
	InternalErrorCode *string `json:"internalErrorCode,omitempty"`


	// Description - A text description of the error
	Description *string `json:"description,omitempty"`

}

func (u *Schedulingprocessingerror) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Schedulingprocessingerror

	

	return json.Marshal(&struct { 
		InternalErrorCode *string `json:"internalErrorCode,omitempty"`
		
		Description *string `json:"description,omitempty"`
		*Alias
	}{ 
		InternalErrorCode: u.InternalErrorCode,
		
		Description: u.Description,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Schedulingprocessingerror) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
