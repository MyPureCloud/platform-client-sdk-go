package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Buagentschedulehistorydroppedchange
type Buagentschedulehistorydroppedchange struct { 
	// Metadata - The metadata of the change, including who and when the change was made
	Metadata *Buagentschedulehistorychangemetadata `json:"metadata,omitempty"`


	// ShiftIds - The IDs of deleted shifts
	ShiftIds *[]string `json:"shiftIds,omitempty"`


	// FullDayTimeOffMarkerDates - The dates of any deleted full day time off markers
	FullDayTimeOffMarkerDates *[]time.Time `json:"fullDayTimeOffMarkerDates,omitempty"`


	// Deletes - The deleted shifts, full day time off markers, or the entire agent schedule
	Deletes *Buagentschedulehistorydeletedchange `json:"deletes,omitempty"`

}

func (o *Buagentschedulehistorydroppedchange) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Buagentschedulehistorydroppedchange
	
	return json.Marshal(&struct { 
		Metadata *Buagentschedulehistorychangemetadata `json:"metadata,omitempty"`
		
		ShiftIds *[]string `json:"shiftIds,omitempty"`
		
		FullDayTimeOffMarkerDates *[]time.Time `json:"fullDayTimeOffMarkerDates,omitempty"`
		
		Deletes *Buagentschedulehistorydeletedchange `json:"deletes,omitempty"`
		*Alias
	}{ 
		Metadata: o.Metadata,
		
		ShiftIds: o.ShiftIds,
		
		FullDayTimeOffMarkerDates: o.FullDayTimeOffMarkerDates,
		
		Deletes: o.Deletes,
		Alias:    (*Alias)(o),
	})
}

func (o *Buagentschedulehistorydroppedchange) UnmarshalJSON(b []byte) error {
	var BuagentschedulehistorydroppedchangeMap map[string]interface{}
	err := json.Unmarshal(b, &BuagentschedulehistorydroppedchangeMap)
	if err != nil {
		return err
	}
	
	if Metadata, ok := BuagentschedulehistorydroppedchangeMap["metadata"].(map[string]interface{}); ok {
		MetadataString, _ := json.Marshal(Metadata)
		json.Unmarshal(MetadataString, &o.Metadata)
	}
	
	if ShiftIds, ok := BuagentschedulehistorydroppedchangeMap["shiftIds"].([]interface{}); ok {
		ShiftIdsString, _ := json.Marshal(ShiftIds)
		json.Unmarshal(ShiftIdsString, &o.ShiftIds)
	}
	
	if FullDayTimeOffMarkerDates, ok := BuagentschedulehistorydroppedchangeMap["fullDayTimeOffMarkerDates"].([]interface{}); ok {
		FullDayTimeOffMarkerDatesString, _ := json.Marshal(FullDayTimeOffMarkerDates)
		json.Unmarshal(FullDayTimeOffMarkerDatesString, &o.FullDayTimeOffMarkerDates)
	}
	
	if Deletes, ok := BuagentschedulehistorydroppedchangeMap["deletes"].(map[string]interface{}); ok {
		DeletesString, _ := json.Marshal(Deletes)
		json.Unmarshal(DeletesString, &o.Deletes)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Buagentschedulehistorydroppedchange) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
