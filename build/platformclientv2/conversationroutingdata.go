package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationroutingdata
type Conversationroutingdata struct { 
	// Queue - The queue to use for routing decisions
	Queue *Addressableentityref `json:"queue,omitempty"`


	// Language - The language to use for routing decisions
	Language *Addressableentityref `json:"language,omitempty"`


	// Priority - The priority of the conversation to use for routing decisions
	Priority *int `json:"priority,omitempty"`


	// Skills - The skills to use for routing decisions
	Skills *[]Addressableentityref `json:"skills,omitempty"`


	// ScoredAgents - A collection of agents and their assigned scores for this conversation (0 - 100, higher being better), for use in routing to preferred agents
	ScoredAgents *[]Scoredagent `json:"scoredAgents,omitempty"`

}

// String returns a JSON representation of the model
func (o *Conversationroutingdata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
