package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Flowdivisionview
type Flowdivisionview struct { 
	// Id - The flow identifier
	Id *string `json:"id,omitempty"`


	// Name - The flow name
	Name *string `json:"name,omitempty"`


	// Division - The division to which this entity belongs.
	Division *Writabledivision `json:"division,omitempty"`


	// VarType
	VarType *string `json:"type,omitempty"`


	// InputSchema - json schema describing the inputs for the flow
	InputSchema *Jsonschemadocument `json:"inputSchema,omitempty"`


	// OutputSchema - json schema describing the outputs for the flow
	OutputSchema *Jsonschemadocument `json:"outputSchema,omitempty"`


	// PublishedVersion - published version information if there is a published version
	PublishedVersion *Flowversion `json:"publishedVersion,omitempty"`


	// DebugVersion - debug version information if there is a debug version
	DebugVersion *Flowversion `json:"debugVersion,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Flowdivisionview) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Flowdivisionview
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Division *Writabledivision `json:"division,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		InputSchema *Jsonschemadocument `json:"inputSchema,omitempty"`
		
		OutputSchema *Jsonschemadocument `json:"outputSchema,omitempty"`
		
		PublishedVersion *Flowversion `json:"publishedVersion,omitempty"`
		
		DebugVersion *Flowversion `json:"debugVersion,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Division: o.Division,
		
		VarType: o.VarType,
		
		InputSchema: o.InputSchema,
		
		OutputSchema: o.OutputSchema,
		
		PublishedVersion: o.PublishedVersion,
		
		DebugVersion: o.DebugVersion,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Flowdivisionview) UnmarshalJSON(b []byte) error {
	var FlowdivisionviewMap map[string]interface{}
	err := json.Unmarshal(b, &FlowdivisionviewMap)
	if err != nil {
		return err
	}
	
	if Id, ok := FlowdivisionviewMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := FlowdivisionviewMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if Division, ok := FlowdivisionviewMap["division"].(map[string]interface{}); ok {
		DivisionString, _ := json.Marshal(Division)
		json.Unmarshal(DivisionString, &o.Division)
	}
	
	if VarType, ok := FlowdivisionviewMap["type"].(string); ok {
		o.VarType = &VarType
	}
	
	if InputSchema, ok := FlowdivisionviewMap["inputSchema"].(map[string]interface{}); ok {
		InputSchemaString, _ := json.Marshal(InputSchema)
		json.Unmarshal(InputSchemaString, &o.InputSchema)
	}
	
	if OutputSchema, ok := FlowdivisionviewMap["outputSchema"].(map[string]interface{}); ok {
		OutputSchemaString, _ := json.Marshal(OutputSchema)
		json.Unmarshal(OutputSchemaString, &o.OutputSchema)
	}
	
	if PublishedVersion, ok := FlowdivisionviewMap["publishedVersion"].(map[string]interface{}); ok {
		PublishedVersionString, _ := json.Marshal(PublishedVersion)
		json.Unmarshal(PublishedVersionString, &o.PublishedVersion)
	}
	
	if DebugVersion, ok := FlowdivisionviewMap["debugVersion"].(map[string]interface{}); ok {
		DebugVersionString, _ := json.Marshal(DebugVersion)
		json.Unmarshal(DebugVersionString, &o.DebugVersion)
	}
	
	if SelfUri, ok := FlowdivisionviewMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Flowdivisionview) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
