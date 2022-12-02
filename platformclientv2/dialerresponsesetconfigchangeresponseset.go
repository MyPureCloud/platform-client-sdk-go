package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Dialerresponsesetconfigchangeresponseset
type Dialerresponsesetconfigchangeresponseset struct { 
	// Responses - Map of disposition identifiers to reactions. For example: {\"disposition.classification.callable.person\": {\"reactionType\": \"transfer\"}}
	Responses *map[string]Dialerresponsesetconfigchangereaction `json:"responses,omitempty"`


	// BeepDetectionEnabled - When beep detection is enabled, answering machine detection will wait for the beep before transferring the call
	BeepDetectionEnabled *bool `json:"beepDetectionEnabled,omitempty"`


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

}

func (o *Dialerresponsesetconfigchangeresponseset) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dialerresponsesetconfigchangeresponseset
	
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
		Responses *map[string]Dialerresponsesetconfigchangereaction `json:"responses,omitempty"`
		
		BeepDetectionEnabled *bool `json:"beepDetectionEnabled,omitempty"`
		
		AdditionalProperties *map[string]interface{} `json:"additionalProperties,omitempty"`
		
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		Version *int `json:"version,omitempty"`
		*Alias
	}{ 
		Responses: o.Responses,
		
		BeepDetectionEnabled: o.BeepDetectionEnabled,
		
		AdditionalProperties: o.AdditionalProperties,
		
		Id: o.Id,
		
		Name: o.Name,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		Version: o.Version,
		Alias:    (*Alias)(o),
	})
}

func (o *Dialerresponsesetconfigchangeresponseset) UnmarshalJSON(b []byte) error {
	var DialerresponsesetconfigchangeresponsesetMap map[string]interface{}
	err := json.Unmarshal(b, &DialerresponsesetconfigchangeresponsesetMap)
	if err != nil {
		return err
	}
	
	if Responses, ok := DialerresponsesetconfigchangeresponsesetMap["responses"].(map[string]interface{}); ok {
		ResponsesString, _ := json.Marshal(Responses)
		json.Unmarshal(ResponsesString, &o.Responses)
	}
	
	if BeepDetectionEnabled, ok := DialerresponsesetconfigchangeresponsesetMap["beepDetectionEnabled"].(bool); ok {
		o.BeepDetectionEnabled = &BeepDetectionEnabled
	}
    
	if AdditionalProperties, ok := DialerresponsesetconfigchangeresponsesetMap["additionalProperties"].(map[string]interface{}); ok {
		AdditionalPropertiesString, _ := json.Marshal(AdditionalProperties)
		json.Unmarshal(AdditionalPropertiesString, &o.AdditionalProperties)
	}
	
	if Id, ok := DialerresponsesetconfigchangeresponsesetMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := DialerresponsesetconfigchangeresponsesetMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if dateCreatedString, ok := DialerresponsesetconfigchangeresponsesetMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := DialerresponsesetconfigchangeresponsesetMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if Version, ok := DialerresponsesetconfigchangeresponsesetMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Dialerresponsesetconfigchangeresponseset) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
