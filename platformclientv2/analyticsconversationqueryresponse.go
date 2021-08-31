package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Analyticsconversationqueryresponse
type Analyticsconversationqueryresponse struct { 
	// Aggregations
	Aggregations *[]Aggregationresult `json:"aggregations,omitempty"`


	// Conversations
	Conversations *[]Analyticsconversationwithoutattributes `json:"conversations,omitempty"`


	// TotalHits
	TotalHits *int `json:"totalHits,omitempty"`

}

func (o *Analyticsconversationqueryresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Analyticsconversationqueryresponse
	
	return json.Marshal(&struct { 
		Aggregations *[]Aggregationresult `json:"aggregations,omitempty"`
		
		Conversations *[]Analyticsconversationwithoutattributes `json:"conversations,omitempty"`
		
		TotalHits *int `json:"totalHits,omitempty"`
		*Alias
	}{ 
		Aggregations: o.Aggregations,
		
		Conversations: o.Conversations,
		
		TotalHits: o.TotalHits,
		Alias:    (*Alias)(o),
	})
}

func (o *Analyticsconversationqueryresponse) UnmarshalJSON(b []byte) error {
	var AnalyticsconversationqueryresponseMap map[string]interface{}
	err := json.Unmarshal(b, &AnalyticsconversationqueryresponseMap)
	if err != nil {
		return err
	}
	
	if Aggregations, ok := AnalyticsconversationqueryresponseMap["aggregations"].([]interface{}); ok {
		AggregationsString, _ := json.Marshal(Aggregations)
		json.Unmarshal(AggregationsString, &o.Aggregations)
	}
	
	if Conversations, ok := AnalyticsconversationqueryresponseMap["conversations"].([]interface{}); ok {
		ConversationsString, _ := json.Marshal(Conversations)
		json.Unmarshal(ConversationsString, &o.Conversations)
	}
	
	if TotalHits, ok := AnalyticsconversationqueryresponseMap["totalHits"].(float64); ok {
		TotalHitsInt := int(TotalHits)
		o.TotalHits = &TotalHitsInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Analyticsconversationqueryresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
