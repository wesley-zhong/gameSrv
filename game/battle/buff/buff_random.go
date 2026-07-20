package buff

import (
	"gameSrv/pkg/utils"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// BuffEffectRandom handles random execution of buff effects with pity system
type BuffEffectRandom struct {
	confCount       int
	curCount        int
	haveSuccessed   bool
	targetRandomNum int
}

// NewBuffEffectRandom creates a new BuffEffectRandom
func NewBuffEffectRandom(confCount int) *BuffEffectRandom {
	targetRandomNum := 100 / confCount
	if targetRandomNum == 0 {
		targetRandomNum = 1
	}
	return &BuffEffectRandom{
		confCount:       confCount,
		curCount:        0,
		haveSuccessed:   false,
		targetRandomNum: targetRandomNum,
	}
}

// reset resets the random state
func (r *BuffEffectRandom) reset() {
	r.curCount = 0
	r.haveSuccessed = false
}

// Random performs a random check with pity system
func (r *BuffEffectRandom) Random() bool {
	r.curCount++

	// Pity system - guaranteed success on last attempt
	if r.curCount == r.confCount {
		if r.haveSuccessed {
			r.reset()
			return false
		}
		r.reset()
		return true
	}

	// Normal random check
	iRandomNum := rand.Intn(100)
	if iRandomNum < r.targetRandomNum {
		r.haveSuccessed = true
		return true
	}
	return false
}

// RandomRange generates a random number in range [min, max)
func RandomRange(min, max int) int {
	if min >= max {
		return min
	}
	return min + rand.Intn(max-min)
}

// GenProcessLongId generates a new unique long ID
func GenProcessLongId() int64 {
	return utils.NextId()
}
