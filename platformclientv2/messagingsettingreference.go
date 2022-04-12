package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Messagingsettingreference - Messaging Setting for messaging platform integrations
type Messagingsettingreference struct { 
	// Id - The messaging Setting unique identifier associated with this integration
	Id *string `json:"id,omitempty"`


	// Name - The messaging Setting profile name
	Name *string `json:"name,omitempty"`


	// SelfUri - The messaging Setting profile URI
	SelfUri *string `json:"selfUri,omitempty"`


	// DateCreated - Date this messaging Setting was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateModified - Date this messaging Setting was modified. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`


	// Version - Version number
	Version *string `json:"version,omitempty"`


	// CreatedBy - User reference that created this Setting
	CreatedBy *Domainentityref `json:"createdBy,omitempty"`


	// UpdatedBy - User reference that modified this Setting
	UpdatedBy *Domainentityref `json:"updatedBy,omitempty"`


	// Content - Settings relating to message contents
	Content *Contentsetting `json:"content,omitempty"`


	// Event - Settings relating to events which may occur
	Event *Eventsetting `json:"event,omitempty"`

}

func (o *Messagingsettingreference) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Messagingsettingreference
	
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
		
		SelfUri *string `json:"selfUri,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		Version *string `json:"version,omitempty"`
		
		CreatedBy *Domainentityref `json:"createdBy,omitempty"`
		
		UpdatedBy *Domainentityref `json:"updatedBy,omitempty"`
		
		Content *Contentsetting `json:"content,omitempty"`
		
		Event *Eventsetting `json:"event,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		SelfUri: o.SelfUri,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		Version: o.Version,
		
		CreatedBy: o.CreatedBy,
		
		UpdatedBy: o.UpdatedBy,
		
		Content: o.Content,
		
		Event: o.Event,
		Alias:    (*Alias)(o),
	})
}

func (o *Messagingsettingreference) UnmarshalJSON(b []byte) error {
	var MessagingsettingreferenceMap map[string]interface{}
	err := json.Unmarshal(b, &MessagingsettingreferenceMap)
	if err != nil {
		return err
	}
	
	if Id, ok := MessagingsettingreferenceMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := MessagingsettingreferenceMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if SelfUri, ok := MessagingsettingreferenceMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	
	if dateCreatedString, ok := MessagingsettingreferenceMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := MessagingsettingreferenceMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if Version, ok := MessagingsettingreferenceMap["version"].(string); ok {
		o.Version = &Version
	}
	
	if CreatedBy, ok := MessagingsettingreferenceMap["createdBy"].(map[string]interface{}); ok {
		CreatedByString, _ := json.Marshal(CreatedBy)
		json.Unmarshal(CreatedByString, &o.CreatedBy)
	}
	
	if UpdatedBy, ok := MessagingsettingreferenceMap["updatedBy"].(map[string]interface{}); ok {
		UpdatedByString, _ := json.Marshal(UpdatedBy)
		json.Unmarshal(UpdatedByString, &o.UpdatedBy)
	}
	
	if Content, ok := MessagingsettingreferenceMap["content"].(map[string]interface{}); ok {
		ContentString, _ := json.Marshal(Content)
		json.Unmarshal(ContentString, &o.Content)
	}
	
	if Event, ok := MessagingsettingreferenceMap["event"].(map[string]interface{}); ok {
		EventString, _ := json.Marshal(Event)
		json.Unmarshal(EventString, &o.Event)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Messagingsettingreference) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
