package platformclientv2
import (
	"github.com/leekchan/timeutil"
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

func (u *Architectpromptnotificationpromptnotification) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Architectpromptnotificationpromptnotification

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		CurrentOperation *Architectpromptnotificationarchitectoperation `json:"currentOperation,omitempty"`
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
func (o *Architectpromptnotificationpromptnotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
