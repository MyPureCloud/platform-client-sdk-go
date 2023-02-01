package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Historicalshrinkageaggregateresponse
type Historicalshrinkageaggregateresponse struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// ScheduledShrinkageSeconds - Aggregated shrinkage value in seconds for scheduled activities
	ScheduledShrinkageSeconds *int `json:"scheduledShrinkageSeconds,omitempty"`

	// ScheduledShrinkagePercent - Aggregated shrinkage value in percent from 0.0 to 100.0 for scheduled activities
	ScheduledShrinkagePercent *float64 `json:"scheduledShrinkagePercent,omitempty"`

	// ActualShrinkageSeconds - Aggregated actual value in seconds for scheduled activities
	ActualShrinkageSeconds *int `json:"actualShrinkageSeconds,omitempty"`

	// ActualShrinkagePercent - Aggregated actual value in percent from 0.0 to 100.0 for scheduled activities
	ActualShrinkagePercent *float64 `json:"actualShrinkagePercent,omitempty"`

	// PaidShrinkageSeconds - Aggregated shrinkage value in seconds for paid activities
	PaidShrinkageSeconds *int `json:"paidShrinkageSeconds,omitempty"`

	// UnpaidShrinkageSeconds - Aggregated shrinkage value in seconds for unpaid activities
	UnpaidShrinkageSeconds *int `json:"unpaidShrinkageSeconds,omitempty"`

	// PlannedShrinkageSeconds - Aggregated shrinkage value in seconds for planned activities
	PlannedShrinkageSeconds *int `json:"plannedShrinkageSeconds,omitempty"`

	// UnplannedShrinkageSeconds - Aggregated shrinkage value in seconds for unplanned activities
	UnplannedShrinkageSeconds *int `json:"unplannedShrinkageSeconds,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Historicalshrinkageaggregateresponse) SetField(field string, fieldValue interface{}) {
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

func (o Historicalshrinkageaggregateresponse) MarshalJSON() ([]byte, error) {
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
	type Alias Historicalshrinkageaggregateresponse
	
	return json.Marshal(&struct { 
		ScheduledShrinkageSeconds *int `json:"scheduledShrinkageSeconds,omitempty"`
		
		ScheduledShrinkagePercent *float64 `json:"scheduledShrinkagePercent,omitempty"`
		
		ActualShrinkageSeconds *int `json:"actualShrinkageSeconds,omitempty"`
		
		ActualShrinkagePercent *float64 `json:"actualShrinkagePercent,omitempty"`
		
		PaidShrinkageSeconds *int `json:"paidShrinkageSeconds,omitempty"`
		
		UnpaidShrinkageSeconds *int `json:"unpaidShrinkageSeconds,omitempty"`
		
		PlannedShrinkageSeconds *int `json:"plannedShrinkageSeconds,omitempty"`
		
		UnplannedShrinkageSeconds *int `json:"unplannedShrinkageSeconds,omitempty"`
		Alias
	}{ 
		ScheduledShrinkageSeconds: o.ScheduledShrinkageSeconds,
		
		ScheduledShrinkagePercent: o.ScheduledShrinkagePercent,
		
		ActualShrinkageSeconds: o.ActualShrinkageSeconds,
		
		ActualShrinkagePercent: o.ActualShrinkagePercent,
		
		PaidShrinkageSeconds: o.PaidShrinkageSeconds,
		
		UnpaidShrinkageSeconds: o.UnpaidShrinkageSeconds,
		
		PlannedShrinkageSeconds: o.PlannedShrinkageSeconds,
		
		UnplannedShrinkageSeconds: o.UnplannedShrinkageSeconds,
		Alias:    (Alias)(o),
	})
}

func (o *Historicalshrinkageaggregateresponse) UnmarshalJSON(b []byte) error {
	var HistoricalshrinkageaggregateresponseMap map[string]interface{}
	err := json.Unmarshal(b, &HistoricalshrinkageaggregateresponseMap)
	if err != nil {
		return err
	}
	
	if ScheduledShrinkageSeconds, ok := HistoricalshrinkageaggregateresponseMap["scheduledShrinkageSeconds"].(float64); ok {
		ScheduledShrinkageSecondsInt := int(ScheduledShrinkageSeconds)
		o.ScheduledShrinkageSeconds = &ScheduledShrinkageSecondsInt
	}
	
	if ScheduledShrinkagePercent, ok := HistoricalshrinkageaggregateresponseMap["scheduledShrinkagePercent"].(float64); ok {
		o.ScheduledShrinkagePercent = &ScheduledShrinkagePercent
	}
    
	if ActualShrinkageSeconds, ok := HistoricalshrinkageaggregateresponseMap["actualShrinkageSeconds"].(float64); ok {
		ActualShrinkageSecondsInt := int(ActualShrinkageSeconds)
		o.ActualShrinkageSeconds = &ActualShrinkageSecondsInt
	}
	
	if ActualShrinkagePercent, ok := HistoricalshrinkageaggregateresponseMap["actualShrinkagePercent"].(float64); ok {
		o.ActualShrinkagePercent = &ActualShrinkagePercent
	}
    
	if PaidShrinkageSeconds, ok := HistoricalshrinkageaggregateresponseMap["paidShrinkageSeconds"].(float64); ok {
		PaidShrinkageSecondsInt := int(PaidShrinkageSeconds)
		o.PaidShrinkageSeconds = &PaidShrinkageSecondsInt
	}
	
	if UnpaidShrinkageSeconds, ok := HistoricalshrinkageaggregateresponseMap["unpaidShrinkageSeconds"].(float64); ok {
		UnpaidShrinkageSecondsInt := int(UnpaidShrinkageSeconds)
		o.UnpaidShrinkageSeconds = &UnpaidShrinkageSecondsInt
	}
	
	if PlannedShrinkageSeconds, ok := HistoricalshrinkageaggregateresponseMap["plannedShrinkageSeconds"].(float64); ok {
		PlannedShrinkageSecondsInt := int(PlannedShrinkageSeconds)
		o.PlannedShrinkageSeconds = &PlannedShrinkageSecondsInt
	}
	
	if UnplannedShrinkageSeconds, ok := HistoricalshrinkageaggregateresponseMap["unplannedShrinkageSeconds"].(float64); ok {
		UnplannedShrinkageSecondsInt := int(UnplannedShrinkageSeconds)
		o.UnplannedShrinkageSeconds = &UnplannedShrinkageSecondsInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Historicalshrinkageaggregateresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
