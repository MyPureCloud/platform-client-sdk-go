package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Routingskill
type Routingskill struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name - The name of the skill.
	Name *string `json:"name,omitempty"`


	// DateModified - Date last modified. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`


	// State - The current state for this skill.
	State *string `json:"state,omitempty"`


	// Version - Required when updating. Version must be the current version. Only the system can assign version.
	Version *string `json:"version,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Routingskill) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Routingskill
	
	DateModified := new(string)
	if o.DateModified != nil {
		
		*DateModified = timeutil.Strftime(o.DateModified, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateModified = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		Version *string `json:"version,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		DateModified: DateModified,
		
		State: o.State,
		
		Version: o.Version,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Routingskill) UnmarshalJSON(b []byte) error {
	var RoutingskillMap map[string]interface{}
	err := json.Unmarshal(b, &RoutingskillMap)
	if err != nil {
		return err
	}
	
	if Id, ok := RoutingskillMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := RoutingskillMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if dateModifiedString, ok := RoutingskillMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if State, ok := RoutingskillMap["state"].(string); ok {
		o.State = &State
	}
    
	if Version, ok := RoutingskillMap["version"].(string); ok {
		o.Version = &Version
	}
    
	if SelfUri, ok := RoutingskillMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Routingskill) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
