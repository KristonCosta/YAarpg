package ecs

type System interface {
	Update(float64)
	RequiredTypes() *[]ComponentType
}
