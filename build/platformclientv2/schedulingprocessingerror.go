package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Schedulingprocessingerror) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
