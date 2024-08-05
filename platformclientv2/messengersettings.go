package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Messengersettings - Settings concerning messenger
type Messengersettings struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Enabled - Whether or not messenger is enabled
	Enabled *bool `json:"enabled,omitempty"`

	// Styles - The style settings for messenger
	Styles *Messengerstyles `json:"styles,omitempty"`

	// LauncherButton - The launcher button settings for messenger
	LauncherButton *Launcherbuttonsettings `json:"launcherButton,omitempty"`

	// FileUpload - The file upload settings for messenger
	FileUpload *Fileuploadsettings `json:"fileUpload,omitempty"`

	// Apps - The apps embedded in the messenger
	Apps *Messengerapps `json:"apps,omitempty"`

	// HomeScreen - The homescreen settings for messenger
	HomeScreen *Messengerhomescreen `json:"homeScreen,omitempty"`

	// SessionPersistenceType - The session persistence type for messenger
	SessionPersistenceType *string `json:"sessionPersistenceType,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Messengersettings) SetField(field string, fieldValue interface{}) {
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

func (o Messengersettings) MarshalJSON() ([]byte, error) {
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
	type Alias Messengersettings
	
	return json.Marshal(&struct { 
		Enabled *bool `json:"enabled,omitempty"`
		
		Styles *Messengerstyles `json:"styles,omitempty"`
		
		LauncherButton *Launcherbuttonsettings `json:"launcherButton,omitempty"`
		
		FileUpload *Fileuploadsettings `json:"fileUpload,omitempty"`
		
		Apps *Messengerapps `json:"apps,omitempty"`
		
		HomeScreen *Messengerhomescreen `json:"homeScreen,omitempty"`
		
		SessionPersistenceType *string `json:"sessionPersistenceType,omitempty"`
		Alias
	}{ 
		Enabled: o.Enabled,
		
		Styles: o.Styles,
		
		LauncherButton: o.LauncherButton,
		
		FileUpload: o.FileUpload,
		
		Apps: o.Apps,
		
		HomeScreen: o.HomeScreen,
		
		SessionPersistenceType: o.SessionPersistenceType,
		Alias:    (Alias)(o),
	})
}

func (o *Messengersettings) UnmarshalJSON(b []byte) error {
	var MessengersettingsMap map[string]interface{}
	err := json.Unmarshal(b, &MessengersettingsMap)
	if err != nil {
		return err
	}
	
	if Enabled, ok := MessengersettingsMap["enabled"].(bool); ok {
		o.Enabled = &Enabled
	}
    
	if Styles, ok := MessengersettingsMap["styles"].(map[string]interface{}); ok {
		StylesString, _ := json.Marshal(Styles)
		json.Unmarshal(StylesString, &o.Styles)
	}
	
	if LauncherButton, ok := MessengersettingsMap["launcherButton"].(map[string]interface{}); ok {
		LauncherButtonString, _ := json.Marshal(LauncherButton)
		json.Unmarshal(LauncherButtonString, &o.LauncherButton)
	}
	
	if FileUpload, ok := MessengersettingsMap["fileUpload"].(map[string]interface{}); ok {
		FileUploadString, _ := json.Marshal(FileUpload)
		json.Unmarshal(FileUploadString, &o.FileUpload)
	}
	
	if Apps, ok := MessengersettingsMap["apps"].(map[string]interface{}); ok {
		AppsString, _ := json.Marshal(Apps)
		json.Unmarshal(AppsString, &o.Apps)
	}
	
	if HomeScreen, ok := MessengersettingsMap["homeScreen"].(map[string]interface{}); ok {
		HomeScreenString, _ := json.Marshal(HomeScreen)
		json.Unmarshal(HomeScreenString, &o.HomeScreen)
	}
	
	if SessionPersistenceType, ok := MessengersettingsMap["sessionPersistenceType"].(string); ok {
		o.SessionPersistenceType = &SessionPersistenceType
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Messengersettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
