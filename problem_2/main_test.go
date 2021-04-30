package main

import (
	"testing"
)

var (
	box Box
)

func TestBoxCase1(t *testing.T) {
	box.Request = []Item{
		Item{
			Name:  "cakes",
			Count: 20,
		},
		Item{
			Name:  "apples",
			Count: 25,
		},
	}
	t.Logf("Request : %v", box.Request)
	totalBox := box.TotalBox()
	t.Logf("Total box : %d", totalBox)
	var expectedresult int64 = 5
	if totalBox != expectedresult {
		t.Errorf("WRONG!! should be : %d", expectedresult)
	}

	items := box.TotalItemInBox()
	t.Logf("Total item in box : %v", items)
	expectedItemsResult := []Item{
		Item{
			Name:  "cakes",
			Count: 4,
		},
		Item{
			Name:  "apples",
			Count: 5,
		},
	}
	for i, v := range items {
		if v.Count != expectedItemsResult[i].Count {
			t.Errorf("WRONG!! %s should be : %d", v.Name, expectedItemsResult[i].Count)
		}
	}
}

func TestBoxCase2(t *testing.T) {
	box.Request = []Item{
		Item{
			Name:  "cakes",
			Count: 20,
		},
		Item{
			Name:  "apples",
			Count: 25,
		},
		Item{
			Name:  "bananas",
			Count: 35,
		},
	}
	t.Logf("Request : %v", box.Request)
	totalBox := box.TotalBox()
	t.Logf("Total box : %d", totalBox)
	var expectedresult int64 = 5
	if totalBox != expectedresult {
		t.Errorf("WRONG!! should be : %d", expectedresult)
	}

	items := box.TotalItemInBox()
	t.Logf("Total item in box : %v", items)
	expectedItemsResult := []Item{
		Item{
			Name:  "cakes",
			Count: 4,
		},
		Item{
			Name:  "apples",
			Count: 5,
		},
		Item{
			Name:  "bananas",
			Count: 7,
		},
	}
	for i, v := range items {
		if v.Count != expectedItemsResult[i].Count {
			t.Errorf("WRONG!! %s should be : %d", v.Name, expectedItemsResult[i].Count)
		}
	}
}

func TestBoxCase3(t *testing.T) {
	box.Request = []Item{
		Item{
			Name:  "cakes",
			Count: 4,
		},
		Item{
			Name:  "apples",
			Count: 6,
		},
		Item{
			Name:  "bananas",
			Count: 2,
		},
	}
	t.Logf("Request : %v", box.Request)
	totalBox := box.TotalBox()
	t.Logf("Total box : %d", totalBox)
	var expectedresult int64 = 2
	if totalBox != expectedresult {
		t.Errorf("WRONG!! should be : %d", expectedresult)
	}

	items := box.TotalItemInBox()
	t.Logf("Total item in box : %v", items)
	expectedItemsResult := []Item{
		Item{
			Name:  "bananas",
			Count: 1,
		},
		Item{
			Name:  "cakes",
			Count: 2,
		},
		Item{
			Name:  "apples",
			Count: 3,
		},
	}
	for i, v := range items {
		if v.Count != expectedItemsResult[i].Count {
			t.Errorf("WRONG!! %s should be : %d", v.Name, expectedItemsResult[i].Count)
		}
	}
}

func TestBoxCase4(t *testing.T) {
	box.Request = []Item{
		Item{
			Name:  "cakes",
			Count: 33,
		},
		Item{
			Name:  "apples",
			Count: 15,
		},
		Item{
			Name:  "bananas",
			Count: 9,
		},
	}
	t.Logf("Request : %v", box.Request)
	totalBox := box.TotalBox()
	t.Logf("Total box : %d", totalBox)
	var expectedresult int64 = 3
	if totalBox != expectedresult {
		t.Errorf("WRONG!! should be : %d", expectedresult)
	}

	items := box.TotalItemInBox()
	t.Logf("Total item in box : %v", items)
	expectedItemsResult := []Item{
		Item{
			Name:  "bananas",
			Count: 3,
		},
		Item{
			Name:  "apples",
			Count: 5,
		},
		Item{
			Name:  "cakes",
			Count: 11,
		},
	}
	for i, v := range items {
		if v.Count != expectedItemsResult[i].Count {
			t.Errorf("WRONG!! %s should be : %d", v.Name, expectedItemsResult[i].Count)
		}
	}
}

func TestBoxCase5(t *testing.T) {
	box.Request = []Item{
		Item{
			Name:  "cakes",
			Count: 33,
		},
		Item{
			Name:  "apples",
			Count: 11,
		},
	}
	t.Logf("Request : %v", box.Request)
	totalBox := box.TotalBox()
	t.Logf("Total box : %d", totalBox)
	var expectedresult int64 = 11
	if totalBox != expectedresult {
		t.Errorf("WRONG!! should be : %d", expectedresult)
	}

	items := box.TotalItemInBox()
	t.Logf("Total item in box : %v", items)
	expectedItemsResult := []Item{
		Item{
			Name:  "apples",
			Count: 1,
		},
		Item{
			Name:  "cakes",
			Count: 3,
		},
	}
	for i, v := range items {
		if v.Count != expectedItemsResult[i].Count {
			t.Errorf("WRONG!! %s should be : %d", v.Name, expectedItemsResult[i].Count)
		}
	}
}

func TestBoxCase6(t *testing.T) {
	box.Request = []Item{
		Item{
			Name:  "cakes",
			Count: 13,
		},
		Item{
			Name:  "apples",
			Count: 2,
		},
		Item{
			Name:  "bananas",
			Count: 15,
		},
	}
	t.Logf("Request : %v", box.Request)
	totalBox := box.TotalBox()
	t.Logf("Total box : %d", totalBox)
	var expectedresult int64 = 1
	if totalBox != expectedresult {
		t.Errorf("WRONG!! should be : %d", expectedresult)
	}

	items := box.TotalItemInBox()
	t.Logf("Total item in box : %v", items)
	expectedItemsResult := []Item{
		Item{
			Name:  "apples",
			Count: 2,
		},
		Item{
			Name:  "cakes",
			Count: 13,
		},
		Item{
			Name:  "bananas",
			Count: 15,
		},
	}
	for i, v := range items {
		if v.Count != expectedItemsResult[i].Count {
			t.Errorf("WRONG!! %s should be : %d", v.Name, expectedItemsResult[i].Count)
		}
	}
}

func TestBoxCase7(t *testing.T) {
	box.Request = []Item{
		Item{
			Name:  "cakes",
			Count: 10,
		},
		Item{
			Name:  "apples",
			Count: 5,
		},
		Item{
			Name:  "bananas",
			Count: 45,
		},
		Item{
			Name:  "candy",
			Count: 15,
		},
	}
	t.Logf("Request : %v", box.Request)
	totalBox := box.TotalBox()
	t.Logf("Total box : %d", totalBox)
	var expectedresult int64 = 5
	if totalBox != expectedresult {
		t.Errorf("WRONG!! should be : %d", expectedresult)
	}

	items := box.TotalItemInBox()
	t.Logf("Total item in box : %v", items)
	expectedItemsResult := []Item{
		Item{
			Name:  "apples",
			Count: 1,
		},
		Item{
			Name:  "cakes",
			Count: 2,
		},
		Item{
			Name:  "candy",
			Count: 3,
		},
		Item{
			Name:  "bananas",
			Count: 9,
		},
	}
	for i, v := range items {
		if v.Count != expectedItemsResult[i].Count {
			t.Errorf("WRONG!! %s should be : %d", v.Name, expectedItemsResult[i].Count)
		}
	}
}
