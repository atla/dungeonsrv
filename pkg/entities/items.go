package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//ItemType type
type ItemType int

//ItemTemplatePropertyType type
type ItemTemplatePropertyType int

//ItemTemplateAttributeType type
type ItemTemplateAttributeType int

//ItemTemplateProperties type
type ItemTemplateProperties []ItemTemplateProperty

//ItemTemplateAttributes type
type ItemTemplateAttributes []ItemTemplateAttribute

const (
	ItemTypeCurrency ItemType = iota + 1
	ItemTypeConsumable
	ItemTypeArmor
	ItemTypeWeapon
	ItemTypeCollectible
)

const (
	ItemTemplatePropertyTypeString = iota + 1
	ItemTemplatePropertyTypeInteger
	ItemTemplatePropertyTypeDouble
)
const (
	ItemTemplateAttributeTypeString = iota + 1
	ItemTemplateAttributeTypeInteger
	ItemTemplateAttributeTypeDouble
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
	Name        string             `yaml:"name"`
	Description string             `yaml:"description"`
	ItemType    ItemType           `yaml:"itemType"`

	// General properties of the item template (interpreted during creation, effects?)
	Properties ItemTemplateProperties `yaml:"properties"`

	// Generic attributes of the item created (copied over from template)
	Attributes ItemTemplateAttributes `yaml:"attributes"`
}

//Item data
type Item struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name        string             `bson:"name,omitempty" json:"name"`
	Description string             `bson:"description,omitempty" json:"description"`
	ItemType    ItemType           `bson:"itemType,omitempty" json:"itemType"`
	Created     time.Time          `bson:"created,omitempty" json:"created,omitempty"`
}
