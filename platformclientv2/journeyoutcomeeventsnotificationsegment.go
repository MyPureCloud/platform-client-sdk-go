package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Journeyoutcomeeventsnotificationsegment
type Journeyoutcomeeventsnotificationsegment struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// SelfUri
	SelfUri *string `json:"selfUri,omitempty"`


	// AssignedDate
	AssignedDate *time.Time `json:"assignedDate,omitempty"`

}

func (o *Journeyoutcomeeventsnotificationsegment) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Journeyoutcomeeventsnotificationsegment
	
	AssignedDate := new(string)
	if o.AssignedDate != nil {
		
		*AssignedDate = timeutil.Strftime(o.AssignedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		AssignedDate = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		
		AssignedDate *string `json:"assignedDate,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		SelfUri: o.SelfUri,
		
		AssignedDate: AssignedDate,
		Alias:    (*Alias)(o),
	})
}

func (o *Journeyoutcomeeventsnotificationsegment) UnmarshalJSON(b []byte) error {
	var JourneyoutcomeeventsnotificationsegmentMap map[string]interface{}
	err := json.Unmarshal(b, &JourneyoutcomeeventsnotificationsegmentMap)
	if err != nil {
		return err
	}
	
	if Id, ok := JourneyoutcomeeventsnotificationsegmentMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if SelfUri, ok := JourneyoutcomeeventsnotificationsegmentMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	
	if assignedDateString, ok := JourneyoutcomeeventsnotificationsegmentMap["assignedDate"].(string); ok {
		AssignedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", assignedDateString)
		o.AssignedDate = &AssignedDate
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Journeyoutcomeeventsnotificationsegment) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
