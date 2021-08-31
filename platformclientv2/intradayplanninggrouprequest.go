package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Intradayplanninggrouprequest
type Intradayplanninggrouprequest struct { 
	// BusinessUnitDate - Requested date in yyyy-MM-dd format. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	BusinessUnitDate *time.Time `json:"businessUnitDate,omitempty"`


	// Categories - The metric categories
	Categories *[]string `json:"categories,omitempty"`


	// PlanningGroupIds - The IDs of the planning groups for which to fetch data.  Omitting or passing an empty list will return all available planning groups
	PlanningGroupIds *[]string `json:"planningGroupIds,omitempty"`


	// IntervalLengthMinutes - The period/interval in minutes for which to aggregate the data. Required, defaults to 15
	IntervalLengthMinutes *int `json:"intervalLengthMinutes,omitempty"`

}

func (o *Intradayplanninggrouprequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Intradayplanninggrouprequest
	
	BusinessUnitDate := new(string)
	if o.BusinessUnitDate != nil {
		*BusinessUnitDate = timeutil.Strftime(o.BusinessUnitDate, "%Y-%m-%d")
	} else {
		BusinessUnitDate = nil
	}
	
	return json.Marshal(&struct { 
		BusinessUnitDate *string `json:"businessUnitDate,omitempty"`
		
		Categories *[]string `json:"categories,omitempty"`
		
		PlanningGroupIds *[]string `json:"planningGroupIds,omitempty"`
		
		IntervalLengthMinutes *int `json:"intervalLengthMinutes,omitempty"`
		*Alias
	}{ 
		BusinessUnitDate: BusinessUnitDate,
		
		Categories: o.Categories,
		
		PlanningGroupIds: o.PlanningGroupIds,
		
		IntervalLengthMinutes: o.IntervalLengthMinutes,
		Alias:    (*Alias)(o),
	})
}

func (o *Intradayplanninggrouprequest) UnmarshalJSON(b []byte) error {
	var IntradayplanninggrouprequestMap map[string]interface{}
	err := json.Unmarshal(b, &IntradayplanninggrouprequestMap)
	if err != nil {
		return err
	}
	
	if businessUnitDateString, ok := IntradayplanninggrouprequestMap["businessUnitDate"].(string); ok {
		BusinessUnitDate, _ := time.Parse("2006-01-02", businessUnitDateString)
		o.BusinessUnitDate = &BusinessUnitDate
	}
	
	if Categories, ok := IntradayplanninggrouprequestMap["categories"].([]interface{}); ok {
		CategoriesString, _ := json.Marshal(Categories)
		json.Unmarshal(CategoriesString, &o.Categories)
	}
	
	if PlanningGroupIds, ok := IntradayplanninggrouprequestMap["planningGroupIds"].([]interface{}); ok {
		PlanningGroupIdsString, _ := json.Marshal(PlanningGroupIds)
		json.Unmarshal(PlanningGroupIdsString, &o.PlanningGroupIds)
	}
	
	if IntervalLengthMinutes, ok := IntradayplanninggrouprequestMap["intervalLengthMinutes"].(float64); ok {
		IntervalLengthMinutesInt := int(IntervalLengthMinutes)
		o.IntervalLengthMinutes = &IntervalLengthMinutesInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Intradayplanninggrouprequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
