package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmbushorttermforecastcopycompletetopicbushorttermforecast
type Wfmbushorttermforecastcopycompletetopicbushorttermforecast struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id
	Id *string `json:"id,omitempty"`

	// WeekDate
	WeekDate *string `json:"weekDate,omitempty"`

	// CreationMethod
	CreationMethod *string `json:"creationMethod,omitempty"`

	// Description
	Description *string `json:"description,omitempty"`

	// Legacy
	Legacy *bool `json:"legacy,omitempty"`

	// ReferenceStartDate
	ReferenceStartDate *time.Time `json:"referenceStartDate,omitempty"`

	// SourceDays
	SourceDays *[]Wfmbushorttermforecastcopycompletetopicforecastsourcedaypointer `json:"sourceDays,omitempty"`

	// Modifications
	Modifications *[]Wfmbushorttermforecastcopycompletetopicbuforecastmodification `json:"modifications,omitempty"`

	// TimeZone
	TimeZone *string `json:"timeZone,omitempty"`

	// PlanningGroupsVersion
	PlanningGroupsVersion *int `json:"planningGroupsVersion,omitempty"`

	// WeekCount
	WeekCount *int `json:"weekCount,omitempty"`

	// Metadata
	Metadata *Wfmbushorttermforecastcopycompletetopicwfmversionedentitymetadata `json:"metadata,omitempty"`

	// CanUseForScheduling
	CanUseForScheduling *bool `json:"canUseForScheduling,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Wfmbushorttermforecastcopycompletetopicbushorttermforecast) SetField(field string, fieldValue interface{}) {
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

func (o Wfmbushorttermforecastcopycompletetopicbushorttermforecast) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "ReferenceStartDate", }
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
	type Alias Wfmbushorttermforecastcopycompletetopicbushorttermforecast
	
	ReferenceStartDate := new(string)
	if o.ReferenceStartDate != nil {
		
		*ReferenceStartDate = timeutil.Strftime(o.ReferenceStartDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ReferenceStartDate = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		WeekDate *string `json:"weekDate,omitempty"`
		
		CreationMethod *string `json:"creationMethod,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		Legacy *bool `json:"legacy,omitempty"`
		
		ReferenceStartDate *string `json:"referenceStartDate,omitempty"`
		
		SourceDays *[]Wfmbushorttermforecastcopycompletetopicforecastsourcedaypointer `json:"sourceDays,omitempty"`
		
		Modifications *[]Wfmbushorttermforecastcopycompletetopicbuforecastmodification `json:"modifications,omitempty"`
		
		TimeZone *string `json:"timeZone,omitempty"`
		
		PlanningGroupsVersion *int `json:"planningGroupsVersion,omitempty"`
		
		WeekCount *int `json:"weekCount,omitempty"`
		
		Metadata *Wfmbushorttermforecastcopycompletetopicwfmversionedentitymetadata `json:"metadata,omitempty"`
		
		CanUseForScheduling *bool `json:"canUseForScheduling,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		WeekDate: o.WeekDate,
		
		CreationMethod: o.CreationMethod,
		
		Description: o.Description,
		
		Legacy: o.Legacy,
		
		ReferenceStartDate: ReferenceStartDate,
		
		SourceDays: o.SourceDays,
		
		Modifications: o.Modifications,
		
		TimeZone: o.TimeZone,
		
		PlanningGroupsVersion: o.PlanningGroupsVersion,
		
		WeekCount: o.WeekCount,
		
		Metadata: o.Metadata,
		
		CanUseForScheduling: o.CanUseForScheduling,
		Alias:    (Alias)(o),
	})
}

func (o *Wfmbushorttermforecastcopycompletetopicbushorttermforecast) UnmarshalJSON(b []byte) error {
	var WfmbushorttermforecastcopycompletetopicbushorttermforecastMap map[string]interface{}
	err := json.Unmarshal(b, &WfmbushorttermforecastcopycompletetopicbushorttermforecastMap)
	if err != nil {
		return err
	}
	
	if Id, ok := WfmbushorttermforecastcopycompletetopicbushorttermforecastMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if WeekDate, ok := WfmbushorttermforecastcopycompletetopicbushorttermforecastMap["weekDate"].(string); ok {
		o.WeekDate = &WeekDate
	}
    
	if CreationMethod, ok := WfmbushorttermforecastcopycompletetopicbushorttermforecastMap["creationMethod"].(string); ok {
		o.CreationMethod = &CreationMethod
	}
    
	if Description, ok := WfmbushorttermforecastcopycompletetopicbushorttermforecastMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if Legacy, ok := WfmbushorttermforecastcopycompletetopicbushorttermforecastMap["legacy"].(bool); ok {
		o.Legacy = &Legacy
	}
    
	if referenceStartDateString, ok := WfmbushorttermforecastcopycompletetopicbushorttermforecastMap["referenceStartDate"].(string); ok {
		ReferenceStartDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", referenceStartDateString)
		o.ReferenceStartDate = &ReferenceStartDate
	}
	
	if SourceDays, ok := WfmbushorttermforecastcopycompletetopicbushorttermforecastMap["sourceDays"].([]interface{}); ok {
		SourceDaysString, _ := json.Marshal(SourceDays)
		json.Unmarshal(SourceDaysString, &o.SourceDays)
	}
	
	if Modifications, ok := WfmbushorttermforecastcopycompletetopicbushorttermforecastMap["modifications"].([]interface{}); ok {
		ModificationsString, _ := json.Marshal(Modifications)
		json.Unmarshal(ModificationsString, &o.Modifications)
	}
	
	if TimeZone, ok := WfmbushorttermforecastcopycompletetopicbushorttermforecastMap["timeZone"].(string); ok {
		o.TimeZone = &TimeZone
	}
    
	if PlanningGroupsVersion, ok := WfmbushorttermforecastcopycompletetopicbushorttermforecastMap["planningGroupsVersion"].(float64); ok {
		PlanningGroupsVersionInt := int(PlanningGroupsVersion)
		o.PlanningGroupsVersion = &PlanningGroupsVersionInt
	}
	
	if WeekCount, ok := WfmbushorttermforecastcopycompletetopicbushorttermforecastMap["weekCount"].(float64); ok {
		WeekCountInt := int(WeekCount)
		o.WeekCount = &WeekCountInt
	}
	
	if Metadata, ok := WfmbushorttermforecastcopycompletetopicbushorttermforecastMap["metadata"].(map[string]interface{}); ok {
		MetadataString, _ := json.Marshal(Metadata)
		json.Unmarshal(MetadataString, &o.Metadata)
	}
	
	if CanUseForScheduling, ok := WfmbushorttermforecastcopycompletetopicbushorttermforecastMap["canUseForScheduling"].(bool); ok {
		o.CanUseForScheduling = &CanUseForScheduling
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmbushorttermforecastcopycompletetopicbushorttermforecast) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
