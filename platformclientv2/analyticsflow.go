package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Analyticsflow
type Analyticsflow struct { 
	// FlowId - The unique identifier of this flow
	FlowId *string `json:"flowId,omitempty"`


	// FlowName - The name of this flow
	FlowName *string `json:"flowName,omitempty"`


	// FlowVersion - The version of this flow
	FlowVersion *string `json:"flowVersion,omitempty"`


	// FlowType - The type of this flow
	FlowType *string `json:"flowType,omitempty"`


	// ExitReason - The exit reason for this flow, e.g. DISCONNECT
	ExitReason *string `json:"exitReason,omitempty"`


	// EntryReason - The particular entry reason for this flow, e.g. an address, userId, or flowId
	EntryReason *string `json:"entryReason,omitempty"`


	// EntryType - The entry type for this flow
	EntryType *string `json:"entryType,omitempty"`


	// TransferType - The type of transfer for flows that ended with a transfer
	TransferType *string `json:"transferType,omitempty"`


	// TransferTargetName - The name of a transfer target
	TransferTargetName *string `json:"transferTargetName,omitempty"`


	// TransferTargetAddress - The address of a transfer target
	TransferTargetAddress *string `json:"transferTargetAddress,omitempty"`


	// IssuedCallback - Flag indicating whether the flow issued a callback
	IssuedCallback *bool `json:"issuedCallback,omitempty"`


	// StartingLanguage - Flow starting language, e.g. en-us
	StartingLanguage *string `json:"startingLanguage,omitempty"`


	// EndingLanguage - Flow ending language, e.g. en-us
	EndingLanguage *string `json:"endingLanguage,omitempty"`


	// Outcomes - Flow outcomes
	Outcomes *[]Analyticsflowoutcome `json:"outcomes,omitempty"`

}

// String returns a JSON representation of the model
func (o *Analyticsflow) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
