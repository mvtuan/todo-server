package models

type CommonFields struct {
	AndCondition []Condition
	OrCondition  []Condition
}

type Condition struct {
	Field    string
	Operator string
	Value    interface{}
}
