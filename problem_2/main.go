package main

import (
	"sort"
)

type Item struct {
	Name  string
	Count int64
}

// func for sorting items
type ByCount []Item

func (a ByCount) Len() int           { return len(a) }
func (a ByCount) Less(i, j int) bool { return a[i].Count < a[j].Count }
func (a ByCount) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

type Box struct {
	Request []Item
}

func (b Box) TotalBox() (res int64) {
	// define variable to calculate box
	var a, devide, start int64

	// should be a request, if empty then return 0
	if len(b.Request) == 0 {
		return
	} else {
		// define minimum box
		res = 1
	}

	// sort the box by item's count
	// because maximum box that ainun can make is the minimum item's count
	sort.Sort(ByCount(b.Request))

	// x loop is for every type of item
	for x := 0; x < len(b.Request); x++ {
		count := b.Request[x].Count
		// 'a' loop is for search maximum item's count that can fit in every items
		for a = count - start; a >= 1; a-- {
			// when x is 0, or its mean the smallest / minimum item's count
			// smallest item's count always become a benchmark for other items
			if x == 0 {
				// get devider for other items
				if count%a == 0 {
					devide = a
				}
				break
			}

			// check if the rest of items is can be devide by devider
			if count%devide == 0 {
				// if every items is fit with the devider
				res = devide
			} else {
				// if an item is not fit with the devider, then go back to find devider by set x to 0, and start from next devider
				x = -1
				start++
				break
			}
		}
	}

	return
}

func (b Box) TotalItemInBox() (res []Item) {
	// should be a request, if empty then return empty item
	if len(b.Request) == 0 {
		return
	}

	// get maximum box
	totalBox := b.TotalBox()

	// devide every item's count with maximum box.
	res = b.Request
	for i, v := range res {
		res[i].Count = v.Count / totalBox
	}

	return
}
