package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Analyticsconversationasyncqueryresponse
type Analyticsconversationasyncqueryresponse struct { 
	// Cursor - Optional cursor to indicate where to resume the results
	Cursor *string `json:"cursor,omitempty"`


	// DataAvailabilityDate - Data available up to at least this datetime. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DataAvailabilityDate *time.Time `json:"dataAvailabilityDate,omitempty"`


	// Conversations
	Conversations *[]Analyticsconversation `json:"conversations,omitempty"`

}

func (o *Analyticsconversationasyncqueryresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Analyticsconversationasyncqueryresponse
	
	DataAvailabilityDate := new(string)
	if o.DataAvailabilityDate != nil {
		
		*DataAvailabilityDate = timeutil.Strftime(o.DataAvailabilityDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DataAvailabilityDate = nil
	}
	
	return json.Marshal(&struct { 
		Cursor *string `json:"cursor,omitempty"`
		
		DataAvailabilityDate *string `json:"dataAvailabilityDate,omitempty"`
		
		Conversations *[]Analyticsconversation `json:"conversations,omitempty"`
		*Alias
	}{ 
		Cursor: o.Cursor,
		
		DataAvailabilityDate: DataAvailabilityDate,
		
		Conversations: o.Conversations,
		Alias:    (*Alias)(o),
	})
}

func (o *Analyticsconversationasyncqueryresponse) UnmarshalJSON(b []byte) error {
	var AnalyticsconversationasyncqueryresponseMap map[string]interface{}
	err := json.Unmarshal(b, &AnalyticsconversationasyncqueryresponseMap)
	if err != nil {
		return err
	}
	
	if Cursor, ok := AnalyticsconversationasyncqueryresponseMap["cursor"].(string); ok {
		o.Cursor = &Cursor
	}
	
	if dataAvailabilityDateString, ok := AnalyticsconversationasyncqueryresponseMap["dataAvailabilityDate"].(string); ok {
		DataAvailabilityDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", dataAvailabilityDateString)
		o.DataAvailabilityDate = &DataAvailabilityDate
	}
	
	if Conversations, ok := AnalyticsconversationasyncqueryresponseMap["conversations"].([]interface{}); ok {
		ConversationsString, _ := json.Marshal(Conversations)
		json.Unmarshal(ConversationsString, &o.Conversations)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Analyticsconversationasyncqueryresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
