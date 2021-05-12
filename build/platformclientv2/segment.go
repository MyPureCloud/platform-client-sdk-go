package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Segment
type Segment struct { 
	// StartTime - The timestamp when this segment began. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	StartTime *time.Time `json:"startTime,omitempty"`


	// EndTime - The timestamp when this segment ended. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	EndTime *time.Time `json:"endTime,omitempty"`


	// VarType - The activity taking place for the participant in the segment.
	VarType *string `json:"type,omitempty"`


	// HowEnded - A description of the event that ended the segment.
	HowEnded *string `json:"howEnded,omitempty"`


	// DisconnectType - A description of the event that disconnected the segment
	DisconnectType *string `json:"disconnectType,omitempty"`

}

// String returns a JSON representation of the model
func (o *Segment) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
