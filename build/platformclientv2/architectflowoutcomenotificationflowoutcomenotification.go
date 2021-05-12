package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Architectflowoutcomenotificationflowoutcomenotification
type Architectflowoutcomenotificationflowoutcomenotification struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Description
	Description *string `json:"description,omitempty"`


	// CurrentOperation
	CurrentOperation *Architectflowoutcomenotificationarchitectoperation `json:"currentOperation,omitempty"`

}

// String returns a JSON representation of the model
func (o *Architectflowoutcomenotificationflowoutcomenotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
