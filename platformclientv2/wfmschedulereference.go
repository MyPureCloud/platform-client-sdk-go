package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmschedulereference
type Wfmschedulereference struct { 
	// Id - The ID of the WFM schedule
	Id *string `json:"id,omitempty"`


	// BusinessUnit - A reference to a Workforce Management Business Unit
	BusinessUnit *Wfmbusinessunitreference `json:"businessUnit,omitempty"`


	// WeekDate - The start week date for this schedule. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	WeekDate *time.Time `json:"weekDate,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Wfmschedulereference) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmschedulereference
	
	WeekDate := new(string)
	if o.WeekDate != nil {
		*WeekDate = timeutil.Strftime(o.WeekDate, "%Y-%m-%d")
	} else {
		WeekDate = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		BusinessUnit *Wfmbusinessunitreference `json:"businessUnit,omitempty"`
		
		WeekDate *string `json:"weekDate,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		BusinessUnit: o.BusinessUnit,
		
		WeekDate: WeekDate,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Wfmschedulereference) UnmarshalJSON(b []byte) error {
	var WfmschedulereferenceMap map[string]interface{}
	err := json.Unmarshal(b, &WfmschedulereferenceMap)
	if err != nil {
		return err
	}
	
	if Id, ok := WfmschedulereferenceMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if BusinessUnit, ok := WfmschedulereferenceMap["businessUnit"].(map[string]interface{}); ok {
		BusinessUnitString, _ := json.Marshal(BusinessUnit)
		json.Unmarshal(BusinessUnitString, &o.BusinessUnit)
	}
	
	if weekDateString, ok := WfmschedulereferenceMap["weekDate"].(string); ok {
		WeekDate, _ := time.Parse("2006-01-02", weekDateString)
		o.WeekDate = &WeekDate
	}
	
	if SelfUri, ok := WfmschedulereferenceMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmschedulereference) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
