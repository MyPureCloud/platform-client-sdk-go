package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Journeyoutcomeeventsnotificationdevice
type Journeyoutcomeeventsnotificationdevice struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// VarType
	VarType *string `json:"type,omitempty"`

	// IsMobile
	IsMobile *bool `json:"isMobile,omitempty"`

	// ScreenHeight
	ScreenHeight *int `json:"screenHeight,omitempty"`

	// ScreenWidth
	ScreenWidth *int `json:"screenWidth,omitempty"`

	// Fingerprint
	Fingerprint *string `json:"fingerprint,omitempty"`

	// OsFamily
	OsFamily *string `json:"osFamily,omitempty"`

	// OsVersion
	OsVersion *string `json:"osVersion,omitempty"`

	// Category
	Category *string `json:"category,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Journeyoutcomeeventsnotificationdevice) SetField(field string, fieldValue interface{}) {
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

func (o Journeyoutcomeeventsnotificationdevice) MarshalJSON() ([]byte, error) {
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
	type Alias Journeyoutcomeeventsnotificationdevice
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		IsMobile *bool `json:"isMobile,omitempty"`
		
		ScreenHeight *int `json:"screenHeight,omitempty"`
		
		ScreenWidth *int `json:"screenWidth,omitempty"`
		
		Fingerprint *string `json:"fingerprint,omitempty"`
		
		OsFamily *string `json:"osFamily,omitempty"`
		
		OsVersion *string `json:"osVersion,omitempty"`
		
		Category *string `json:"category,omitempty"`
		Alias
	}{ 
		VarType: o.VarType,
		
		IsMobile: o.IsMobile,
		
		ScreenHeight: o.ScreenHeight,
		
		ScreenWidth: o.ScreenWidth,
		
		Fingerprint: o.Fingerprint,
		
		OsFamily: o.OsFamily,
		
		OsVersion: o.OsVersion,
		
		Category: o.Category,
		Alias:    (Alias)(o),
	})
}

func (o *Journeyoutcomeeventsnotificationdevice) UnmarshalJSON(b []byte) error {
	var JourneyoutcomeeventsnotificationdeviceMap map[string]interface{}
	err := json.Unmarshal(b, &JourneyoutcomeeventsnotificationdeviceMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := JourneyoutcomeeventsnotificationdeviceMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if IsMobile, ok := JourneyoutcomeeventsnotificationdeviceMap["isMobile"].(bool); ok {
		o.IsMobile = &IsMobile
	}
    
	if ScreenHeight, ok := JourneyoutcomeeventsnotificationdeviceMap["screenHeight"].(float64); ok {
		ScreenHeightInt := int(ScreenHeight)
		o.ScreenHeight = &ScreenHeightInt
	}
	
	if ScreenWidth, ok := JourneyoutcomeeventsnotificationdeviceMap["screenWidth"].(float64); ok {
		ScreenWidthInt := int(ScreenWidth)
		o.ScreenWidth = &ScreenWidthInt
	}
	
	if Fingerprint, ok := JourneyoutcomeeventsnotificationdeviceMap["fingerprint"].(string); ok {
		o.Fingerprint = &Fingerprint
	}
    
	if OsFamily, ok := JourneyoutcomeeventsnotificationdeviceMap["osFamily"].(string); ok {
		o.OsFamily = &OsFamily
	}
    
	if OsVersion, ok := JourneyoutcomeeventsnotificationdeviceMap["osVersion"].(string); ok {
		o.OsVersion = &OsVersion
	}
    
	if Category, ok := JourneyoutcomeeventsnotificationdeviceMap["category"].(string); ok {
		o.Category = &Category
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Journeyoutcomeeventsnotificationdevice) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
