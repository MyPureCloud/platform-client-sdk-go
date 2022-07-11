package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Userschedule
type Userschedule struct { 
	// Shifts - The shifts that belong to this schedule
	Shifts *[]Userscheduleshift `json:"shifts,omitempty"`


	// FullDayTimeOffMarkers - Markers to indicate a full day time off request, relative to the management unit time zone
	FullDayTimeOffMarkers *[]Userschedulefulldaytimeoffmarker `json:"fullDayTimeOffMarkers,omitempty"`


	// Delete - If marked true for updating an existing user schedule, it will be deleted
	Delete *bool `json:"delete,omitempty"`


	// Metadata - Version metadata for this schedule
	Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`


	// WorkPlanId - ID of the work plan associated with the user during schedule creation
	WorkPlanId *string `json:"workPlanId,omitempty"`

}

func (o *Userschedule) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Userschedule
	
	return json.Marshal(&struct { 
		Shifts *[]Userscheduleshift `json:"shifts,omitempty"`
		
		FullDayTimeOffMarkers *[]Userschedulefulldaytimeoffmarker `json:"fullDayTimeOffMarkers,omitempty"`
		
		Delete *bool `json:"delete,omitempty"`
		
		Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`
		
		WorkPlanId *string `json:"workPlanId,omitempty"`
		*Alias
	}{ 
		Shifts: o.Shifts,
		
		FullDayTimeOffMarkers: o.FullDayTimeOffMarkers,
		
		Delete: o.Delete,
		
		Metadata: o.Metadata,
		
		WorkPlanId: o.WorkPlanId,
		Alias:    (*Alias)(o),
	})
}

func (o *Userschedule) UnmarshalJSON(b []byte) error {
	var UserscheduleMap map[string]interface{}
	err := json.Unmarshal(b, &UserscheduleMap)
	if err != nil {
		return err
	}
	
	if Shifts, ok := UserscheduleMap["shifts"].([]interface{}); ok {
		ShiftsString, _ := json.Marshal(Shifts)
		json.Unmarshal(ShiftsString, &o.Shifts)
	}
	
	if FullDayTimeOffMarkers, ok := UserscheduleMap["fullDayTimeOffMarkers"].([]interface{}); ok {
		FullDayTimeOffMarkersString, _ := json.Marshal(FullDayTimeOffMarkers)
		json.Unmarshal(FullDayTimeOffMarkersString, &o.FullDayTimeOffMarkers)
	}
	
	if Delete, ok := UserscheduleMap["delete"].(bool); ok {
		o.Delete = &Delete
	}
    
	if Metadata, ok := UserscheduleMap["metadata"].(map[string]interface{}); ok {
		MetadataString, _ := json.Marshal(Metadata)
		json.Unmarshal(MetadataString, &o.Metadata)
	}
	
	if WorkPlanId, ok := UserscheduleMap["workPlanId"].(string); ok {
		o.WorkPlanId = &WorkPlanId
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Userschedule) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
