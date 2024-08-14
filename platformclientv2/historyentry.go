package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Historyentry
type Historyentry struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Action - The action performed
	Action *string `json:"action,omitempty"`

	// Resource - For actions performed not on the item itself, but on a sub-item, this field identifies the sub-item by name.  For example, for actions performed on prompt resources, this will be the prompt resource name.
	Resource *string `json:"resource,omitempty"`

	// Timestamp - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	Timestamp *time.Time `json:"timestamp,omitempty"`

	// User - User associated with this entry.
	User *User `json:"user,omitempty"`

	// Client - OAuth client associated with this entry.
	Client *Domainentityref `json:"client,omitempty"`

	// Version
	Version *string `json:"version,omitempty"`

	// Secure
	Secure *bool `json:"secure,omitempty"`

	// VirtualAgentEnabled
	VirtualAgentEnabled *bool `json:"virtualAgentEnabled,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Historyentry) SetField(field string, fieldValue interface{}) {
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

func (o Historyentry) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "Timestamp", }
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
	type Alias Historyentry
	
	Timestamp := new(string)
	if o.Timestamp != nil {
		
		*Timestamp = timeutil.Strftime(o.Timestamp, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		Timestamp = nil
	}
	
	return json.Marshal(&struct { 
		Action *string `json:"action,omitempty"`
		
		Resource *string `json:"resource,omitempty"`
		
		Timestamp *string `json:"timestamp,omitempty"`
		
		User *User `json:"user,omitempty"`
		
		Client *Domainentityref `json:"client,omitempty"`
		
		Version *string `json:"version,omitempty"`
		
		Secure *bool `json:"secure,omitempty"`
		
		VirtualAgentEnabled *bool `json:"virtualAgentEnabled,omitempty"`
		Alias
	}{ 
		Action: o.Action,
		
		Resource: o.Resource,
		
		Timestamp: Timestamp,
		
		User: o.User,
		
		Client: o.Client,
		
		Version: o.Version,
		
		Secure: o.Secure,
		
		VirtualAgentEnabled: o.VirtualAgentEnabled,
		Alias:    (Alias)(o),
	})
}

func (o *Historyentry) UnmarshalJSON(b []byte) error {
	var HistoryentryMap map[string]interface{}
	err := json.Unmarshal(b, &HistoryentryMap)
	if err != nil {
		return err
	}
	
	if Action, ok := HistoryentryMap["action"].(string); ok {
		o.Action = &Action
	}
    
	if Resource, ok := HistoryentryMap["resource"].(string); ok {
		o.Resource = &Resource
	}
    
	if timestampString, ok := HistoryentryMap["timestamp"].(string); ok {
		Timestamp, _ := time.Parse("2006-01-02T15:04:05.999999Z", timestampString)
		o.Timestamp = &Timestamp
	}
	
	if User, ok := HistoryentryMap["user"].(map[string]interface{}); ok {
		UserString, _ := json.Marshal(User)
		json.Unmarshal(UserString, &o.User)
	}
	
	if Client, ok := HistoryentryMap["client"].(map[string]interface{}); ok {
		ClientString, _ := json.Marshal(Client)
		json.Unmarshal(ClientString, &o.Client)
	}
	
	if Version, ok := HistoryentryMap["version"].(string); ok {
		o.Version = &Version
	}
    
	if Secure, ok := HistoryentryMap["secure"].(bool); ok {
		o.Secure = &Secure
	}
    
	if VirtualAgentEnabled, ok := HistoryentryMap["virtualAgentEnabled"].(bool); ok {
		o.VirtualAgentEnabled = &VirtualAgentEnabled
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Historyentry) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
