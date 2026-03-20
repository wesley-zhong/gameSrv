package quadtree

type Quadtree[T any] struct {
	Bounds     Bounds
	MaxObjects int
	MaxLevels  int
	Level      int
	Objects    []Entity[T]
	Nodes      []*Quadtree[T]
	Total      int
}

type Bounds struct {
	X      float64
	Y      float64
	Width  float64
	Height float64
}

type Entity[T any] struct {
	Bounds
	obj *T
}

func (b *Bounds) IsPoint() bool {
	return b.Width == 0 && b.Height == 0
}

func (b *Bounds) Intersects(a Bounds) bool {
	aMaxX := a.X + a.Width
	aMaxY := a.Y + a.Height
	bMaxX := b.X + b.Width
	bMaxY := b.Y + b.Height

	return !(aMaxX < b.X || a.X > bMaxX || aMaxY < b.Y || a.Y > bMaxY)
}

func (qt *Quadtree[T]) TotalNodes() int {
	total := 0
	for _, node := range qt.Nodes {
		total += 1 + node.TotalNodes()
	}
	return total
}

func (qt *Quadtree[T]) split() {
	if len(qt.Nodes) == 4 {
		return
	}

	nextLevel := qt.Level + 1
	subWidth := qt.Bounds.Width / 2
	subHeight := qt.Bounds.Height / 2
	x := qt.Bounds.X
	y := qt.Bounds.Y

	qt.Nodes = append(qt.Nodes,
		&Quadtree[T]{
			Bounds: Bounds{X: x + subWidth, Y: y, Width: subWidth, Height: subHeight},
			MaxObjects: qt.MaxObjects, MaxLevels: qt.MaxLevels,
			Level: nextLevel, Objects: make([]Entity[T], 0),
		},
		&Quadtree[T]{
			Bounds: Bounds{X: x, Y: y, Width: subWidth, Height: subHeight},
			MaxObjects: qt.MaxObjects, MaxLevels: qt.MaxLevels,
			Level: nextLevel, Objects: make([]Entity[T], 0),
		},
		&Quadtree[T]{
			Bounds: Bounds{X: x, Y: y + subHeight, Width: subWidth, Height: subHeight},
			MaxObjects: qt.MaxObjects, MaxLevels: qt.MaxLevels,
			Level: nextLevel, Objects: make([]Entity[T], 0),
		},
		&Quadtree[T]{
			Bounds: Bounds{X: x + subWidth, Y: y + subHeight, Width: subWidth, Height: subHeight},
			MaxObjects: qt.MaxObjects, MaxLevels: qt.MaxLevels,
			Level: nextLevel, Objects: make([]Entity[T], 0),
		},
	)
}

func (qt *Quadtree[T]) getIndex(pRect Bounds) int {
	verticalMidpoint := qt.Bounds.X + (qt.Bounds.Width / 2)
	horizontalMidpoint := qt.Bounds.Y + (qt.Bounds.Height / 2)

	topQuadrant := (pRect.Y < horizontalMidpoint) && (pRect.Y+pRect.Height < horizontalMidpoint)
	bottomQuadrant := pRect.Y > horizontalMidpoint

	if (pRect.X < verticalMidpoint) && (pRect.X+pRect.Width < verticalMidpoint) {
		if topQuadrant {
			return 1
		}
		if bottomQuadrant {
			return 2
		}
	} else if pRect.X > verticalMidpoint {
		if topQuadrant {
			return 0
		}
		if bottomQuadrant {
			return 3
		}
	}
	return -1
}

func (qt *Quadtree[T]) Insert(pRect Entity[T]) {
	qt.Total++

	if len(qt.Nodes) > 0 {
		index := qt.getIndex(pRect.Bounds)
		if index != -1 {
			qt.Nodes[index].Insert(pRect)
			return
		}
	}

	qt.Objects = append(qt.Objects, pRect)

	if len(qt.Objects) > qt.MaxObjects && qt.Level < qt.MaxLevels {
		if len(qt.Nodes) == 0 {
			qt.split()
		}

		for i := 0; i < len(qt.Objects); {
			index := qt.getIndex(qt.Objects[i].Bounds)
			if index != -1 {
				splice := qt.Objects[i]
				qt.Objects = append(qt.Objects[:i], qt.Objects[i+1:]...)
				qt.Nodes[index].Insert(splice)
			} else {
				i++
			}
		}
	}
}

func (qt *Quadtree[T]) Retrieve(pRect Bounds) []Entity[T] {
	index := qt.getIndex(pRect)
	result := qt.Objects

	if len(qt.Nodes) > 0 {
		if index != -1 {
			result = append(result, qt.Nodes[index].Retrieve(pRect)...)
		} else {
			for _, node := range qt.Nodes {
				result = append(result, node.Retrieve(pRect)...)
			}
		}
	}
	return result
}

func (qt *Quadtree[T]) RetrievePoints(find Bounds) []Bounds {
	var found []Bounds
	potentials := qt.Retrieve(find)
	for _, p := range potentials {
		if p.X == find.X && p.Y == find.Y && p.IsPoint() {
			found = append(found, find)
		}
	}
	return found
}

func (qt *Quadtree[T]) RetrieveIntersections(find Bounds) []Entity[T] {
	var result []Entity[T]
	potentials := qt.Retrieve(find)
	for _, p := range potentials {
		if p.Intersects(find) {
			result = append(result, p)
		}
	}
	return result
}

func (qt *Quadtree[T]) Clear() {
	qt.Objects = qt.Objects[:0]
	for _, node := range qt.Nodes {
		node.Clear()
	}
	qt.Nodes = qt.Nodes[:0]
	qt.Total = 0
}