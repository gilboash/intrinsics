// THESE PACKAGES ARE FOR DEMONSTRATION PURPOSES ONLY!
//
// THEY DO NOT NOT CONTAIN WORKING INTRINSICS!
//
// See https://github.com/klauspost/intrinsics
package misc

import "github.com/klauspost/intrinsics/x86"

var _ = x86.M64{}  // Make sure we use x86 package


// AddcarryU32: Add unsigned 32-bit integers 'a' and 'b' with unsigned 8-bit
// carry-in 'c_in' (carry flag), and store the unsigned 32-bit result in 'out',
// and the carry-out in 'dst' (carry or overflow flag). 
//
//		dst:out[31:0] := a[31:0] + b[31:0] + c_in;
//
// Instruction: 'ADC'. Intrinsic: '_addcarry_u32'.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func AddcarryU32(c_in uint8, a uint32, b uint32, out *uint32) uint8 {
	panic("not implemented")
}


// AddcarryU64: Add unsigned 64-bit integers 'a' and 'b' with unsigned 8-bit
// carry-in 'c_in' (carry flag), and store the unsigned 64-bit result in 'out',
// and the carry-out in 'dst' (carry or overflow flag). 
//
//		dst:out[63:0] := a[63:0] + b[63:0] + c_in;
//
// Instruction: 'ADC'. Intrinsic: '_addcarry_u64'.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func AddcarryU64(c_in uint8, a uint64, b uint64, out *uint64) uint8 {
	panic("not implemented")
}


// AllowCpuFeatures: Treat the processor-specific feature(s) specified in 'a'
// as available. Multiple features may be OR'd together. See the valid feature
// flags below: 
//
//		_FEATURE_GENERIC_IA32
//		_FEATURE_FPU
//		_FEATURE_CMOV
//		_FEATURE_MMX
//		_FEATURE_FXSAVE
//		_FEATURE_SSE
//		_FEATURE_SSE2
//		_FEATURE_SSE3
//		_FEATURE_SSSE3
//		_FEATURE_SSE4_1
//		_FEATURE_SSE4_2
//		_FEATURE_MOVBE
//		_FEATURE_POPCNT
//		_FEATURE_PCLMULQDQ
//		_FEATURE_AES
//		_FEATURE_F16C
//		_FEATURE_AVX
//		_FEATURE_RDRND
//		_FEATURE_FMA
//		_FEATURE_BMI
//		_FEATURE_LZCNT
//		_FEATURE_HLE
//		_FEATURE_RTM
//		_FEATURE_AVX2
//		_FEATURE_KNCNI
//		_FEATURE_AVX512F
//		_FEATURE_ADX
//		_FEATURE_RDSEED
//		_FEATURE_AVX512ER
//		_FEATURE_AVX512PF
//		_FEATURE_AVX512CD
//		_FEATURE_SHA
//		_FEATURE_MPX
//
// Instruction: '...'. Intrinsic: '_allow_cpu_features'.
func AllowCpuFeatures(a uint64)  {
	panic("not implemented")
}


// BitScanForward: Set 'dst' to the index of the lowest set bit in 32-bit
// integer 'a'. If no bits are set in 'a' then 'dst' is undefined. 
//
//		tmp := 0
//		IF a = 0
//			dst := undefined
//		ELSE
//			DO WHILE ((tmp < 32) AND a[tmp] = 0)
//				tmp := tmp + 1
//				dst := tmp
//			OD
//		FI
//
// Instruction: 'BSF'. Intrinsic: '_bit_scan_forward'.
func BitScanForward(a int) int {
	panic("not implemented")
}


// BitScanReverse: Set 'dst' to the index of the highest set bit in 32-bit
// integer 'a'. If no bits are set in 'a' then 'dst' is undefined. 
//
//		tmp := 31
//		IF a = 0
//			dst := undefined
//		ELSE
//			DO WHILE ((tmp > 0) AND a[tmp] = 0)
//				tmp := tmp - 1
//				dst := tmp
//			OD
//		FI
//
// Instruction: 'BSR'. Intrinsic: '_bit_scan_reverse'.
func BitScanReverse(a int) int {
	panic("not implemented")
}


// BitScanForward2: Set 'index' to the index of the lowest set bit in 32-bit
// integer 'mask'. If no bits are set in 'mask', then set 'dst' to 0, otherwise
// set 'dst' to 1. 
//
//		tmp := 0
//		IF mask = 0
//			dst := 0
//		ELSE
//			DO WHILE ((tmp < 32) AND mask[tmp] = 0)
//				tmp := tmp + 1
//				index := tmp
//				dst := 1
//			OD
//		FI
//
// Instruction: 'BSF'. Intrinsic: '_BitScanForward'.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func BitScanForward2(index *uint32, mask uint32) uint8 {
	panic("not implemented")
}


// BitScanForward64: Set 'index' to the index of the lowest set bit in 64-bit
// integer 'mask'. If no bits are set in 'mask', then set 'dst' to 0, otherwise
// set 'dst' to 1. 
//
//		tmp := 0
//		IF mask = 0
//			dst := 0
//		ELSE
//			DO WHILE ((tmp < 64) AND mask[tmp] = 0)
//				tmp := tmp + 1
//				index := tmp
//				dst := 1
//			OD
//		FI
//
// Instruction: 'BSF'. Intrinsic: '_BitScanForward64'.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func BitScanForward64(index *uint32, mask uint64) uint8 {
	panic("not implemented")
}


// BitScanReverse2: Set 'index' to the index of the highest set bit in 32-bit
// integer 'mask'. If no bits are set in 'mask', then set 'dst' to 0, otherwise
// set 'dst' to 1. 
//
//		tmp := 31
//		IF mask = 0
//			dst := 0
//		ELSE
//			DO WHILE ((tmp > 0) AND mask[tmp] = 0)
//				tmp := tmp - 1
//				index := tmp
//				dst := 1
//			OD
//		FI
//
// Instruction: 'BSR'. Intrinsic: '_BitScanReverse'.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func BitScanReverse2(index *uint32, mask uint32) uint8 {
	panic("not implemented")
}


// BitScanReverse64: Set 'index' to the index of the highest set bit in 64-bit
// integer 'mask'. If no bits are set in 'mask', then set 'dst' to 0, otherwise
// set 'dst' to 1. 
//
//		tmp := 31
//		IF mask = 0
//			dst := 0
//		ELSE
//			DO WHILE ((tmp > 0) AND mask[tmp] = 0)
//				tmp := tmp - 1
//				index := tmp
//				dst := 1
//			OD
//		FI
//
// Instruction: 'BSR'. Intrinsic: '_BitScanReverse64'.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func BitScanReverse64(index *uint32, mask uint64) uint8 {
	panic("not implemented")
}


// Bittest: Return the bit at index 'b' of 32-bit integer 'a'. 
//
//		dst := a[b]
//
// Instruction: 'BT'. Intrinsic: '_bittest'.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Bittest(a *int32, b int32) uint8 {
	panic("not implemented")
}


// Bittest64: Return the bit at index 'b' of 64-bit integer 'a'. 
//
//		dst := a[b]
//
// Instruction: 'BT'. Intrinsic: '_bittest64'.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Bittest64(a *int64, b int64) uint8 {
	panic("not implemented")
}


// Bittestandcomplement: Return the bit at index 'b' of 32-bit integer 'a', and
// set that bit to its complement. 
//
//		dst := a[b]
//		a[b] := ~a[b]
//
// Instruction: 'BTC'. Intrinsic: '_bittestandcomplement'.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Bittestandcomplement(a *int32, b int32) uint8 {
	panic("not implemented")
}


// Bittestandcomplement64: Return the bit at index 'b' of 64-bit integer 'a',
// and set that bit to its complement. 
//
//		dst := a[b]
//		a[b] := ~a[b]
//
// Instruction: 'BTC'. Intrinsic: '_bittestandcomplement64'.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Bittestandcomplement64(a *int64, b int64) uint8 {
	panic("not implemented")
}


// Bittestandreset: Return the bit at index 'b' of 32-bit integer 'a', and set
// that bit to zero. 
//
//		dst := a[b]
//		a[b] := 0
//
// Instruction: 'BTR'. Intrinsic: '_bittestandreset'.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Bittestandreset(a *int32, b int32) uint8 {
	panic("not implemented")
}


// Bittestandreset64: Return the bit at index 'b' of 64-bit integer 'a', and
// set that bit to zero. 
//
//		dst := a[b]
//		a[b] := 0
//
// Instruction: 'BTR'. Intrinsic: '_bittestandreset64'.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Bittestandreset64(a *int64, b int64) uint8 {
	panic("not implemented")
}


// Bittestandset: Return the bit at index 'b' of 32-bit integer 'a', and set
// that bit to one. 
//
//		dst := a[b]
//		a[b] := 1
//
// Instruction: 'BTS'. Intrinsic: '_bittestandset'.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Bittestandset(a *int32, b int32) uint8 {
	panic("not implemented")
}


// Bittestandset64: Return the bit at index 'b' of 64-bit integer 'a', and set
// that bit to one. 
//
//		dst := a[b]
//		a[b] := 1
//
// Instruction: 'BTS'. Intrinsic: '_bittestandset64'.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Bittestandset64(a *int64, b int64) uint8 {
	panic("not implemented")
}


// Bswap: Reverse the byte order of 32-bit integer 'a', and store the result in
// 'dst'. This intrinsic is provided for conversion between little and big
// endian values. 
//
//		dst[7:0] := a[31:24]
//		dst[15:8] := a[23:16]
//		dst[23:16] := a[15:8]
//		dst[31:24] := a[7:0]
//
// Instruction: 'BSWAP'. Intrinsic: '_bswap'.
func Bswap(a int) int {
	panic("not implemented")
}


// Bswap64: Reverse the byte order of 64-bit integer 'a', and store the result
// in 'dst'. This intrinsic is provided for conversion between little and big
// endian values. 
//
//		dst[7:0] := a[63:56]
//		dst[15:8] := a[55:48]
//		dst[23:16] := a[47:40]
//		dst[31:24] := a[39:32]
//		dst[39:32] := a[31:24]
//		dst[47:40] := a[23:16]
//		dst[55:48] := a[15:8]
//		dst[63:56] := a[7:0]
//
// Instruction: 'BSWAP'. Intrinsic: '_bswap64'.
func Bswap64(a int64) int64 {
	panic("not implemented")
}


// Castf32U32: Cast from type float to type unsigned __int32 without
// conversion.
// 	This intrinsic is only used for compilation and does not generate any
// instructions, thus it has zero latency. 
//
//		
//
// Instruction: ''. Intrinsic: '_castf32_u32'.
func Castf32U32(a float32) uint32 {
	panic("not implemented")
}


// Castf64U64: Cast from type double to type unsigned __int64 without
// conversion.
// 	This intrinsic is only used for compilation and does not generate any
// instructions, thus it has zero latency. 
//
//		
//
// Instruction: ''. Intrinsic: '_castf64_u64'.
func Castf64U64(a float64) uint64 {
	panic("not implemented")
}


// Castu32F32: Cast from type unsigned __int32 to type float without
// conversion.
// 	This intrinsic is only used for compilation and does not generate any
// instructions, thus it has zero latency. 
//
//		
//
// Instruction: ''. Intrinsic: '_castu32_f32'.
func Castu32F32(a uint32) float32 {
	panic("not implemented")
}


// Castu64F64: Cast from type unsigned __int64 to type double without
// conversion.
// 	This intrinsic is only used for compilation and does not generate any
// instructions, thus it has zero latency. 
//
//		
//
// Instruction: ''. Intrinsic: '_castu64_f64'.
func Castu64F64(a uint64) float64 {
	panic("not implemented")
}


// CvtshSs: Convert the half-precision (16-bit) floating-point value 'a' to a
// single-precision (32-bit) floating-point value, and store the result in
// 'dst'. 
//
//		dst[31:0] := Convert_FP16_To_FP32(a[15:0])
//
// Instruction: '...'. Intrinsic: '_cvtsh_ss'.
func CvtshSs(a uint16) float32 {
	panic("not implemented")
}


// CvtssSh: Convert the single-precision (32-bit) floating-point value 'a' to a
// half-precision (16-bit) floating-point value, and store the result in 'dst'. 
//
//		dst[15:0] := Convert_FP32_To_FP16(a[31:0])
//
// Instruction: '...'. Intrinsic: '_cvtss_sh'.
//
// FIXME: Requires compiler support (has immediate)
func CvtssSh(a float32, imm8 byte) uint16 {
	panic("not implemented")
}


// Skipped: _loadbe_i16. Contains pointer parameter.


// Skipped: _loadbe_i32. Contains pointer parameter.


// Skipped: _loadbe_i64. Contains pointer parameter.


// Skipped: _mm_loadu_si16. Contains pointer parameter.


// Skipped: _mm_loadu_si32. Contains pointer parameter.


// Skipped: _mm_loadu_si64. Contains pointer parameter.


// Lrotl: Shift the bits of unsigned 64-bit integer 'a' left by the number of
// bits specified in 'shift', rotating the most-significant bit to the
// least-significant bit location, and store the unsigned result in 'dst'. 
//
//		dst := a
//		count := shift BITWISE AND 63
//		DO WHILE (count > 0)
//			tmp[0] := dst[63]
//			dst := (dst << 1) OR tmp[0]
//			count := count - 1
//		OD
//
// Instruction: 'ROL'. Intrinsic: '_lrotl'.
func Lrotl(a uint32, shift int) uint32 {
	panic("not implemented")
}


// Lrotr: Shift the bits of unsigned 64-bit integer 'a' right by the number of
// bits specified in 'shift', rotating the least-significant bit to the
// most-significant bit location, and store the unsigned result in 'dst'. 
//
//		dst := a
//		count := shift BITWISE AND 63
//		DO WHILE (count > 0)
//			tmp[63] := dst[0]
//			dst := (dst >> 1) OR tmp[63]
//			count := count - 1
//		OD
//
// Instruction: 'ROR'. Intrinsic: '_lrotr'.
func Lrotr(a uint32, shift int) uint32 {
	panic("not implemented")
}


// MayIUseCpuFeature: Dynamically query the processor to determine if the
// processor-specific feature(s) specified in 'a' are available, and return
// true or false (1 or 0) if the set of features is available. Multiple
// features may be OR'd together. This intrinsic does not check the processor
// vendor. See the valid feature flags below: 
//
//		_FEATURE_GENERIC_IA32
//		_FEATURE_FPU
//		_FEATURE_CMOV
//		_FEATURE_MMX
//		_FEATURE_FXSAVE
//		_FEATURE_SSE
//		_FEATURE_SSE2
//		_FEATURE_SSE3
//		_FEATURE_SSSE3
//		_FEATURE_SSE4_1
//		_FEATURE_SSE4_2
//		_FEATURE_MOVBE
//		_FEATURE_POPCNT
//		_FEATURE_PCLMULQDQ
//		_FEATURE_AES
//		_FEATURE_F16C
//		_FEATURE_AVX
//		_FEATURE_RDRND
//		_FEATURE_FMA
//		_FEATURE_BMI
//		_FEATURE_LZCNT
//		_FEATURE_HLE
//		_FEATURE_RTM
//		_FEATURE_AVX2
//		_FEATURE_KNCNI
//		_FEATURE_AVX512F
//		_FEATURE_ADX
//		_FEATURE_RDSEED
//		_FEATURE_AVX512ER
//		_FEATURE_AVX512PF
//		_FEATURE_AVX512CD
//		_FEATURE_SHA
//		_FEATURE_MPX
//
// Instruction: '...'. Intrinsic: '_may_i_use_cpu_feature'.
func MayIUseCpuFeature(a uint64) int {
	panic("not implemented")
}


// Rdpmc: Read the Performance Monitor Counter (PMC) specified by 'a', and
// store up to 64-bits in 'dst'. The width of performance counters is
// implementation specific. 
//
//		dst[63:0] := ReadPMC(a)
//
// Instruction: 'RDPMC'. Intrinsic: '_rdpmc'.
func Rdpmc(a int) int64 {
	panic("not implemented")
}


// Rotl: Shift the bits of unsigned 32-bit integer 'a' left by the number of
// bits specified in 'shift', rotating the most-significant bit to the
// least-significant bit location, and store the unsigned result in 'dst'. 
//
//		dst := a
//		count := shift BITWISE AND 31
//		DO WHILE (count > 0)
//			tmp[0] := dst[31]
//			dst := (dst << 1) OR tmp[0]
//			count := count - 1
//		OD
//
// Instruction: 'ROL'. Intrinsic: '_rotl'.
func Rotl(a uint32, shift int) uint32 {
	panic("not implemented")
}


// Rotr: Shift the bits of unsigned 32-bit integer 'a' right by the number of
// bits specified in 'shift', rotating the least-significant bit to the
// most-significant bit location, and store the unsigned result in 'dst'. 
//
//		dst := a
//		count := shift BITWISE AND 31
//		DO WHILE (count > 0)
//			tmp[31] := dst[0]
//			dst := (dst >> 1) OR tmp[31]
//			count := count - 1
//		OD
//
// Instruction: 'ROR'. Intrinsic: '_rotr'.
func Rotr(a uint32, shift int) uint32 {
	panic("not implemented")
}


// Rotwl: Shift the bits of unsigned 16-bit integer 'a' left by the number of
// bits specified in 'shift', rotating the most-significant bit to the
// least-significant bit location, and store the unsigned result in 'dst'. 
//
//		dst := a
//		count := shift BITWISE AND 15
//		DO WHILE (count > 0)
//			tmp[0] := dst[15]
//			dst := (dst << 1) OR tmp[0]
//			count := count - 1
//		OD
//
// Instruction: 'ROL'. Intrinsic: '_rotwl'.
func Rotwl(a uint16, shift int) uint16 {
	panic("not implemented")
}


// Rotwr: Shift the bits of unsigned 16-bit integer 'a' right by the number of
// bits specified in 'shift', rotating the least-significant bit to the
// most-significant bit location, and store the unsigned result in 'dst'. 
//
//		dst := a
//		count := shift BITWISE AND 15
//		DO WHILE (count > 0)
//			tmp[15] := dst[0]
//			dst := (dst >> 1) OR tmp[15]
//			count := count - 1
//		OD
//
// Instruction: 'ROR'. Intrinsic: '_rotwr'.
func Rotwr(a uint16, shift int) uint16 {
	panic("not implemented")
}


// Skipped: _storebe_i16. Contains pointer parameter.


// Skipped: _storebe_i32. Contains pointer parameter.


// Skipped: _storebe_i64. Contains pointer parameter.


// Skipped: _mm_storeu_si16. Contains pointer parameter.


// Skipped: _mm_storeu_si32. Contains pointer parameter.


// Skipped: _mm_storeu_si64. Contains pointer parameter.


// SubborrowU32: Add unsigned 8-bit borrow 'b_in' (carry flag) to unsigned
// 32-bit integer 'a', and subtract the result from unsigned 32-bit integer
// 'b'. Store the unsigned 32-bit result in 'out', and the carry-out in 'dst'
// (carry or overflow flag). 
//
//		dst:out[31:0] := (b[31:0] - (a[31:0] + b_in));
//
// Instruction: 'SBB'. Intrinsic: '_subborrow_u32'.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func SubborrowU32(b_in uint8, a uint32, b uint32, out *uint32) uint8 {
	panic("not implemented")
}


// SubborrowU64: Add unsigned 8-bit borrow 'b_in' (carry flag) to unsigned
// 64-bit integer 'a', and subtract the result from unsigned 64-bit integer
// 'b'. Store the unsigned 64-bit result in 'out', and the carry-out in 'dst'
// (carry or overflow flag). 
//
//		dst:out[63:0] := (b[63:0] - (a[63:0] + b_in));
//
// Instruction: 'SBB'. Intrinsic: '_subborrow_u64'.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func SubborrowU64(b_in uint8, a uint64, b uint64, out *uint64) uint8 {
	panic("not implemented")
}

