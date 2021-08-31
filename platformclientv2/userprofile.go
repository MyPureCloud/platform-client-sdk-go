package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Userprofile
type Userprofile struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// State - The state of the user resource
	State *string `json:"state,omitempty"`


	// DateModified - Datetime of the last modification. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`


	// Version - The version of the group resource
	Version *int `json:"version,omitempty"`


	// Expands - User information expansions
	Expands *Userexpands `json:"expands,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Userprofile) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Userprofile
	
	DateModified := new(string)
	if o.DateModified != nil {
		
		*DateModified = timeutil.Strftime(o.DateModified, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateModified = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		Version *int `json:"version,omitempty"`
		
		Expands *Userexpands `json:"expands,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		State: o.State,
		
		DateModified: DateModified,
		
		Version: o.Version,
		
		Expands: o.Expands,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Userprofile) UnmarshalJSON(b []byte) error {
	var UserprofileMap map[string]interface{}
	err := json.Unmarshal(b, &UserprofileMap)
	if err != nil {
		return err
	}
	
	if Id, ok := UserprofileMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := UserprofileMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if State, ok := UserprofileMap["state"].(string); ok {
		o.State = &State
	}
	
	if dateModifiedString, ok := UserprofileMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if Version, ok := UserprofileMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	
	if Expands, ok := UserprofileMap["expands"].(map[string]interface{}); ok {
		ExpandsString, _ := json.Marshal(Expands)
		json.Unmarshal(ExpandsString, &o.Expands)
	}
	
	if SelfUri, ok := UserprofileMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Userprofile) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
