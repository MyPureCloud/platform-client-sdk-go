package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Shiftstartvariance - Variance in minutes among start times of shifts in work plan
type Shiftstartvariance struct { 
	// ApplicableDays - Days for which shift start variance is configured
	ApplicableDays *[]string `json:"applicableDays,omitempty"`


	// MaxShiftStartVarianceMinutes - Maximum variance in minutes across shift starts
	MaxShiftStartVarianceMinutes *int `json:"maxShiftStartVarianceMinutes,omitempty"`

}

func (u *Shiftstartvariance) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Shiftstartvariance

	

	return json.Marshal(&struct { 
		ApplicableDays *[]string `json:"applicableDays,omitempty"`
		
		MaxShiftStartVarianceMinutes *int `json:"maxShiftStartVarianceMinutes,omitempty"`
		*Alias
	}{ 
		ApplicableDays: u.ApplicableDays,
		
		MaxShiftStartVarianceMinutes: u.MaxShiftStartVarianceMinutes,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Shiftstartvariance) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
