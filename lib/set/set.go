package set

type Set map[interface{}]struct{}

func NewSet(items ...interface{}) (result Set) {
	result = make(Set)
	for _, item := range items {
		result.Add(item)
	}
	return result
}

func (this Set) Len() int {
	return len(this)
}

func (this Set) Items() (result []interface{}) {
	for item := range this {
		result = append(result, item)
	}
	return result
}

func (this Set) Add(items ...interface{}) {
	for _, item := range items {
		if !this.Contains(item) {
			this[item] = struct{}{}
		}
	}
}

func (this Set) Contains(item interface{}) bool {
	_, found := this[item]
	return found
}

func (this Set) IsSubset(that Set) bool {
	for item := range this {
		if !that.Contains(item) {
			return false
		}
	}
	return true
}

func (this Set) IsSuperset(that Set) bool {
	for item := range that {
		if !this.Contains(item) {
			return false
		}
	}
	return true
}

func (this Set) Union(that Set) (result Set) {
	result = make(Set)
	for item := range this {
		result.Add(item)
	}
	for item := range that {
		result.Add(item)
	}
	return result
}

func (this Set) Intersection(that Set) (result Set) {
	result = make(Set)
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

func (this Set) Difference(that Set) (result Set) {
	result = make(Set)
	for item := range this {
		if !that.Contains(item) {
			result.Add(item)
		}
	}
	return result
}

func (this Set) SymmetricDifference(that Set) (result Set) {
	result = make(Set)
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
