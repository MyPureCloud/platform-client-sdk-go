package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Metricdefinition
type Metricdefinition struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name
	Name *string `json:"name,omitempty"`

	// UnitType - The type of associated metric unit
	UnitType *string `json:"unitType,omitempty"`

	// ShortName - An alternate name for this metric definition, often abbreviation
	ShortName *string `json:"shortName,omitempty"`

	// DividendMetrics - Metric names used as dividend
	DividendMetrics *[]string `json:"dividendMetrics,omitempty"`

	// DivisorMetrics - Metric names used as divisor
	DivisorMetrics *[]string `json:"divisorMetrics,omitempty"`

	// DefaultObjective - A predefined default objective for this metric
	DefaultObjective *Defaultobjective `json:"defaultObjective,omitempty"`

	// LockTemplateId - An optional field to specify if this metric definition is locked to certain template. e.g. punctuality
	LockTemplateId *string `json:"lockTemplateId,omitempty"`

	// MediaTypeFilteringAllowed - Flag to indicate if this metricDefinition allows filter based on media types
	MediaTypeFilteringAllowed *bool `json:"mediaTypeFilteringAllowed,omitempty"`

	// QueueFilteringAllowed - Flag to indicate if this metricDefinition allows filter based on queues
	QueueFilteringAllowed *bool `json:"queueFilteringAllowed,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Metricdefinition) SetField(field string, fieldValue interface{}) {
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

func (o Metricdefinition) MarshalJSON() ([]byte, error) {
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
	type Alias Metricdefinition
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		UnitType *string `json:"unitType,omitempty"`
		
		ShortName *string `json:"shortName,omitempty"`
		
		DividendMetrics *[]string `json:"dividendMetrics,omitempty"`
		
		DivisorMetrics *[]string `json:"divisorMetrics,omitempty"`
		
		DefaultObjective *Defaultobjective `json:"defaultObjective,omitempty"`
		
		LockTemplateId *string `json:"lockTemplateId,omitempty"`
		
		MediaTypeFilteringAllowed *bool `json:"mediaTypeFilteringAllowed,omitempty"`
		
		QueueFilteringAllowed *bool `json:"queueFilteringAllowed,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		UnitType: o.UnitType,
		
		ShortName: o.ShortName,
		
		DividendMetrics: o.DividendMetrics,
		
		DivisorMetrics: o.DivisorMetrics,
		
		DefaultObjective: o.DefaultObjective,
		
		LockTemplateId: o.LockTemplateId,
		
		MediaTypeFilteringAllowed: o.MediaTypeFilteringAllowed,
		
		QueueFilteringAllowed: o.QueueFilteringAllowed,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Metricdefinition) UnmarshalJSON(b []byte) error {
	var MetricdefinitionMap map[string]interface{}
	err := json.Unmarshal(b, &MetricdefinitionMap)
	if err != nil {
		return err
	}
	
	if Id, ok := MetricdefinitionMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := MetricdefinitionMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if UnitType, ok := MetricdefinitionMap["unitType"].(string); ok {
		o.UnitType = &UnitType
	}
    
	if ShortName, ok := MetricdefinitionMap["shortName"].(string); ok {
		o.ShortName = &ShortName
	}
    
	if DividendMetrics, ok := MetricdefinitionMap["dividendMetrics"].([]interface{}); ok {
		DividendMetricsString, _ := json.Marshal(DividendMetrics)
		json.Unmarshal(DividendMetricsString, &o.DividendMetrics)
	}
	
	if DivisorMetrics, ok := MetricdefinitionMap["divisorMetrics"].([]interface{}); ok {
		DivisorMetricsString, _ := json.Marshal(DivisorMetrics)
		json.Unmarshal(DivisorMetricsString, &o.DivisorMetrics)
	}
	
	if DefaultObjective, ok := MetricdefinitionMap["defaultObjective"].(map[string]interface{}); ok {
		DefaultObjectiveString, _ := json.Marshal(DefaultObjective)
		json.Unmarshal(DefaultObjectiveString, &o.DefaultObjective)
	}
	
	if LockTemplateId, ok := MetricdefinitionMap["lockTemplateId"].(string); ok {
		o.LockTemplateId = &LockTemplateId
	}
    
	if MediaTypeFilteringAllowed, ok := MetricdefinitionMap["mediaTypeFilteringAllowed"].(bool); ok {
		o.MediaTypeFilteringAllowed = &MediaTypeFilteringAllowed
	}
    
	if QueueFilteringAllowed, ok := MetricdefinitionMap["queueFilteringAllowed"].(bool); ok {
		o.QueueFilteringAllowed = &QueueFilteringAllowed
	}
    
	if SelfUri, ok := MetricdefinitionMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Metricdefinition) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
