package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Planningperiodsettings
type Planningperiodsettings struct { 
	// WeekCount - Planning period length in weeks
	WeekCount *int `json:"weekCount,omitempty"`


	// StartDate - Start date of the planning period in yyyy-MM-dd format. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	StartDate *time.Time `json:"startDate,omitempty"`

}

func (o *Planningperiodsettings) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Planningperiodsettings
	
	StartDate := new(string)
	if o.StartDate != nil {
		*StartDate = timeutil.Strftime(o.StartDate, "%Y-%m-%d")
	} else {
		StartDate = nil
	}
	
	return json.Marshal(&struct { 
		WeekCount *int `json:"weekCount,omitempty"`
		
		StartDate *string `json:"startDate,omitempty"`
		*Alias
	}{ 
		WeekCount: o.WeekCount,
		
		StartDate: StartDate,
		Alias:    (*Alias)(o),
	})
}

func (o *Planningperiodsettings) UnmarshalJSON(b []byte) error {
	var PlanningperiodsettingsMap map[string]interface{}
	err := json.Unmarshal(b, &PlanningperiodsettingsMap)
	if err != nil {
		return err
	}
	
	if WeekCount, ok := PlanningperiodsettingsMap["weekCount"].(float64); ok {
		WeekCountInt := int(WeekCount)
		o.WeekCount = &WeekCountInt
	}
	
	if startDateString, ok := PlanningperiodsettingsMap["startDate"].(string); ok {
		StartDate, _ := time.Parse("2006-01-02", startDateString)
		o.StartDate = &StartDate
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Planningperiodsettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
