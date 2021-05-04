package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Agentownedmappingpreview
type Agentownedmappingpreview struct { 
	// AgentOwnedColumn - The raw value of the agent-owned column
	AgentOwnedColumn *string `json:"agentOwnedColumn,omitempty"`


	// Email - The email address of the user, if it exists
	Email *string `json:"email,omitempty"`


	// UserId - The id of the user, if it exists
	UserId *string `json:"userId,omitempty"`


	// Exists - Whether the user exists
	Exists *bool `json:"exists,omitempty"`


	// IsQueueMember - Whether the user is a member of the campaign's queue
	IsQueueMember *bool `json:"isQueueMember,omitempty"`


	// RecordCount - The number of contact records whose agent-owned column matches the raw value
	RecordCount *int `json:"recordCount,omitempty"`

}

// String returns a JSON representation of the model
func (o *Agentownedmappingpreview) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
