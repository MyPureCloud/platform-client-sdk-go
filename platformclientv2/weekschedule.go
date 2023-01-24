package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Weekschedule
type Weekschedule struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

	// WeekDate - First day of this week schedule in yyyy-MM-dd format
	WeekDate *string `json:"weekDate,omitempty"`

	// Description - Description of the week schedule
	Description *string `json:"description,omitempty"`

	// Published - Whether the week schedule is published
	Published *bool `json:"published,omitempty"`

	// GenerationResults - Summary of the results from the schedule run
	GenerationResults *Weekschedulegenerationresult `json:"generationResults,omitempty"`

	// ShortTermForecast - Short term forecast associated with this schedule
	ShortTermForecast *Shorttermforecastreference `json:"shortTermForecast,omitempty"`

	// Metadata - Version metadata for this work plan
	Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`

	// UserSchedules - User schedules in the week
	UserSchedules *map[string]Userschedule `json:"userSchedules,omitempty"`

	// HeadcountForecast - Headcount information for the week schedule
	HeadcountForecast *Headcountforecast `json:"headcountForecast,omitempty"`

	// AgentSchedulesVersion - Version of agent schedules in the week schedule
	AgentSchedulesVersion *int `json:"agentSchedulesVersion,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Weekschedule) SetField(field string, fieldValue interface{}) {
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

func (o Weekschedule) MarshalJSON() ([]byte, error) {
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
	type Alias Weekschedule
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		
		WeekDate *string `json:"weekDate,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		Published *bool `json:"published,omitempty"`
		
		GenerationResults *Weekschedulegenerationresult `json:"generationResults,omitempty"`
		
		ShortTermForecast *Shorttermforecastreference `json:"shortTermForecast,omitempty"`
		
		Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`
		
		UserSchedules *map[string]Userschedule `json:"userSchedules,omitempty"`
		
		HeadcountForecast *Headcountforecast `json:"headcountForecast,omitempty"`
		
		AgentSchedulesVersion *int `json:"agentSchedulesVersion,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		SelfUri: o.SelfUri,
		
		WeekDate: o.WeekDate,
		
		Description: o.Description,
		
		Published: o.Published,
		
		GenerationResults: o.GenerationResults,
		
		ShortTermForecast: o.ShortTermForecast,
		
		Metadata: o.Metadata,
		
		UserSchedules: o.UserSchedules,
		
		HeadcountForecast: o.HeadcountForecast,
		
		AgentSchedulesVersion: o.AgentSchedulesVersion,
		Alias:    (Alias)(o),
	})
}

func (o *Weekschedule) UnmarshalJSON(b []byte) error {
	var WeekscheduleMap map[string]interface{}
	err := json.Unmarshal(b, &WeekscheduleMap)
	if err != nil {
		return err
	}
	
	if Id, ok := WeekscheduleMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if SelfUri, ok := WeekscheduleMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    
	if WeekDate, ok := WeekscheduleMap["weekDate"].(string); ok {
		o.WeekDate = &WeekDate
	}
    
	if Description, ok := WeekscheduleMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if Published, ok := WeekscheduleMap["published"].(bool); ok {
		o.Published = &Published
	}
    
	if GenerationResults, ok := WeekscheduleMap["generationResults"].(map[string]interface{}); ok {
		GenerationResultsString, _ := json.Marshal(GenerationResults)
		json.Unmarshal(GenerationResultsString, &o.GenerationResults)
	}
	
	if ShortTermForecast, ok := WeekscheduleMap["shortTermForecast"].(map[string]interface{}); ok {
		ShortTermForecastString, _ := json.Marshal(ShortTermForecast)
		json.Unmarshal(ShortTermForecastString, &o.ShortTermForecast)
	}
	
	if Metadata, ok := WeekscheduleMap["metadata"].(map[string]interface{}); ok {
		MetadataString, _ := json.Marshal(Metadata)
		json.Unmarshal(MetadataString, &o.Metadata)
	}
	
	if UserSchedules, ok := WeekscheduleMap["userSchedules"].(map[string]interface{}); ok {
		UserSchedulesString, _ := json.Marshal(UserSchedules)
		json.Unmarshal(UserSchedulesString, &o.UserSchedules)
	}
	
	if HeadcountForecast, ok := WeekscheduleMap["headcountForecast"].(map[string]interface{}); ok {
		HeadcountForecastString, _ := json.Marshal(HeadcountForecast)
		json.Unmarshal(HeadcountForecastString, &o.HeadcountForecast)
	}
	
	if AgentSchedulesVersion, ok := WeekscheduleMap["agentSchedulesVersion"].(float64); ok {
		AgentSchedulesVersionInt := int(AgentSchedulesVersion)
		o.AgentSchedulesVersion = &AgentSchedulesVersionInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Weekschedule) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
