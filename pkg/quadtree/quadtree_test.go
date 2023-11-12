package quadtree

import "testing"

type MapUnit struct {
	Uid int64
	Bounds
}

func Test_add_test(t *testing.T) {
	qt := Quadtree[MapUnit]{
		Bounds:     Bounds{0, 0, 307200, 307200},
		MaxObjects: 128,
		MaxLevels:  10,
		Level:      1,
		Objects:    make([]Entity[MapUnit], 128),
		Nodes:      nil,
		Total:      0,
	}
	mapUnit := &MapUnit{
		Uid:    0,
		Bounds: Bounds{100, 100, 20, 30},
	}
	entity := Entity[MapUnit]{
		obj:    mapUnit,
		Bounds: mapUnit.Bounds,
	}
	qt.Insert(entity)
	if qt.Total != 1 {
		t.Errorf("expect %d but =%d", 1, qt.Total)
	}
}
