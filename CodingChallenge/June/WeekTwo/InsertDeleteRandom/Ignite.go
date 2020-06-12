package main

import (
	"fmt"
	"math/rand"
)

func main() {
	RandomSet := Constructor()
	RandomSet.Insert(0)
	RandomSet.Insert(2)
	RandomSet.Insert(3)
	RandomSet.Insert(4)
	fmt.Println(RandomSet.GetRandom())
	RandomSet.Remove(0)

}

/*
Insert Delete GetRandom O(1)

Solution
Design a data structure that supports all following operations in average O(1) time.

insert(val): Inserts an item val to the set if not already present.
remove(val): Removes an item val from the set if present.
getRandom: Returns a random element from current set of elements. Each element must have the same probability of being returned.
*/

//RandomizedSet ...
type RandomizedSet struct {
	Hash map[int]int
	List []int
}

// Constructor ...
/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{Hash: make(map[int]int), List: make([]int, 0)}
}

// Insert ...
/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (r *RandomizedSet) Insert(val int) bool {
	if _, isthere := r.Hash[val]; isthere {
		return false
	}
	r.Hash[val] = len(r.List)
	r.List = append(r.List, val)
	return true
}

// Remove ...
/** Removes a value from the set. Returns true if the set contained the specified element. */
func (r *RandomizedSet) Remove(val int) bool {
	if _, isthere := r.Hash[val]; !isthere {
		return false
	}
	index := r.Hash[val]
	lastElement := r.List[len(r.List)-1]
	r.List[index] = lastElement
	r.List = r.List[:len(r.List)-1]
	r.Hash[lastElement] = index
	delete(r.Hash, val)
	return true
}

// GetRandom ...
/** Get a random element from the set. */
func (r *RandomizedSet) GetRandom() int {
	return r.List[(rand.Intn(len(r.List)))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
