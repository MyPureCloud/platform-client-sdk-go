package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Schedulingnoforecastoptionsrequest
type Schedulingnoforecastoptionsrequest struct { 
	// ShiftLength - The shift length option to apply if no forecast is supplied
	ShiftLength *string `json:"shiftLength,omitempty"`


	// ShiftStart - The shift start option to apply if no forecast is supplied
	ShiftStart *string `json:"shiftStart,omitempty"`

}

func (o *Schedulingnoforecastoptionsrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Schedulingnoforecastoptionsrequest
	
	return json.Marshal(&struct { 
		ShiftLength *string `json:"shiftLength,omitempty"`
		
		ShiftStart *string `json:"shiftStart,omitempty"`
		*Alias
	}{ 
		ShiftLength: o.ShiftLength,
		
		ShiftStart: o.ShiftStart,
		Alias:    (*Alias)(o),
	})
}

func (o *Schedulingnoforecastoptionsrequest) UnmarshalJSON(b []byte) error {
	var SchedulingnoforecastoptionsrequestMap map[string]interface{}
	err := json.Unmarshal(b, &SchedulingnoforecastoptionsrequestMap)
	if err != nil {
		return err
	}
	
	if ShiftLength, ok := SchedulingnoforecastoptionsrequestMap["shiftLength"].(string); ok {
		o.ShiftLength = &ShiftLength
	}
	
	if ShiftStart, ok := SchedulingnoforecastoptionsrequestMap["shiftStart"].(string); ok {
		o.ShiftStart = &ShiftStart
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Schedulingnoforecastoptionsrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
