package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Analyticsflow
type Analyticsflow struct { 
	// EndingLanguage - Flow ending language, e.g. en-us
	EndingLanguage *string `json:"endingLanguage,omitempty"`


	// EntryReason - The particular entry reason for this flow, e.g. an address, userId, or flowId
	EntryReason *string `json:"entryReason,omitempty"`


	// EntryType - The entry type for this flow, e.g. dnis, dialer, agent, flow, or direct
	EntryType *string `json:"entryType,omitempty"`


	// ExitReason - The exit reason for this flow, e.g. DISCONNECT
	ExitReason *string `json:"exitReason,omitempty"`


	// FlowId - The unique identifier of this flow
	FlowId *string `json:"flowId,omitempty"`


	// FlowName - The name of this flow at the time of flow execution
	FlowName *string `json:"flowName,omitempty"`


	// FlowType - The type of this flow
	FlowType *string `json:"flowType,omitempty"`


	// FlowVersion - The version of this flow
	FlowVersion *string `json:"flowVersion,omitempty"`


	// IssuedCallback - Flag indicating whether the flow issued a callback
	IssuedCallback *bool `json:"issuedCallback,omitempty"`


	// RecognitionFailureReason - The recognition failure reason causing to exit/disconnect
	RecognitionFailureReason *string `json:"recognitionFailureReason,omitempty"`


	// StartingLanguage - Flow starting language, e.g. en-us
	StartingLanguage *string `json:"startingLanguage,omitempty"`


	// TransferTargetAddress - The address of a flow transfer target, e.g. a phone number, an email address, or a queueId
	TransferTargetAddress *string `json:"transferTargetAddress,omitempty"`


	// TransferTargetName - The name of a flow transfer target
	TransferTargetName *string `json:"transferTargetName,omitempty"`


	// TransferType - The type of transfer for flows that ended with a transfer
	TransferType *string `json:"transferType,omitempty"`


	// Outcomes - Flow outcomes
	Outcomes *[]Analyticsflowoutcome `json:"outcomes,omitempty"`

}

// String returns a JSON representation of the model
func (o *Analyticsflow) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
