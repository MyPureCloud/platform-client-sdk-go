package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Dialerattemptlimitsconfigchangeattemptlimits
type Dialerattemptlimitsconfigchangeattemptlimits struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// MaxAttemptsPerContact
	MaxAttemptsPerContact *int `json:"maxAttemptsPerContact,omitempty"`

	// MaxAttemptsPerNumber
	MaxAttemptsPerNumber *int `json:"maxAttemptsPerNumber,omitempty"`

	// TimeZoneId - The timezone is necessary to define when \"today\" starts and ends
	TimeZoneId *string `json:"timeZoneId,omitempty"`

	// ResetPeriod - After how long the number of attempts will be set back to 0
	ResetPeriod *string `json:"resetPeriod,omitempty"`

	// RecallEntries - Configuration for recall attempts
	RecallEntries *map[string]Dialerattemptlimitsconfigchangerecallentry `json:"recallEntries,omitempty"`

	// BreadthFirstRecalls - Whether recalls are performed before considering other numbers (true) or after (false)
	BreadthFirstRecalls *bool `json:"breadthFirstRecalls,omitempty"`

	// AdditionalProperties
	AdditionalProperties *map[string]interface{} `json:"additionalProperties,omitempty"`

	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name - The UI-visible name of the object
	Name *string `json:"name,omitempty"`

	// DateCreated - Creation time of the entity
	DateCreated *time.Time `json:"dateCreated,omitempty"`

	// DateModified - Last modified time of the entity
	DateModified *time.Time `json:"dateModified,omitempty"`

	// Version - Required for updates, must match the version number of the most recent update
	Version *int `json:"version,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Dialerattemptlimitsconfigchangeattemptlimits) SetField(field string, fieldValue interface{}) {
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

func (o Dialerattemptlimitsconfigchangeattemptlimits) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateCreated","DateModified", }
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
	type Alias Dialerattemptlimitsconfigchangeattemptlimits
	
	DateCreated := new(string)
	if o.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(o.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	DateModified := new(string)
	if o.DateModified != nil {
		
		*DateModified = timeutil.Strftime(o.DateModified, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateModified = nil
	}
	
	return json.Marshal(&struct { 
		MaxAttemptsPerContact *int `json:"maxAttemptsPerContact,omitempty"`
		
		MaxAttemptsPerNumber *int `json:"maxAttemptsPerNumber,omitempty"`
		
		TimeZoneId *string `json:"timeZoneId,omitempty"`
		
		ResetPeriod *string `json:"resetPeriod,omitempty"`
		
		RecallEntries *map[string]Dialerattemptlimitsconfigchangerecallentry `json:"recallEntries,omitempty"`
		
		BreadthFirstRecalls *bool `json:"breadthFirstRecalls,omitempty"`
		
		AdditionalProperties *map[string]interface{} `json:"additionalProperties,omitempty"`
		
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		Version *int `json:"version,omitempty"`
		Alias
	}{ 
		MaxAttemptsPerContact: o.MaxAttemptsPerContact,
		
		MaxAttemptsPerNumber: o.MaxAttemptsPerNumber,
		
		TimeZoneId: o.TimeZoneId,
		
		ResetPeriod: o.ResetPeriod,
		
		RecallEntries: o.RecallEntries,
		
		BreadthFirstRecalls: o.BreadthFirstRecalls,
		
		AdditionalProperties: o.AdditionalProperties,
		
		Id: o.Id,
		
		Name: o.Name,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		Version: o.Version,
		Alias:    (Alias)(o),
	})
}

func (o *Dialerattemptlimitsconfigchangeattemptlimits) UnmarshalJSON(b []byte) error {
	var DialerattemptlimitsconfigchangeattemptlimitsMap map[string]interface{}
	err := json.Unmarshal(b, &DialerattemptlimitsconfigchangeattemptlimitsMap)
	if err != nil {
		return err
	}
	
	if MaxAttemptsPerContact, ok := DialerattemptlimitsconfigchangeattemptlimitsMap["maxAttemptsPerContact"].(float64); ok {
		MaxAttemptsPerContactInt := int(MaxAttemptsPerContact)
		o.MaxAttemptsPerContact = &MaxAttemptsPerContactInt
	}
	
	if MaxAttemptsPerNumber, ok := DialerattemptlimitsconfigchangeattemptlimitsMap["maxAttemptsPerNumber"].(float64); ok {
		MaxAttemptsPerNumberInt := int(MaxAttemptsPerNumber)
		o.MaxAttemptsPerNumber = &MaxAttemptsPerNumberInt
	}
	
	if TimeZoneId, ok := DialerattemptlimitsconfigchangeattemptlimitsMap["timeZoneId"].(string); ok {
		o.TimeZoneId = &TimeZoneId
	}
    
	if ResetPeriod, ok := DialerattemptlimitsconfigchangeattemptlimitsMap["resetPeriod"].(string); ok {
		o.ResetPeriod = &ResetPeriod
	}
    
	if RecallEntries, ok := DialerattemptlimitsconfigchangeattemptlimitsMap["recallEntries"].(map[string]interface{}); ok {
		RecallEntriesString, _ := json.Marshal(RecallEntries)
		json.Unmarshal(RecallEntriesString, &o.RecallEntries)
	}
	
	if BreadthFirstRecalls, ok := DialerattemptlimitsconfigchangeattemptlimitsMap["breadthFirstRecalls"].(bool); ok {
		o.BreadthFirstRecalls = &BreadthFirstRecalls
	}
    
	if AdditionalProperties, ok := DialerattemptlimitsconfigchangeattemptlimitsMap["additionalProperties"].(map[string]interface{}); ok {
		AdditionalPropertiesString, _ := json.Marshal(AdditionalProperties)
		json.Unmarshal(AdditionalPropertiesString, &o.AdditionalProperties)
	}
	
	if Id, ok := DialerattemptlimitsconfigchangeattemptlimitsMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := DialerattemptlimitsconfigchangeattemptlimitsMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if dateCreatedString, ok := DialerattemptlimitsconfigchangeattemptlimitsMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := DialerattemptlimitsconfigchangeattemptlimitsMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if Version, ok := DialerattemptlimitsconfigchangeattemptlimitsMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Dialerattemptlimitsconfigchangeattemptlimits) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
