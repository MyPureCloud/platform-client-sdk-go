package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Messagingsetting - Messaging setting for messaging platform integrations
type Messagingsetting struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The messaging Setting unique identifier associated with this integration
	Id *string `json:"id,omitempty"`

	// Name - The messaging Setting profile name
	Name *string `json:"name,omitempty"`

	// DateCreated - Date this messaging setting was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`

	// DateModified - Date this messaging setting was modified. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`

	// Version - Version number
	Version *string `json:"version,omitempty"`

	// CreatedBy - User reference that created this Setting
	CreatedBy *Domainentityref `json:"createdBy,omitempty"`

	// UpdatedBy - User reference that modified this Setting
	UpdatedBy *Domainentityref `json:"updatedBy,omitempty"`

	// Content - Configuration relating to message contents
	Content *Contentsetting `json:"content,omitempty"`

	// Event - Configuration relating to events which may occur
	Event *Eventsetting `json:"event,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Messagingsetting) SetField(field string, fieldValue interface{}) {
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

func (o Messagingsetting) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateCreated","DateModified", }
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
	type Alias Messagingsetting
	
	DateCreated := new(string)
	if o.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(o.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	DateModified := new(string)
	if o.DateModified != nil {
		
		*DateModified = timeutil.Strftime(o.DateModified, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateModified = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		Version *string `json:"version,omitempty"`
		
		CreatedBy *Domainentityref `json:"createdBy,omitempty"`
		
		UpdatedBy *Domainentityref `json:"updatedBy,omitempty"`
		
		Content *Contentsetting `json:"content,omitempty"`
		
		Event *Eventsetting `json:"event,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		Version: o.Version,
		
		CreatedBy: o.CreatedBy,
		
		UpdatedBy: o.UpdatedBy,
		
		Content: o.Content,
		
		Event: o.Event,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Messagingsetting) UnmarshalJSON(b []byte) error {
	var MessagingsettingMap map[string]interface{}
	err := json.Unmarshal(b, &MessagingsettingMap)
	if err != nil {
		return err
	}
	
	if Id, ok := MessagingsettingMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := MessagingsettingMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if dateCreatedString, ok := MessagingsettingMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := MessagingsettingMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if Version, ok := MessagingsettingMap["version"].(string); ok {
		o.Version = &Version
	}
    
	if CreatedBy, ok := MessagingsettingMap["createdBy"].(map[string]interface{}); ok {
		CreatedByString, _ := json.Marshal(CreatedBy)
		json.Unmarshal(CreatedByString, &o.CreatedBy)
	}
	
	if UpdatedBy, ok := MessagingsettingMap["updatedBy"].(map[string]interface{}); ok {
		UpdatedByString, _ := json.Marshal(UpdatedBy)
		json.Unmarshal(UpdatedByString, &o.UpdatedBy)
	}
	
	if Content, ok := MessagingsettingMap["content"].(map[string]interface{}); ok {
		ContentString, _ := json.Marshal(Content)
		json.Unmarshal(ContentString, &o.Content)
	}
	
	if Event, ok := MessagingsettingMap["event"].(map[string]interface{}); ok {
		EventString, _ := json.Marshal(Event)
		json.Unmarshal(EventString, &o.Event)
	}
	
	if SelfUri, ok := MessagingsettingMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Messagingsetting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
