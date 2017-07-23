// THESE PACKAGES ARE FOR DEMONSTRATION PURPOSES ONLY!
//
// THEY DO NOT NOT CONTAIN WORKING INTRINSICS!
//
// See https://github.com/klauspost/intrinsics
package lzcnt

import "github.com/klauspost/intrinsics/x86"

var _ = x86.M64{}  // Make sure we use x86 package


// LzcntU32: Count the number of leading zero bits in unsigned 32-bit integer
// 'a', and return that count in 'dst'. 
//
//		tmp := 31
//		dst := 0
//		DO WHILE (tmp >= 0 AND a[tmp] = 0)
//			tmp := tmp - 1
//			dst := dst + 1
//		OD
//
// Instruction: 'LZCNT'. Intrinsic: '_lzcnt_u32'.
// Requires LZCNT.
func LzcntU32(a uint32) uint32 {
	panic("not implemented")
}


// LzcntU64: Count the number of leading zero bits in unsigned 64-bit integer
// 'a', and return that count in 'dst'. 
//
//		tmp := 63
//		dst := 0
//		DO WHILE (tmp >= 0 AND a[tmp] = 0)
//			tmp := tmp - 1
//			dst := dst + 1
//		OD
//
// Instruction: 'LZCNT'. Intrinsic: '_lzcnt_u64'.
// Requires LZCNT.
func LzcntU64(a uint64) uint64 {
	panic("not implemented")
}

