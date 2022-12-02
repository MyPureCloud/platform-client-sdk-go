package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Dialerrulesetconfigchangeruleset
type Dialerrulesetconfigchangeruleset struct { 
	// ContactList
	ContactList *Dialerrulesetconfigchangeurireference `json:"contactList,omitempty"`


	// Queue - A UriReference for a resource
	Queue *Dialerrulesetconfigchangeurireference `json:"queue,omitempty"`


	// Rules
	Rules *[]Dialerrulesetconfigchangerule `json:"rules,omitempty"`


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

func (o *Dialerrulesetconfigchangeruleset) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dialerrulesetconfigchangeruleset
	
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
		ContactList *Dialerrulesetconfigchangeurireference `json:"contactList,omitempty"`
		
		Queue *Dialerrulesetconfigchangeurireference `json:"queue,omitempty"`
		
		Rules *[]Dialerrulesetconfigchangerule `json:"rules,omitempty"`
		
		AdditionalProperties *map[string]interface{} `json:"additionalProperties,omitempty"`
		
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		Version *int `json:"version,omitempty"`
		*Alias
	}{ 
		ContactList: o.ContactList,
		
		Queue: o.Queue,
		
		Rules: o.Rules,
		
		AdditionalProperties: o.AdditionalProperties,
		
		Id: o.Id,
		
		Name: o.Name,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		Version: o.Version,
		Alias:    (*Alias)(o),
	})
}

func (o *Dialerrulesetconfigchangeruleset) UnmarshalJSON(b []byte) error {
	var DialerrulesetconfigchangerulesetMap map[string]interface{}
	err := json.Unmarshal(b, &DialerrulesetconfigchangerulesetMap)
	if err != nil {
		return err
	}
	
	if ContactList, ok := DialerrulesetconfigchangerulesetMap["contactList"].(map[string]interface{}); ok {
		ContactListString, _ := json.Marshal(ContactList)
		json.Unmarshal(ContactListString, &o.ContactList)
	}
	
	if Queue, ok := DialerrulesetconfigchangerulesetMap["queue"].(map[string]interface{}); ok {
		QueueString, _ := json.Marshal(Queue)
		json.Unmarshal(QueueString, &o.Queue)
	}
	
	if Rules, ok := DialerrulesetconfigchangerulesetMap["rules"].([]interface{}); ok {
		RulesString, _ := json.Marshal(Rules)
		json.Unmarshal(RulesString, &o.Rules)
	}
	
	if AdditionalProperties, ok := DialerrulesetconfigchangerulesetMap["additionalProperties"].(map[string]interface{}); ok {
		AdditionalPropertiesString, _ := json.Marshal(AdditionalProperties)
		json.Unmarshal(AdditionalPropertiesString, &o.AdditionalProperties)
	}
	
	if Id, ok := DialerrulesetconfigchangerulesetMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := DialerrulesetconfigchangerulesetMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if dateCreatedString, ok := DialerrulesetconfigchangerulesetMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := DialerrulesetconfigchangerulesetMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if Version, ok := DialerrulesetconfigchangerulesetMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Dialerrulesetconfigchangeruleset) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
