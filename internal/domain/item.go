package domain

type ItemClass string
type Symbol string
type Uri string

const (
	ITEM_CLASS_CURRENCY        ItemClass = "currency"
	ITEM_CLASS_FRAGMENTS       ItemClass = "fragments"
	ITEM_CLASS_DIVIATION_CARDS ItemClass = "diviation_cards"
	ITEM_CLASS_ARTIFACTS       ItemClass = "artifacts"
	ITEM_CLASS_OILS            ItemClass = "oils"
	ITEM_CLASS_SCARABS         ItemClass = "scarabs"
)

type Item struct {
	Symbol Symbol
	Name   string
	Icon   Uri
	Class  ItemClass
}

type ItemsRepository interface {
}
