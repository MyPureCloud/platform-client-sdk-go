package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wemcoachingappointmenttopiccoachingappointmentexternallink
type Wemcoachingappointmenttopiccoachingappointmentexternallink struct { 
	// ExternalLink
	ExternalLink *string `json:"externalLink,omitempty"`


	// Action
	Action *string `json:"action,omitempty"`

}

func (u *Wemcoachingappointmenttopiccoachingappointmentexternallink) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wemcoachingappointmenttopiccoachingappointmentexternallink

	

	return json.Marshal(&struct { 
		ExternalLink *string `json:"externalLink,omitempty"`
		
		Action *string `json:"action,omitempty"`
		*Alias
	}{ 
		ExternalLink: u.ExternalLink,
		
		Action: u.Action,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Wemcoachingappointmenttopiccoachingappointmentexternallink) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
