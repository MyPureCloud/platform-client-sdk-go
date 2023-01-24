package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmintradaydataupdatetopicintradayforecastdata
type Wfmintradaydataupdatetopicintradayforecastdata struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Offered
	Offered *float32 `json:"offered,omitempty"`

	// AverageTalkTimeSeconds
	AverageTalkTimeSeconds *float32 `json:"averageTalkTimeSeconds,omitempty"`

	// AverageAfterCallWorkSeconds
	AverageAfterCallWorkSeconds *float32 `json:"averageAfterCallWorkSeconds,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Wfmintradaydataupdatetopicintradayforecastdata) SetField(field string, fieldValue interface{}) {
	// Get Value object for field
	target := reflect.ValueOf(o)
	targetField := reflect.Indirect(target).FieldByName(field)

	// Set value
	if fieldValue != nil {
		targetField.Set(reflect.ValueOf(fieldValue))
	} else {
		// Must create a new Value (creates **type) then get its element (*type), which will be nil pointer of the appropriate type
		x := reflect.Indirect(reflect.New(targetField.Type()))
		targetField.Set(x)
	}

	// Add field to set field names list
	if o.SetFieldNames == nil {
		o.SetFieldNames = make(map[string]bool)
	}
	o.SetFieldNames[field] = true
}

func (o Wfmintradaydataupdatetopicintradayforecastdata) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{  }
		localDateTimeFields := []string{  }
		dateFields := []string{  }

		// Construct object
		newObj := make(map[string]interface{})
		for fieldName := range o.SetFieldNames {
			// Get initial field value
			fieldValue := val.FieldByName(fieldName).Interface()

			// Apply value formatting overrides
			if contains(dateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%fZ")
			} else if contains(localDateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%f")
			} else if contains(dateFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%d")
			}

			// Assign value to field using JSON tag name
			newObj[getFieldName(reflect.TypeOf(&o), fieldName)] = fieldValue
		}

		// Marshal and return dynamically constructed interface
		return json.Marshal(newObj)
	}

	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmintradaydataupdatetopicintradayforecastdata
	
	return json.Marshal(&struct { 
		Offered *float32 `json:"offered,omitempty"`
		
		AverageTalkTimeSeconds *float32 `json:"averageTalkTimeSeconds,omitempty"`
		
		AverageAfterCallWorkSeconds *float32 `json:"averageAfterCallWorkSeconds,omitempty"`
		Alias
	}{ 
		Offered: o.Offered,
		
		AverageTalkTimeSeconds: o.AverageTalkTimeSeconds,
		
		AverageAfterCallWorkSeconds: o.AverageAfterCallWorkSeconds,
		Alias:    (Alias)(o),
	})
}

func (o *Wfmintradaydataupdatetopicintradayforecastdata) UnmarshalJSON(b []byte) error {
	var WfmintradaydataupdatetopicintradayforecastdataMap map[string]interface{}
	err := json.Unmarshal(b, &WfmintradaydataupdatetopicintradayforecastdataMap)
	if err != nil {
		return err
	}
	
	if Offered, ok := WfmintradaydataupdatetopicintradayforecastdataMap["offered"].(float64); ok {
		OfferedFloat32 := float32(Offered)
		o.Offered = &OfferedFloat32
	}
    
	if AverageTalkTimeSeconds, ok := WfmintradaydataupdatetopicintradayforecastdataMap["averageTalkTimeSeconds"].(float64); ok {
		AverageTalkTimeSecondsFloat32 := float32(AverageTalkTimeSeconds)
		o.AverageTalkTimeSeconds = &AverageTalkTimeSecondsFloat32
	}
    
	if AverageAfterCallWorkSeconds, ok := WfmintradaydataupdatetopicintradayforecastdataMap["averageAfterCallWorkSeconds"].(float64); ok {
		AverageAfterCallWorkSecondsFloat32 := float32(AverageAfterCallWorkSeconds)
		o.AverageAfterCallWorkSeconds = &AverageAfterCallWorkSecondsFloat32
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmintradaydataupdatetopicintradayforecastdata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
