package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Routepathrequest - Route path configuration
type Routepathrequest struct { 
	// QueueId - The ID of the queue to associate with the route path
	QueueId *string `json:"queueId,omitempty"`


	// MediaType - The media type of the given queue to associate with the route path
	MediaType *string `json:"mediaType,omitempty"`


	// LanguageId - The ID of the language to associate with the route path
	LanguageId *string `json:"languageId,omitempty"`


	// SkillIds - The set of skill IDs to associate with the route path
	SkillIds *[]string `json:"skillIds,omitempty"`


	// SourcePlanningGroup - The planning group from which to copy route paths
	SourcePlanningGroup *Sourceplanninggrouprequest `json:"sourcePlanningGroup,omitempty"`

}

func (u *Routepathrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Routepathrequest

	

	return json.Marshal(&struct { 
		QueueId *string `json:"queueId,omitempty"`
		
		MediaType *string `json:"mediaType,omitempty"`
		
		LanguageId *string `json:"languageId,omitempty"`
		
		SkillIds *[]string `json:"skillIds,omitempty"`
		
		SourcePlanningGroup *Sourceplanninggrouprequest `json:"sourcePlanningGroup,omitempty"`
		*Alias
	}{ 
		QueueId: u.QueueId,
		
		MediaType: u.MediaType,
		
		LanguageId: u.LanguageId,
		
		SkillIds: u.SkillIds,
		
		SourcePlanningGroup: u.SourcePlanningGroup,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Routepathrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
