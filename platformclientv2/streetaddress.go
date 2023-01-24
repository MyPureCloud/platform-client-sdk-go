package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Streetaddress
type Streetaddress struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Country - 2 Letter Country code, like US or GB
	Country *string `json:"country,omitempty"`

	// A1 - State or Province
	A1 *string `json:"A1,omitempty"`

	// A3 - City or township
	A3 *string `json:"A3,omitempty"`

	// RD - Number and street
	RD *string `json:"RD,omitempty"`

	// HNO - House Number
	HNO *string `json:"HNO,omitempty"`

	// LOC - extra location info like suite 300
	LOC *string `json:"LOC,omitempty"`

	// NAM - Name of the customer
	NAM *string `json:"NAM,omitempty"`

	// PC - Postal code
	PC *string `json:"PC,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Streetaddress) SetField(field string, fieldValue interface{}) {
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

func (o Streetaddress) MarshalJSON() ([]byte, error) {
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
	type Alias Streetaddress
	
	return json.Marshal(&struct { 
		Country *string `json:"country,omitempty"`
		
		A1 *string `json:"A1,omitempty"`
		
		A3 *string `json:"A3,omitempty"`
		
		RD *string `json:"RD,omitempty"`
		
		HNO *string `json:"HNO,omitempty"`
		
		LOC *string `json:"LOC,omitempty"`
		
		NAM *string `json:"NAM,omitempty"`
		
		PC *string `json:"PC,omitempty"`
		Alias
	}{ 
		Country: o.Country,
		
		A1: o.A1,
		
		A3: o.A3,
		
		RD: o.RD,
		
		HNO: o.HNO,
		
		LOC: o.LOC,
		
		NAM: o.NAM,
		
		PC: o.PC,
		Alias:    (Alias)(o),
	})
}

func (o *Streetaddress) UnmarshalJSON(b []byte) error {
	var StreetaddressMap map[string]interface{}
	err := json.Unmarshal(b, &StreetaddressMap)
	if err != nil {
		return err
	}
	
	if Country, ok := StreetaddressMap["country"].(string); ok {
		o.Country = &Country
	}
    
	if A1, ok := StreetaddressMap["A1"].(string); ok {
		o.A1 = &A1
	}
    
	if A3, ok := StreetaddressMap["A3"].(string); ok {
		o.A3 = &A3
	}
    
	if RD, ok := StreetaddressMap["RD"].(string); ok {
		o.RD = &RD
	}
    
	if HNO, ok := StreetaddressMap["HNO"].(string); ok {
		o.HNO = &HNO
	}
    
	if LOC, ok := StreetaddressMap["LOC"].(string); ok {
		o.LOC = &LOC
	}
    
	if NAM, ok := StreetaddressMap["NAM"].(string); ok {
		o.NAM = &NAM
	}
    
	if PC, ok := StreetaddressMap["PC"].(string); ok {
		o.PC = &PC
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Streetaddress) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
