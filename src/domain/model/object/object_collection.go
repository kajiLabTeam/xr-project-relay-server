package object_model_domain

type ObjectCollection []Object

func (oc *ObjectCollection) AddObject(o *Object) {
	*oc = append(*oc, *o)
}
