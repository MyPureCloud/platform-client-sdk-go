package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// V2mobiusrulestopicentityproperties
type V2mobiusrulestopicentityproperties struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// EntityType
	EntityType *string `json:"entityType,omitempty"`

	// User
	User *V2mobiusrulestopicaddressableentityref `json:"user,omitempty"`

	// Group
	Group *V2mobiusrulestopicaddressableentityref `json:"group,omitempty"`

	// Queue
	Queue *V2mobiusrulestopicaddressableentityref `json:"queue,omitempty"`

	// Team
	Team *V2mobiusrulestopicaddressableentityref `json:"team,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *V2mobiusrulestopicentityproperties) SetField(field string, fieldValue interface{}) {
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

func (o V2mobiusrulestopicentityproperties) MarshalJSON() ([]byte, error) {
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
	type Alias V2mobiusrulestopicentityproperties
	
	return json.Marshal(&struct { 
		EntityType *string `json:"entityType,omitempty"`
		
		User *V2mobiusrulestopicaddressableentityref `json:"user,omitempty"`
		
		Group *V2mobiusrulestopicaddressableentityref `json:"group,omitempty"`
		
		Queue *V2mobiusrulestopicaddressableentityref `json:"queue,omitempty"`
		
		Team *V2mobiusrulestopicaddressableentityref `json:"team,omitempty"`
		Alias
	}{ 
		EntityType: o.EntityType,
		
		User: o.User,
		
		Group: o.Group,
		
		Queue: o.Queue,
		
		Team: o.Team,
		Alias:    (Alias)(o),
	})
}

func (o *V2mobiusrulestopicentityproperties) UnmarshalJSON(b []byte) error {
	var V2mobiusrulestopicentitypropertiesMap map[string]interface{}
	err := json.Unmarshal(b, &V2mobiusrulestopicentitypropertiesMap)
	if err != nil {
		return err
	}
	
	if EntityType, ok := V2mobiusrulestopicentitypropertiesMap["entityType"].(string); ok {
		o.EntityType = &EntityType
	}
    
	if User, ok := V2mobiusrulestopicentitypropertiesMap["user"].(map[string]interface{}); ok {
		UserString, _ := json.Marshal(User)
		json.Unmarshal(UserString, &o.User)
	}
	
	if Group, ok := V2mobiusrulestopicentitypropertiesMap["group"].(map[string]interface{}); ok {
		GroupString, _ := json.Marshal(Group)
		json.Unmarshal(GroupString, &o.Group)
	}
	
	if Queue, ok := V2mobiusrulestopicentitypropertiesMap["queue"].(map[string]interface{}); ok {
		QueueString, _ := json.Marshal(Queue)
		json.Unmarshal(QueueString, &o.Queue)
	}
	
	if Team, ok := V2mobiusrulestopicentitypropertiesMap["team"].(map[string]interface{}); ok {
		TeamString, _ := json.Marshal(Team)
		json.Unmarshal(TeamString, &o.Team)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *V2mobiusrulestopicentityproperties) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
