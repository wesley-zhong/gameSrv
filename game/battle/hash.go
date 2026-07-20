package battle

// SparseArrayHasher provides hash computation for sparse arrays
type SparseArrayHasher struct{}

var globalHasher = SparseArrayHasher{}

// PositionPrimes are position-sensitive primes for hashing
var positionPrimes = []int{
	2, 3, 5, 7, 11, 13, 17, 23, 29, 31, 37, 41, 43, 47,
	53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109,
	113, 127, 131, 137, 139, 149, 151, 157, 163, 167, 173, 179, 181, 191,
	193, 197, 199, 211, 223, 227, 229, 233, 239, 241, 251, 257, 263, 269,
	271, 277, 281, 283, 293, 307, 311, 313, 317, 331, 337, 347, 349, 353,
	359, 367, 373, 379, 383, 389, 397, 401, 409, 419, 421, 431, 433, 439,
	443, 449, 457, 461, 463, 467, 479, 487, 491, 499, 503, 509, 521, 523,
	541, 547, 557, 563, 569, 571, 577, 587, 593, 599, 601, 607, 613, 617,
	619, 631, 641, 643, 647, 653, 659, 661, 673, 677, 683, 691, 701, 709,
	719, 727, 733, 739, 743, 751, 757, 761, 769, 773, 787, 797, 809, 811,
	821, 823, 827, 829, 839, 853, 857, 859, 863, 877, 881, 883, 887, 907,
	911, 919, 929, 937, 941, 947, 953, 967, 971, 977, 983, 991, 997, 1009,
}

// processNonZeroValues processes non-zero values with position weights
func processNonZeroValues(data []int, hash int64, count int) int64 {
	for i := 0; i < count; i++ {
		if data[i] != 0 {
			pos := 691
			if i < len(positionPrimes) {
				pos = positionPrimes[i]
			}
			// Position encoding: element value × position prime × magic
			elementHash := int64(data[i])
			elementHash *= int64(pos)
			magic := []int64{
				0x243F6A8885A308D3, 0x13198A2E03707344,
				0x04093822299F31D0, 0x082EFA98EC4E6C89,
			}
			elementHash ^= magic[i%4]

			// Spiral mixing
			hash = (hash << 17) ^ (hash >> 47) ^ elementHash
			// Use smaller multiplier to avoid overflow
			hash += (elementHash % 1000000) * 0x9E3779B7 / 1000
		}
	}
	return hash
}

// ComputeHash computes hash for sparse array data
func (h *SparseArrayHasher) ComputeHash(data []int, count int) int64 {
	hash := int64(0x0CBF29CE48) // Use smaller initial value to avoid overflow

	// Phase 1: Process non-zero values
	hash = processNonZeroValues(data, hash, count)

	// Phase 2: Process zero value distribution pattern (commented in original)
	// Phase 3: Enhanced mixing (commented in original)
	return hash
}

// ComputeHash is the global function for hash computation
func ComputeHash(data []int, count int) int64 {
	return globalHasher.ComputeHash(data, count)
}