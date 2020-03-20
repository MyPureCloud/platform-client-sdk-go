package platformclientv2
import (
	"encoding/json"
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


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Flowdivisionview) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
