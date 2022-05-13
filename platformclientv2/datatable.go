package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Datatable - Contains a metadata representation for a JSON schema stored in DataTables along with an optional field for the schema itself
type Datatable struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Division - The division to which this entity belongs.
	Division *Writabledivision `json:"division,omitempty"`


	// Description - The description from the JSON schema (equates to the Description field on the JSON schema.)
	Description *string `json:"description,omitempty"`


	// Schema - the schema as stored in the system.
	Schema *Jsonschemadocument `json:"schema,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Datatable) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Datatable
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Division *Writabledivision `json:"division,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		Schema *Jsonschemadocument `json:"schema,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Division: o.Division,
		
		Description: o.Description,
		
		Schema: o.Schema,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Datatable) UnmarshalJSON(b []byte) error {
	var DatatableMap map[string]interface{}
	err := json.Unmarshal(b, &DatatableMap)
	if err != nil {
		return err
	}
	
	if Id, ok := DatatableMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := DatatableMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Division, ok := DatatableMap["division"].(map[string]interface{}); ok {
		DivisionString, _ := json.Marshal(Division)
		json.Unmarshal(DivisionString, &o.Division)
	}
	
	if Description, ok := DatatableMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if Schema, ok := DatatableMap["schema"].(map[string]interface{}); ok {
		SchemaString, _ := json.Marshal(Schema)
		json.Unmarshal(SchemaString, &o.Schema)
	}
	
	if SelfUri, ok := DatatableMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Datatable) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
