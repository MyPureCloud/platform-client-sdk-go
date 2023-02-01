package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Burescheduleresult
type Burescheduleresult struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// GenerationResults - The generation results.  Note the result will always be delivered via the generationResultsDownloadUrl; however the schema is included for documentation
	GenerationResults *Schedulegenerationresult `json:"generationResults,omitempty"`

	// GenerationResultsDownloadUrl - The download URL from which to fetch the generation results for the rescheduling run
	GenerationResultsDownloadUrl *string `json:"generationResultsDownloadUrl,omitempty"`

	// HeadcountForecast - The headcount forecast.  Note the result will always be delivered via the headcountForecastDownloadUrl; however the schema is included for documentation
	HeadcountForecast *Buheadcountforecast `json:"headcountForecast,omitempty"`

	// HeadcountForecastDownloadUrl - The download URL from which to fetch the headcount forecast for the rescheduling run
	HeadcountForecastDownloadUrl *string `json:"headcountForecastDownloadUrl,omitempty"`

	// AgentSchedules - List of download links for agent schedules produced by the rescheduling run
	AgentSchedules *[]Burescheduleagentscheduleresult `json:"agentSchedules,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Burescheduleresult) SetField(field string, fieldValue interface{}) {
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

func (o Burescheduleresult) MarshalJSON() ([]byte, error) {
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
	type Alias Burescheduleresult
	
	return json.Marshal(&struct { 
		GenerationResults *Schedulegenerationresult `json:"generationResults,omitempty"`
		
		GenerationResultsDownloadUrl *string `json:"generationResultsDownloadUrl,omitempty"`
		
		HeadcountForecast *Buheadcountforecast `json:"headcountForecast,omitempty"`
		
		HeadcountForecastDownloadUrl *string `json:"headcountForecastDownloadUrl,omitempty"`
		
		AgentSchedules *[]Burescheduleagentscheduleresult `json:"agentSchedules,omitempty"`
		Alias
	}{ 
		GenerationResults: o.GenerationResults,
		
		GenerationResultsDownloadUrl: o.GenerationResultsDownloadUrl,
		
		HeadcountForecast: o.HeadcountForecast,
		
		HeadcountForecastDownloadUrl: o.HeadcountForecastDownloadUrl,
		
		AgentSchedules: o.AgentSchedules,
		Alias:    (Alias)(o),
	})
}

func (o *Burescheduleresult) UnmarshalJSON(b []byte) error {
	var BurescheduleresultMap map[string]interface{}
	err := json.Unmarshal(b, &BurescheduleresultMap)
	if err != nil {
		return err
	}
	
	if GenerationResults, ok := BurescheduleresultMap["generationResults"].(map[string]interface{}); ok {
		GenerationResultsString, _ := json.Marshal(GenerationResults)
		json.Unmarshal(GenerationResultsString, &o.GenerationResults)
	}
	
	if GenerationResultsDownloadUrl, ok := BurescheduleresultMap["generationResultsDownloadUrl"].(string); ok {
		o.GenerationResultsDownloadUrl = &GenerationResultsDownloadUrl
	}
    
	if HeadcountForecast, ok := BurescheduleresultMap["headcountForecast"].(map[string]interface{}); ok {
		HeadcountForecastString, _ := json.Marshal(HeadcountForecast)
		json.Unmarshal(HeadcountForecastString, &o.HeadcountForecast)
	}
	
	if HeadcountForecastDownloadUrl, ok := BurescheduleresultMap["headcountForecastDownloadUrl"].(string); ok {
		o.HeadcountForecastDownloadUrl = &HeadcountForecastDownloadUrl
	}
    
	if AgentSchedules, ok := BurescheduleresultMap["agentSchedules"].([]interface{}); ok {
		AgentSchedulesString, _ := json.Marshal(AgentSchedules)
		json.Unmarshal(AgentSchedulesString, &o.AgentSchedules)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Burescheduleresult) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
