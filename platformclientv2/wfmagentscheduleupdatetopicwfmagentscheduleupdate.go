package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmagentscheduleupdatetopicwfmagentscheduleupdate
type Wfmagentscheduleupdatetopicwfmagentscheduleupdate struct { 
	// UpdateType
	UpdateType *string `json:"updateType,omitempty"`


	// ShiftStartDates
	ShiftStartDates *[]time.Time `json:"shiftStartDates,omitempty"`

}

func (o *Wfmagentscheduleupdatetopicwfmagentscheduleupdate) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmagentscheduleupdatetopicwfmagentscheduleupdate
	
	return json.Marshal(&struct { 
		UpdateType *string `json:"updateType,omitempty"`
		
		ShiftStartDates *[]time.Time `json:"shiftStartDates,omitempty"`
		*Alias
	}{ 
		UpdateType: o.UpdateType,
		
		ShiftStartDates: o.ShiftStartDates,
		Alias:    (*Alias)(o),
	})
}

func (o *Wfmagentscheduleupdatetopicwfmagentscheduleupdate) UnmarshalJSON(b []byte) error {
	var WfmagentscheduleupdatetopicwfmagentscheduleupdateMap map[string]interface{}
	err := json.Unmarshal(b, &WfmagentscheduleupdatetopicwfmagentscheduleupdateMap)
	if err != nil {
		return err
	}
	
	if UpdateType, ok := WfmagentscheduleupdatetopicwfmagentscheduleupdateMap["updateType"].(string); ok {
		o.UpdateType = &UpdateType
	}
    
	if ShiftStartDates, ok := WfmagentscheduleupdatetopicwfmagentscheduleupdateMap["shiftStartDates"].([]interface{}); ok {
		ShiftStartDatesString, _ := json.Marshal(ShiftStartDates)
		json.Unmarshal(ShiftStartDatesString, &o.ShiftStartDates)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmagentscheduleupdatetopicwfmagentscheduleupdate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
