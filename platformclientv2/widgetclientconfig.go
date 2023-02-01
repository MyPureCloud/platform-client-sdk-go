package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Widgetclientconfig
type Widgetclientconfig struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// V1
	V1 *Widgetclientconfigv1 `json:"v1,omitempty"`

	// V2
	V2 *interface{} `json:"v2,omitempty"`

	// V1Http
	V1Http *Widgetclientconfigv1http `json:"v1-http,omitempty"`

	// ThirdParty
	ThirdParty *interface{} `json:"third-party,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Widgetclientconfig) SetField(field string, fieldValue interface{}) {
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

func (o Widgetclientconfig) MarshalJSON() ([]byte, error) {
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
	type Alias Widgetclientconfig
	
	return json.Marshal(&struct { 
		V1 *Widgetclientconfigv1 `json:"v1,omitempty"`
		
		V2 *interface{} `json:"v2,omitempty"`
		
		V1Http *Widgetclientconfigv1http `json:"v1-http,omitempty"`
		
		ThirdParty *interface{} `json:"third-party,omitempty"`
		Alias
	}{ 
		V1: o.V1,
		
		V2: o.V2,
		
		V1Http: o.V1Http,
		
		ThirdParty: o.ThirdParty,
		Alias:    (Alias)(o),
	})
}

func (o *Widgetclientconfig) UnmarshalJSON(b []byte) error {
	var WidgetclientconfigMap map[string]interface{}
	err := json.Unmarshal(b, &WidgetclientconfigMap)
	if err != nil {
		return err
	}
	
	if V1, ok := WidgetclientconfigMap["v1"].(map[string]interface{}); ok {
		V1String, _ := json.Marshal(V1)
		json.Unmarshal(V1String, &o.V1)
	}
	
	if V2, ok := WidgetclientconfigMap["v2"].(map[string]interface{}); ok {
		V2String, _ := json.Marshal(V2)
		json.Unmarshal(V2String, &o.V2)
	}
	
	if V1Http, ok := WidgetclientconfigMap["v1-http"].(map[string]interface{}); ok {
		V1HttpString, _ := json.Marshal(V1Http)
		json.Unmarshal(V1HttpString, &o.V1Http)
	}
	
	if ThirdParty, ok := WidgetclientconfigMap["third-party"].(map[string]interface{}); ok {
		ThirdPartyString, _ := json.Marshal(ThirdParty)
		json.Unmarshal(ThirdPartyString, &o.ThirdParty)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Widgetclientconfig) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
