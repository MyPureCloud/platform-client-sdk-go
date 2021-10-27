package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Journeyoutcomeeventsnotificationoutcomeattributionmessage
type Journeyoutcomeeventsnotificationoutcomeattributionmessage struct { 
	// Outcome
	Outcome *Journeyoutcomeeventsnotificationoutcome `json:"outcome,omitempty"`


	// OutcomeTouchpoints
	OutcomeTouchpoints *[]Journeyoutcomeeventsnotificationoutcometouchpoint `json:"outcomeTouchpoints,omitempty"`


	// SegmentAssignments
	SegmentAssignments *[]Journeyoutcomeeventsnotificationsegment `json:"segmentAssignments,omitempty"`

}

func (o *Journeyoutcomeeventsnotificationoutcomeattributionmessage) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Journeyoutcomeeventsnotificationoutcomeattributionmessage
	
	return json.Marshal(&struct { 
		Outcome *Journeyoutcomeeventsnotificationoutcome `json:"outcome,omitempty"`
		
		OutcomeTouchpoints *[]Journeyoutcomeeventsnotificationoutcometouchpoint `json:"outcomeTouchpoints,omitempty"`
		
		SegmentAssignments *[]Journeyoutcomeeventsnotificationsegment `json:"segmentAssignments,omitempty"`
		*Alias
	}{ 
		Outcome: o.Outcome,
		
		OutcomeTouchpoints: o.OutcomeTouchpoints,
		
		SegmentAssignments: o.SegmentAssignments,
		Alias:    (*Alias)(o),
	})
}

func (o *Journeyoutcomeeventsnotificationoutcomeattributionmessage) UnmarshalJSON(b []byte) error {
	var JourneyoutcomeeventsnotificationoutcomeattributionmessageMap map[string]interface{}
	err := json.Unmarshal(b, &JourneyoutcomeeventsnotificationoutcomeattributionmessageMap)
	if err != nil {
		return err
	}
	
	if Outcome, ok := JourneyoutcomeeventsnotificationoutcomeattributionmessageMap["outcome"].(map[string]interface{}); ok {
		OutcomeString, _ := json.Marshal(Outcome)
		json.Unmarshal(OutcomeString, &o.Outcome)
	}
	
	if OutcomeTouchpoints, ok := JourneyoutcomeeventsnotificationoutcomeattributionmessageMap["outcomeTouchpoints"].([]interface{}); ok {
		OutcomeTouchpointsString, _ := json.Marshal(OutcomeTouchpoints)
		json.Unmarshal(OutcomeTouchpointsString, &o.OutcomeTouchpoints)
	}
	
	if SegmentAssignments, ok := JourneyoutcomeeventsnotificationoutcomeattributionmessageMap["segmentAssignments"].([]interface{}); ok {
		SegmentAssignmentsString, _ := json.Marshal(SegmentAssignments)
		json.Unmarshal(SegmentAssignmentsString, &o.SegmentAssignments)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Journeyoutcomeeventsnotificationoutcomeattributionmessage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
