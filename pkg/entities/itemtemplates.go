package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

//ItemTemplatePropertyType type
type ItemTemplatePropertyType int

//ItemTemplateAttributeType type
type ItemTemplateAttributeType int

//ItemTemplateProperties type
type ItemTemplateProperties []ItemTemplateProperty

//ItemTemplateAttributes type
type ItemTemplateAttributes []ItemTemplateAttribute

const (
	itemTemplatePropertyTypeString = iota + 1
	itemTemplatePropertyTypeInteger
	itemTemplatePropertyTypeDouble
)
const (
	itemTemplateAttributeTypeString = iota + 1
	itemTemplateAttributeTypeInteger
	itemTemplateAttributeTypeDouble
)

//ItemTemplateProperty data
type ItemTemplateProperty struct {
	ID    primitive.ObjectID       `bson:"_id,omitempty" json:"id,omitempty"`
	Value string                   `bson:"value" json:"value,omitempty"`
	Type  ItemTemplatePropertyType `bson:"type,omitempty" json:"type,omitempty"`
}

//ItemTemplateAttribute data
type ItemTemplateAttribute struct {
	ID    primitive.ObjectID       `bson:"_id,omitempty" json:"id,omitempty"`
	Value string                   `bson:"value" json:"value,omitempty"`
	Type  ItemTemplatePropertyType `bson:"type,omitempty" json:"type,omitempty"`
}

//ItemTemplate data
type ItemTemplate struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	TemplateID  string             `yaml:"templateID,omitempty" json:"templateID,omitempty"`
	Name        string             `yaml:"name" json:"name"`
	Description string             `yaml:"description" json:"description"`
	ItemType    ItemType           `yaml:"itemType" json:"itemType"`

	// General properties of the item template (interpreted during creation, effects?)
	Properties ItemTemplateProperties `yaml:"properties"`

	// Generic attributes of the item created (copied over from template)
	Attributes ItemTemplateAttributes `yaml:"attributes"`
}

//ItemTemplates type
type ItemTemplates []*ItemTemplate
