package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wemcoachingappointmenttopiccoachingappointmentdocument
type Wemcoachingappointmenttopiccoachingappointmentdocument struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Action
	Action *string `json:"action,omitempty"`

}

func (u *Wemcoachingappointmenttopiccoachingappointmentdocument) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wemcoachingappointmenttopiccoachingappointmentdocument

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Action *string `json:"action,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Action: u.Action,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Wemcoachingappointmenttopiccoachingappointmentdocument) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
