package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Flowmilestone
type Flowmilestone struct { 
	// Id - The flow milestone identifier
	Id *string `json:"id,omitempty"`


	// Name - The flow milestone name.
	Name *string `json:"name,omitempty"`


	// Division - The division to which this entity belongs.
	Division *Writabledivision `json:"division,omitempty"`


	// Description - The flow milestone description.
	Description *string `json:"description,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Flowmilestone) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Flowmilestone
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Division *Writabledivision `json:"division,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Division: o.Division,
		
		Description: o.Description,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Flowmilestone) UnmarshalJSON(b []byte) error {
	var FlowmilestoneMap map[string]interface{}
	err := json.Unmarshal(b, &FlowmilestoneMap)
	if err != nil {
		return err
	}
	
	if Id, ok := FlowmilestoneMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := FlowmilestoneMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if Division, ok := FlowmilestoneMap["division"].(map[string]interface{}); ok {
		DivisionString, _ := json.Marshal(Division)
		json.Unmarshal(DivisionString, &o.Division)
	}
	
	if Description, ok := FlowmilestoneMap["description"].(string); ok {
		o.Description = &Description
	}
	
	if SelfUri, ok := FlowmilestoneMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Flowmilestone) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
