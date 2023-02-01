package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Vmpairinginfo
type Vmpairinginfo struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// MetaData - This is to be used to complete the setup process of a locally deployed virtual edge device.
	MetaData *Metadata `json:"meta-data,omitempty"`

	// EdgeId
	EdgeId *string `json:"edge-id,omitempty"`

	// AuthToken
	AuthToken *string `json:"auth-token,omitempty"`

	// OrgId
	OrgId *string `json:"org-id,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Vmpairinginfo) SetField(field string, fieldValue interface{}) {
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

func (o Vmpairinginfo) MarshalJSON() ([]byte, error) {
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
	type Alias Vmpairinginfo
	
	return json.Marshal(&struct { 
		MetaData *Metadata `json:"meta-data,omitempty"`
		
		EdgeId *string `json:"edge-id,omitempty"`
		
		AuthToken *string `json:"auth-token,omitempty"`
		
		OrgId *string `json:"org-id,omitempty"`
		Alias
	}{ 
		MetaData: o.MetaData,
		
		EdgeId: o.EdgeId,
		
		AuthToken: o.AuthToken,
		
		OrgId: o.OrgId,
		Alias:    (Alias)(o),
	})
}

func (o *Vmpairinginfo) UnmarshalJSON(b []byte) error {
	var VmpairinginfoMap map[string]interface{}
	err := json.Unmarshal(b, &VmpairinginfoMap)
	if err != nil {
		return err
	}
	
	if MetaData, ok := VmpairinginfoMap["meta-data"].(map[string]interface{}); ok {
		MetaDataString, _ := json.Marshal(MetaData)
		json.Unmarshal(MetaDataString, &o.MetaData)
	}
	
	if EdgeId, ok := VmpairinginfoMap["edge-id"].(string); ok {
		o.EdgeId = &EdgeId
	}
    
	if AuthToken, ok := VmpairinginfoMap["auth-token"].(string); ok {
		o.AuthToken = &AuthToken
	}
    
	if OrgId, ok := VmpairinginfoMap["org-id"].(string); ok {
		o.OrgId = &OrgId
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Vmpairinginfo) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
