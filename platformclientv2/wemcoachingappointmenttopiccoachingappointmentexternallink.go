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

func (o *Wemcoachingappointmenttopiccoachingappointmentexternallink) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wemcoachingappointmenttopiccoachingappointmentexternallink
	
	return json.Marshal(&struct { 
		ExternalLink *string `json:"externalLink,omitempty"`
		
		Action *string `json:"action,omitempty"`
		*Alias
	}{ 
		ExternalLink: o.ExternalLink,
		
		Action: o.Action,
		Alias:    (*Alias)(o),
	})
}

func (o *Wemcoachingappointmenttopiccoachingappointmentexternallink) UnmarshalJSON(b []byte) error {
	var WemcoachingappointmenttopiccoachingappointmentexternallinkMap map[string]interface{}
	err := json.Unmarshal(b, &WemcoachingappointmenttopiccoachingappointmentexternallinkMap)
	if err != nil {
		return err
	}
	
	if ExternalLink, ok := WemcoachingappointmenttopiccoachingappointmentexternallinkMap["externalLink"].(string); ok {
		o.ExternalLink = &ExternalLink
	}
	
	if Action, ok := WemcoachingappointmenttopiccoachingappointmentexternallinkMap["action"].(string); ok {
		o.Action = &Action
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Wemcoachingappointmenttopiccoachingappointmentexternallink) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
