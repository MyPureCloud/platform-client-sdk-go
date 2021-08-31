package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Analyticsuserdetailsqueryresponse
type Analyticsuserdetailsqueryresponse struct { 
	// UserDetails
	UserDetails *[]Analyticsuserdetail `json:"userDetails,omitempty"`


	// Aggregations
	Aggregations *[]Aggregationresult `json:"aggregations,omitempty"`


	// TotalHits
	TotalHits *int `json:"totalHits,omitempty"`

}

func (o *Analyticsuserdetailsqueryresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Analyticsuserdetailsqueryresponse
	
	return json.Marshal(&struct { 
		UserDetails *[]Analyticsuserdetail `json:"userDetails,omitempty"`
		
		Aggregations *[]Aggregationresult `json:"aggregations,omitempty"`
		
		TotalHits *int `json:"totalHits,omitempty"`
		*Alias
	}{ 
		UserDetails: o.UserDetails,
		
		Aggregations: o.Aggregations,
		
		TotalHits: o.TotalHits,
		Alias:    (*Alias)(o),
	})
}

func (o *Analyticsuserdetailsqueryresponse) UnmarshalJSON(b []byte) error {
	var AnalyticsuserdetailsqueryresponseMap map[string]interface{}
	err := json.Unmarshal(b, &AnalyticsuserdetailsqueryresponseMap)
	if err != nil {
		return err
	}
	
	if UserDetails, ok := AnalyticsuserdetailsqueryresponseMap["userDetails"].([]interface{}); ok {
		UserDetailsString, _ := json.Marshal(UserDetails)
		json.Unmarshal(UserDetailsString, &o.UserDetails)
	}
	
	if Aggregations, ok := AnalyticsuserdetailsqueryresponseMap["aggregations"].([]interface{}); ok {
		AggregationsString, _ := json.Marshal(Aggregations)
		json.Unmarshal(AggregationsString, &o.Aggregations)
	}
	
	if TotalHits, ok := AnalyticsuserdetailsqueryresponseMap["totalHits"].(float64); ok {
		TotalHitsInt := int(TotalHits)
		o.TotalHits = &TotalHitsInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Analyticsuserdetailsqueryresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
