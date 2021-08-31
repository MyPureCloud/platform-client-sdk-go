package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Calendarurlresponse
type Calendarurlresponse struct { 
	// CalendarUrl - The calendar url for the user to subscribe with supported clients
	CalendarUrl *string `json:"calendarUrl,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Calendarurlresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Calendarurlresponse
	
	return json.Marshal(&struct { 
		CalendarUrl *string `json:"calendarUrl,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		CalendarUrl: o.CalendarUrl,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Calendarurlresponse) UnmarshalJSON(b []byte) error {
	var CalendarurlresponseMap map[string]interface{}
	err := json.Unmarshal(b, &CalendarurlresponseMap)
	if err != nil {
		return err
	}
	
	if CalendarUrl, ok := CalendarurlresponseMap["calendarUrl"].(string); ok {
		o.CalendarUrl = &CalendarUrl
	}
	
	if SelfUri, ok := CalendarurlresponseMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Calendarurlresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
