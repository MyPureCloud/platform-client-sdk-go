package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Team
type Team struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name - The team name
	Name *string `json:"name,omitempty"`


	// Division - The division to which this entity belongs.
	Division *Writabledivision `json:"division,omitempty"`


	// Description - Team information.
	Description *string `json:"description,omitempty"`


	// DateModified - Last modified datetime. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`


	// MemberCount - Number of members in a team
	MemberCount *int `json:"memberCount,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Team) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Team
	
	DateModified := new(string)
	if o.DateModified != nil {
		
		*DateModified = timeutil.Strftime(o.DateModified, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateModified = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Division *Writabledivision `json:"division,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		MemberCount *int `json:"memberCount,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Division: o.Division,
		
		Description: o.Description,
		
		DateModified: DateModified,
		
		MemberCount: o.MemberCount,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Team) UnmarshalJSON(b []byte) error {
	var TeamMap map[string]interface{}
	err := json.Unmarshal(b, &TeamMap)
	if err != nil {
		return err
	}
	
	if Id, ok := TeamMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := TeamMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if Division, ok := TeamMap["division"].(map[string]interface{}); ok {
		DivisionString, _ := json.Marshal(Division)
		json.Unmarshal(DivisionString, &o.Division)
	}
	
	if Description, ok := TeamMap["description"].(string); ok {
		o.Description = &Description
	}
	
	if dateModifiedString, ok := TeamMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if MemberCount, ok := TeamMap["memberCount"].(float64); ok {
		MemberCountInt := int(MemberCount)
		o.MemberCount = &MemberCountInt
	}
	
	if SelfUri, ok := TeamMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Team) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
