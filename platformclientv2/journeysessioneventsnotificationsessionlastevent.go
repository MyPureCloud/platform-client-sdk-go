package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Journeysessioneventsnotificationsessionlastevent
type Journeysessioneventsnotificationsessionlastevent struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// EventName
	EventName *string `json:"eventName,omitempty"`


	// CreatedDate
	CreatedDate *time.Time `json:"createdDate,omitempty"`

}

func (o *Journeysessioneventsnotificationsessionlastevent) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Journeysessioneventsnotificationsessionlastevent
	
	CreatedDate := new(string)
	if o.CreatedDate != nil {
		
		*CreatedDate = timeutil.Strftime(o.CreatedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		CreatedDate = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		EventName *string `json:"eventName,omitempty"`
		
		CreatedDate *string `json:"createdDate,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		EventName: o.EventName,
		
		CreatedDate: CreatedDate,
		Alias:    (*Alias)(o),
	})
}

func (o *Journeysessioneventsnotificationsessionlastevent) UnmarshalJSON(b []byte) error {
	var JourneysessioneventsnotificationsessionlasteventMap map[string]interface{}
	err := json.Unmarshal(b, &JourneysessioneventsnotificationsessionlasteventMap)
	if err != nil {
		return err
	}
	
	if Id, ok := JourneysessioneventsnotificationsessionlasteventMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if EventName, ok := JourneysessioneventsnotificationsessionlasteventMap["eventName"].(string); ok {
		o.EventName = &EventName
	}
    
	if createdDateString, ok := JourneysessioneventsnotificationsessionlasteventMap["createdDate"].(string); ok {
		CreatedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", createdDateString)
		o.CreatedDate = &CreatedDate
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Journeysessioneventsnotificationsessionlastevent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
