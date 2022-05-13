package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Sessionlastevent
type Sessionlastevent struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// EventName - The name of the event.
	EventName *string `json:"eventName,omitempty"`


	// CreatedDate - Timestamp indicating when the event was published. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	CreatedDate *time.Time `json:"createdDate,omitempty"`

}

func (o *Sessionlastevent) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Sessionlastevent
	
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

func (o *Sessionlastevent) UnmarshalJSON(b []byte) error {
	var SessionlasteventMap map[string]interface{}
	err := json.Unmarshal(b, &SessionlasteventMap)
	if err != nil {
		return err
	}
	
	if Id, ok := SessionlasteventMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if EventName, ok := SessionlasteventMap["eventName"].(string); ok {
		o.EventName = &EventName
	}
    
	if createdDateString, ok := SessionlasteventMap["createdDate"].(string); ok {
		CreatedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", createdDateString)
		o.CreatedDate = &CreatedDate
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Sessionlastevent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
