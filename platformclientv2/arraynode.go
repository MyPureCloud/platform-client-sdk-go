package platformclientv2
import (
	"encoding/json"
)

// Arraynode
type Arraynode struct { 
	// NodeType
	NodeType *string `json:"nodeType,omitempty"`


	// Float
	Float *bool `json:"float,omitempty"`


	// Number
	Number *bool `json:"number,omitempty"`


	// Boolean
	Boolean *bool `json:"boolean,omitempty"`


	// Object
	Object *bool `json:"object,omitempty"`


	// MissingNode
	MissingNode *bool `json:"missingNode,omitempty"`


	// ValueNode
	ValueNode *bool `json:"valueNode,omitempty"`


	// Binary
	Binary *bool `json:"binary,omitempty"`


	// Pojo
	Pojo *bool `json:"pojo,omitempty"`


	// ContainerNode
	ContainerNode *bool `json:"containerNode,omitempty"`


	// Int
	Int *bool `json:"int,omitempty"`


	// IntegralNumber
	IntegralNumber *bool `json:"integralNumber,omitempty"`


	// Short
	Short *bool `json:"short,omitempty"`


	// Long
	Long *bool `json:"long,omitempty"`


	// FloatingPointNumber
	FloatingPointNumber *bool `json:"floatingPointNumber,omitempty"`


	// Double
	Double *bool `json:"double,omitempty"`


	// BigDecimal
	BigDecimal *bool `json:"bigDecimal,omitempty"`


	// BigInteger
	BigInteger *bool `json:"bigInteger,omitempty"`


	// Textual
	Textual *bool `json:"textual,omitempty"`


	// Array
	Array *bool `json:"array,omitempty"`


	// Null
	Null *bool `json:"null,omitempty"`

}

// String returns a JSON representation of the model
func (o *Arraynode) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
