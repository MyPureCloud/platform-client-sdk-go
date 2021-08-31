package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Attendancestatuslisting
type Attendancestatuslisting struct { 
	// Entities
	Entities *[]Attendancestatus `json:"entities,omitempty"`

}

func (o *Attendancestatuslisting) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Attendancestatuslisting
	
	return json.Marshal(&struct { 
		Entities *[]Attendancestatus `json:"entities,omitempty"`
		*Alias
	}{ 
		Entities: o.Entities,
		Alias:    (*Alias)(o),
	})
}

func (o *Attendancestatuslisting) UnmarshalJSON(b []byte) error {
	var AttendancestatuslistingMap map[string]interface{}
	err := json.Unmarshal(b, &AttendancestatuslistingMap)
	if err != nil {
		return err
	}
	
	if Entities, ok := AttendancestatuslistingMap["entities"].([]interface{}); ok {
		EntitiesString, _ := json.Marshal(Entities)
		json.Unmarshal(EntitiesString, &o.Entities)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Attendancestatuslisting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
