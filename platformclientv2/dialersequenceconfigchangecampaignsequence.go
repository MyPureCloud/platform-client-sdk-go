package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Dialersequenceconfigchangecampaignsequence
type Dialersequenceconfigchangecampaignsequence struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Campaigns - the ordered list of campaign identifiers
	Campaigns *[]Dialersequenceconfigchangeurireference `json:"campaigns,omitempty"`

	// CurrentCampaign - the zero-based index of the current campaign in the campaigns list
	CurrentCampaign *int `json:"currentCampaign,omitempty"`

	// Status
	Status *string `json:"status,omitempty"`

	// StopMessage - if a sequence has unexpectedly stopped, this message provides the reason
	StopMessage *string `json:"stopMessage,omitempty"`

	// Repeat - indicates if a sequence is to repeat from the beginning after the last campaign completes; default is false
	Repeat *bool `json:"repeat,omitempty"`

	// AdditionalProperties
	AdditionalProperties *map[string]interface{} `json:"additionalProperties,omitempty"`

	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name - The UI-visible name of the object
	Name *string `json:"name,omitempty"`

	// DateCreated - Creation time of the entity
	DateCreated *time.Time `json:"dateCreated,omitempty"`

	// DateModified - Last modified time of the entity
	DateModified *time.Time `json:"dateModified,omitempty"`

	// Version - Required for updates, must match the version number of the most recent update
	Version *int `json:"version,omitempty"`

	// GetAdditionalProperties
	GetAdditionalProperties *map[string]interface{} `json:"getAdditionalProperties,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Dialersequenceconfigchangecampaignsequence) SetField(field string, fieldValue interface{}) {
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

func (o Dialersequenceconfigchangecampaignsequence) MarshalJSON() ([]byte, error) {
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
	type Alias Dialersequenceconfigchangecampaignsequence
	
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
		Campaigns *[]Dialersequenceconfigchangeurireference `json:"campaigns,omitempty"`
		
		CurrentCampaign *int `json:"currentCampaign,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		StopMessage *string `json:"stopMessage,omitempty"`
		
		Repeat *bool `json:"repeat,omitempty"`
		
		AdditionalProperties *map[string]interface{} `json:"additionalProperties,omitempty"`
		
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		Version *int `json:"version,omitempty"`
		
		GetAdditionalProperties *map[string]interface{} `json:"getAdditionalProperties,omitempty"`
		Alias
	}{ 
		Campaigns: o.Campaigns,
		
		CurrentCampaign: o.CurrentCampaign,
		
		Status: o.Status,
		
		StopMessage: o.StopMessage,
		
		Repeat: o.Repeat,
		
		AdditionalProperties: o.AdditionalProperties,
		
		Id: o.Id,
		
		Name: o.Name,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		Version: o.Version,
		
		GetAdditionalProperties: o.GetAdditionalProperties,
		Alias:    (Alias)(o),
	})
}

func (o *Dialersequenceconfigchangecampaignsequence) UnmarshalJSON(b []byte) error {
	var DialersequenceconfigchangecampaignsequenceMap map[string]interface{}
	err := json.Unmarshal(b, &DialersequenceconfigchangecampaignsequenceMap)
	if err != nil {
		return err
	}
	
	if Campaigns, ok := DialersequenceconfigchangecampaignsequenceMap["campaigns"].([]interface{}); ok {
		CampaignsString, _ := json.Marshal(Campaigns)
		json.Unmarshal(CampaignsString, &o.Campaigns)
	}
	
	if CurrentCampaign, ok := DialersequenceconfigchangecampaignsequenceMap["currentCampaign"].(float64); ok {
		CurrentCampaignInt := int(CurrentCampaign)
		o.CurrentCampaign = &CurrentCampaignInt
	}
	
	if Status, ok := DialersequenceconfigchangecampaignsequenceMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if StopMessage, ok := DialersequenceconfigchangecampaignsequenceMap["stopMessage"].(string); ok {
		o.StopMessage = &StopMessage
	}
    
	if Repeat, ok := DialersequenceconfigchangecampaignsequenceMap["repeat"].(bool); ok {
		o.Repeat = &Repeat
	}
    
	if AdditionalProperties, ok := DialersequenceconfigchangecampaignsequenceMap["additionalProperties"].(map[string]interface{}); ok {
		AdditionalPropertiesString, _ := json.Marshal(AdditionalProperties)
		json.Unmarshal(AdditionalPropertiesString, &o.AdditionalProperties)
	}
	
	if Id, ok := DialersequenceconfigchangecampaignsequenceMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := DialersequenceconfigchangecampaignsequenceMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if dateCreatedString, ok := DialersequenceconfigchangecampaignsequenceMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := DialersequenceconfigchangecampaignsequenceMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if Version, ok := DialersequenceconfigchangecampaignsequenceMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	
	if GetAdditionalProperties, ok := DialersequenceconfigchangecampaignsequenceMap["getAdditionalProperties"].(map[string]interface{}); ok {
		GetAdditionalPropertiesString, _ := json.Marshal(GetAdditionalProperties)
		json.Unmarshal(GetAdditionalPropertiesString, &o.GetAdditionalProperties)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Dialersequenceconfigchangecampaignsequence) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
