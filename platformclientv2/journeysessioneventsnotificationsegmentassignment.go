package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Journeysessioneventsnotificationsegmentassignment
type Journeysessioneventsnotificationsegmentassignment struct { 
	// Segment
	Segment *Journeysessioneventsnotificationsegment `json:"segment,omitempty"`


	// AssignedDate
	AssignedDate *time.Time `json:"assignedDate,omitempty"`

}

func (o *Journeysessioneventsnotificationsegmentassignment) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Journeysessioneventsnotificationsegmentassignment
	
	AssignedDate := new(string)
	if o.AssignedDate != nil {
		
		*AssignedDate = timeutil.Strftime(o.AssignedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		AssignedDate = nil
	}
	
	return json.Marshal(&struct { 
		Segment *Journeysessioneventsnotificationsegment `json:"segment,omitempty"`
		
		AssignedDate *string `json:"assignedDate,omitempty"`
		*Alias
	}{ 
		Segment: o.Segment,
		
		AssignedDate: AssignedDate,
		Alias:    (*Alias)(o),
	})
}

func (o *Journeysessioneventsnotificationsegmentassignment) UnmarshalJSON(b []byte) error {
	var JourneysessioneventsnotificationsegmentassignmentMap map[string]interface{}
	err := json.Unmarshal(b, &JourneysessioneventsnotificationsegmentassignmentMap)
	if err != nil {
		return err
	}
	
	if Segment, ok := JourneysessioneventsnotificationsegmentassignmentMap["segment"].(map[string]interface{}); ok {
		SegmentString, _ := json.Marshal(Segment)
		json.Unmarshal(SegmentString, &o.Segment)
	}
	
	if assignedDateString, ok := JourneysessioneventsnotificationsegmentassignmentMap["assignedDate"].(string); ok {
		AssignedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", assignedDateString)
		o.AssignedDate = &AssignedDate
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Journeysessioneventsnotificationsegmentassignment) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
