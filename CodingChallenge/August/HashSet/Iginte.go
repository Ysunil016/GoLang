package main

func main() {

}

// MyHashSet ...
type MyHashSet struct {
	values []bool
}

// Constructor ...
func Constructor() MyHashSet {

}

// Add ...
func (this *MyHashSet) Add(key int) {
	if key >= len(this.values) {
		value := make([]bool, key+2)
		value = append(value, this.values...)
		this.values = value
	}
	this.values[key] = true
}

// Remove ...
func (this *MyHashSet) Remove(key int) {
	if key < len(this.values) {
		this.values[key] = false
	}
}

func (this *MyHashSet) Contains(key int) bool {
	if key >= len(this.values) {
		return false
	}
	return this.values[key]
}
