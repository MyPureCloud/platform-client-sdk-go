package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Createservicegoaltemplate
type Createservicegoaltemplate struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Name - The name of the service goal template.
	Name *string `json:"name,omitempty"`

	// ServiceLevel - Service level targets for this service goal template
	ServiceLevel *Buservicelevel `json:"serviceLevel,omitempty"`

	// AverageSpeedOfAnswer - Average speed of answer targets for this service goal template
	AverageSpeedOfAnswer *Buaveragespeedofanswer `json:"averageSpeedOfAnswer,omitempty"`

	// AbandonRate - Abandon rate targets for this service goal template
	AbandonRate *Buabandonrate `json:"abandonRate,omitempty"`

	// ImpactOverride - Settings controlling max percent increase and decrease of service goals for this service goal template
	ImpactOverride *Servicegoaltemplateimpactoverride `json:"impactOverride,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Createservicegoaltemplate) SetField(field string, fieldValue interface{}) {
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

func (o Createservicegoaltemplate) MarshalJSON() ([]byte, error) {
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
	type Alias Createservicegoaltemplate
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		ServiceLevel *Buservicelevel `json:"serviceLevel,omitempty"`
		
		AverageSpeedOfAnswer *Buaveragespeedofanswer `json:"averageSpeedOfAnswer,omitempty"`
		
		AbandonRate *Buabandonrate `json:"abandonRate,omitempty"`
		
		ImpactOverride *Servicegoaltemplateimpactoverride `json:"impactOverride,omitempty"`
		Alias
	}{ 
		Name: o.Name,
		
		ServiceLevel: o.ServiceLevel,
		
		AverageSpeedOfAnswer: o.AverageSpeedOfAnswer,
		
		AbandonRate: o.AbandonRate,
		
		ImpactOverride: o.ImpactOverride,
		Alias:    (Alias)(o),
	})
}

func (o *Createservicegoaltemplate) UnmarshalJSON(b []byte) error {
	var CreateservicegoaltemplateMap map[string]interface{}
	err := json.Unmarshal(b, &CreateservicegoaltemplateMap)
	if err != nil {
		return err
	}
	
	if Name, ok := CreateservicegoaltemplateMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if ServiceLevel, ok := CreateservicegoaltemplateMap["serviceLevel"].(map[string]interface{}); ok {
		ServiceLevelString, _ := json.Marshal(ServiceLevel)
		json.Unmarshal(ServiceLevelString, &o.ServiceLevel)
	}
	
	if AverageSpeedOfAnswer, ok := CreateservicegoaltemplateMap["averageSpeedOfAnswer"].(map[string]interface{}); ok {
		AverageSpeedOfAnswerString, _ := json.Marshal(AverageSpeedOfAnswer)
		json.Unmarshal(AverageSpeedOfAnswerString, &o.AverageSpeedOfAnswer)
	}
	
	if AbandonRate, ok := CreateservicegoaltemplateMap["abandonRate"].(map[string]interface{}); ok {
		AbandonRateString, _ := json.Marshal(AbandonRate)
		json.Unmarshal(AbandonRateString, &o.AbandonRate)
	}
	
	if ImpactOverride, ok := CreateservicegoaltemplateMap["impactOverride"].(map[string]interface{}); ok {
		ImpactOverrideString, _ := json.Marshal(ImpactOverride)
		json.Unmarshal(ImpactOverrideString, &o.ImpactOverride)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Createservicegoaltemplate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
