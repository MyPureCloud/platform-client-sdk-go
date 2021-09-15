package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Sessionsegmentassignment
type Sessionsegmentassignment struct { 
	// Segment - The segment that was assigned.
	Segment *Assignedsegment `json:"segment,omitempty"`


	// AssignedDate - Timestamp indicating when the segment was assigned. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	AssignedDate *time.Time `json:"assignedDate,omitempty"`

}

func (o *Sessionsegmentassignment) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Sessionsegmentassignment
	
	AssignedDate := new(string)
	if o.AssignedDate != nil {
		
		*AssignedDate = timeutil.Strftime(o.AssignedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		AssignedDate = nil
	}
	
	return json.Marshal(&struct { 
		Segment *Assignedsegment `json:"segment,omitempty"`
		
		AssignedDate *string `json:"assignedDate,omitempty"`
		*Alias
	}{ 
		Segment: o.Segment,
		
		AssignedDate: AssignedDate,
		Alias:    (*Alias)(o),
	})
}

func (o *Sessionsegmentassignment) UnmarshalJSON(b []byte) error {
	var SessionsegmentassignmentMap map[string]interface{}
	err := json.Unmarshal(b, &SessionsegmentassignmentMap)
	if err != nil {
		return err
	}
	
	if Segment, ok := SessionsegmentassignmentMap["segment"].(map[string]interface{}); ok {
		SegmentString, _ := json.Marshal(Segment)
		json.Unmarshal(SegmentString, &o.Segment)
	}
	
	if assignedDateString, ok := SessionsegmentassignmentMap["assignedDate"].(string); ok {
		AssignedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", assignedDateString)
		o.AssignedDate = &AssignedDate
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Sessionsegmentassignment) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
