package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Webchatroutingtarget
type Webchatroutingtarget struct { 
	// TargetType - The target type of the routing target, such as 'QUEUE'.
	TargetType *string `json:"targetType,omitempty"`


	// TargetAddress - The target of the route, in the format appropriate given the 'targetType'.
	TargetAddress *string `json:"targetAddress,omitempty"`


	// Skills - The list of skill names to use for routing.
	Skills *[]string `json:"skills,omitempty"`


	// Language - The language name to use for routing.
	Language *string `json:"language,omitempty"`


	// Priority - The priority to assign to the conversation for routing.
	Priority *int `json:"priority,omitempty"`

}

func (o *Webchatroutingtarget) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Webchatroutingtarget
	
	return json.Marshal(&struct { 
		TargetType *string `json:"targetType,omitempty"`
		
		TargetAddress *string `json:"targetAddress,omitempty"`
		
		Skills *[]string `json:"skills,omitempty"`
		
		Language *string `json:"language,omitempty"`
		
		Priority *int `json:"priority,omitempty"`
		*Alias
	}{ 
		TargetType: o.TargetType,
		
		TargetAddress: o.TargetAddress,
		
		Skills: o.Skills,
		
		Language: o.Language,
		
		Priority: o.Priority,
		Alias:    (*Alias)(o),
	})
}

func (o *Webchatroutingtarget) UnmarshalJSON(b []byte) error {
	var WebchatroutingtargetMap map[string]interface{}
	err := json.Unmarshal(b, &WebchatroutingtargetMap)
	if err != nil {
		return err
	}
	
	if TargetType, ok := WebchatroutingtargetMap["targetType"].(string); ok {
		o.TargetType = &TargetType
	}
	
	if TargetAddress, ok := WebchatroutingtargetMap["targetAddress"].(string); ok {
		o.TargetAddress = &TargetAddress
	}
	
	if Skills, ok := WebchatroutingtargetMap["skills"].([]interface{}); ok {
		SkillsString, _ := json.Marshal(Skills)
		json.Unmarshal(SkillsString, &o.Skills)
	}
	
	if Language, ok := WebchatroutingtargetMap["language"].(string); ok {
		o.Language = &Language
	}
	
	if Priority, ok := WebchatroutingtargetMap["priority"].(float64); ok {
		PriorityInt := int(Priority)
		o.Priority = &PriorityInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Webchatroutingtarget) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
