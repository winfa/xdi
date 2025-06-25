package xdi

type AutoConstructdStruct struct {
	Name  string `inject:""`
	Age   int    `inject:""`
	Email string `inject:""`
}

type NestedAutoConstructdStruct struct {
	Injected AutoConstructdStruct `inject:""`
	Address  string               `inject:""`
}

type AutoConstructdStructWithNonInjectFields struct {
	Name  string `inject:""`
	Age   int    `inject:""`
	Email string `inject:""`
	Other string // Field without inject tag
}
