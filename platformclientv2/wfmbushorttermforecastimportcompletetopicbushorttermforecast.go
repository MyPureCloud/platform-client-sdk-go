package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmbushorttermforecastimportcompletetopicbushorttermforecast
type Wfmbushorttermforecastimportcompletetopicbushorttermforecast struct { 
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
	SourceDays *[]Wfmbushorttermforecastimportcompletetopicforecastsourcedaypointer `json:"sourceDays,omitempty"`

	// Modifications
	Modifications *[]Wfmbushorttermforecastimportcompletetopicbuforecastmodification `json:"modifications,omitempty"`

	// TimeZone
	TimeZone *string `json:"timeZone,omitempty"`

	// PlanningGroupsVersion
	PlanningGroupsVersion *int `json:"planningGroupsVersion,omitempty"`

	// WeekCount
	WeekCount *int `json:"weekCount,omitempty"`

	// Metadata
	Metadata *Wfmbushorttermforecastimportcompletetopicwfmversionedentitymetadata `json:"metadata,omitempty"`

	// CanUseForScheduling
	CanUseForScheduling *bool `json:"canUseForScheduling,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Wfmbushorttermforecastimportcompletetopicbushorttermforecast) SetField(field string, fieldValue interface{}) {
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

func (o Wfmbushorttermforecastimportcompletetopicbushorttermforecast) MarshalJSON() ([]byte, error) {
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
	type Alias Wfmbushorttermforecastimportcompletetopicbushorttermforecast
	
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
		
		SourceDays *[]Wfmbushorttermforecastimportcompletetopicforecastsourcedaypointer `json:"sourceDays,omitempty"`
		
		Modifications *[]Wfmbushorttermforecastimportcompletetopicbuforecastmodification `json:"modifications,omitempty"`
		
		TimeZone *string `json:"timeZone,omitempty"`
		
		PlanningGroupsVersion *int `json:"planningGroupsVersion,omitempty"`
		
		WeekCount *int `json:"weekCount,omitempty"`
		
		Metadata *Wfmbushorttermforecastimportcompletetopicwfmversionedentitymetadata `json:"metadata,omitempty"`
		
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

func (o *Wfmbushorttermforecastimportcompletetopicbushorttermforecast) UnmarshalJSON(b []byte) error {
	var WfmbushorttermforecastimportcompletetopicbushorttermforecastMap map[string]interface{}
	err := json.Unmarshal(b, &WfmbushorttermforecastimportcompletetopicbushorttermforecastMap)
	if err != nil {
		return err
	}
	
	if Id, ok := WfmbushorttermforecastimportcompletetopicbushorttermforecastMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if WeekDate, ok := WfmbushorttermforecastimportcompletetopicbushorttermforecastMap["weekDate"].(string); ok {
		o.WeekDate = &WeekDate
	}
    
	if CreationMethod, ok := WfmbushorttermforecastimportcompletetopicbushorttermforecastMap["creationMethod"].(string); ok {
		o.CreationMethod = &CreationMethod
	}
    
	if Description, ok := WfmbushorttermforecastimportcompletetopicbushorttermforecastMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if Legacy, ok := WfmbushorttermforecastimportcompletetopicbushorttermforecastMap["legacy"].(bool); ok {
		o.Legacy = &Legacy
	}
    
	if referenceStartDateString, ok := WfmbushorttermforecastimportcompletetopicbushorttermforecastMap["referenceStartDate"].(string); ok {
		ReferenceStartDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", referenceStartDateString)
		o.ReferenceStartDate = &ReferenceStartDate
	}
	
	if SourceDays, ok := WfmbushorttermforecastimportcompletetopicbushorttermforecastMap["sourceDays"].([]interface{}); ok {
		SourceDaysString, _ := json.Marshal(SourceDays)
		json.Unmarshal(SourceDaysString, &o.SourceDays)
	}
	
	if Modifications, ok := WfmbushorttermforecastimportcompletetopicbushorttermforecastMap["modifications"].([]interface{}); ok {
		ModificationsString, _ := json.Marshal(Modifications)
		json.Unmarshal(ModificationsString, &o.Modifications)
	}
	
	if TimeZone, ok := WfmbushorttermforecastimportcompletetopicbushorttermforecastMap["timeZone"].(string); ok {
		o.TimeZone = &TimeZone
	}
    
	if PlanningGroupsVersion, ok := WfmbushorttermforecastimportcompletetopicbushorttermforecastMap["planningGroupsVersion"].(float64); ok {
		PlanningGroupsVersionInt := int(PlanningGroupsVersion)
		o.PlanningGroupsVersion = &PlanningGroupsVersionInt
	}
	
	if WeekCount, ok := WfmbushorttermforecastimportcompletetopicbushorttermforecastMap["weekCount"].(float64); ok {
		WeekCountInt := int(WeekCount)
		o.WeekCount = &WeekCountInt
	}
	
	if Metadata, ok := WfmbushorttermforecastimportcompletetopicbushorttermforecastMap["metadata"].(map[string]interface{}); ok {
		MetadataString, _ := json.Marshal(Metadata)
		json.Unmarshal(MetadataString, &o.Metadata)
	}
	
	if CanUseForScheduling, ok := WfmbushorttermforecastimportcompletetopicbushorttermforecastMap["canUseForScheduling"].(bool); ok {
		o.CanUseForScheduling = &CanUseForScheduling
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmbushorttermforecastimportcompletetopicbushorttermforecast) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
