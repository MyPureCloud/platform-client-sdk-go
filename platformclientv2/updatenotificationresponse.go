package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Updatenotificationresponse
type Updatenotificationresponse struct { 
	// MutableGroupId - The mutableGroupId of the notification
	MutableGroupId *string `json:"mutableGroupId,omitempty"`


	// Id - The id of the notification for mapping the potentially new mutableGroupId
	Id *string `json:"id,omitempty"`

}

func (u *Updatenotificationresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Updatenotificationresponse

	

	return json.Marshal(&struct { 
		MutableGroupId *string `json:"mutableGroupId,omitempty"`
		
		Id *string `json:"id,omitempty"`
		*Alias
	}{ 
		MutableGroupId: u.MutableGroupId,
		
		Id: u.Id,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Updatenotificationresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
