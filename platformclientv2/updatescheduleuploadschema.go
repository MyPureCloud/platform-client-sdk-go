package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Updatescheduleuploadschema
type Updatescheduleuploadschema struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Description - The description to set for the schedule
	Description *string `json:"description,omitempty"`

	// Published - Whether to publish the schedule. Note: a schedule cannot be un-published unless another schedule is published over it
	Published *bool `json:"published,omitempty"`

	// ShortTermForecast - The short term forecast to associate with the schedule
	ShortTermForecast *Bushorttermforecastreference `json:"shortTermForecast,omitempty"`

	// HeadcountForecast - The headcount forecast to associate with the schedule
	HeadcountForecast *Buheadcountforecast `json:"headcountForecast,omitempty"`

	// AgentSchedules - Individual agent schedules
	AgentSchedules *[]Buupdateagentscheduleuploadschema `json:"agentSchedules,omitempty"`

	// Metadata - Version metadata for this schedule
	Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Updatescheduleuploadschema) SetField(field string, fieldValue interface{}) {
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

func (o Updatescheduleuploadschema) MarshalJSON() ([]byte, error) {
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
	type Alias Updatescheduleuploadschema
	
	return json.Marshal(&struct { 
		Description *string `json:"description,omitempty"`
		
		Published *bool `json:"published,omitempty"`
		
		ShortTermForecast *Bushorttermforecastreference `json:"shortTermForecast,omitempty"`
		
		HeadcountForecast *Buheadcountforecast `json:"headcountForecast,omitempty"`
		
		AgentSchedules *[]Buupdateagentscheduleuploadschema `json:"agentSchedules,omitempty"`
		
		Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`
		Alias
	}{ 
		Description: o.Description,
		
		Published: o.Published,
		
		ShortTermForecast: o.ShortTermForecast,
		
		HeadcountForecast: o.HeadcountForecast,
		
		AgentSchedules: o.AgentSchedules,
		
		Metadata: o.Metadata,
		Alias:    (Alias)(o),
	})
}

func (o *Updatescheduleuploadschema) UnmarshalJSON(b []byte) error {
	var UpdatescheduleuploadschemaMap map[string]interface{}
	err := json.Unmarshal(b, &UpdatescheduleuploadschemaMap)
	if err != nil {
		return err
	}
	
	if Description, ok := UpdatescheduleuploadschemaMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if Published, ok := UpdatescheduleuploadschemaMap["published"].(bool); ok {
		o.Published = &Published
	}
    
	if ShortTermForecast, ok := UpdatescheduleuploadschemaMap["shortTermForecast"].(map[string]interface{}); ok {
		ShortTermForecastString, _ := json.Marshal(ShortTermForecast)
		json.Unmarshal(ShortTermForecastString, &o.ShortTermForecast)
	}
	
	if HeadcountForecast, ok := UpdatescheduleuploadschemaMap["headcountForecast"].(map[string]interface{}); ok {
		HeadcountForecastString, _ := json.Marshal(HeadcountForecast)
		json.Unmarshal(HeadcountForecastString, &o.HeadcountForecast)
	}
	
	if AgentSchedules, ok := UpdatescheduleuploadschemaMap["agentSchedules"].([]interface{}); ok {
		AgentSchedulesString, _ := json.Marshal(AgentSchedules)
		json.Unmarshal(AgentSchedulesString, &o.AgentSchedules)
	}
	
	if Metadata, ok := UpdatescheduleuploadschemaMap["metadata"].(map[string]interface{}); ok {
		MetadataString, _ := json.Marshal(Metadata)
		json.Unmarshal(MetadataString, &o.Metadata)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Updatescheduleuploadschema) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
