package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Userstations
type Userstations struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// AssociatedStation - Current associated station for this user.
	AssociatedStation *Userstation `json:"associatedStation,omitempty"`

	// EffectiveStation - The station where the user can be reached based on their default and associated station.
	EffectiveStation *Userstation `json:"effectiveStation,omitempty"`

	// DefaultStation - Default station to be used if not associated with a station.
	DefaultStation *Userstation `json:"defaultStation,omitempty"`

	// LastAssociatedStation - Last associated station for this user.
	LastAssociatedStation *Userstation `json:"lastAssociatedStation,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Userstations) SetField(field string, fieldValue interface{}) {
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

func (o Userstations) MarshalJSON() ([]byte, error) {
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
	type Alias Userstations
	
	return json.Marshal(&struct { 
		AssociatedStation *Userstation `json:"associatedStation,omitempty"`
		
		EffectiveStation *Userstation `json:"effectiveStation,omitempty"`
		
		DefaultStation *Userstation `json:"defaultStation,omitempty"`
		
		LastAssociatedStation *Userstation `json:"lastAssociatedStation,omitempty"`
		Alias
	}{ 
		AssociatedStation: o.AssociatedStation,
		
		EffectiveStation: o.EffectiveStation,
		
		DefaultStation: o.DefaultStation,
		
		LastAssociatedStation: o.LastAssociatedStation,
		Alias:    (Alias)(o),
	})
}

func (o *Userstations) UnmarshalJSON(b []byte) error {
	var UserstationsMap map[string]interface{}
	err := json.Unmarshal(b, &UserstationsMap)
	if err != nil {
		return err
	}
	
	if AssociatedStation, ok := UserstationsMap["associatedStation"].(map[string]interface{}); ok {
		AssociatedStationString, _ := json.Marshal(AssociatedStation)
		json.Unmarshal(AssociatedStationString, &o.AssociatedStation)
	}
	
	if EffectiveStation, ok := UserstationsMap["effectiveStation"].(map[string]interface{}); ok {
		EffectiveStationString, _ := json.Marshal(EffectiveStation)
		json.Unmarshal(EffectiveStationString, &o.EffectiveStation)
	}
	
	if DefaultStation, ok := UserstationsMap["defaultStation"].(map[string]interface{}); ok {
		DefaultStationString, _ := json.Marshal(DefaultStation)
		json.Unmarshal(DefaultStationString, &o.DefaultStation)
	}
	
	if LastAssociatedStation, ok := UserstationsMap["lastAssociatedStation"].(map[string]interface{}); ok {
		LastAssociatedStationString, _ := json.Marshal(LastAssociatedStation)
		json.Unmarshal(LastAssociatedStationString, &o.LastAssociatedStation)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Userstations) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
