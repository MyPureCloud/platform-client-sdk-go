package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Buagentschedulehistorydeletedchange
type Buagentschedulehistorydeletedchange struct { 
	// ShiftIds - The IDs of deleted shifts
	ShiftIds *[]string `json:"shiftIds,omitempty"`


	// FullDayTimeOffMarkerDates - The dates of any deleted full day time off markers
	FullDayTimeOffMarkerDates *[]time.Time `json:"fullDayTimeOffMarkerDates,omitempty"`


	// AgentSchedule - Whether the entire agent schedule was deleted
	AgentSchedule *bool `json:"agentSchedule,omitempty"`

}

func (o *Buagentschedulehistorydeletedchange) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Buagentschedulehistorydeletedchange
	
	return json.Marshal(&struct { 
		ShiftIds *[]string `json:"shiftIds,omitempty"`
		
		FullDayTimeOffMarkerDates *[]time.Time `json:"fullDayTimeOffMarkerDates,omitempty"`
		
		AgentSchedule *bool `json:"agentSchedule,omitempty"`
		*Alias
	}{ 
		ShiftIds: o.ShiftIds,
		
		FullDayTimeOffMarkerDates: o.FullDayTimeOffMarkerDates,
		
		AgentSchedule: o.AgentSchedule,
		Alias:    (*Alias)(o),
	})
}

func (o *Buagentschedulehistorydeletedchange) UnmarshalJSON(b []byte) error {
	var BuagentschedulehistorydeletedchangeMap map[string]interface{}
	err := json.Unmarshal(b, &BuagentschedulehistorydeletedchangeMap)
	if err != nil {
		return err
	}
	
	if ShiftIds, ok := BuagentschedulehistorydeletedchangeMap["shiftIds"].([]interface{}); ok {
		ShiftIdsString, _ := json.Marshal(ShiftIds)
		json.Unmarshal(ShiftIdsString, &o.ShiftIds)
	}
	
	if FullDayTimeOffMarkerDates, ok := BuagentschedulehistorydeletedchangeMap["fullDayTimeOffMarkerDates"].([]interface{}); ok {
		FullDayTimeOffMarkerDatesString, _ := json.Marshal(FullDayTimeOffMarkerDates)
		json.Unmarshal(FullDayTimeOffMarkerDatesString, &o.FullDayTimeOffMarkerDates)
	}
	
	if AgentSchedule, ok := BuagentschedulehistorydeletedchangeMap["agentSchedule"].(bool); ok {
		o.AgentSchedule = &AgentSchedule
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Buagentschedulehistorydeletedchange) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
