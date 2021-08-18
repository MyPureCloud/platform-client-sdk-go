package platformclientv2
import (
	"github.com/leekchan/timeutil"
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
	InputSchema *interface{} `json:"inputSchema,omitempty"`


	// OutputSchema - json schema describing the outputs for the flow
	OutputSchema *interface{} `json:"outputSchema,omitempty"`


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


	// CompatibleFlowTypes - Compatible flow types designate which flow types are allowed to embed a flow’s configuration within their own flow configuration.  Currently the only flows that can be embedded are Common Module flows and the embedding flow can invoke them using the Call Common Module action.
	CompatibleFlowTypes *[]string `json:"compatibleFlowTypes,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (u *Flow) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Flow

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Division *Writabledivision `json:"division,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		LockedUser *User `json:"lockedUser,omitempty"`
		
		LockedClient *Domainentityref `json:"lockedClient,omitempty"`
		
		Active *bool `json:"active,omitempty"`
		
		System *bool `json:"system,omitempty"`
		
		Deleted *bool `json:"deleted,omitempty"`
		
		PublishedVersion *Flowversion `json:"publishedVersion,omitempty"`
		
		SavedVersion *Flowversion `json:"savedVersion,omitempty"`
		
		InputSchema *interface{} `json:"inputSchema,omitempty"`
		
		OutputSchema *interface{} `json:"outputSchema,omitempty"`
		
		CheckedInVersion *Flowversion `json:"checkedInVersion,omitempty"`
		
		DebugVersion *Flowversion `json:"debugVersion,omitempty"`
		
		PublishedBy *User `json:"publishedBy,omitempty"`
		
		CurrentOperation *Operation `json:"currentOperation,omitempty"`
		
		NluInfo *Nluinfo `json:"nluInfo,omitempty"`
		
		SupportedLanguages *[]Supportedlanguage `json:"supportedLanguages,omitempty"`
		
		CompatibleFlowTypes *[]string `json:"compatibleFlowTypes,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		Division: u.Division,
		
		Description: u.Description,
		
		VarType: u.VarType,
		
		LockedUser: u.LockedUser,
		
		LockedClient: u.LockedClient,
		
		Active: u.Active,
		
		System: u.System,
		
		Deleted: u.Deleted,
		
		PublishedVersion: u.PublishedVersion,
		
		SavedVersion: u.SavedVersion,
		
		InputSchema: u.InputSchema,
		
		OutputSchema: u.OutputSchema,
		
		CheckedInVersion: u.CheckedInVersion,
		
		DebugVersion: u.DebugVersion,
		
		PublishedBy: u.PublishedBy,
		
		CurrentOperation: u.CurrentOperation,
		
		NluInfo: u.NluInfo,
		
		SupportedLanguages: u.SupportedLanguages,
		
		CompatibleFlowTypes: u.CompatibleFlowTypes,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Flow) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
