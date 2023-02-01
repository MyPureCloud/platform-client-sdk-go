package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Weekschedulelistitemresponse
type Weekschedulelistitemresponse struct { 
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
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Weekschedulelistitemresponse) SetField(field string, fieldValue interface{}) {
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

func (o Weekschedulelistitemresponse) MarshalJSON() ([]byte, error) {
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
	type Alias Weekschedulelistitemresponse
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		
		WeekDate *string `json:"weekDate,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		Published *bool `json:"published,omitempty"`
		
		GenerationResults *Weekschedulegenerationresult `json:"generationResults,omitempty"`
		
		ShortTermForecast *Shorttermforecastreference `json:"shortTermForecast,omitempty"`
		
		Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`
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
		Alias:    (Alias)(o),
	})
}

func (o *Weekschedulelistitemresponse) UnmarshalJSON(b []byte) error {
	var WeekschedulelistitemresponseMap map[string]interface{}
	err := json.Unmarshal(b, &WeekschedulelistitemresponseMap)
	if err != nil {
		return err
	}
	
	if Id, ok := WeekschedulelistitemresponseMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if SelfUri, ok := WeekschedulelistitemresponseMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    
	if WeekDate, ok := WeekschedulelistitemresponseMap["weekDate"].(string); ok {
		o.WeekDate = &WeekDate
	}
    
	if Description, ok := WeekschedulelistitemresponseMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if Published, ok := WeekschedulelistitemresponseMap["published"].(bool); ok {
		o.Published = &Published
	}
    
	if GenerationResults, ok := WeekschedulelistitemresponseMap["generationResults"].(map[string]interface{}); ok {
		GenerationResultsString, _ := json.Marshal(GenerationResults)
		json.Unmarshal(GenerationResultsString, &o.GenerationResults)
	}
	
	if ShortTermForecast, ok := WeekschedulelistitemresponseMap["shortTermForecast"].(map[string]interface{}); ok {
		ShortTermForecastString, _ := json.Marshal(ShortTermForecast)
		json.Unmarshal(ShortTermForecastString, &o.ShortTermForecast)
	}
	
	if Metadata, ok := WeekschedulelistitemresponseMap["metadata"].(map[string]interface{}); ok {
		MetadataString, _ := json.Marshal(Metadata)
		json.Unmarshal(MetadataString, &o.Metadata)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Weekschedulelistitemresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
