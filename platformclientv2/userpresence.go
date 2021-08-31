package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Userpresence
type Userpresence struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Source - Represents the source where the Presence was set. Some examples are: PURECLOUD, LYNC, OUTLOOK, etc.
	Source *string `json:"source,omitempty"`


	// Primary - A boolean used to tell whether or not to set this presence source as the primary on a PATCH
	Primary *bool `json:"primary,omitempty"`


	// PresenceDefinition
	PresenceDefinition *Presencedefinition `json:"presenceDefinition,omitempty"`


	// Message
	Message *string `json:"message,omitempty"`


	// ModifiedDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ModifiedDate *time.Time `json:"modifiedDate,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Userpresence) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Userpresence
	
	ModifiedDate := new(string)
	if o.ModifiedDate != nil {
		
		*ModifiedDate = timeutil.Strftime(o.ModifiedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ModifiedDate = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Source *string `json:"source,omitempty"`
		
		Primary *bool `json:"primary,omitempty"`
		
		PresenceDefinition *Presencedefinition `json:"presenceDefinition,omitempty"`
		
		Message *string `json:"message,omitempty"`
		
		ModifiedDate *string `json:"modifiedDate,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Source: o.Source,
		
		Primary: o.Primary,
		
		PresenceDefinition: o.PresenceDefinition,
		
		Message: o.Message,
		
		ModifiedDate: ModifiedDate,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Userpresence) UnmarshalJSON(b []byte) error {
	var UserpresenceMap map[string]interface{}
	err := json.Unmarshal(b, &UserpresenceMap)
	if err != nil {
		return err
	}
	
	if Id, ok := UserpresenceMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := UserpresenceMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if Source, ok := UserpresenceMap["source"].(string); ok {
		o.Source = &Source
	}
	
	if Primary, ok := UserpresenceMap["primary"].(bool); ok {
		o.Primary = &Primary
	}
	
	if PresenceDefinition, ok := UserpresenceMap["presenceDefinition"].(map[string]interface{}); ok {
		PresenceDefinitionString, _ := json.Marshal(PresenceDefinition)
		json.Unmarshal(PresenceDefinitionString, &o.PresenceDefinition)
	}
	
	if Message, ok := UserpresenceMap["message"].(string); ok {
		o.Message = &Message
	}
	
	if modifiedDateString, ok := UserpresenceMap["modifiedDate"].(string); ok {
		ModifiedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", modifiedDateString)
		o.ModifiedDate = &ModifiedDate
	}
	
	if SelfUri, ok := UserpresenceMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Userpresence) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
