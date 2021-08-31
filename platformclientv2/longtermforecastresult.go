package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Longtermforecastresult
type Longtermforecastresult struct { 
	// PlanningGroups - The forecast data broken up by planning group
	PlanningGroups *[]Longtermforecastplanninggroupdata `json:"planningGroups,omitempty"`


	// ReferenceStartDate - The reference start date relative to the business unit time zone in this forecast. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	ReferenceStartDate *time.Time `json:"referenceStartDate,omitempty"`


	// WeekCount - The number of weeks in this forecast
	WeekCount *int `json:"weekCount,omitempty"`

}

func (o *Longtermforecastresult) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Longtermforecastresult
	
	ReferenceStartDate := new(string)
	if o.ReferenceStartDate != nil {
		*ReferenceStartDate = timeutil.Strftime(o.ReferenceStartDate, "%Y-%m-%d")
	} else {
		ReferenceStartDate = nil
	}
	
	return json.Marshal(&struct { 
		PlanningGroups *[]Longtermforecastplanninggroupdata `json:"planningGroups,omitempty"`
		
		ReferenceStartDate *string `json:"referenceStartDate,omitempty"`
		
		WeekCount *int `json:"weekCount,omitempty"`
		*Alias
	}{ 
		PlanningGroups: o.PlanningGroups,
		
		ReferenceStartDate: ReferenceStartDate,
		
		WeekCount: o.WeekCount,
		Alias:    (*Alias)(o),
	})
}

func (o *Longtermforecastresult) UnmarshalJSON(b []byte) error {
	var LongtermforecastresultMap map[string]interface{}
	err := json.Unmarshal(b, &LongtermforecastresultMap)
	if err != nil {
		return err
	}
	
	if PlanningGroups, ok := LongtermforecastresultMap["planningGroups"].([]interface{}); ok {
		PlanningGroupsString, _ := json.Marshal(PlanningGroups)
		json.Unmarshal(PlanningGroupsString, &o.PlanningGroups)
	}
	
	if referenceStartDateString, ok := LongtermforecastresultMap["referenceStartDate"].(string); ok {
		ReferenceStartDate, _ := time.Parse("2006-01-02", referenceStartDateString)
		o.ReferenceStartDate = &ReferenceStartDate
	}
	
	if WeekCount, ok := LongtermforecastresultMap["weekCount"].(float64); ok {
		WeekCountInt := int(WeekCount)
		o.WeekCount = &WeekCountInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Longtermforecastresult) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
