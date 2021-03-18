package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Architectpromptnotificationpromptnotification
type Architectpromptnotificationpromptnotification struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Description
	Description *string `json:"description,omitempty"`


	// CurrentOperation
	CurrentOperation *Architectpromptnotificationarchitectoperation `json:"currentOperation,omitempty"`

}

// String returns a JSON representation of the model
func (o *Architectpromptnotificationpromptnotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
