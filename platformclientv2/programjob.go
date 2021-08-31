package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Programjob
type Programjob struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// State
	State *string `json:"state,omitempty"`


	// Programs
	Programs *[]Baseprogramentity `json:"programs,omitempty"`


	// CreatedBy
	CreatedBy *Addressableentityref `json:"createdBy,omitempty"`


	// DateCreated - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateModified - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Programjob) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Programjob
	
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
		
		State *string `json:"state,omitempty"`
		
		Programs *[]Baseprogramentity `json:"programs,omitempty"`
		
		CreatedBy *Addressableentityref `json:"createdBy,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		State: o.State,
		
		Programs: o.Programs,
		
		CreatedBy: o.CreatedBy,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Programjob) UnmarshalJSON(b []byte) error {
	var ProgramjobMap map[string]interface{}
	err := json.Unmarshal(b, &ProgramjobMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ProgramjobMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if State, ok := ProgramjobMap["state"].(string); ok {
		o.State = &State
	}
	
	if Programs, ok := ProgramjobMap["programs"].([]interface{}); ok {
		ProgramsString, _ := json.Marshal(Programs)
		json.Unmarshal(ProgramsString, &o.Programs)
	}
	
	if CreatedBy, ok := ProgramjobMap["createdBy"].(map[string]interface{}); ok {
		CreatedByString, _ := json.Marshal(CreatedBy)
		json.Unmarshal(CreatedByString, &o.CreatedBy)
	}
	
	if dateCreatedString, ok := ProgramjobMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := ProgramjobMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if SelfUri, ok := ProgramjobMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Programjob) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
