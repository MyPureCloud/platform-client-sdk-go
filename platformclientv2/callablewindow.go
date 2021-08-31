package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Callablewindow
type Callablewindow struct { 
	// Mapped - The time interval to place outbound calls, for contacts that can be mapped to a time zone.
	Mapped *Atzmtimeslot `json:"mapped,omitempty"`


	// Unmapped - The time interval and time zone to place outbound calls, for contacts that cannot be mapped to a time zone.
	Unmapped *Atzmtimeslotwithtimezone `json:"unmapped,omitempty"`

}

func (o *Callablewindow) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Callablewindow
	
	return json.Marshal(&struct { 
		Mapped *Atzmtimeslot `json:"mapped,omitempty"`
		
		Unmapped *Atzmtimeslotwithtimezone `json:"unmapped,omitempty"`
		*Alias
	}{ 
		Mapped: o.Mapped,
		
		Unmapped: o.Unmapped,
		Alias:    (*Alias)(o),
	})
}

func (o *Callablewindow) UnmarshalJSON(b []byte) error {
	var CallablewindowMap map[string]interface{}
	err := json.Unmarshal(b, &CallablewindowMap)
	if err != nil {
		return err
	}
	
	if Mapped, ok := CallablewindowMap["mapped"].(map[string]interface{}); ok {
		MappedString, _ := json.Marshal(Mapped)
		json.Unmarshal(MappedString, &o.Mapped)
	}
	
	if Unmapped, ok := CallablewindowMap["unmapped"].(map[string]interface{}); ok {
		UnmappedString, _ := json.Marshal(Unmapped)
		json.Unmarshal(UnmappedString, &o.Unmapped)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Callablewindow) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
