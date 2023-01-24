package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Trunkmetricstopictrunkmetrics
type Trunkmetricstopictrunkmetrics struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Calls
	Calls *Trunkmetricstopictrunkmetricscalls `json:"calls,omitempty"`

	// EventTime
	EventTime *time.Time `json:"eventTime,omitempty"`

	// Qos
	Qos *Trunkmetricstopictrunkmetricsqos `json:"qos,omitempty"`

	// Trunk
	Trunk *Trunkmetricstopicurireference `json:"trunk,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Trunkmetricstopictrunkmetrics) SetField(field string, fieldValue interface{}) {
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

func (o Trunkmetricstopictrunkmetrics) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "EventTime", }
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
	type Alias Trunkmetricstopictrunkmetrics
	
	EventTime := new(string)
	if o.EventTime != nil {
		
		*EventTime = timeutil.Strftime(o.EventTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		EventTime = nil
	}
	
	return json.Marshal(&struct { 
		Calls *Trunkmetricstopictrunkmetricscalls `json:"calls,omitempty"`
		
		EventTime *string `json:"eventTime,omitempty"`
		
		Qos *Trunkmetricstopictrunkmetricsqos `json:"qos,omitempty"`
		
		Trunk *Trunkmetricstopicurireference `json:"trunk,omitempty"`
		Alias
	}{ 
		Calls: o.Calls,
		
		EventTime: EventTime,
		
		Qos: o.Qos,
		
		Trunk: o.Trunk,
		Alias:    (Alias)(o),
	})
}

func (o *Trunkmetricstopictrunkmetrics) UnmarshalJSON(b []byte) error {
	var TrunkmetricstopictrunkmetricsMap map[string]interface{}
	err := json.Unmarshal(b, &TrunkmetricstopictrunkmetricsMap)
	if err != nil {
		return err
	}
	
	if Calls, ok := TrunkmetricstopictrunkmetricsMap["calls"].(map[string]interface{}); ok {
		CallsString, _ := json.Marshal(Calls)
		json.Unmarshal(CallsString, &o.Calls)
	}
	
	if eventTimeString, ok := TrunkmetricstopictrunkmetricsMap["eventTime"].(string); ok {
		EventTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", eventTimeString)
		o.EventTime = &EventTime
	}
	
	if Qos, ok := TrunkmetricstopictrunkmetricsMap["qos"].(map[string]interface{}); ok {
		QosString, _ := json.Marshal(Qos)
		json.Unmarshal(QosString, &o.Qos)
	}
	
	if Trunk, ok := TrunkmetricstopictrunkmetricsMap["trunk"].(map[string]interface{}); ok {
		TrunkString, _ := json.Marshal(Trunk)
		json.Unmarshal(TrunkString, &o.Trunk)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Trunkmetricstopictrunkmetrics) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
