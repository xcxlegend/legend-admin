package models

type Base struct{}

func (this *Base) Options() []map[string]string {
	return []map[string]string{}
}

func (this *Base) PreAdd() {}
func (this *Base) PreUpdate() {}
func (this *Base) PreDelete() {}