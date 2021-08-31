package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Analyticsresolution
type Analyticsresolution struct { 
	// EventTime - Specifies when an event occurred. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	EventTime *time.Time `json:"eventTime,omitempty"`


	// QueueId - The ID of the last queue on which the conversation was handled.
	QueueId *string `json:"queueId,omitempty"`


	// UserId - The ID of the last user who handled the conversation.
	UserId *string `json:"userId,omitempty"`


	// NNextContactAvoided
	NNextContactAvoided *int `json:"nNextContactAvoided,omitempty"`

}

func (o *Analyticsresolution) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Analyticsresolution
	
	EventTime := new(string)
	if o.EventTime != nil {
		
		*EventTime = timeutil.Strftime(o.EventTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		EventTime = nil
	}
	
	return json.Marshal(&struct { 
		EventTime *string `json:"eventTime,omitempty"`
		
		QueueId *string `json:"queueId,omitempty"`
		
		UserId *string `json:"userId,omitempty"`
		
		NNextContactAvoided *int `json:"nNextContactAvoided,omitempty"`
		*Alias
	}{ 
		EventTime: EventTime,
		
		QueueId: o.QueueId,
		
		UserId: o.UserId,
		
		NNextContactAvoided: o.NNextContactAvoided,
		Alias:    (*Alias)(o),
	})
}

func (o *Analyticsresolution) UnmarshalJSON(b []byte) error {
	var AnalyticsresolutionMap map[string]interface{}
	err := json.Unmarshal(b, &AnalyticsresolutionMap)
	if err != nil {
		return err
	}
	
	if eventTimeString, ok := AnalyticsresolutionMap["eventTime"].(string); ok {
		EventTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", eventTimeString)
		o.EventTime = &EventTime
	}
	
	if QueueId, ok := AnalyticsresolutionMap["queueId"].(string); ok {
		o.QueueId = &QueueId
	}
	
	if UserId, ok := AnalyticsresolutionMap["userId"].(string); ok {
		o.UserId = &UserId
	}
	
	if NNextContactAvoided, ok := AnalyticsresolutionMap["nNextContactAvoided"].(float64); ok {
		NNextContactAvoidedInt := int(NNextContactAvoided)
		o.NNextContactAvoided = &NNextContactAvoidedInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Analyticsresolution) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
