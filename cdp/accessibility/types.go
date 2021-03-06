package accessibility

// AUTOGENERATED. DO NOT EDIT.

import (
	"errors"

	cdp "github.com/knq/chromedp/cdp"
	"github.com/mailru/easyjson"
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

// AXNodeID unique accessibility node identifier.
type AXNodeID string

// String returns the AXNodeID as string value.
func (t AXNodeID) String() string {
	return string(t)
}

// AXValueType enum of possible property types.
type AXValueType string

// String returns the AXValueType as string value.
func (t AXValueType) String() string {
	return string(t)
}

// AXValueType values.
const (
	AXValueTypeBoolean            AXValueType = "boolean"
	AXValueTypeTristate           AXValueType = "tristate"
	AXValueTypeBooleanOrUndefined AXValueType = "booleanOrUndefined"
	AXValueTypeIdref              AXValueType = "idref"
	AXValueTypeIdrefList          AXValueType = "idrefList"
	AXValueTypeInteger            AXValueType = "integer"
	AXValueTypeNode               AXValueType = "node"
	AXValueTypeNodeList           AXValueType = "nodeList"
	AXValueTypeNumber             AXValueType = "number"
	AXValueTypeString             AXValueType = "string"
	AXValueTypeComputedString     AXValueType = "computedString"
	AXValueTypeToken              AXValueType = "token"
	AXValueTypeTokenList          AXValueType = "tokenList"
	AXValueTypeDomRelation        AXValueType = "domRelation"
	AXValueTypeRole               AXValueType = "role"
	AXValueTypeInternalRole       AXValueType = "internalRole"
	AXValueTypeValueUndefined     AXValueType = "valueUndefined"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t AXValueType) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t AXValueType) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *AXValueType) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch AXValueType(in.String()) {
	case AXValueTypeBoolean:
		*t = AXValueTypeBoolean
	case AXValueTypeTristate:
		*t = AXValueTypeTristate
	case AXValueTypeBooleanOrUndefined:
		*t = AXValueTypeBooleanOrUndefined
	case AXValueTypeIdref:
		*t = AXValueTypeIdref
	case AXValueTypeIdrefList:
		*t = AXValueTypeIdrefList
	case AXValueTypeInteger:
		*t = AXValueTypeInteger
	case AXValueTypeNode:
		*t = AXValueTypeNode
	case AXValueTypeNodeList:
		*t = AXValueTypeNodeList
	case AXValueTypeNumber:
		*t = AXValueTypeNumber
	case AXValueTypeString:
		*t = AXValueTypeString
	case AXValueTypeComputedString:
		*t = AXValueTypeComputedString
	case AXValueTypeToken:
		*t = AXValueTypeToken
	case AXValueTypeTokenList:
		*t = AXValueTypeTokenList
	case AXValueTypeDomRelation:
		*t = AXValueTypeDomRelation
	case AXValueTypeRole:
		*t = AXValueTypeRole
	case AXValueTypeInternalRole:
		*t = AXValueTypeInternalRole
	case AXValueTypeValueUndefined:
		*t = AXValueTypeValueUndefined

	default:
		in.AddError(errors.New("unknown AXValueType value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *AXValueType) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// AXValueSourceType enum of possible property sources.
type AXValueSourceType string

// String returns the AXValueSourceType as string value.
func (t AXValueSourceType) String() string {
	return string(t)
}

// AXValueSourceType values.
const (
	AXValueSourceTypeAttribute      AXValueSourceType = "attribute"
	AXValueSourceTypeImplicit       AXValueSourceType = "implicit"
	AXValueSourceTypeStyle          AXValueSourceType = "style"
	AXValueSourceTypeContents       AXValueSourceType = "contents"
	AXValueSourceTypePlaceholder    AXValueSourceType = "placeholder"
	AXValueSourceTypeRelatedElement AXValueSourceType = "relatedElement"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t AXValueSourceType) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t AXValueSourceType) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *AXValueSourceType) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch AXValueSourceType(in.String()) {
	case AXValueSourceTypeAttribute:
		*t = AXValueSourceTypeAttribute
	case AXValueSourceTypeImplicit:
		*t = AXValueSourceTypeImplicit
	case AXValueSourceTypeStyle:
		*t = AXValueSourceTypeStyle
	case AXValueSourceTypeContents:
		*t = AXValueSourceTypeContents
	case AXValueSourceTypePlaceholder:
		*t = AXValueSourceTypePlaceholder
	case AXValueSourceTypeRelatedElement:
		*t = AXValueSourceTypeRelatedElement

	default:
		in.AddError(errors.New("unknown AXValueSourceType value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *AXValueSourceType) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// AXValueNativeSourceType enum of possible native property sources (as a
// subtype of a particular AXValueSourceType).
type AXValueNativeSourceType string

// String returns the AXValueNativeSourceType as string value.
func (t AXValueNativeSourceType) String() string {
	return string(t)
}

// AXValueNativeSourceType values.
const (
	AXValueNativeSourceTypeFigcaption   AXValueNativeSourceType = "figcaption"
	AXValueNativeSourceTypeLabel        AXValueNativeSourceType = "label"
	AXValueNativeSourceTypeLabelfor     AXValueNativeSourceType = "labelfor"
	AXValueNativeSourceTypeLabelwrapped AXValueNativeSourceType = "labelwrapped"
	AXValueNativeSourceTypeLegend       AXValueNativeSourceType = "legend"
	AXValueNativeSourceTypeTablecaption AXValueNativeSourceType = "tablecaption"
	AXValueNativeSourceTypeTitle        AXValueNativeSourceType = "title"
	AXValueNativeSourceTypeOther        AXValueNativeSourceType = "other"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t AXValueNativeSourceType) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t AXValueNativeSourceType) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *AXValueNativeSourceType) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch AXValueNativeSourceType(in.String()) {
	case AXValueNativeSourceTypeFigcaption:
		*t = AXValueNativeSourceTypeFigcaption
	case AXValueNativeSourceTypeLabel:
		*t = AXValueNativeSourceTypeLabel
	case AXValueNativeSourceTypeLabelfor:
		*t = AXValueNativeSourceTypeLabelfor
	case AXValueNativeSourceTypeLabelwrapped:
		*t = AXValueNativeSourceTypeLabelwrapped
	case AXValueNativeSourceTypeLegend:
		*t = AXValueNativeSourceTypeLegend
	case AXValueNativeSourceTypeTablecaption:
		*t = AXValueNativeSourceTypeTablecaption
	case AXValueNativeSourceTypeTitle:
		*t = AXValueNativeSourceTypeTitle
	case AXValueNativeSourceTypeOther:
		*t = AXValueNativeSourceTypeOther

	default:
		in.AddError(errors.New("unknown AXValueNativeSourceType value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *AXValueNativeSourceType) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// AXValueSource a single source for a computed AX property.
type AXValueSource struct {
	Type              AXValueSourceType       `json:"type,omitempty"`              // What type of source this is.
	Value             *AXValue                `json:"value,omitempty"`             // The value of this property source.
	Attribute         string                  `json:"attribute,omitempty"`         // The name of the relevant attribute, if any.
	AttributeValue    *AXValue                `json:"attributeValue,omitempty"`    // The value of the relevant attribute, if any.
	Superseded        bool                    `json:"superseded,omitempty"`        // Whether this source is superseded by a higher priority source.
	NativeSource      AXValueNativeSourceType `json:"nativeSource,omitempty"`      // The native markup source for this value, e.g. a <label> element.
	NativeSourceValue *AXValue                `json:"nativeSourceValue,omitempty"` // The value, such as a node or node list, of the native source.
	Invalid           bool                    `json:"invalid,omitempty"`           // Whether the value for this property is invalid.
	InvalidReason     string                  `json:"invalidReason,omitempty"`     // Reason for the value being invalid, if it is.
}

// AXRelatedNode [no description].
type AXRelatedNode struct {
	BackendDOMNodeID cdp.BackendNodeID `json:"backendDOMNodeId,omitempty"` // The BackendNodeId of the related DOM node.
	Idref            string            `json:"idref,omitempty"`            // The IDRef value provided, if any.
	Text             string            `json:"text,omitempty"`             // The text alternative of this node in the current context.
}

// AXProperty [no description].
type AXProperty struct {
	Name  string   `json:"name,omitempty"`  // The name of this property.
	Value *AXValue `json:"value,omitempty"` // The value of this property.
}

// AXValue a single computed AX property.
type AXValue struct {
	Type         AXValueType         `json:"type,omitempty"`         // The type of this value.
	Value        easyjson.RawMessage `json:"value,omitempty"`        // The computed value of this property.
	RelatedNodes []*AXRelatedNode    `json:"relatedNodes,omitempty"` // One or more related nodes, if applicable.
	Sources      []*AXValueSource    `json:"sources,omitempty"`      // The sources which contributed to the computation of this property.
}

// AXGlobalStates states which apply to every AX node.
type AXGlobalStates string

// String returns the AXGlobalStates as string value.
func (t AXGlobalStates) String() string {
	return string(t)
}

// AXGlobalStates values.
const (
	AXGlobalStatesDisabled   AXGlobalStates = "disabled"
	AXGlobalStatesHidden     AXGlobalStates = "hidden"
	AXGlobalStatesHiddenRoot AXGlobalStates = "hiddenRoot"
	AXGlobalStatesInvalid    AXGlobalStates = "invalid"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t AXGlobalStates) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t AXGlobalStates) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *AXGlobalStates) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch AXGlobalStates(in.String()) {
	case AXGlobalStatesDisabled:
		*t = AXGlobalStatesDisabled
	case AXGlobalStatesHidden:
		*t = AXGlobalStatesHidden
	case AXGlobalStatesHiddenRoot:
		*t = AXGlobalStatesHiddenRoot
	case AXGlobalStatesInvalid:
		*t = AXGlobalStatesInvalid

	default:
		in.AddError(errors.New("unknown AXGlobalStates value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *AXGlobalStates) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// AXLiveRegionAttributes attributes which apply to nodes in live regions.
type AXLiveRegionAttributes string

// String returns the AXLiveRegionAttributes as string value.
func (t AXLiveRegionAttributes) String() string {
	return string(t)
}

// AXLiveRegionAttributes values.
const (
	AXLiveRegionAttributesLive     AXLiveRegionAttributes = "live"
	AXLiveRegionAttributesAtomic   AXLiveRegionAttributes = "atomic"
	AXLiveRegionAttributesRelevant AXLiveRegionAttributes = "relevant"
	AXLiveRegionAttributesBusy     AXLiveRegionAttributes = "busy"
	AXLiveRegionAttributesRoot     AXLiveRegionAttributes = "root"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t AXLiveRegionAttributes) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t AXLiveRegionAttributes) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *AXLiveRegionAttributes) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch AXLiveRegionAttributes(in.String()) {
	case AXLiveRegionAttributesLive:
		*t = AXLiveRegionAttributesLive
	case AXLiveRegionAttributesAtomic:
		*t = AXLiveRegionAttributesAtomic
	case AXLiveRegionAttributesRelevant:
		*t = AXLiveRegionAttributesRelevant
	case AXLiveRegionAttributesBusy:
		*t = AXLiveRegionAttributesBusy
	case AXLiveRegionAttributesRoot:
		*t = AXLiveRegionAttributesRoot

	default:
		in.AddError(errors.New("unknown AXLiveRegionAttributes value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *AXLiveRegionAttributes) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// AXWidgetAttributes attributes which apply to widgets.
type AXWidgetAttributes string

// String returns the AXWidgetAttributes as string value.
func (t AXWidgetAttributes) String() string {
	return string(t)
}

// AXWidgetAttributes values.
const (
	AXWidgetAttributesAutocomplete    AXWidgetAttributes = "autocomplete"
	AXWidgetAttributesHaspopup        AXWidgetAttributes = "haspopup"
	AXWidgetAttributesLevel           AXWidgetAttributes = "level"
	AXWidgetAttributesMultiselectable AXWidgetAttributes = "multiselectable"
	AXWidgetAttributesOrientation     AXWidgetAttributes = "orientation"
	AXWidgetAttributesMultiline       AXWidgetAttributes = "multiline"
	AXWidgetAttributesReadonly        AXWidgetAttributes = "readonly"
	AXWidgetAttributesRequired        AXWidgetAttributes = "required"
	AXWidgetAttributesValuemin        AXWidgetAttributes = "valuemin"
	AXWidgetAttributesValuemax        AXWidgetAttributes = "valuemax"
	AXWidgetAttributesValuetext       AXWidgetAttributes = "valuetext"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t AXWidgetAttributes) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t AXWidgetAttributes) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *AXWidgetAttributes) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch AXWidgetAttributes(in.String()) {
	case AXWidgetAttributesAutocomplete:
		*t = AXWidgetAttributesAutocomplete
	case AXWidgetAttributesHaspopup:
		*t = AXWidgetAttributesHaspopup
	case AXWidgetAttributesLevel:
		*t = AXWidgetAttributesLevel
	case AXWidgetAttributesMultiselectable:
		*t = AXWidgetAttributesMultiselectable
	case AXWidgetAttributesOrientation:
		*t = AXWidgetAttributesOrientation
	case AXWidgetAttributesMultiline:
		*t = AXWidgetAttributesMultiline
	case AXWidgetAttributesReadonly:
		*t = AXWidgetAttributesReadonly
	case AXWidgetAttributesRequired:
		*t = AXWidgetAttributesRequired
	case AXWidgetAttributesValuemin:
		*t = AXWidgetAttributesValuemin
	case AXWidgetAttributesValuemax:
		*t = AXWidgetAttributesValuemax
	case AXWidgetAttributesValuetext:
		*t = AXWidgetAttributesValuetext

	default:
		in.AddError(errors.New("unknown AXWidgetAttributes value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *AXWidgetAttributes) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// AXWidgetStates states which apply to widgets.
type AXWidgetStates string

// String returns the AXWidgetStates as string value.
func (t AXWidgetStates) String() string {
	return string(t)
}

// AXWidgetStates values.
const (
	AXWidgetStatesChecked  AXWidgetStates = "checked"
	AXWidgetStatesExpanded AXWidgetStates = "expanded"
	AXWidgetStatesPressed  AXWidgetStates = "pressed"
	AXWidgetStatesSelected AXWidgetStates = "selected"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t AXWidgetStates) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t AXWidgetStates) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *AXWidgetStates) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch AXWidgetStates(in.String()) {
	case AXWidgetStatesChecked:
		*t = AXWidgetStatesChecked
	case AXWidgetStatesExpanded:
		*t = AXWidgetStatesExpanded
	case AXWidgetStatesPressed:
		*t = AXWidgetStatesPressed
	case AXWidgetStatesSelected:
		*t = AXWidgetStatesSelected

	default:
		in.AddError(errors.New("unknown AXWidgetStates value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *AXWidgetStates) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// AXRelationshipAttributes relationships between elements other than
// parent/child/sibling.
type AXRelationshipAttributes string

// String returns the AXRelationshipAttributes as string value.
func (t AXRelationshipAttributes) String() string {
	return string(t)
}

// AXRelationshipAttributes values.
const (
	AXRelationshipAttributesActivedescendant AXRelationshipAttributes = "activedescendant"
	AXRelationshipAttributesFlowto           AXRelationshipAttributes = "flowto"
	AXRelationshipAttributesControls         AXRelationshipAttributes = "controls"
	AXRelationshipAttributesDescribedby      AXRelationshipAttributes = "describedby"
	AXRelationshipAttributesLabelledby       AXRelationshipAttributes = "labelledby"
	AXRelationshipAttributesOwns             AXRelationshipAttributes = "owns"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t AXRelationshipAttributes) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t AXRelationshipAttributes) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *AXRelationshipAttributes) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch AXRelationshipAttributes(in.String()) {
	case AXRelationshipAttributesActivedescendant:
		*t = AXRelationshipAttributesActivedescendant
	case AXRelationshipAttributesFlowto:
		*t = AXRelationshipAttributesFlowto
	case AXRelationshipAttributesControls:
		*t = AXRelationshipAttributesControls
	case AXRelationshipAttributesDescribedby:
		*t = AXRelationshipAttributesDescribedby
	case AXRelationshipAttributesLabelledby:
		*t = AXRelationshipAttributesLabelledby
	case AXRelationshipAttributesOwns:
		*t = AXRelationshipAttributesOwns

	default:
		in.AddError(errors.New("unknown AXRelationshipAttributes value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *AXRelationshipAttributes) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// AXNode a node in the accessibility tree.
type AXNode struct {
	NodeID           AXNodeID          `json:"nodeId,omitempty"`           // Unique identifier for this node.
	Ignored          bool              `json:"ignored,omitempty"`          // Whether this node is ignored for accessibility
	IgnoredReasons   []*AXProperty     `json:"ignoredReasons,omitempty"`   // Collection of reasons why this node is hidden.
	Role             *AXValue          `json:"role,omitempty"`             // This Node's role, whether explicit or implicit.
	Name             *AXValue          `json:"name,omitempty"`             // The accessible name for this Node.
	Description      *AXValue          `json:"description,omitempty"`      // The accessible description for this Node.
	Value            *AXValue          `json:"value,omitempty"`            // The value for this Node.
	Properties       []*AXProperty     `json:"properties,omitempty"`       // All other properties
	ChildIds         []AXNodeID        `json:"childIds,omitempty"`         // IDs for each of this node's child nodes.
	BackendDOMNodeID cdp.BackendNodeID `json:"backendDOMNodeId,omitempty"` // The backend ID for the associated DOM node, if any.
}
