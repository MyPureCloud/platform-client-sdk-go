package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Messagingintegration
type Messagingintegration struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - A unique Integration Id
	Id *string `json:"id,omitempty"`

	// Name - The name of the Integration
	Name *string `json:"name,omitempty"`

	// SupportedContent - Defines the SupportedContent profile configured for an integration
	SupportedContent *Supportedcontentreference `json:"supportedContent,omitempty"`

	// MessagingSetting
	MessagingSetting *Messagingsettingreference `json:"messagingSetting,omitempty"`

	// Status - The status of the Integration
	Status *string `json:"status,omitempty"`

	// MessengerType - The type of Messaging Integration
	MessengerType *string `json:"messengerType,omitempty"`

	// Recipient - The recipient associated to the Integration. This recipient is used to associate a flow to an integration
	Recipient *Domainentityref `json:"recipient,omitempty"`

	// DateCreated - Date this Integration was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`

	// DateModified - Date this Integration was modified. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`

	// CreatedBy - User reference that created this Integration
	CreatedBy *Domainentityref `json:"createdBy,omitempty"`

	// ModifiedBy - User reference that last modified this Integration
	ModifiedBy *Domainentityref `json:"modifiedBy,omitempty"`

	// Version - Version number required for updates.
	Version *int `json:"version,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Messagingintegration) SetField(field string, fieldValue interface{}) {
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

func (o Messagingintegration) MarshalJSON() ([]byte, error) {
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
	type Alias Messagingintegration
	
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
		
		SupportedContent *Supportedcontentreference `json:"supportedContent,omitempty"`
		
		MessagingSetting *Messagingsettingreference `json:"messagingSetting,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		MessengerType *string `json:"messengerType,omitempty"`
		
		Recipient *Domainentityref `json:"recipient,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		CreatedBy *Domainentityref `json:"createdBy,omitempty"`
		
		ModifiedBy *Domainentityref `json:"modifiedBy,omitempty"`
		
		Version *int `json:"version,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		SupportedContent: o.SupportedContent,
		
		MessagingSetting: o.MessagingSetting,
		
		Status: o.Status,
		
		MessengerType: o.MessengerType,
		
		Recipient: o.Recipient,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		CreatedBy: o.CreatedBy,
		
		ModifiedBy: o.ModifiedBy,
		
		Version: o.Version,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Messagingintegration) UnmarshalJSON(b []byte) error {
	var MessagingintegrationMap map[string]interface{}
	err := json.Unmarshal(b, &MessagingintegrationMap)
	if err != nil {
		return err
	}
	
	if Id, ok := MessagingintegrationMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := MessagingintegrationMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if SupportedContent, ok := MessagingintegrationMap["supportedContent"].(map[string]interface{}); ok {
		SupportedContentString, _ := json.Marshal(SupportedContent)
		json.Unmarshal(SupportedContentString, &o.SupportedContent)
	}
	
	if MessagingSetting, ok := MessagingintegrationMap["messagingSetting"].(map[string]interface{}); ok {
		MessagingSettingString, _ := json.Marshal(MessagingSetting)
		json.Unmarshal(MessagingSettingString, &o.MessagingSetting)
	}
	
	if Status, ok := MessagingintegrationMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if MessengerType, ok := MessagingintegrationMap["messengerType"].(string); ok {
		o.MessengerType = &MessengerType
	}
    
	if Recipient, ok := MessagingintegrationMap["recipient"].(map[string]interface{}); ok {
		RecipientString, _ := json.Marshal(Recipient)
		json.Unmarshal(RecipientString, &o.Recipient)
	}
	
	if dateCreatedString, ok := MessagingintegrationMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := MessagingintegrationMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if CreatedBy, ok := MessagingintegrationMap["createdBy"].(map[string]interface{}); ok {
		CreatedByString, _ := json.Marshal(CreatedBy)
		json.Unmarshal(CreatedByString, &o.CreatedBy)
	}
	
	if ModifiedBy, ok := MessagingintegrationMap["modifiedBy"].(map[string]interface{}); ok {
		ModifiedByString, _ := json.Marshal(ModifiedBy)
		json.Unmarshal(ModifiedByString, &o.ModifiedBy)
	}
	
	if Version, ok := MessagingintegrationMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	
	if SelfUri, ok := MessagingintegrationMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Messagingintegration) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
