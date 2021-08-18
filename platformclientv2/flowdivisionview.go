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

func (u *Flowdivisionview) MarshalJSON() ([]byte, error) {
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
		Id: u.Id,
		
		Name: u.Name,
		
		Division: u.Division,
		
		VarType: u.VarType,
		
		InputSchema: u.InputSchema,
		
		OutputSchema: u.OutputSchema,
		
		PublishedVersion: u.PublishedVersion,
		
		DebugVersion: u.DebugVersion,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Flowdivisionview) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
