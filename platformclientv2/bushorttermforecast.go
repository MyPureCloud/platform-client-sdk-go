package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Bushorttermforecast
type Bushorttermforecast struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// WeekDate - The start week date of this forecast in yyyy-MM-dd.  Must fall on the start day of week for the associated business unit. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	WeekDate *time.Time `json:"weekDate,omitempty"`

	// WeekCount - The number of weeks this forecast covers
	WeekCount *int `json:"weekCount,omitempty"`

	// CreationMethod - The method by which this forecast was created
	CreationMethod *string `json:"creationMethod,omitempty"`

	// Description - The description of this forecast
	Description *string `json:"description,omitempty"`

	// Legacy - Whether this forecast contains modifications on legacy metrics
	Legacy *bool `json:"legacy,omitempty"`

	// Metadata - Metadata for this forecast
	Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`

	// CanUseForScheduling - Whether this forecast can be used for scheduling
	CanUseForScheduling *bool `json:"canUseForScheduling,omitempty"`

	// ReferenceStartDate - The reference start date for interval-based data for this forecast. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ReferenceStartDate *time.Time `json:"referenceStartDate,omitempty"`

	// SourceDays - The source day pointers for this forecast
	SourceDays *[]Forecastsourcedaypointer `json:"sourceDays,omitempty"`

	// Modifications - Any manual modifications applied to this forecast
	Modifications *[]Buforecastmodification `json:"modifications,omitempty"`

	// GenerationResults - Generation result metadata
	GenerationResults *Buforecastgenerationresult `json:"generationResults,omitempty"`

	// TimeZone - The time zone for this forecast
	TimeZone *string `json:"timeZone,omitempty"`

	// PlanningGroupsVersion - The version of the planning groups that was used for this forecast
	PlanningGroupsVersion *int `json:"planningGroupsVersion,omitempty"`

	// PlanningGroups - A snapshot of the planning groups used for this forecast as of the version number indicated
	PlanningGroups *Forecastplanninggroupsresponse `json:"planningGroups,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Bushorttermforecast) SetField(field string, fieldValue interface{}) {
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

func (o Bushorttermforecast) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "ReferenceStartDate", }
		localDateTimeFields := []string{  }
		dateFields := []string{ "WeekDate", }

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
	type Alias Bushorttermforecast
	
	WeekDate := new(string)
	if o.WeekDate != nil {
		*WeekDate = timeutil.Strftime(o.WeekDate, "%Y-%m-%d")
	} else {
		WeekDate = nil
	}
	
	ReferenceStartDate := new(string)
	if o.ReferenceStartDate != nil {
		
		*ReferenceStartDate = timeutil.Strftime(o.ReferenceStartDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ReferenceStartDate = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		WeekDate *string `json:"weekDate,omitempty"`
		
		WeekCount *int `json:"weekCount,omitempty"`
		
		CreationMethod *string `json:"creationMethod,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		Legacy *bool `json:"legacy,omitempty"`
		
		Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`
		
		CanUseForScheduling *bool `json:"canUseForScheduling,omitempty"`
		
		ReferenceStartDate *string `json:"referenceStartDate,omitempty"`
		
		SourceDays *[]Forecastsourcedaypointer `json:"sourceDays,omitempty"`
		
		Modifications *[]Buforecastmodification `json:"modifications,omitempty"`
		
		GenerationResults *Buforecastgenerationresult `json:"generationResults,omitempty"`
		
		TimeZone *string `json:"timeZone,omitempty"`
		
		PlanningGroupsVersion *int `json:"planningGroupsVersion,omitempty"`
		
		PlanningGroups *Forecastplanninggroupsresponse `json:"planningGroups,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		WeekDate: WeekDate,
		
		WeekCount: o.WeekCount,
		
		CreationMethod: o.CreationMethod,
		
		Description: o.Description,
		
		Legacy: o.Legacy,
		
		Metadata: o.Metadata,
		
		CanUseForScheduling: o.CanUseForScheduling,
		
		ReferenceStartDate: ReferenceStartDate,
		
		SourceDays: o.SourceDays,
		
		Modifications: o.Modifications,
		
		GenerationResults: o.GenerationResults,
		
		TimeZone: o.TimeZone,
		
		PlanningGroupsVersion: o.PlanningGroupsVersion,
		
		PlanningGroups: o.PlanningGroups,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Bushorttermforecast) UnmarshalJSON(b []byte) error {
	var BushorttermforecastMap map[string]interface{}
	err := json.Unmarshal(b, &BushorttermforecastMap)
	if err != nil {
		return err
	}
	
	if Id, ok := BushorttermforecastMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if weekDateString, ok := BushorttermforecastMap["weekDate"].(string); ok {
		WeekDate, _ := time.Parse("2006-01-02", weekDateString)
		o.WeekDate = &WeekDate
	}
	
	if WeekCount, ok := BushorttermforecastMap["weekCount"].(float64); ok {
		WeekCountInt := int(WeekCount)
		o.WeekCount = &WeekCountInt
	}
	
	if CreationMethod, ok := BushorttermforecastMap["creationMethod"].(string); ok {
		o.CreationMethod = &CreationMethod
	}
    
	if Description, ok := BushorttermforecastMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if Legacy, ok := BushorttermforecastMap["legacy"].(bool); ok {
		o.Legacy = &Legacy
	}
    
	if Metadata, ok := BushorttermforecastMap["metadata"].(map[string]interface{}); ok {
		MetadataString, _ := json.Marshal(Metadata)
		json.Unmarshal(MetadataString, &o.Metadata)
	}
	
	if CanUseForScheduling, ok := BushorttermforecastMap["canUseForScheduling"].(bool); ok {
		o.CanUseForScheduling = &CanUseForScheduling
	}
    
	if referenceStartDateString, ok := BushorttermforecastMap["referenceStartDate"].(string); ok {
		ReferenceStartDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", referenceStartDateString)
		o.ReferenceStartDate = &ReferenceStartDate
	}
	
	if SourceDays, ok := BushorttermforecastMap["sourceDays"].([]interface{}); ok {
		SourceDaysString, _ := json.Marshal(SourceDays)
		json.Unmarshal(SourceDaysString, &o.SourceDays)
	}
	
	if Modifications, ok := BushorttermforecastMap["modifications"].([]interface{}); ok {
		ModificationsString, _ := json.Marshal(Modifications)
		json.Unmarshal(ModificationsString, &o.Modifications)
	}
	
	if GenerationResults, ok := BushorttermforecastMap["generationResults"].(map[string]interface{}); ok {
		GenerationResultsString, _ := json.Marshal(GenerationResults)
		json.Unmarshal(GenerationResultsString, &o.GenerationResults)
	}
	
	if TimeZone, ok := BushorttermforecastMap["timeZone"].(string); ok {
		o.TimeZone = &TimeZone
	}
    
	if PlanningGroupsVersion, ok := BushorttermforecastMap["planningGroupsVersion"].(float64); ok {
		PlanningGroupsVersionInt := int(PlanningGroupsVersion)
		o.PlanningGroupsVersion = &PlanningGroupsVersionInt
	}
	
	if PlanningGroups, ok := BushorttermforecastMap["planningGroups"].(map[string]interface{}); ok {
		PlanningGroupsString, _ := json.Marshal(PlanningGroups)
		json.Unmarshal(PlanningGroupsString, &o.PlanningGroups)
	}
	
	if SelfUri, ok := BushorttermforecastMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Bushorttermforecast) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
