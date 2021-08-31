package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Generatebuforecastrequest
type Generatebuforecastrequest struct { 
	// Description - The description for the forecast
	Description *string `json:"description,omitempty"`


	// WeekCount - The number of weeks this forecast covers
	WeekCount *int `json:"weekCount,omitempty"`


	// CanUseForScheduling - Whether this forecast can be used for scheduling
	CanUseForScheduling *bool `json:"canUseForScheduling,omitempty"`

}

func (o *Generatebuforecastrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Generatebuforecastrequest
	
	return json.Marshal(&struct { 
		Description *string `json:"description,omitempty"`
		
		WeekCount *int `json:"weekCount,omitempty"`
		
		CanUseForScheduling *bool `json:"canUseForScheduling,omitempty"`
		*Alias
	}{ 
		Description: o.Description,
		
		WeekCount: o.WeekCount,
		
		CanUseForScheduling: o.CanUseForScheduling,
		Alias:    (*Alias)(o),
	})
}

func (o *Generatebuforecastrequest) UnmarshalJSON(b []byte) error {
	var GeneratebuforecastrequestMap map[string]interface{}
	err := json.Unmarshal(b, &GeneratebuforecastrequestMap)
	if err != nil {
		return err
	}
	
	if Description, ok := GeneratebuforecastrequestMap["description"].(string); ok {
		o.Description = &Description
	}
	
	if WeekCount, ok := GeneratebuforecastrequestMap["weekCount"].(float64); ok {
		WeekCountInt := int(WeekCount)
		o.WeekCount = &WeekCountInt
	}
	
	if CanUseForScheduling, ok := GeneratebuforecastrequestMap["canUseForScheduling"].(bool); ok {
		o.CanUseForScheduling = &CanUseForScheduling
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Generatebuforecastrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
