package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmbuscheduletopicbuschedulemetadata
type Wfmbuscheduletopicbuschedulemetadata struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id
	Id *string `json:"id,omitempty"`

	// WeekDate
	WeekDate *time.Time `json:"weekDate,omitempty"`

	// WeekCount
	WeekCount *int `json:"weekCount,omitempty"`

	// Description
	Description *string `json:"description,omitempty"`

	// Published
	Published *bool `json:"published,omitempty"`

	// ShortTermForecast
	ShortTermForecast *Wfmbuscheduletopicbushorttermforecastreference `json:"shortTermForecast,omitempty"`

	// ManagementUnits
	ManagementUnits *[]Wfmbuscheduletopicbumanagementunitschedulesummary `json:"managementUnits,omitempty"`

	// GenerationResults
	GenerationResults *Wfmbuscheduletopicbuschedulegenerationresultsummary `json:"generationResults,omitempty"`

	// Metadata
	Metadata *Wfmbuscheduletopicwfmversionedentitymetadata `json:"metadata,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Wfmbuscheduletopicbuschedulemetadata) SetField(field string, fieldValue interface{}) {
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

func (o Wfmbuscheduletopicbuschedulemetadata) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "WeekDate", }
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
	type Alias Wfmbuscheduletopicbuschedulemetadata
	
	WeekDate := new(string)
	if o.WeekDate != nil {
		
		*WeekDate = timeutil.Strftime(o.WeekDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		WeekDate = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		WeekDate *string `json:"weekDate,omitempty"`
		
		WeekCount *int `json:"weekCount,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		Published *bool `json:"published,omitempty"`
		
		ShortTermForecast *Wfmbuscheduletopicbushorttermforecastreference `json:"shortTermForecast,omitempty"`
		
		ManagementUnits *[]Wfmbuscheduletopicbumanagementunitschedulesummary `json:"managementUnits,omitempty"`
		
		GenerationResults *Wfmbuscheduletopicbuschedulegenerationresultsummary `json:"generationResults,omitempty"`
		
		Metadata *Wfmbuscheduletopicwfmversionedentitymetadata `json:"metadata,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		WeekDate: WeekDate,
		
		WeekCount: o.WeekCount,
		
		Description: o.Description,
		
		Published: o.Published,
		
		ShortTermForecast: o.ShortTermForecast,
		
		ManagementUnits: o.ManagementUnits,
		
		GenerationResults: o.GenerationResults,
		
		Metadata: o.Metadata,
		Alias:    (Alias)(o),
	})
}

func (o *Wfmbuscheduletopicbuschedulemetadata) UnmarshalJSON(b []byte) error {
	var WfmbuscheduletopicbuschedulemetadataMap map[string]interface{}
	err := json.Unmarshal(b, &WfmbuscheduletopicbuschedulemetadataMap)
	if err != nil {
		return err
	}
	
	if Id, ok := WfmbuscheduletopicbuschedulemetadataMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if weekDateString, ok := WfmbuscheduletopicbuschedulemetadataMap["weekDate"].(string); ok {
		WeekDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", weekDateString)
		o.WeekDate = &WeekDate
	}
	
	if WeekCount, ok := WfmbuscheduletopicbuschedulemetadataMap["weekCount"].(float64); ok {
		WeekCountInt := int(WeekCount)
		o.WeekCount = &WeekCountInt
	}
	
	if Description, ok := WfmbuscheduletopicbuschedulemetadataMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if Published, ok := WfmbuscheduletopicbuschedulemetadataMap["published"].(bool); ok {
		o.Published = &Published
	}
    
	if ShortTermForecast, ok := WfmbuscheduletopicbuschedulemetadataMap["shortTermForecast"].(map[string]interface{}); ok {
		ShortTermForecastString, _ := json.Marshal(ShortTermForecast)
		json.Unmarshal(ShortTermForecastString, &o.ShortTermForecast)
	}
	
	if ManagementUnits, ok := WfmbuscheduletopicbuschedulemetadataMap["managementUnits"].([]interface{}); ok {
		ManagementUnitsString, _ := json.Marshal(ManagementUnits)
		json.Unmarshal(ManagementUnitsString, &o.ManagementUnits)
	}
	
	if GenerationResults, ok := WfmbuscheduletopicbuschedulemetadataMap["generationResults"].(map[string]interface{}); ok {
		GenerationResultsString, _ := json.Marshal(GenerationResults)
		json.Unmarshal(GenerationResultsString, &o.GenerationResults)
	}
	
	if Metadata, ok := WfmbuscheduletopicbuschedulemetadataMap["metadata"].(map[string]interface{}); ok {
		MetadataString, _ := json.Marshal(Metadata)
		json.Unmarshal(MetadataString, &o.Metadata)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmbuscheduletopicbuschedulemetadata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
