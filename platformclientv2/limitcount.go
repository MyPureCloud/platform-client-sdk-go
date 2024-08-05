package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Limitcount
type Limitcount struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Name - The name of the limit.
	Name *string `json:"name,omitempty"`

	// EstimatedCount - The total used count of the limit.
	EstimatedCount *int `json:"estimatedCount,omitempty"`

	// Max - The maximum value of the limit.
	Max *int `json:"max,omitempty"`

	// EntityId - The entity which makes this count unique. The context of what the entity is would be dependant on the limit. May not be applicable for all limits.
	EntityId *string `json:"entityId,omitempty"`

	// UserId - The user which makes this count unique. May not be applicable for all limits.
	UserId *string `json:"userId,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Limitcount) SetField(field string, fieldValue interface{}) {
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

func (o Limitcount) MarshalJSON() ([]byte, error) {
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
	type Alias Limitcount
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		EstimatedCount *int `json:"estimatedCount,omitempty"`
		
		Max *int `json:"max,omitempty"`
		
		EntityId *string `json:"entityId,omitempty"`
		
		UserId *string `json:"userId,omitempty"`
		Alias
	}{ 
		Name: o.Name,
		
		EstimatedCount: o.EstimatedCount,
		
		Max: o.Max,
		
		EntityId: o.EntityId,
		
		UserId: o.UserId,
		Alias:    (Alias)(o),
	})
}

func (o *Limitcount) UnmarshalJSON(b []byte) error {
	var LimitcountMap map[string]interface{}
	err := json.Unmarshal(b, &LimitcountMap)
	if err != nil {
		return err
	}
	
	if Name, ok := LimitcountMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if EstimatedCount, ok := LimitcountMap["estimatedCount"].(float64); ok {
		EstimatedCountInt := int(EstimatedCount)
		o.EstimatedCount = &EstimatedCountInt
	}
	
	if Max, ok := LimitcountMap["max"].(float64); ok {
		MaxInt := int(Max)
		o.Max = &MaxInt
	}
	
	if EntityId, ok := LimitcountMap["entityId"].(string); ok {
		o.EntityId = &EntityId
	}
    
	if UserId, ok := LimitcountMap["userId"].(string); ok {
		o.UserId = &UserId
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Limitcount) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
