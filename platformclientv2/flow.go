package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Flow
type Flow struct { 
	// Id - The flow identifier
	Id *string `json:"id,omitempty"`


	// Name - The flow name
	Name *string `json:"name,omitempty"`


	// Division - The division to which this entity belongs.
	Division *Writabledivision `json:"division,omitempty"`


	// Description
	Description *string `json:"description,omitempty"`


	// VarType
	VarType *string `json:"type,omitempty"`


	// LockedUser - User that has the flow locked.
	LockedUser *User `json:"lockedUser,omitempty"`


	// LockedClient - OAuth client that has the flow locked.
	LockedClient *Domainentityref `json:"lockedClient,omitempty"`


	// Active
	Active *bool `json:"active,omitempty"`


	// System
	System *bool `json:"system,omitempty"`


	// Deleted
	Deleted *bool `json:"deleted,omitempty"`


	// PublishedVersion
	PublishedVersion *Flowversion `json:"publishedVersion,omitempty"`


	// SavedVersion
	SavedVersion *Flowversion `json:"savedVersion,omitempty"`


	// InputSchema - json schema describing the inputs for the flow
	InputSchema *map[string]interface{} `json:"inputSchema,omitempty"`


	// OutputSchema - json schema describing the outputs for the flow
	OutputSchema *map[string]interface{} `json:"outputSchema,omitempty"`


	// CheckedInVersion
	CheckedInVersion *Flowversion `json:"checkedInVersion,omitempty"`


	// DebugVersion
	DebugVersion *Flowversion `json:"debugVersion,omitempty"`


	// PublishedBy
	PublishedBy *User `json:"publishedBy,omitempty"`


	// CurrentOperation
	CurrentOperation *Operation `json:"currentOperation,omitempty"`


	// NluInfo - Information about the natural language understanding configuration for the published version of the flow
	NluInfo *Nluinfo `json:"nluInfo,omitempty"`


	// SupportedLanguages - List of supported languages for the published version of the flow.
	SupportedLanguages *[]Supportedlanguage `json:"supportedLanguages,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Flow) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
