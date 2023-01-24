package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Keyperformanceindicator
type Keyperformanceindicator struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name - The name of the Key Performance Indicator.
	Name *string `json:"name,omitempty"`

	// OptimizationType - The optimization type of the Key Performance Indicator.
	OptimizationType *string `json:"optimizationType,omitempty"`

	// DateCreated - DateTime indicating when the Key Performance Indicator was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`

	// DateModified - DateTime indicating when the Key Performance Indicator was modified. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`

	// Description - The description of the Key Performance Indicator.
	Description *string `json:"description,omitempty"`

	// KpiType - The type of Key Performance Indicator.
	KpiType *string `json:"kpiType,omitempty"`

	// Source - Source of values for Key Performance Indicator.
	Source *string `json:"source,omitempty"`

	// WrapUpCodeConfig - Defines what wrap up codes are mapped to Key Performance Indicator.
	WrapUpCodeConfig *Wrapupcodeconfig `json:"wrapUpCodeConfig,omitempty"`

	// Status - The status of the Key Performance Indicator.
	Status *string `json:"status,omitempty"`

	// KpiGroup - The group the Key Performance Indicator belongs to.
	KpiGroup *string `json:"kpiGroup,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Keyperformanceindicator) SetField(field string, fieldValue interface{}) {
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

func (o Keyperformanceindicator) MarshalJSON() ([]byte, error) {
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
	type Alias Keyperformanceindicator
	
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
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		OptimizationType *string `json:"optimizationType,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		KpiType *string `json:"kpiType,omitempty"`
		
		Source *string `json:"source,omitempty"`
		
		WrapUpCodeConfig *Wrapupcodeconfig `json:"wrapUpCodeConfig,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		KpiGroup *string `json:"kpiGroup,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		OptimizationType: o.OptimizationType,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		Description: o.Description,
		
		KpiType: o.KpiType,
		
		Source: o.Source,
		
		WrapUpCodeConfig: o.WrapUpCodeConfig,
		
		Status: o.Status,
		
		KpiGroup: o.KpiGroup,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Keyperformanceindicator) UnmarshalJSON(b []byte) error {
	var KeyperformanceindicatorMap map[string]interface{}
	err := json.Unmarshal(b, &KeyperformanceindicatorMap)
	if err != nil {
		return err
	}
	
	if Id, ok := KeyperformanceindicatorMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := KeyperformanceindicatorMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if OptimizationType, ok := KeyperformanceindicatorMap["optimizationType"].(string); ok {
		o.OptimizationType = &OptimizationType
	}
    
	if dateCreatedString, ok := KeyperformanceindicatorMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := KeyperformanceindicatorMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if Description, ok := KeyperformanceindicatorMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if KpiType, ok := KeyperformanceindicatorMap["kpiType"].(string); ok {
		o.KpiType = &KpiType
	}
    
	if Source, ok := KeyperformanceindicatorMap["source"].(string); ok {
		o.Source = &Source
	}
    
	if WrapUpCodeConfig, ok := KeyperformanceindicatorMap["wrapUpCodeConfig"].(map[string]interface{}); ok {
		WrapUpCodeConfigString, _ := json.Marshal(WrapUpCodeConfig)
		json.Unmarshal(WrapUpCodeConfigString, &o.WrapUpCodeConfig)
	}
	
	if Status, ok := KeyperformanceindicatorMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if KpiGroup, ok := KeyperformanceindicatorMap["kpiGroup"].(string); ok {
		o.KpiGroup = &KpiGroup
	}
    
	if SelfUri, ok := KeyperformanceindicatorMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Keyperformanceindicator) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
