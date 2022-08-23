package set

type Set[T comparable] map[T]nothing

func New[T comparable](size int) Set[T] {
	return make(Set[T], size)
}
func From[T comparable](items ...T) (result Set[T]) {
	result = New[T](len(items))
	result.Add(items...)
	return result
}
func (this Set[T]) Len() int {
	return len(this)
}
func (this Set[T]) Slice() (result []T) {
	result = make([]T, 0, len(this))
	for item := range this {
		result = append(result, item)
	}
	return result
}
func (this Set[T]) Add(items ...T) {
	for _, item := range items {
		this[item] = nothing{}
	}
}
func (this Set[T]) Remove(items ...T) {
	for _, item := range items {
		delete(this, item)
	}
}
func (this Set[T]) Clear() {
	for item := range this {
		delete(this, item)
	}
}
func (this Set[T]) Contains(item T) bool {
	_, found := this[item]
	return found
}
func (this Set[T]) Equal(that Set[T]) bool {
	if len(this) != len(that) {
		return false
	}
	for item := range this {
		if !that.Contains(item) {
			return false
		}
	}
	return true
}
func (this Set[T]) IsSubset(that Set[T]) bool {
	for item := range this {
		if !that.Contains(item) {
			return false
		}
	}
	return true
}
func (this Set[T]) IsSuperset(that Set[T]) bool {
	for item := range that {
		if !this.Contains(item) {
			return false
		}
	}
	return true
}
func (this Set[T]) Union(that Set[T]) (result Set[T]) {
	result = make(Set[T])
	for item := range this {
		result.Add(item)
	}
	for item := range that {
		result.Add(item)
	}
	return result
}
func (this Set[T]) Intersection(that Set[T]) (result Set[T]) {
	result = make(Set[T])
	for item := range this {
		if that.Contains(item) {
			result.Add(item)
		}
	}
	for item := range that {
		if this.Contains(item) {
			result.Add(item)
		}
	}
	return result
}
func (this Set[T]) Difference(that Set[T]) (result Set[T]) {
	result = make(Set[T])
	for item := range this {
		if !that.Contains(item) {
			result.Add(item)
		}
	}
	return result
}
func (this Set[T]) SymmetricDifference(that Set[T]) (result Set[T]) {
	result = make(Set[T])
	for item := range this {
		if !that.Contains(item) {
			result.Add(item)
		}
	}
	for item := range that {
		if !this.Contains(item) {
			result.Add(item)
		}
	}
	return result
}

type nothing struct{}
