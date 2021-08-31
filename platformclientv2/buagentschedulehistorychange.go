package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Buagentschedulehistorychange
type Buagentschedulehistorychange struct { 
	// Metadata - The metadata of the change, including who and when the change was made
	Metadata *Buagentschedulehistorychangemetadata `json:"metadata,omitempty"`


	// Shifts - The list of changed shifts
	Shifts *[]Buagentscheduleshift `json:"shifts,omitempty"`


	// FullDayTimeOffMarkers - The list of changed full day time off markers
	FullDayTimeOffMarkers *[]Bufulldaytimeoffmarker `json:"fullDayTimeOffMarkers,omitempty"`


	// Deletes - The deleted shifts, full day time off markers, or the entire agent schedule
	Deletes *Buagentschedulehistorydeletedchange `json:"deletes,omitempty"`

}

func (o *Buagentschedulehistorychange) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Buagentschedulehistorychange
	
	return json.Marshal(&struct { 
		Metadata *Buagentschedulehistorychangemetadata `json:"metadata,omitempty"`
		
		Shifts *[]Buagentscheduleshift `json:"shifts,omitempty"`
		
		FullDayTimeOffMarkers *[]Bufulldaytimeoffmarker `json:"fullDayTimeOffMarkers,omitempty"`
		
		Deletes *Buagentschedulehistorydeletedchange `json:"deletes,omitempty"`
		*Alias
	}{ 
		Metadata: o.Metadata,
		
		Shifts: o.Shifts,
		
		FullDayTimeOffMarkers: o.FullDayTimeOffMarkers,
		
		Deletes: o.Deletes,
		Alias:    (*Alias)(o),
	})
}

func (o *Buagentschedulehistorychange) UnmarshalJSON(b []byte) error {
	var BuagentschedulehistorychangeMap map[string]interface{}
	err := json.Unmarshal(b, &BuagentschedulehistorychangeMap)
	if err != nil {
		return err
	}
	
	if Metadata, ok := BuagentschedulehistorychangeMap["metadata"].(map[string]interface{}); ok {
		MetadataString, _ := json.Marshal(Metadata)
		json.Unmarshal(MetadataString, &o.Metadata)
	}
	
	if Shifts, ok := BuagentschedulehistorychangeMap["shifts"].([]interface{}); ok {
		ShiftsString, _ := json.Marshal(Shifts)
		json.Unmarshal(ShiftsString, &o.Shifts)
	}
	
	if FullDayTimeOffMarkers, ok := BuagentschedulehistorychangeMap["fullDayTimeOffMarkers"].([]interface{}); ok {
		FullDayTimeOffMarkersString, _ := json.Marshal(FullDayTimeOffMarkers)
		json.Unmarshal(FullDayTimeOffMarkersString, &o.FullDayTimeOffMarkers)
	}
	
	if Deletes, ok := BuagentschedulehistorychangeMap["deletes"].(map[string]interface{}); ok {
		DeletesString, _ := json.Marshal(Deletes)
		json.Unmarshal(DeletesString, &o.Deletes)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Buagentschedulehistorychange) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
