package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Employeeperformanceexternalmetricsdefinitionexternalmetricsdefinition
type Employeeperformanceexternalmetricsdefinitionexternalmetricsdefinition struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id
	Id *string `json:"id,omitempty"`

	// Name
	Name *string `json:"name,omitempty"`

	// Unit
	Unit *string `json:"unit,omitempty"`

	// UnitDefinition
	UnitDefinition *string `json:"unitDefinition,omitempty"`

	// Precision
	Precision *int `json:"precision,omitempty"`

	// DefaultObjectiveType
	DefaultObjectiveType *string `json:"defaultObjectiveType,omitempty"`

	// RetentionMonths
	RetentionMonths *int `json:"retentionMonths,omitempty"`

	// Enabled
	Enabled *bool `json:"enabled,omitempty"`

	// InUse
	InUse *bool `json:"inUse,omitempty"`

	// DateLastRefreshed
	DateLastRefreshed *time.Time `json:"dateLastRefreshed,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Employeeperformanceexternalmetricsdefinitionexternalmetricsdefinition) SetField(field string, fieldValue interface{}) {
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

func (o Employeeperformanceexternalmetricsdefinitionexternalmetricsdefinition) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateLastRefreshed", }
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
	type Alias Employeeperformanceexternalmetricsdefinitionexternalmetricsdefinition
	
	DateLastRefreshed := new(string)
	if o.DateLastRefreshed != nil {
		
		*DateLastRefreshed = timeutil.Strftime(o.DateLastRefreshed, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateLastRefreshed = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Unit *string `json:"unit,omitempty"`
		
		UnitDefinition *string `json:"unitDefinition,omitempty"`
		
		Precision *int `json:"precision,omitempty"`
		
		DefaultObjectiveType *string `json:"defaultObjectiveType,omitempty"`
		
		RetentionMonths *int `json:"retentionMonths,omitempty"`
		
		Enabled *bool `json:"enabled,omitempty"`
		
		InUse *bool `json:"inUse,omitempty"`
		
		DateLastRefreshed *string `json:"dateLastRefreshed,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Unit: o.Unit,
		
		UnitDefinition: o.UnitDefinition,
		
		Precision: o.Precision,
		
		DefaultObjectiveType: o.DefaultObjectiveType,
		
		RetentionMonths: o.RetentionMonths,
		
		Enabled: o.Enabled,
		
		InUse: o.InUse,
		
		DateLastRefreshed: DateLastRefreshed,
		Alias:    (Alias)(o),
	})
}

func (o *Employeeperformanceexternalmetricsdefinitionexternalmetricsdefinition) UnmarshalJSON(b []byte) error {
	var EmployeeperformanceexternalmetricsdefinitionexternalmetricsdefinitionMap map[string]interface{}
	err := json.Unmarshal(b, &EmployeeperformanceexternalmetricsdefinitionexternalmetricsdefinitionMap)
	if err != nil {
		return err
	}
	
	if Id, ok := EmployeeperformanceexternalmetricsdefinitionexternalmetricsdefinitionMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := EmployeeperformanceexternalmetricsdefinitionexternalmetricsdefinitionMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Unit, ok := EmployeeperformanceexternalmetricsdefinitionexternalmetricsdefinitionMap["unit"].(string); ok {
		o.Unit = &Unit
	}
    
	if UnitDefinition, ok := EmployeeperformanceexternalmetricsdefinitionexternalmetricsdefinitionMap["unitDefinition"].(string); ok {
		o.UnitDefinition = &UnitDefinition
	}
    
	if Precision, ok := EmployeeperformanceexternalmetricsdefinitionexternalmetricsdefinitionMap["precision"].(float64); ok {
		PrecisionInt := int(Precision)
		o.Precision = &PrecisionInt
	}
	
	if DefaultObjectiveType, ok := EmployeeperformanceexternalmetricsdefinitionexternalmetricsdefinitionMap["defaultObjectiveType"].(string); ok {
		o.DefaultObjectiveType = &DefaultObjectiveType
	}
    
	if RetentionMonths, ok := EmployeeperformanceexternalmetricsdefinitionexternalmetricsdefinitionMap["retentionMonths"].(float64); ok {
		RetentionMonthsInt := int(RetentionMonths)
		o.RetentionMonths = &RetentionMonthsInt
	}
	
	if Enabled, ok := EmployeeperformanceexternalmetricsdefinitionexternalmetricsdefinitionMap["enabled"].(bool); ok {
		o.Enabled = &Enabled
	}
    
	if InUse, ok := EmployeeperformanceexternalmetricsdefinitionexternalmetricsdefinitionMap["inUse"].(bool); ok {
		o.InUse = &InUse
	}
    
	if dateLastRefreshedString, ok := EmployeeperformanceexternalmetricsdefinitionexternalmetricsdefinitionMap["dateLastRefreshed"].(string); ok {
		DateLastRefreshed, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateLastRefreshedString)
		o.DateLastRefreshed = &DateLastRefreshed
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Employeeperformanceexternalmetricsdefinitionexternalmetricsdefinition) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
