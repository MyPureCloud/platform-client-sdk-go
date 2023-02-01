package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Edgemetricstopicedgemetrics
type Edgemetricstopicedgemetrics struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Edge
	Edge *Edgemetricstopicurireference `json:"edge,omitempty"`

	// EventTime
	EventTime *time.Time `json:"eventTime,omitempty"`

	// UpTimeMsec
	UpTimeMsec *int `json:"upTimeMsec,omitempty"`

	// Processors
	Processors *[]Edgemetricstopicedgemetricprocessor `json:"processors,omitempty"`

	// Memory
	Memory *[]Edgemetricstopicedgemetricmemory `json:"memory,omitempty"`

	// Disks
	Disks *[]Edgemetricstopicedgemetricdisk `json:"disks,omitempty"`

	// Subsystems
	Subsystems *[]Edgemetricstopicedgemetricsubsystem `json:"subsystems,omitempty"`

	// Networks
	Networks *[]Edgemetricstopicedgemetricnetworks `json:"networks,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Edgemetricstopicedgemetrics) SetField(field string, fieldValue interface{}) {
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

func (o Edgemetricstopicedgemetrics) MarshalJSON() ([]byte, error) {
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
			if fieldValue == nil || reflect.ValueOf(fieldValue).IsNil()  {
				// Do nothing. Just catching this case to avoid trying to custom serialize a nil value
			} else if contains(dateTimeFields, fieldName) {
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
	type Alias Edgemetricstopicedgemetrics
	
	EventTime := new(string)
	if o.EventTime != nil {
		
		*EventTime = timeutil.Strftime(o.EventTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		EventTime = nil
	}
	
	return json.Marshal(&struct { 
		Edge *Edgemetricstopicurireference `json:"edge,omitempty"`
		
		EventTime *string `json:"eventTime,omitempty"`
		
		UpTimeMsec *int `json:"upTimeMsec,omitempty"`
		
		Processors *[]Edgemetricstopicedgemetricprocessor `json:"processors,omitempty"`
		
		Memory *[]Edgemetricstopicedgemetricmemory `json:"memory,omitempty"`
		
		Disks *[]Edgemetricstopicedgemetricdisk `json:"disks,omitempty"`
		
		Subsystems *[]Edgemetricstopicedgemetricsubsystem `json:"subsystems,omitempty"`
		
		Networks *[]Edgemetricstopicedgemetricnetworks `json:"networks,omitempty"`
		Alias
	}{ 
		Edge: o.Edge,
		
		EventTime: EventTime,
		
		UpTimeMsec: o.UpTimeMsec,
		
		Processors: o.Processors,
		
		Memory: o.Memory,
		
		Disks: o.Disks,
		
		Subsystems: o.Subsystems,
		
		Networks: o.Networks,
		Alias:    (Alias)(o),
	})
}

func (o *Edgemetricstopicedgemetrics) UnmarshalJSON(b []byte) error {
	var EdgemetricstopicedgemetricsMap map[string]interface{}
	err := json.Unmarshal(b, &EdgemetricstopicedgemetricsMap)
	if err != nil {
		return err
	}
	
	if Edge, ok := EdgemetricstopicedgemetricsMap["edge"].(map[string]interface{}); ok {
		EdgeString, _ := json.Marshal(Edge)
		json.Unmarshal(EdgeString, &o.Edge)
	}
	
	if eventTimeString, ok := EdgemetricstopicedgemetricsMap["eventTime"].(string); ok {
		EventTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", eventTimeString)
		o.EventTime = &EventTime
	}
	
	if UpTimeMsec, ok := EdgemetricstopicedgemetricsMap["upTimeMsec"].(float64); ok {
		UpTimeMsecInt := int(UpTimeMsec)
		o.UpTimeMsec = &UpTimeMsecInt
	}
	
	if Processors, ok := EdgemetricstopicedgemetricsMap["processors"].([]interface{}); ok {
		ProcessorsString, _ := json.Marshal(Processors)
		json.Unmarshal(ProcessorsString, &o.Processors)
	}
	
	if Memory, ok := EdgemetricstopicedgemetricsMap["memory"].([]interface{}); ok {
		MemoryString, _ := json.Marshal(Memory)
		json.Unmarshal(MemoryString, &o.Memory)
	}
	
	if Disks, ok := EdgemetricstopicedgemetricsMap["disks"].([]interface{}); ok {
		DisksString, _ := json.Marshal(Disks)
		json.Unmarshal(DisksString, &o.Disks)
	}
	
	if Subsystems, ok := EdgemetricstopicedgemetricsMap["subsystems"].([]interface{}); ok {
		SubsystemsString, _ := json.Marshal(Subsystems)
		json.Unmarshal(SubsystemsString, &o.Subsystems)
	}
	
	if Networks, ok := EdgemetricstopicedgemetricsMap["networks"].([]interface{}); ok {
		NetworksString, _ := json.Marshal(Networks)
		json.Unmarshal(NetworksString, &o.Networks)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Edgemetricstopicedgemetrics) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
