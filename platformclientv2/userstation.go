package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Userstation
type Userstation struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - A globally unique identifier for this station
	Id *string `json:"id,omitempty"`

	// Name
	Name *string `json:"name,omitempty"`

	// VarType
	VarType *string `json:"type,omitempty"`

	// AssociatedUser
	AssociatedUser *User `json:"associatedUser,omitempty"`

	// AssociatedDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	AssociatedDate *time.Time `json:"associatedDate,omitempty"`

	// DefaultUser
	DefaultUser *User `json:"defaultUser,omitempty"`

	// ProviderInfo - Provider-specific info for this station, e.g. { \"edgeGroupId\": \"ffe7b15c-a9cc-4f4c-88f5-781327819a49\" }
	ProviderInfo *map[string]string `json:"providerInfo,omitempty"`

	// WebRtcCallAppearances - The number of call appearances on the station.
	WebRtcCallAppearances *int `json:"webRtcCallAppearances,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Userstation) SetField(field string, fieldValue interface{}) {
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

func (o Userstation) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "AssociatedDate", }
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
	type Alias Userstation
	
	AssociatedDate := new(string)
	if o.AssociatedDate != nil {
		
		*AssociatedDate = timeutil.Strftime(o.AssociatedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		AssociatedDate = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		AssociatedUser *User `json:"associatedUser,omitempty"`
		
		AssociatedDate *string `json:"associatedDate,omitempty"`
		
		DefaultUser *User `json:"defaultUser,omitempty"`
		
		ProviderInfo *map[string]string `json:"providerInfo,omitempty"`
		
		WebRtcCallAppearances *int `json:"webRtcCallAppearances,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		VarType: o.VarType,
		
		AssociatedUser: o.AssociatedUser,
		
		AssociatedDate: AssociatedDate,
		
		DefaultUser: o.DefaultUser,
		
		ProviderInfo: o.ProviderInfo,
		
		WebRtcCallAppearances: o.WebRtcCallAppearances,
		Alias:    (Alias)(o),
	})
}

func (o *Userstation) UnmarshalJSON(b []byte) error {
	var UserstationMap map[string]interface{}
	err := json.Unmarshal(b, &UserstationMap)
	if err != nil {
		return err
	}
	
	if Id, ok := UserstationMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := UserstationMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if VarType, ok := UserstationMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if AssociatedUser, ok := UserstationMap["associatedUser"].(map[string]interface{}); ok {
		AssociatedUserString, _ := json.Marshal(AssociatedUser)
		json.Unmarshal(AssociatedUserString, &o.AssociatedUser)
	}
	
	if associatedDateString, ok := UserstationMap["associatedDate"].(string); ok {
		AssociatedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", associatedDateString)
		o.AssociatedDate = &AssociatedDate
	}
	
	if DefaultUser, ok := UserstationMap["defaultUser"].(map[string]interface{}); ok {
		DefaultUserString, _ := json.Marshal(DefaultUser)
		json.Unmarshal(DefaultUserString, &o.DefaultUser)
	}
	
	if ProviderInfo, ok := UserstationMap["providerInfo"].(map[string]interface{}); ok {
		ProviderInfoString, _ := json.Marshal(ProviderInfo)
		json.Unmarshal(ProviderInfoString, &o.ProviderInfo)
	}
	
	if WebRtcCallAppearances, ok := UserstationMap["webRtcCallAppearances"].(float64); ok {
		WebRtcCallAppearancesInt := int(WebRtcCallAppearances)
		o.WebRtcCallAppearances = &WebRtcCallAppearancesInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Userstation) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
