package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Shiftstartvariance
type Shiftstartvariance struct { 
	// ApplicableDays - Days for which shift start variance is configured
	ApplicableDays *[]string `json:"applicableDays,omitempty"`


	// MaxShiftStartVarianceMinutes - Maximum variance in minutes across shift starts
	MaxShiftStartVarianceMinutes *int `json:"maxShiftStartVarianceMinutes,omitempty"`

}

func (o *Shiftstartvariance) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Shiftstartvariance
	
	return json.Marshal(&struct { 
		ApplicableDays *[]string `json:"applicableDays,omitempty"`
		
		MaxShiftStartVarianceMinutes *int `json:"maxShiftStartVarianceMinutes,omitempty"`
		*Alias
	}{ 
		ApplicableDays: o.ApplicableDays,
		
		MaxShiftStartVarianceMinutes: o.MaxShiftStartVarianceMinutes,
		Alias:    (*Alias)(o),
	})
}

func (o *Shiftstartvariance) UnmarshalJSON(b []byte) error {
	var ShiftstartvarianceMap map[string]interface{}
	err := json.Unmarshal(b, &ShiftstartvarianceMap)
	if err != nil {
		return err
	}
	
	if ApplicableDays, ok := ShiftstartvarianceMap["applicableDays"].([]interface{}); ok {
		ApplicableDaysString, _ := json.Marshal(ApplicableDays)
		json.Unmarshal(ApplicableDaysString, &o.ApplicableDays)
	}
	
	if MaxShiftStartVarianceMinutes, ok := ShiftstartvarianceMap["maxShiftStartVarianceMinutes"].(float64); ok {
		MaxShiftStartVarianceMinutesInt := int(MaxShiftStartVarianceMinutes)
		o.MaxShiftStartVarianceMinutes = &MaxShiftStartVarianceMinutesInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Shiftstartvariance) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
