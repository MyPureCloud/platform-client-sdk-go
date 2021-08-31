package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Reportinginterval
type Reportinginterval struct { 
	// IntervalType - The granularity of the reporting interval period
	IntervalType *string `json:"intervalType,omitempty"`


	// IntervalValue - The value of the reporting interval period for a given interval type
	IntervalValue *int `json:"intervalValue,omitempty"`

}

func (o *Reportinginterval) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Reportinginterval
	
	return json.Marshal(&struct { 
		IntervalType *string `json:"intervalType,omitempty"`
		
		IntervalValue *int `json:"intervalValue,omitempty"`
		*Alias
	}{ 
		IntervalType: o.IntervalType,
		
		IntervalValue: o.IntervalValue,
		Alias:    (*Alias)(o),
	})
}

func (o *Reportinginterval) UnmarshalJSON(b []byte) error {
	var ReportingintervalMap map[string]interface{}
	err := json.Unmarshal(b, &ReportingintervalMap)
	if err != nil {
		return err
	}
	
	if IntervalType, ok := ReportingintervalMap["intervalType"].(string); ok {
		o.IntervalType = &IntervalType
	}
	
	if IntervalValue, ok := ReportingintervalMap["intervalValue"].(float64); ok {
		IntervalValueInt := int(IntervalValue)
		o.IntervalValue = &IntervalValueInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Reportinginterval) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
