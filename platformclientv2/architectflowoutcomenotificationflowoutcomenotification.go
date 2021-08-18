package platformclientv2
import (
	"github.com/leekchan/timeutil"
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

func (u *Architectflowoutcomenotificationflowoutcomenotification) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Architectflowoutcomenotificationflowoutcomenotification

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		CurrentOperation *Architectflowoutcomenotificationarchitectoperation `json:"currentOperation,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		Description: u.Description,
		
		CurrentOperation: u.CurrentOperation,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Architectflowoutcomenotificationflowoutcomenotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
