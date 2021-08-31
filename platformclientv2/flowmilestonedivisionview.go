package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Flowmilestonedivisionview
type Flowmilestonedivisionview struct { 
	// Id - The flow milestone identifier
	Id *string `json:"id,omitempty"`


	// Name - The flow milestone name
	Name *string `json:"name,omitempty"`


	// Division - The division to which this entity belongs.
	Division *Writabledivision `json:"division,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Flowmilestonedivisionview) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Flowmilestonedivisionview
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Division *Writabledivision `json:"division,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Division: o.Division,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Flowmilestonedivisionview) UnmarshalJSON(b []byte) error {
	var FlowmilestonedivisionviewMap map[string]interface{}
	err := json.Unmarshal(b, &FlowmilestonedivisionviewMap)
	if err != nil {
		return err
	}
	
	if Id, ok := FlowmilestonedivisionviewMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := FlowmilestonedivisionviewMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if Division, ok := FlowmilestonedivisionviewMap["division"].(map[string]interface{}); ok {
		DivisionString, _ := json.Marshal(Division)
		json.Unmarshal(DivisionString, &o.Division)
	}
	
	if SelfUri, ok := FlowmilestonedivisionviewMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Flowmilestonedivisionview) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
