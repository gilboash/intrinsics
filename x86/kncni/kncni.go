// THESE PACKAGES ARE FOR DEMONSTRATION PURPOSES ONLY!
//
// THEY DO NOT NOT CONTAIN WORKING INTRINSICS!
//
// See https://github.com/klauspost/intrinsics
package kncni

import "github.com/klauspost/intrinsics/x86"

var _ = x86.M64{}  // Make sure we use x86 package


// M512AbsPd: Finds the absolute value of each packed double-precision (64-bit)
// floating-point element in 'v2', storing the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*64
//			dst[i+63:i] := ABS(v2[i+63:i])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPANDQ'. Intrinsic: '_mm512_abs_pd'.
// Requires KNCNI.
func M512AbsPd(v2 x86.M512d) (dst x86.M512d) {
	panic("not implemented")
}


// M512MaskAbsPd: Finds the absolute value of each packed double-precision
// (64-bit) floating-point element in 'v2', storing the results in 'dst' using
// writemask 'k' (elements are copied from 'src' when the corresponding mask
// bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := ABS(v2[i+63:i])
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPANDQ'. Intrinsic: '_mm512_mask_abs_pd'.
// Requires KNCNI.
func M512MaskAbsPd(src x86.M512d, k x86.Mmask8, v2 x86.M512d) (dst x86.M512d) {
	panic("not implemented")
}


// M512AbsPs: Finds the absolute value of each packed single-precision (32-bit)
// floating-point element in 'v2', storing the results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := ABS(v2[i+31:i])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPANDD'. Intrinsic: '_mm512_abs_ps'.
// Requires KNCNI.
func M512AbsPs(v2 x86.M512) (dst x86.M512) {
	panic("not implemented")
}


// M512MaskAbsPs: Finds the absolute value of each packed single-precision
// (32-bit) floating-point element in 'v2', storing the results in 'dst' using
// writemask 'k' (elements are copied from 'src' when the corresponding mask
// bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := ABS(v2[i+31:i])
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPANDD'. Intrinsic: '_mm512_mask_abs_ps'.
// Requires KNCNI.
func M512MaskAbsPs(src x86.M512, k x86.Mmask16, v2 x86.M512) (dst x86.M512) {
	panic("not implemented")
}


// M512AdcEpi32: Performs element-by-element addition of packed 32-bit integers
// in 'v2' and 'v3' and the corresponding bit in 'k2', storing the result of
// the addition in 'dst' and the result of the carry in 'k2_res'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			k2_res[j]   := Carry(v2[i+31:i] + v3[i+31:i] + k2[j])
//			dst[i+31:i] := v2[i+31:i] + v3[i+31:i] + k2[j]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPADCD'. Intrinsic: '_mm512_adc_epi32'.
// Requires KNCNI.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func M512AdcEpi32(v2 x86.M512i, k2 x86.Mmask16, v3 x86.M512i, k2_res *x86.Mmask16) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskAdcEpi32: Performs element-by-element addition of packed 32-bit
// integers in 'v2' and 'v3' and the corresponding bit in 'k2', storing the
// result of the addition in 'dst' and the result of the carry in 'k2_res'
// using writemask 'k1' (elements are copied from 'v2' when the corresponding
// mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k1[j]
//				k2_res[j]   := Carry(v2[i+31:i] + v3[i+31:i] + k2[j])
//				dst[i+31:i] := v2[i+31:i] + v3[i+31:i] + k2[j]
//			ELSE
//				dst[i+31:i] := v2[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPADCD'. Intrinsic: '_mm512_mask_adc_epi32'.
// Requires KNCNI.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func M512MaskAdcEpi32(v2 x86.M512i, k1 x86.Mmask16, k2 x86.Mmask16, v3 x86.M512i, k2_res *x86.Mmask16) (dst x86.M512i) {
	panic("not implemented")
}


// M512AddEpi32: Add packed 32-bit integers in 'a' and 'b', and store the
// results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := a[i+31:i] + b[i+31:i]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPADDD'. Intrinsic: '_mm512_add_epi32'.
// Requires KNCNI.
func M512AddEpi32(a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskAddEpi32: Add packed 32-bit integers in 'a' and 'b', and store the
// results in 'dst' using writemask 'k' (elements are copied from 'src' when
// the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := a[i+31:i] + b[i+31:i]
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPADDD'. Intrinsic: '_mm512_mask_add_epi32'.
// Requires KNCNI.
func M512MaskAddEpi32(src x86.M512i, k x86.Mmask16, a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512AddPd: Add packed double-precision (64-bit) floating-point elements in
// 'a' and 'b', and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*64
//			dst[i+63:i] := a[i+63:i] + b[i+63:i]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VADDPD'. Intrinsic: '_mm512_add_pd'.
// Requires KNCNI.
func M512AddPd(a x86.M512d, b x86.M512d) (dst x86.M512d) {
	panic("not implemented")
}


// M512MaskAddPd: Add packed double-precision (64-bit) floating-point elements
// in 'a' and 'b', and store the results in 'dst' using writemask 'k' (elements
// are copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := a[i+63:i] + b[i+63:i]
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VADDPD'. Intrinsic: '_mm512_mask_add_pd'.
// Requires KNCNI.
func M512MaskAddPd(src x86.M512d, k x86.Mmask8, a x86.M512d, b x86.M512d) (dst x86.M512d) {
	panic("not implemented")
}


// M512AddPs: Add packed single-precision (32-bit) floating-point elements in
// 'a' and 'b', and store the results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := a[i+31:i] + b[i+31:i]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VADDPS'. Intrinsic: '_mm512_add_ps'.
// Requires KNCNI.
func M512AddPs(a x86.M512, b x86.M512) (dst x86.M512) {
	panic("not implemented")
}


// M512MaskAddPs: Add packed single-precision (32-bit) floating-point elements
// in 'a' and 'b', and store the results in 'dst' using writemask 'k' (elements
// are copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := a[i+31:i] + b[i+31:i]
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VADDPS'. Intrinsic: '_mm512_mask_add_ps'.
// Requires KNCNI.
func M512MaskAddPs(src x86.M512, k x86.Mmask16, a x86.M512, b x86.M512) (dst x86.M512) {
	panic("not implemented")
}


// M512AddRoundPd: Add packed double-precision (64-bit) floating-point elements
// in 'a' and 'b', and store the results in 'dst'.
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 7
//			i := j*64
//			dst[i+63:i] := a[i+63:i] + b[i+63:i]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VADDPD'. Intrinsic: '_mm512_add_round_pd'.
// Requires KNCNI.
func M512AddRoundPd(a x86.M512d, b x86.M512d, rounding int) (dst x86.M512d) {
	panic("not implemented")
}


// M512MaskAddRoundPd: Add packed double-precision (64-bit) floating-point
// elements in 'a' and 'b', and store the results in 'dst' using writemask 'k'
// (elements are copied from 'src' when the corresponding mask bit is not set).
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := a[i+63:i] + b[i+63:i]
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VADDPD'. Intrinsic: '_mm512_mask_add_round_pd'.
// Requires KNCNI.
func M512MaskAddRoundPd(src x86.M512d, k x86.Mmask8, a x86.M512d, b x86.M512d, rounding int) (dst x86.M512d) {
	panic("not implemented")
}


// M512AddRoundPs: Add packed single-precision (32-bit) floating-point elements
// in 'a' and 'b', and store the results in 'dst'.
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := a[i+31:i] + b[i+31:i]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VADDPS'. Intrinsic: '_mm512_add_round_ps'.
// Requires KNCNI.
func M512AddRoundPs(a x86.M512, b x86.M512, rounding int) (dst x86.M512) {
	panic("not implemented")
}


// M512MaskAddRoundPs: Add packed single-precision (32-bit) floating-point
// elements in 'a' and 'b', and store the results in 'dst' using writemask 'k'
// (elements are copied from 'src' when the corresponding mask bit is not set).
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := a[i+31:i] + b[i+31:i]
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VADDPS'. Intrinsic: '_mm512_mask_add_round_ps'.
// Requires KNCNI.
func M512MaskAddRoundPs(src x86.M512, k x86.Mmask16, a x86.M512, b x86.M512, rounding int) (dst x86.M512) {
	panic("not implemented")
}


// M512AddnPd: Performs element-by-element addition between packed
// double-precision (64-bit) floating-point elements in 'v2' and 'v3' and
// negates their sum, storing the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*64
//			dst[i+63:i] := -(v2[i+63:i] + v3[i+63:i])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VADDNPD'. Intrinsic: '_mm512_addn_pd'.
// Requires KNCNI.
func M512AddnPd(v2 x86.M512d, v3 x86.M512d) (dst x86.M512d) {
	panic("not implemented")
}


// M512MaskAddnPd: Performs element-by-element addition between packed
// double-precision (64-bit) floating-point elements in 'v2' and 'v3' and
// negates their sum, storing the results in 'dst' using writemask 'k'
// (elements are copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := -(v2[i+63:i] + v3[i+63:i])
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VADDNPD'. Intrinsic: '_mm512_mask_addn_pd'.
// Requires KNCNI.
func M512MaskAddnPd(src x86.M512d, k x86.Mmask8, v2 x86.M512d, v3 x86.M512d) (dst x86.M512d) {
	panic("not implemented")
}


// M512AddnPs: Performs element-by-element addition between packed
// single-precision (32-bit) floating-point elements in 'v2' and 'v3' and
// negates their sum, storing the results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := -(v2[i+31:i] + v3[i+31:i])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VADDNPS'. Intrinsic: '_mm512_addn_ps'.
// Requires KNCNI.
func M512AddnPs(v2 x86.M512, v3 x86.M512) (dst x86.M512) {
	panic("not implemented")
}


// M512MaskAddnPs: Performs element-by-element addition between packed
// single-precision (32-bit) floating-point elements in 'v2' and 'v3' and
// negates their sum, storing the results in 'dst' using writemask 'k'
// (elements are copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := -(v2[i+31:i] + v3[i+31:i])
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VADDNPS'. Intrinsic: '_mm512_mask_addn_ps'.
// Requires KNCNI.
func M512MaskAddnPs(src x86.M512, k x86.Mmask16, v2 x86.M512, v3 x86.M512) (dst x86.M512) {
	panic("not implemented")
}


// M512AddnRoundPd: Performs element by element addition between packed
// double-precision (64-bit) floating-point elements in 'v2' and 'v3' and
// negates the sum, storing the result in 'dst'.
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 7
//			i := j*64
//			dst[i+63:i] := -(v2[i+63:i] + v3[i+63:i])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VADDNPD'. Intrinsic: '_mm512_addn_round_pd'.
// Requires KNCNI.
func M512AddnRoundPd(v2 x86.M512d, v3 x86.M512d, rounding int) (dst x86.M512d) {
	panic("not implemented")
}


// M512MaskAddnRoundPd: Performs element by element addition between packed
// double-precision (64-bit) floating-point elements in 'v2' and 'v3' and
// negates the sum, storing the result in 'dst' using writemask 'k' (elements
// are copied from 'src' when the corresponding mask bit is not set).
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := -(v2[i+63:i] + v3[i+63:i])
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VADDNPD'. Intrinsic: '_mm512_mask_addn_round_pd'.
// Requires KNCNI.
func M512MaskAddnRoundPd(src x86.M512d, k x86.Mmask8, v2 x86.M512d, v3 x86.M512d, rounding int) (dst x86.M512d) {
	panic("not implemented")
}


// M512AddnRoundPs: Performs element by element addition between packed
// single-precision (32-bit) floating-point elements in 'v2' and 'v3' and
// negates the sum, storing the result in 'dst'.
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := -(v2[i+31:i] + v3[i+31:i])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VADDNPS'. Intrinsic: '_mm512_addn_round_ps'.
// Requires KNCNI.
func M512AddnRoundPs(v2 x86.M512, v3 x86.M512, rounding int) (dst x86.M512) {
	panic("not implemented")
}


// M512MaskAddnRoundPs: Performs element by element addition between packed
// single-precision (32-bit) floating-point elements in 'v2' and 'v3' and
// negates the sum, storing the result in 'dst' using writemask 'k' (elements
// are copied from 'src' when the corresponding mask bit is not set).
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := -(v2[i+31:i] + v3[i+31:i])
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VADDNPS'. Intrinsic: '_mm512_mask_addn_round_ps'.
// Requires KNCNI.
func M512MaskAddnRoundPs(src x86.M512, k x86.Mmask16, v2 x86.M512, v3 x86.M512, rounding int) (dst x86.M512) {
	panic("not implemented")
}


// M512AddsetcEpi32: Performs element-by-element addition of packed 32-bit
// integer elements in 'v2' and 'v3', storing the resultant carry in 'k2_res'
// (carry flag) and the addition results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := v2[i+31:i] + v3[i+31:i]
//			k2_res[j] := Carry(v2[i+31:i] + v3[i+31:i])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPADDSETCD'. Intrinsic: '_mm512_addsetc_epi32'.
// Requires KNCNI.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func M512AddsetcEpi32(v2 x86.M512i, v3 x86.M512i, k2_res *x86.Mmask16) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskAddsetcEpi32: Performs element-by-element addition of packed 32-bit
// integer elements in 'v2' and 'v3', storing the resultant carry in 'k2_res'
// (carry flag) and the addition results in 'dst' using writemask 'k' (elements
// are copied from 'v2' and 'k_old' when the corresponding mask bit is not
// set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := v2[i+31:i] + v3[i+31:i]
//			ELSE
//				dst[i+31:i] := v2[i+31:i]
//				k2_res[j] := k_old[j]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPADDSETCD'. Intrinsic: '_mm512_mask_addsetc_epi32'.
// Requires KNCNI.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func M512MaskAddsetcEpi32(v2 x86.M512i, k x86.Mmask16, k_old x86.Mmask16, v3 x86.M512i, k2_res *x86.Mmask16) (dst x86.M512i) {
	panic("not implemented")
}


// M512AddsetsEpi32: Performs an element-by-element addition of packed 32-bit
// integer elements in 'v2' and 'v3', storing the results in 'dst' and the sign
// of the sum in 'sign' (sign flag). 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := v2[i+31:i] + v3[i+31:i]
//			sign[j] := v2[i+31:i] & v3[i+31:i] & 0x80000000
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPADDSETSD'. Intrinsic: '_mm512_addsets_epi32'.
// Requires KNCNI.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func M512AddsetsEpi32(v2 x86.M512i, v3 x86.M512i, sign *x86.Mmask16) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskAddsetsEpi32: Performs an element-by-element addition of packed
// 32-bit integer elements in 'v2' and 'v3', storing the results in 'dst' and
// the sign of the sum in 'sign' (sign flag). Results are stored using
// writemask 'k' (elements are copied from 'src' when the corresponding mask
// bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := v2[i+31:i] + v3[i+31:i]
//				sign[j] := v2[i+31:i] & v3[i+31:i] & 0x80000000
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPADDSETSD'. Intrinsic: '_mm512_mask_addsets_epi32'.
// Requires KNCNI.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func M512MaskAddsetsEpi32(src x86.M512i, k x86.Mmask16, v2 x86.M512i, v3 x86.M512i, sign *x86.Mmask16) (dst x86.M512i) {
	panic("not implemented")
}


// M512AddsetsPs: Performs an element-by-element addition of packed
// single-precision (32-bit) floating-point elements in 'v2' and 'v3', storing
// the results in 'dst' and the sign of the sum in 'sign' (sign flag). 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := v2[i+31:i] + v3[i+31:i]
//			sign[j] := v2[i+31:i] & v3[i+31:i] & 0x80000000
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VADDSETSPS'. Intrinsic: '_mm512_addsets_ps'.
// Requires KNCNI.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func M512AddsetsPs(v2 x86.M512, v3 x86.M512, sign *x86.Mmask16) (dst x86.M512) {
	panic("not implemented")
}


// M512MaskAddsetsPs: Performs an element-by-element addition of packed
// single-precision (32-bit) floating-point elements in 'v2' and 'v3', storing
// the results in 'dst' and the sign of the sum in 'sign' (sign flag). Results
// are stored using writemask 'k' (elements are copied from 'src' when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := v2[i+31:i] + v3[i+31:i]
//				sign[j] := v2[i+31:i] & v3[i+31:i] & 0x80000000
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VADDSETSPS'. Intrinsic: '_mm512_mask_addsets_ps'.
// Requires KNCNI.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func M512MaskAddsetsPs(src x86.M512, k x86.Mmask16, v2 x86.M512, v3 x86.M512, sign *x86.Mmask16) (dst x86.M512) {
	panic("not implemented")
}


// M512AddsetsRoundPs: Performs an element-by-element addition of packed
// single-precision (32-bit) floating-point elements in 'v2' and 'v3', storing
// the results in 'dst' and the sign of the sum in 'sign' (sign flag).
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := v2[i+31:i] + v3[i+31:i]
//			sign[j] := v2[i+31:i] & v3[i+31:i] & 0x80000000
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VADDSETSPS'. Intrinsic: '_mm512_addsets_round_ps'.
// Requires KNCNI.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func M512AddsetsRoundPs(v2 x86.M512, v3 x86.M512, sign *x86.Mmask16, rounding int) (dst x86.M512) {
	panic("not implemented")
}


// M512MaskAddsetsRoundPs: Performs an element-by-element addition of packed
// single-precision (32-bit) floating-point elements in 'v2' and 'v3', storing
// the results in 'dst' and the sign of the sum in 'sign' (sign flag). Results
// are stored using writemask 'k' (elements are copied from 'src' when the
// corresponding mask bit is not set).
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := v2[i+31:i] + v3[i+31:i]
//				sign[j] := v2[i+31:i] & v3[i+31:i] & 0x80000000
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VADDSETSPS'. Intrinsic: '_mm512_mask_addsets_round_ps'.
// Requires KNCNI.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func M512MaskAddsetsRoundPs(src x86.M512, k x86.Mmask16, v2 x86.M512, v3 x86.M512, sign *x86.Mmask16, rounding int) (dst x86.M512) {
	panic("not implemented")
}


// M512AlignrEpi32: Concatenate 'a' and 'b' into a 128-byte immediate result,
// shift the result right by 'count' 32-bit elements, and store the low 64
// bytes (16 elements) in 'dst'. 
//
//		temp[1023:512] := a[511:0]
//		temp[511:0] := b[511:0]
//		temp[1023:0] := temp[1023:0] >> (32*count)
//		dst[511:0] := temp[511:0]
//		dst[MAX:512] := 0
//
// Instruction: 'VALIGND'. Intrinsic: '_mm512_alignr_epi32'.
// Requires KNCNI.
func M512AlignrEpi32(a x86.M512i, b x86.M512i, count int) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskAlignrEpi32: Concatenate 'a' and 'b' into a 128-byte immediate
// result, shift the result right by 'count' 32-bit elements, and store the low
// 64 bytes (16 elements) in 'dst' using writemask 'k' (elements are copied
// from 'src' when the corresponding mask bit is not set). 
//
//		temp[1023:512] := a[511:0]
//		temp[511:0] := b[511:0]
//		temp[1023:0] := temp[1023:0] >> (32*count)
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := temp[i+31:i]
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VALIGND'. Intrinsic: '_mm512_mask_alignr_epi32'.
// Requires KNCNI.
func M512MaskAlignrEpi32(src x86.M512i, k x86.Mmask16, a x86.M512i, b x86.M512i, count int) (dst x86.M512i) {
	panic("not implemented")
}


// M512AndEpi32: Compute the bitwise AND of packed 32-bit integers in 'a' and
// 'b', and store the results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := a[i+31:i] BITWISE AND b[i+31:i]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPANDD'. Intrinsic: '_mm512_and_epi32'.
// Requires KNCNI.
func M512AndEpi32(a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskAndEpi32: Performs element-by-element bitwise AND between packed
// 32-bit integer elements of 'v2' and 'v3', storing the results in 'dst' using
// writemask 'k' (elements are copied from 'src' when the corresponding mask
// bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := v2[i+31:i] & v3[i+31:i]
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPANDD'. Intrinsic: '_mm512_mask_and_epi32'.
// Requires KNCNI.
func M512MaskAndEpi32(src x86.M512i, k x86.Mmask16, v2 x86.M512i, v3 x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512AndEpi64: Compute the bitwise AND of 512 bits (composed of packed 64-bit
// integers) in 'a' and 'b', and store the results in 'dst'. 
//
//		dst[511:0] := (a[511:0] AND b[511:0])
//		dst[MAX:512] := 0
//
// Instruction: 'VPANDQ'. Intrinsic: '_mm512_and_epi64'.
// Requires KNCNI.
func M512AndEpi64(a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskAndEpi64: Compute the bitwise AND of packed 64-bit integers in 'a'
// and 'b', and store the results in 'dst' using writemask 'k' (elements are
// copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := a[i+63:i] AND b[i+63:i]
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPANDQ'. Intrinsic: '_mm512_mask_and_epi64'.
// Requires KNCNI.
func M512MaskAndEpi64(src x86.M512i, k x86.Mmask8, a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512AndSi512: Compute the bitwise AND of 512 bits (representing integer
// data) in 'a' and 'b', and store the result in 'dst'. 
//
//		dst[511:0] := (a[511:0] AND b[511:0])
//		dst[MAX:512] := 0
//
// Instruction: 'VPANDD'. Intrinsic: '_mm512_and_si512'.
// Requires KNCNI.
func M512AndSi512(a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512AndnotEpi32: Compute the bitwise AND NOT of packed 32-bit integers in
// 'a' and 'b', and store the results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := (NOT a[i+31:i]) AND b[i+31:i]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPANDND'. Intrinsic: '_mm512_andnot_epi32'.
// Requires KNCNI.
func M512AndnotEpi32(a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskAndnotEpi32: Compute the bitwise AND NOT of packed 32-bit integers
// in 'a' and 'b', and store the results in 'dst' using writemask 'k' (elements
// are copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := ((NOT a[i+31:i]) AND b[i+31:i])
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPANDND'. Intrinsic: '_mm512_mask_andnot_epi32'.
// Requires KNCNI.
func M512MaskAndnotEpi32(src x86.M512i, k x86.Mmask16, a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512AndnotEpi64: Compute the bitwise AND NOT of 512 bits (composed of packed
// 64-bit integers) in 'a' and 'b', and store the results in 'dst'. 
//
//		dst[511:0] := ((NOT a[511:0]) AND b[511:0])
//		dst[MAX:512] := 0
//
// Instruction: 'VPANDNQ'. Intrinsic: '_mm512_andnot_epi64'.
// Requires KNCNI.
func M512AndnotEpi64(a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskAndnotEpi64: Compute the bitwise AND NOT of packed 64-bit integers
// in 'a' and 'b', and store the results in 'dst' using writemask 'k' (elements
// are copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := ((NOT a[i+63:i]) AND b[i+63:i])
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPANDNQ'. Intrinsic: '_mm512_mask_andnot_epi64'.
// Requires KNCNI.
func M512MaskAndnotEpi64(src x86.M512i, k x86.Mmask8, a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512AndnotSi512: Compute the bitwise AND NOT of 512 bits (representing
// integer data) in 'a' and 'b', and store the result in 'dst'. 
//
//		dst[511:0] := ((NOT a[511:0]) AND b[511:0])
//		dst[MAX:512] := 0
//
// Instruction: 'VPANDND'. Intrinsic: '_mm512_andnot_si512'.
// Requires KNCNI.
func M512AndnotSi512(a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskBlendEpi32: Blend packed 32-bit integers from 'a' and 'b' using
// control mask 'k', and store the results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := b[i+31:i]
//			ELSE
//				dst[i+31:i] := a[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPBLENDMD'. Intrinsic: '_mm512_mask_blend_epi32'.
// Requires KNCNI.
func M512MaskBlendEpi32(k x86.Mmask16, a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskBlendEpi64: Blend packed 64-bit integers from 'a' and 'b' using
// control mask 'k', and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := b[i+63:i]
//			ELSE
//				dst[i+63:i] := a[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPBLENDMQ'. Intrinsic: '_mm512_mask_blend_epi64'.
// Requires KNCNI.
func M512MaskBlendEpi64(k x86.Mmask8, a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskBlendPd: Blend packed double-precision (64-bit) floating-point
// elements from 'a' and 'b' using control mask 'k', and store the results in
// 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := b[i+63:i]
//			ELSE
//				dst[i+63:i] := a[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VBLENDMPD'. Intrinsic: '_mm512_mask_blend_pd'.
// Requires KNCNI.
func M512MaskBlendPd(k x86.Mmask8, a x86.M512d, b x86.M512d) (dst x86.M512d) {
	panic("not implemented")
}


// M512MaskBlendPs: Blend packed single-precision (32-bit) floating-point
// elements from 'a' and 'b' using control mask 'k', and store the results in
// 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := b[i+31:i]
//			ELSE
//				dst[i+31:i] := a[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VBLENDMPS'. Intrinsic: '_mm512_mask_blend_ps'.
// Requires KNCNI.
func M512MaskBlendPs(k x86.Mmask16, a x86.M512, b x86.M512) (dst x86.M512) {
	panic("not implemented")
}


// M512CastpdPs: Cast vector of type __m512d to type __m512.
// 	This intrinsic is only used for compilation and does not generate any
// instructions, thus it has zero latency. 
//
//		
//
// Instruction: ''. Intrinsic: '_mm512_castpd_ps'.
// Requires KNCNI.
func M512CastpdPs(a x86.M512d) (dst x86.M512) {
	panic("not implemented")
}


// M512CastpdSi512: Cast vector of type __m512d to type __m512i.
// 	This intrinsic is only used for compilation and does not generate any
// instructions, thus it has zero latency. 
//
//		
//
// Instruction: ''. Intrinsic: '_mm512_castpd_si512'.
// Requires KNCNI.
func M512CastpdSi512(a x86.M512d) (dst x86.M512i) {
	panic("not implemented")
}


// M512CastpsPd: Cast vector of type __m512 to type __m512d.
// 	This intrinsic is only used for compilation and does not generate any
// instructions, thus it has zero latency. 
//
//		
//
// Instruction: ''. Intrinsic: '_mm512_castps_pd'.
// Requires KNCNI.
func M512CastpsPd(a x86.M512) (dst x86.M512d) {
	panic("not implemented")
}


// M512CastpsSi512: Cast vector of type __m512 to type __m512i.
// 	This intrinsic is only used for compilation and does not generate any
// instructions, thus it has zero latency. 
//
//		
//
// Instruction: ''. Intrinsic: '_mm512_castps_si512'.
// Requires KNCNI.
func M512CastpsSi512(a x86.M512) (dst x86.M512i) {
	panic("not implemented")
}


// M512Castsi512Pd: Cast vector of type __m512i to type __m512d.
// 	This intrinsic is only used for compilation and does not generate any
// instructions, thus it has zero latency. 
//
//		
//
// Instruction: ''. Intrinsic: '_mm512_castsi512_pd'.
// Requires KNCNI.
func M512Castsi512Pd(a x86.M512i) (dst x86.M512d) {
	panic("not implemented")
}


// M512Castsi512Ps: Cast vector of type __m512i to type __m512.
// 	This intrinsic is only used for compilation and does not generate any
// instructions, thus it has zero latency. 
//
//		
//
// Instruction: ''. Intrinsic: '_mm512_castsi512_ps'.
// Requires KNCNI.
func M512Castsi512Ps(a x86.M512i) (dst x86.M512) {
	panic("not implemented")
}


// Skipped: _mm_clevict. Contains pointer parameter.


// M512CmpEpi32Mask: Compare packed 32-bit integers in 'a' and 'b' based on the
// comparison operand specified by 'imm8', and store the results in mask vector
// 'k'. 
//
//		CASE (imm8[7:0]) OF
//		0: OP := _MM_CMPINT_EQ
//		1: OP := _MM_CMPINT_LT
//		2: OP := _MM_CMPINT_LE
//		3: OP := _MM_CMPINT_FALSE
//		4: OP := _MM_CMPINT_NEQ
//		5: OP := _MM_CMPINT_NLT
//		6: OP := _MM_CMPINT_NLE
//		7: OP := _MM_CMPINT_TRUE
//		ESAC
//		FOR j := 0 to 15
//			i := j*32
//			k[j] := ( a[i+31:i] OP b[i+31:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPD'. Intrinsic: '_mm512_cmp_epi32_mask'.
// Requires KNCNI.
//
// FIXME: Requires compiler support (has immediate)
func M512CmpEpi32Mask(a x86.M512i, b x86.M512i, imm8 byte) (dst x86.Mmask16) {
	panic("not implemented")
}


// M512MaskCmpEpi32Mask: Compare packed 32-bit integers in 'a' and 'b' based on
// the comparison operand specified by 'imm8', and store the results in mask
// vector 'k1' using zeromask 'k' (elements are zeroed out when the
// corresponding mask bit is not set). 
//
//		CASE (imm8[7:0]) OF
//		0: OP := _MM_CMPINT_EQ
//		1: OP := _MM_CMPINT_LT
//		2: OP := _MM_CMPINT_LE
//		3: OP := _MM_CMPINT_FALSE
//		4: OP := _MM_CMPINT_NEQ
//		5: OP := _MM_CMPINT_NLT
//		6: OP := _MM_CMPINT_NLE
//		7: OP := _MM_CMPINT_TRUE
//		ESAC
//		FOR j := 0 to 15
//			i := j*32
//			IF k1[j]
//				k[j] := ( a[i+31:i] OP b[i+31:i] ) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPD'. Intrinsic: '_mm512_mask_cmp_epi32_mask'.
// Requires KNCNI.
//
// FIXME: Requires compiler support (has immediate)
func M512MaskCmpEpi32Mask(k1 x86.Mmask16, a x86.M512i, b x86.M512i, imm8 byte) (dst x86.Mmask16) {
	panic("not implemented")
}


// M512CmpEpu32Mask: Compare packed unsigned 32-bit integers in 'a' and 'b'
// based on the comparison operand specified by 'imm8', and store the results
// in mask vector 'k'. 
//
//		CASE (imm8[7:0]) OF
//		0: OP := _MM_CMPINT_EQ
//		1: OP := _MM_CMPINT_LT
//		2: OP := _MM_CMPINT_LE
//		3: OP := _MM_CMPINT_FALSE
//		4: OP := _MM_CMPINT_NEQ
//		5: OP := _MM_CMPINT_NLT
//		6: OP := _MM_CMPINT_NLE
//		7: OP := _MM_CMPINT_TRUE
//		ESAC
//		FOR j := 0 to 15
//			i := j*32
//			k[j] := ( a[i+31:i] OP b[i+31:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPUD'. Intrinsic: '_mm512_cmp_epu32_mask'.
// Requires KNCNI.
//
// FIXME: Requires compiler support (has immediate)
func M512CmpEpu32Mask(a x86.M512i, b x86.M512i, imm8 byte) (dst x86.Mmask16) {
	panic("not implemented")
}


// M512MaskCmpEpu32Mask: Compare packed unsigned 32-bit integers in 'a' and 'b'
// based on the comparison operand specified by 'imm8', and store the results
// in mask vector 'k1' using zeromask 'k' (elements are zeroed out when the
// corresponding mask bit is not set). 
//
//		CASE (imm8[7:0]) OF
//		0: OP := _MM_CMPINT_EQ
//		1: OP := _MM_CMPINT_LT
//		2: OP := _MM_CMPINT_LE
//		3: OP := _MM_CMPINT_FALSE
//		4: OP := _MM_CMPINT_NEQ
//		5: OP := _MM_CMPINT_NLT
//		6: OP := _MM_CMPINT_NLE
//		7: OP := _MM_CMPINT_TRUE
//		ESAC
//		FOR j := 0 to 15
//			i := j*32
//			IF k1[j]
//				k[j] := ( a[i+31:i] OP b[i+31:i] ) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPUD'. Intrinsic: '_mm512_mask_cmp_epu32_mask'.
// Requires KNCNI.
//
// FIXME: Requires compiler support (has immediate)
func M512MaskCmpEpu32Mask(k1 x86.Mmask16, a x86.M512i, b x86.M512i, imm8 byte) (dst x86.Mmask16) {
	panic("not implemented")
}


// M512CmpPdMask: Compare packed double-precision (64-bit) floating-point
// elements in 'a' and 'b' based on the comparison operand specified by 'imm8',
// and store the results in mask vector 'k'. 
//
//		CASE (imm8[7:0]) OF
//		0: OP := _CMP_EQ_OQ
//		1: OP := _CMP_LT_OS
//		2: OP := _CMP_LE_OS
//		3: OP := _CMP_UNORD_Q 
//		4: OP := _CMP_NEQ_UQ
//		5: OP := _CMP_NLT_US
//		6: OP := _CMP_NLE_US
//		7: OP := _CMP_ORD_Q
//		8: OP := _CMP_EQ_UQ
//		9: OP := _CMP_NGE_US
//		10: OP := _CMP_NGT_US
//		11: OP := _CMP_FALSE_OQ
//		12: OP := _CMP_NEQ_OQ
//		13: OP := _CMP_GE_OS
//		14: OP := _CMP_GT_OS
//		15: OP := _CMP_TRUE_UQ
//		16: OP := _CMP_EQ_OS
//		17: OP := _CMP_LT_OQ
//		18: OP := _CMP_LE_OQ
//		19: OP := _CMP_UNORD_S
//		20: OP := _CMP_NEQ_US
//		21: OP := _CMP_NLT_UQ
//		22: OP := _CMP_NLE_UQ
//		23: OP := _CMP_ORD_S
//		24: OP := _CMP_EQ_US
//		25: OP := _CMP_NGE_UQ 
//		26: OP := _CMP_NGT_UQ 
//		27: OP := _CMP_FALSE_OS 
//		28: OP := _CMP_NEQ_OS 
//		29: OP := _CMP_GE_OQ
//		30: OP := _CMP_GT_OQ
//		31: OP := _CMP_TRUE_US
//		ESAC
//		FOR j := 0 to 7
//			i := j*64
//			k[j] := (a[i+63:i] OP b[i+63:i]) ? 1 : 0
//		ENDFOR
//		k[MAX:8] := 0
//
// Instruction: 'VCMPPD'. Intrinsic: '_mm512_cmp_pd_mask'.
// Requires KNCNI.
//
// FIXME: Requires compiler support (has immediate)
func M512CmpPdMask(a x86.M512d, b x86.M512d, imm8 byte) (dst x86.Mmask8) {
	panic("not implemented")
}


// M512MaskCmpPdMask: Compare packed double-precision (64-bit) floating-point
// elements in 'a' and 'b' based on the comparison operand specified by 'imm8',
// and store the results in mask vector 'k' using zeromask 'k1' (elements are
// zeroed out when the corresponding mask bit is not set). 
//
//		CASE (imm8[7:0]) OF
//		0: OP := _CMP_EQ_OQ
//		1: OP := _CMP_LT_OS
//		2: OP := _CMP_LE_OS
//		3: OP := _CMP_UNORD_Q 
//		4: OP := _CMP_NEQ_UQ
//		5: OP := _CMP_NLT_US
//		6: OP := _CMP_NLE_US
//		7: OP := _CMP_ORD_Q
//		8: OP := _CMP_EQ_UQ
//		9: OP := _CMP_NGE_US
//		10: OP := _CMP_NGT_US
//		11: OP := _CMP_FALSE_OQ
//		12: OP := _CMP_NEQ_OQ
//		13: OP := _CMP_GE_OS
//		14: OP := _CMP_GT_OS
//		15: OP := _CMP_TRUE_UQ
//		16: OP := _CMP_EQ_OS
//		17: OP := _CMP_LT_OQ
//		18: OP := _CMP_LE_OQ
//		19: OP := _CMP_UNORD_S
//		20: OP := _CMP_NEQ_US
//		21: OP := _CMP_NLT_UQ
//		22: OP := _CMP_NLE_UQ
//		23: OP := _CMP_ORD_S
//		24: OP := _CMP_EQ_US
//		25: OP := _CMP_NGE_UQ 
//		26: OP := _CMP_NGT_UQ 
//		27: OP := _CMP_FALSE_OS 
//		28: OP := _CMP_NEQ_OS 
//		29: OP := _CMP_GE_OQ
//		30: OP := _CMP_GT_OQ
//		31: OP := _CMP_TRUE_US
//		ESAC
//		FOR j := 0 to 7
//			i := j*64
//			IF k1[j]
//				k[j] := ( a[i+63:i] OP b[i+63:i] ) ? 1 : 0
//			ELSE
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:8] := 0
//
// Instruction: 'VCMPPD'. Intrinsic: '_mm512_mask_cmp_pd_mask'.
// Requires KNCNI.
//
// FIXME: Requires compiler support (has immediate)
func M512MaskCmpPdMask(k1 x86.Mmask8, a x86.M512d, b x86.M512d, imm8 byte) (dst x86.Mmask8) {
	panic("not implemented")
}


// M512CmpPsMask: Compare packed single-precision (32-bit) floating-point
// elements in 'a' and 'b' based on the comparison operand specified by 'imm8',
// and store the results in mask vector 'k'. 
//
//		CASE (imm8[7:0]) OF
//		0: OP := _CMP_EQ_OQ
//		1: OP := _CMP_LT_OS
//		2: OP := _CMP_LE_OS
//		3: OP := _CMP_UNORD_Q 
//		4: OP := _CMP_NEQ_UQ
//		5: OP := _CMP_NLT_US
//		6: OP := _CMP_NLE_US
//		7: OP := _CMP_ORD_Q
//		8: OP := _CMP_EQ_UQ
//		9: OP := _CMP_NGE_US
//		10: OP := _CMP_NGT_US
//		11: OP := _CMP_FALSE_OQ
//		12: OP := _CMP_NEQ_OQ
//		13: OP := _CMP_GE_OS
//		14: OP := _CMP_GT_OS
//		15: OP := _CMP_TRUE_UQ
//		16: OP := _CMP_EQ_OS
//		17: OP := _CMP_LT_OQ
//		18: OP := _CMP_LE_OQ
//		19: OP := _CMP_UNORD_S
//		20: OP := _CMP_NEQ_US
//		21: OP := _CMP_NLT_UQ
//		22: OP := _CMP_NLE_UQ
//		23: OP := _CMP_ORD_S
//		24: OP := _CMP_EQ_US
//		25: OP := _CMP_NGE_UQ 
//		26: OP := _CMP_NGT_UQ 
//		27: OP := _CMP_FALSE_OS 
//		28: OP := _CMP_NEQ_OS 
//		29: OP := _CMP_GE_OQ
//		30: OP := _CMP_GT_OQ
//		31: OP := _CMP_TRUE_US
//		ESAC
//		FOR j := 0 to 15
//			i := j*32
//			k[j] := (a[i+31:i] OP b[i+31:i]) ? 1 : 0
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VCMPPS'. Intrinsic: '_mm512_cmp_ps_mask'.
// Requires KNCNI.
//
// FIXME: Requires compiler support (has immediate)
func M512CmpPsMask(a x86.M512, b x86.M512, imm8 byte) (dst x86.Mmask16) {
	panic("not implemented")
}


// M512MaskCmpPsMask: Compare packed single-precision (32-bit) floating-point
// elements in 'a' and 'b' based on the comparison operand specified by 'imm8',
// and store the results in mask vector 'k' using zeromask 'k1' (elements are
// zeroed out when the corresponding mask bit is not set). 
//
//		CASE (imm8[7:0]) OF
//		0: OP := _CMP_EQ_OQ
//		1: OP := _CMP_LT_OS
//		2: OP := _CMP_LE_OS
//		3: OP := _CMP_UNORD_Q 
//		4: OP := _CMP_NEQ_UQ
//		5: OP := _CMP_NLT_US
//		6: OP := _CMP_NLE_US
//		7: OP := _CMP_ORD_Q
//		8: OP := _CMP_EQ_UQ
//		9: OP := _CMP_NGE_US
//		10: OP := _CMP_NGT_US
//		11: OP := _CMP_FALSE_OQ
//		12: OP := _CMP_NEQ_OQ
//		13: OP := _CMP_GE_OS
//		14: OP := _CMP_GT_OS
//		15: OP := _CMP_TRUE_UQ
//		16: OP := _CMP_EQ_OS
//		17: OP := _CMP_LT_OQ
//		18: OP := _CMP_LE_OQ
//		19: OP := _CMP_UNORD_S
//		20: OP := _CMP_NEQ_US
//		21: OP := _CMP_NLT_UQ
//		22: OP := _CMP_NLE_UQ
//		23: OP := _CMP_ORD_S
//		24: OP := _CMP_EQ_US
//		25: OP := _CMP_NGE_UQ 
//		26: OP := _CMP_NGT_UQ 
//		27: OP := _CMP_FALSE_OS 
//		28: OP := _CMP_NEQ_OS 
//		29: OP := _CMP_GE_OQ
//		30: OP := _CMP_GT_OQ
//		31: OP := _CMP_TRUE_US
//		ESAC
//		FOR j := 0 to 15
//			i := j*32
//			IF k1[j]
//				k[j] := ( a[i+31:i] OP b[i+31:i] ) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VCMPPS'. Intrinsic: '_mm512_mask_cmp_ps_mask'.
// Requires KNCNI.
//
// FIXME: Requires compiler support (has immediate)
func M512MaskCmpPsMask(k1 x86.Mmask16, a x86.M512, b x86.M512, imm8 byte) (dst x86.Mmask16) {
	panic("not implemented")
}


// M512CmpRoundPdMask: Compare packed double-precision (64-bit) floating-point
// elements in 'a' and 'b' based on the comparison operand specified by 'imm8',
// and store the results in mask vector 'k'.
// 	Pass __MM_FROUND_NO_EXC to 'sae' to suppress all exceptions. 
//
//		CASE (imm8[7:0]) OF
//		0: OP := _CMP_EQ_OQ
//		1: OP := _CMP_LT_OS
//		2: OP := _CMP_LE_OS
//		3: OP := _CMP_UNORD_Q 
//		4: OP := _CMP_NEQ_UQ
//		5: OP := _CMP_NLT_US
//		6: OP := _CMP_NLE_US
//		7: OP := _CMP_ORD_Q
//		8: OP := _CMP_EQ_UQ
//		9: OP := _CMP_NGE_US
//		10: OP := _CMP_NGT_US
//		11: OP := _CMP_FALSE_OQ
//		12: OP := _CMP_NEQ_OQ
//		13: OP := _CMP_GE_OS
//		14: OP := _CMP_GT_OS
//		15: OP := _CMP_TRUE_UQ
//		16: OP := _CMP_EQ_OS
//		17: OP := _CMP_LT_OQ
//		18: OP := _CMP_LE_OQ
//		19: OP := _CMP_UNORD_S
//		20: OP := _CMP_NEQ_US
//		21: OP := _CMP_NLT_UQ
//		22: OP := _CMP_NLE_UQ
//		23: OP := _CMP_ORD_S
//		24: OP := _CMP_EQ_US
//		25: OP := _CMP_NGE_UQ 
//		26: OP := _CMP_NGT_UQ 
//		27: OP := _CMP_FALSE_OS 
//		28: OP := _CMP_NEQ_OS 
//		29: OP := _CMP_GE_OQ
//		30: OP := _CMP_GT_OQ
//		31: OP := _CMP_TRUE_US
//		ESAC
//		FOR j := 0 to 7
//			i := j*64
//			k[j] := (a[i+63:i] OP b[i+63:i]) ? 1 : 0
//		ENDFOR
//		k[MAX:8] := 0
//
// Instruction: 'VCMPPD'. Intrinsic: '_mm512_cmp_round_pd_mask'.
// Requires KNCNI.
//
// FIXME: Requires compiler support (has immediate)
func M512CmpRoundPdMask(a x86.M512d, b x86.M512d, imm8 byte, sae int) (dst x86.Mmask8) {
	panic("not implemented")
}


// M512MaskCmpRoundPdMask: Compare packed double-precision (64-bit)
// floating-point elements in 'a' and 'b' based on the comparison operand
// specified by 'imm8', and store the results in mask vector 'k' using zeromask
// 'k1' (elements are zeroed out when the corresponding mask bit is not set). 
// 	Pass __MM_FROUND_NO_EXC to 'sae' to suppress all exceptions. 
//
//		CASE (imm8[7:0]) OF
//		0: OP := _CMP_EQ_OQ
//		1: OP := _CMP_LT_OS
//		2: OP := _CMP_LE_OS
//		3: OP := _CMP_UNORD_Q 
//		4: OP := _CMP_NEQ_UQ
//		5: OP := _CMP_NLT_US
//		6: OP := _CMP_NLE_US
//		7: OP := _CMP_ORD_Q
//		8: OP := _CMP_EQ_UQ
//		9: OP := _CMP_NGE_US
//		10: OP := _CMP_NGT_US
//		11: OP := _CMP_FALSE_OQ
//		12: OP := _CMP_NEQ_OQ
//		13: OP := _CMP_GE_OS
//		14: OP := _CMP_GT_OS
//		15: OP := _CMP_TRUE_UQ
//		16: OP := _CMP_EQ_OS
//		17: OP := _CMP_LT_OQ
//		18: OP := _CMP_LE_OQ
//		19: OP := _CMP_UNORD_S
//		20: OP := _CMP_NEQ_US
//		21: OP := _CMP_NLT_UQ
//		22: OP := _CMP_NLE_UQ
//		23: OP := _CMP_ORD_S
//		24: OP := _CMP_EQ_US
//		25: OP := _CMP_NGE_UQ 
//		26: OP := _CMP_NGT_UQ 
//		27: OP := _CMP_FALSE_OS 
//		28: OP := _CMP_NEQ_OS 
//		29: OP := _CMP_GE_OQ
//		30: OP := _CMP_GT_OQ
//		31: OP := _CMP_TRUE_US
//		ESAC
//		FOR j := 0 to 7
//			i := j*64
//			IF k1[j]
//				k[j] := ( a[i+63:i] OP b[i+63:i] ) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:8] := 0
//
// Instruction: 'VCMPPD'. Intrinsic: '_mm512_mask_cmp_round_pd_mask'.
// Requires KNCNI.
//
// FIXME: Requires compiler support (has immediate)
func M512MaskCmpRoundPdMask(k1 x86.Mmask8, a x86.M512d, b x86.M512d, imm8 byte, sae int) (dst x86.Mmask8) {
	panic("not implemented")
}


// M512CmpRoundPsMask: Compare packed single-precision (32-bit) floating-point
// elements in 'a' and 'b' based on the comparison operand specified by 'imm8',
// and store the results in mask vector 'k'.
// 	Pass __MM_FROUND_NO_EXC to 'sae' to suppress all exceptions. 
//
//		CASE (imm8[7:0]) OF
//		0: OP := _CMP_EQ_OQ
//		1: OP := _CMP_LT_OS
//		2: OP := _CMP_LE_OS
//		3: OP := _CMP_UNORD_Q 
//		4: OP := _CMP_NEQ_UQ
//		5: OP := _CMP_NLT_US
//		6: OP := _CMP_NLE_US
//		7: OP := _CMP_ORD_Q
//		8: OP := _CMP_EQ_UQ
//		9: OP := _CMP_NGE_US
//		10: OP := _CMP_NGT_US
//		11: OP := _CMP_FALSE_OQ
//		12: OP := _CMP_NEQ_OQ
//		13: OP := _CMP_GE_OS
//		14: OP := _CMP_GT_OS
//		15: OP := _CMP_TRUE_UQ
//		16: OP := _CMP_EQ_OS
//		17: OP := _CMP_LT_OQ
//		18: OP := _CMP_LE_OQ
//		19: OP := _CMP_UNORD_S
//		20: OP := _CMP_NEQ_US
//		21: OP := _CMP_NLT_UQ
//		22: OP := _CMP_NLE_UQ
//		23: OP := _CMP_ORD_S
//		24: OP := _CMP_EQ_US
//		25: OP := _CMP_NGE_UQ 
//		26: OP := _CMP_NGT_UQ 
//		27: OP := _CMP_FALSE_OS 
//		28: OP := _CMP_NEQ_OS 
//		29: OP := _CMP_GE_OQ
//		30: OP := _CMP_GT_OQ
//		31: OP := _CMP_TRUE_US
//		ESAC
//		FOR j := 0 to 15
//			i := j*32
//			k[j] := (a[i+31:i] OP b[i+31:i]) ? 1 : 0
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VCMPPS'. Intrinsic: '_mm512_cmp_round_ps_mask'.
// Requires KNCNI.
//
// FIXME: Requires compiler support (has immediate)
func M512CmpRoundPsMask(a x86.M512, b x86.M512, imm8 byte, sae int) (dst x86.Mmask16) {
	panic("not implemented")
}


// M512MaskCmpRoundPsMask: Compare packed single-precision (32-bit)
// floating-point elements in 'a' and 'b' based on the comparison operand
// specified by 'imm8', and store the results in mask vector 'k' using zeromask
// 'k1' (elements are zeroed out when the corresponding mask bit is not set). 
// 	Pass __MM_FROUND_NO_EXC to 'sae' to suppress all exceptions. 
//
//		CASE (imm8[7:0]) OF
//		0: OP := _CMP_EQ_OQ
//		1: OP := _CMP_LT_OS
//		2: OP := _CMP_LE_OS
//		3: OP := _CMP_UNORD_Q 
//		4: OP := _CMP_NEQ_UQ
//		5: OP := _CMP_NLT_US
//		6: OP := _CMP_NLE_US
//		7: OP := _CMP_ORD_Q
//		8: OP := _CMP_EQ_UQ
//		9: OP := _CMP_NGE_US
//		10: OP := _CMP_NGT_US
//		11: OP := _CMP_FALSE_OQ
//		12: OP := _CMP_NEQ_OQ
//		13: OP := _CMP_GE_OS
//		14: OP := _CMP_GT_OS
//		15: OP := _CMP_TRUE_UQ
//		16: OP := _CMP_EQ_OS
//		17: OP := _CMP_LT_OQ
//		18: OP := _CMP_LE_OQ
//		19: OP := _CMP_UNORD_S
//		20: OP := _CMP_NEQ_US
//		21: OP := _CMP_NLT_UQ
//		22: OP := _CMP_NLE_UQ
//		23: OP := _CMP_ORD_S
//		24: OP := _CMP_EQ_US
//		25: OP := _CMP_NGE_UQ 
//		26: OP := _CMP_NGT_UQ 
//		27: OP := _CMP_FALSE_OS 
//		28: OP := _CMP_NEQ_OS 
//		29: OP := _CMP_GE_OQ
//		30: OP := _CMP_GT_OQ
//		31: OP := _CMP_TRUE_US
//		ESAC
//		FOR j := 0 to 15
//			i := j*32
//			IF k1[j]
//				k[j] := ( a[i+31:i] OP b[i+31:i] ) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VCMPPS'. Intrinsic: '_mm512_mask_cmp_round_ps_mask'.
// Requires KNCNI.
//
// FIXME: Requires compiler support (has immediate)
func M512MaskCmpRoundPsMask(k1 x86.Mmask16, a x86.M512, b x86.M512, imm8 byte, sae int) (dst x86.Mmask16) {
	panic("not implemented")
}


// M512CmpeqEpi32Mask: Compare packed 32-bit integers in 'a' and 'b' for
// equality, and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			k[j] := ( a[i+31:i] == b[i+31:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPEQD'. Intrinsic: '_mm512_cmpeq_epi32_mask'.
// Requires KNCNI.
func M512CmpeqEpi32Mask(a x86.M512i, b x86.M512i) (dst x86.Mmask16) {
	panic("not implemented")
}


// M512MaskCmpeqEpi32Mask: Compare packed 32-bit integers in 'a' and 'b' for
// equality, and store the results in mask vector 'k1' using zeromask 'k'
// (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k1[j]
//				k[j] := ( a[i+31:i] == b[i+31:i] ) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPEQD'. Intrinsic: '_mm512_mask_cmpeq_epi32_mask'.
// Requires KNCNI.
func M512MaskCmpeqEpi32Mask(k1 x86.Mmask16, a x86.M512i, b x86.M512i) (dst x86.Mmask16) {
	panic("not implemented")
}


// M512CmpeqEpu32Mask: Compare packed unsigned 32-bit integers in 'a' and 'b'
// for equality, and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			k[j] := ( a[i+31:i] == b[i+31:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPUD'. Intrinsic: '_mm512_cmpeq_epu32_mask'.
// Requires KNCNI.
func M512CmpeqEpu32Mask(a x86.M512i, b x86.M512i) (dst x86.Mmask16) {
	panic("not implemented")
}


// M512MaskCmpeqEpu32Mask: Compare packed unsigned 32-bit integers in 'a' and
// 'b' for equality, and store the results in mask vector 'k1' using zeromask
// 'k' (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k1[j]
//				k[j] := ( a[i+31:i] == b[i+31:i] ) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPUD'. Intrinsic: '_mm512_mask_cmpeq_epu32_mask'.
// Requires KNCNI.
func M512MaskCmpeqEpu32Mask(k1 x86.Mmask16, a x86.M512i, b x86.M512i) (dst x86.Mmask16) {
	panic("not implemented")
}


// M512CmpeqPdMask: Compare packed double-precision (64-bit) floating-point
// elements in 'a' and 'b' for equality, and store the results in mask vector
// 'k'. 
//
//		FOR j := 0 to 7
//			i := j*64
//			k[j] := (a[i+63:i] == b[i+63:i]) ? 1 : 0
//		ENDFOR	
//		k[MAX:8] := 0
//
// Instruction: 'VCMPPD'. Intrinsic: '_mm512_cmpeq_pd_mask'.
// Requires KNCNI.
func M512CmpeqPdMask(a x86.M512d, b x86.M512d) (dst x86.Mmask8) {
	panic("not implemented")
}


// M512MaskCmpeqPdMask: Compare packed double-precision (64-bit) floating-point
// elements in 'a' and 'b' for equality, and store the results in mask vector
// 'k' using zeromask 'k1' (elements are zeroed out when the corresponding mask
// bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k1[j]
//				k[j] := (a[i+63:i] == b[i+63:i]) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR	
//		k[MAX:8] := 0
//
// Instruction: 'VCMPPD'. Intrinsic: '_mm512_mask_cmpeq_pd_mask'.
// Requires KNCNI.
func M512MaskCmpeqPdMask(k1 x86.Mmask8, a x86.M512d, b x86.M512d) (dst x86.Mmask8) {
	panic("not implemented")
}


// M512CmpeqPsMask: Compare packed single-precision (32-bit) floating-point
// elements in 'a' and 'b' for equality, and store the results in mask vector
// 'k'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			k[j] := (a[i+31:i] == b[i+31:i]) ? 1 : 0
//		ENDFOR	
//		k[MAX:16] := 0
//
// Instruction: 'VCMPPS'. Intrinsic: '_mm512_cmpeq_ps_mask'.
// Requires KNCNI.
func M512CmpeqPsMask(a x86.M512, b x86.M512) (dst x86.Mmask16) {
	panic("not implemented")
}


// M512MaskCmpeqPsMask: Compare packed single-precision (32-bit) floating-point
// elements in 'a' and 'b' for equality, and store the results in mask vector
// 'k' using zeromask 'k1' (elements are zeroed out when the corresponding mask
// bit is not set). 
//
//		y
//		FOR j := 0 to 15
//			i := j*32
//			IF k1[j]
//				k[j] := (a[i+31:i] == b[i+31:i]) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR		
//		k[MAX:16] := 0
//
// Instruction: 'VCMPPS'. Intrinsic: '_mm512_mask_cmpeq_ps_mask'.
// Requires KNCNI.
func M512MaskCmpeqPsMask(k1 x86.Mmask16, a x86.M512, b x86.M512) (dst x86.Mmask16) {
	panic("not implemented")
}


// M512CmpgeEpi32Mask: Compare packed 32-bit integers in 'a' and 'b' for
// greater-than-or-equal, and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			k[j] := ( a[i+31:i] >= b[i+31:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPD'. Intrinsic: '_mm512_cmpge_epi32_mask'.
// Requires KNCNI.
func M512CmpgeEpi32Mask(a x86.M512i, b x86.M512i) (dst x86.Mmask16) {
	panic("not implemented")
}


// M512MaskCmpgeEpi32Mask: Compare packed 32-bit integers in 'a' and 'b' for
// greater-than-or-equal, and store the results in mask vector 'k1' using
// zeromask 'k' (elements are zeroed out when the corresponding mask bit is not
// set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k1[j]
//				k[j] := ( a[i+31:i] >= b[i+31:i] ) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPD'. Intrinsic: '_mm512_mask_cmpge_epi32_mask'.
// Requires KNCNI.
func M512MaskCmpgeEpi32Mask(k1 x86.Mmask16, a x86.M512i, b x86.M512i) (dst x86.Mmask16) {
	panic("not implemented")
}


// M512CmpgeEpu32Mask: Compare packed unsigned 32-bit integers in 'a' and 'b'
// for greater-than-or-equal, and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			k[j] := ( a[i+31:i] >= b[i+31:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPUD'. Intrinsic: '_mm512_cmpge_epu32_mask'.
// Requires KNCNI.
func M512CmpgeEpu32Mask(a x86.M512i, b x86.M512i) (dst x86.Mmask16) {
	panic("not implemented")
}


// M512MaskCmpgeEpu32Mask: Compare packed unsigned 32-bit integers in 'a' and
// 'b' for greater-than-or-equal, and store the results in mask vector 'k1'
// using zeromask 'k' (elements are zeroed out when the corresponding mask bit
// is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k1[j]
//				k[j] := ( a[i+31:i] >= b[i+31:i] ) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPUD'. Intrinsic: '_mm512_mask_cmpge_epu32_mask'.
// Requires KNCNI.
func M512MaskCmpgeEpu32Mask(k1 x86.Mmask16, a x86.M512i, b x86.M512i) (dst x86.Mmask16) {
	panic("not implemented")
}


// M512CmpgtEpi32Mask: Compare packed 32-bit integers in 'a' and 'b' for
// greater-than, and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			k[j] := ( a[i+31:i] > b[i+31:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPGTD'. Intrinsic: '_mm512_cmpgt_epi32_mask'.
// Requires KNCNI.
func M512CmpgtEpi32Mask(a x86.M512i, b x86.M512i) (dst x86.Mmask16) {
	panic("not implemented")
}


// M512MaskCmpgtEpi32Mask: Compare packed 32-bit integers in 'a' and 'b' for
// greater-than, and store the results in mask vector 'k1' using zeromask 'k'
// (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k1[j]
//				k[j] := ( a[i+31:i] > b[i+31:i] ) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPGTD'. Intrinsic: '_mm512_mask_cmpgt_epi32_mask'.
// Requires KNCNI.
func M512MaskCmpgtEpi32Mask(k1 x86.Mmask16, a x86.M512i, b x86.M512i) (dst x86.Mmask16) {
	panic("not implemented")
}


// M512CmpgtEpu32Mask: Compare packed unsigned 32-bit integers in 'a' and 'b'
// for greater-than, and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			k[j] := ( a[i+31:i] > b[i+31:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPUD'. Intrinsic: '_mm512_cmpgt_epu32_mask'.
// Requires KNCNI.
func M512CmpgtEpu32Mask(a x86.M512i, b x86.M512i) (dst x86.Mmask16) {
	panic("not implemented")
}


// M512MaskCmpgtEpu32Mask: Compare packed unsigned 32-bit integers in 'a' and
// 'b' for greater-than, and store the results in mask vector 'k1' using
// zeromask 'k' (elements are zeroed out when the corresponding mask bit is not
// set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k1[j]
//				k[j] := ( a[i+31:i] > b[i+31:i] ) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPUD'. Intrinsic: '_mm512_mask_cmpgt_epu32_mask'.
// Requires KNCNI.
func M512MaskCmpgtEpu32Mask(k1 x86.Mmask16, a x86.M512i, b x86.M512i) (dst x86.Mmask16) {
	panic("not implemented")
}


// M512CmpleEpi32Mask: Compare packed 32-bit integers in 'a' and 'b' for
// less-than-or-equal, and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			k[j] := ( a[i+31:i] <= b[i+31:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPD'. Intrinsic: '_mm512_cmple_epi32_mask'.
// Requires KNCNI.
func M512CmpleEpi32Mask(a x86.M512i, b x86.M512i) (dst x86.Mmask16) {
	panic("not implemented")
}


// M512MaskCmpleEpi32Mask: Compare packed 32-bit integers in 'a' and 'b' for
// less-than, and store the results in mask vector 'k1' using zeromask 'k'
// (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k1[j]
//				k[j] := ( a[i+31:i] < b[i+31:i] ) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPD'. Intrinsic: '_mm512_mask_cmple_epi32_mask'.
// Requires KNCNI.
func M512MaskCmpleEpi32Mask(k1 x86.Mmask16, a x86.M512i, b x86.M512i) (dst x86.Mmask16) {
	panic("not implemented")
}


// M512CmpleEpu32Mask: Compare packed unsigned 32-bit integers in 'a' and 'b'
// for less-than-or-equal, and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			k[j] := ( a[i+31:i] <= b[i+31:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPUD'. Intrinsic: '_mm512_cmple_epu32_mask'.
// Requires KNCNI.
func M512CmpleEpu32Mask(a x86.M512i, b x86.M512i) (dst x86.Mmask16) {
	panic("not implemented")
}


// M512MaskCmpleEpu32Mask: Compare packed unsigned 32-bit integers in 'a' and
// 'b' for less-than, and store the results in mask vector 'k1' using zeromask
// 'k' (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k1[j]
//				k[j] := ( a[i+31:i] < b[i+31:i] ) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPUD'. Intrinsic: '_mm512_mask_cmple_epu32_mask'.
// Requires KNCNI.
func M512MaskCmpleEpu32Mask(k1 x86.Mmask16, a x86.M512i, b x86.M512i) (dst x86.Mmask16) {
	panic("not implemented")
}


// M512CmplePdMask: Compare packed double-precision (64-bit) floating-point
// elements in 'a' and 'b' for less-than-or-equal, and store the results in
// mask vector 'k'. 
//
//		FOR j := 0 to 7
//			i := j*64
//			k[j] := (a[i+63:i] <= b[i+63:i]) ? 1 : 0
//		ENDFOR
//		k[MAX:8] := 0
//
// Instruction: 'VCMPPD'. Intrinsic: '_mm512_cmple_pd_mask'.
// Requires KNCNI.
func M512CmplePdMask(a x86.M512d, b x86.M512d) (dst x86.Mmask8) {
	panic("not implemented")
}


// M512MaskCmplePdMask: Compare packed double-precision (64-bit) floating-point
// elements in 'a' and 'b' for less-than-or-equal, and store the results in
// mask vector 'k' using zeromask 'k1' (elements are zeroed out when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k1[j]
//				k[j] := (a[i+63:i] <= b[i+63:i]) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:8] := 0
//
// Instruction: 'VCMPPD'. Intrinsic: '_mm512_mask_cmple_pd_mask'.
// Requires KNCNI.
func M512MaskCmplePdMask(k1 x86.Mmask8, a x86.M512d, b x86.M512d) (dst x86.Mmask8) {
	panic("not implemented")
}


// M512CmplePsMask: Compare packed single-precision (32-bit) floating-point
// elements in 'a' and 'b' for less-than-or-equal, and store the results in
// mask vector 'k'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			k[j] := (a[i+31:i] <= b[i+31:i]) ? 1 : 0
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VCMPPS'. Intrinsic: '_mm512_cmple_ps_mask'.
// Requires KNCNI.
func M512CmplePsMask(a x86.M512, b x86.M512) (dst x86.Mmask16) {
	panic("not implemented")
}


// M512MaskCmplePsMask: Compare packed single-precision (32-bit) floating-point
// elements in 'a' and 'b' for less-than-or-equal, and store the results in
// mask vector 'k' using zeromask 'k1' (elements are zeroed out when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k1[j]
//				k[j] := (a[i+31:i] <= b[i+31:i]) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VCMPPS'. Intrinsic: '_mm512_mask_cmple_ps_mask'.
// Requires KNCNI.
func M512MaskCmplePsMask(k1 x86.Mmask16, a x86.M512, b x86.M512) (dst x86.Mmask16) {
	panic("not implemented")
}


// M512CmpltEpi32Mask: Compare packed 32-bit integers in 'a' and 'b' for
// less-than, and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			k[j] := ( a[i+31:i] < b[i+31:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPLTD'. Intrinsic: '_mm512_cmplt_epi32_mask'.
// Requires KNCNI.
func M512CmpltEpi32Mask(a x86.M512i, b x86.M512i) (dst x86.Mmask16) {
	panic("not implemented")
}


// M512MaskCmpltEpi32Mask: Compare packed 32-bit integers in 'a' and 'b' for
// less-than-or-equal, and store the results in mask vector 'k1' using zeromask
// 'k' (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k1[j]
//				k[j] := ( a[i+31:i] <= b[i+31:i] ) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPLTD'. Intrinsic: '_mm512_mask_cmplt_epi32_mask'.
// Requires KNCNI.
func M512MaskCmpltEpi32Mask(k1 x86.Mmask16, a x86.M512i, b x86.M512i) (dst x86.Mmask16) {
	panic("not implemented")
}


// M512CmpltEpu32Mask: Compare packed unsigned 32-bit integers in 'a' and 'b'
// for less-than, and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			k[j] := ( a[i+31:i] < b[i+31:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPUD'. Intrinsic: '_mm512_cmplt_epu32_mask'.
// Requires KNCNI.
func M512CmpltEpu32Mask(a x86.M512i, b x86.M512i) (dst x86.Mmask16) {
	panic("not implemented")
}


// M512MaskCmpltEpu32Mask: Compare packed unsigned 32-bit integers in 'a' and
// 'b' for less-than-or-equal, and store the results in mask vector 'k1' using
// zeromask 'k' (elements are zeroed out when the corresponding mask bit is not
// set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k1[j]
//				k[j] := ( a[i+31:i] <= b[i+31:i] ) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPUD'. Intrinsic: '_mm512_mask_cmplt_epu32_mask'.
// Requires KNCNI.
func M512MaskCmpltEpu32Mask(k1 x86.Mmask16, a x86.M512i, b x86.M512i) (dst x86.Mmask16) {
	panic("not implemented")
}


// M512CmpltPdMask: Compare packed double-precision (64-bit) floating-point
// elements in 'a' and 'b' for less-than, and store the results in mask vector
// 'k'. 
//
//		FOR j := 0 to 7
//			i := j*64
//			k[j] := (a[i+63:i] < b[i+63:i]) ? 1 : 0
//		ENDFOR
//		k[MAX:8] := 0
//
// Instruction: 'VCMPPD'. Intrinsic: '_mm512_cmplt_pd_mask'.
// Requires KNCNI.
func M512CmpltPdMask(a x86.M512d, b x86.M512d) (dst x86.Mmask8) {
	panic("not implemented")
}


// M512MaskCmpltPdMask: Compare packed double-precision (64-bit) floating-point
// elements in 'a' and 'b' for less-than, and store the results in mask vector
// 'k' using zeromask 'k1' (elements are zeroed out when the corresponding mask
// bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k1[j]
//				k[j] := (a[i+63:i] < b[i+63:i]) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:8] := 0
//
// Instruction: 'VCMPPD'. Intrinsic: '_mm512_mask_cmplt_pd_mask'.
// Requires KNCNI.
func M512MaskCmpltPdMask(k1 x86.Mmask8, a x86.M512d, b x86.M512d) (dst x86.Mmask8) {
	panic("not implemented")
}


// M512CmpltPsMask: Compare packed single-precision (32-bit) floating-point
// elements in 'a' and 'b' for less-than, and store the results in mask vector
// 'k'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			k[j] := (a[i+31:i] < b[i+31:i]) ? 1 : 0
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VCMPPS'. Intrinsic: '_mm512_cmplt_ps_mask'.
// Requires KNCNI.
func M512CmpltPsMask(a x86.M512, b x86.M512) (dst x86.Mmask16) {
	panic("not implemented")
}


// M512MaskCmpltPsMask: Compare packed single-precision (32-bit) floating-point
// elements in 'a' and 'b' for less-than, and store the results in mask vector
// 'k' using zeromask 'k1' (elements are zeroed out when the corresponding mask
// bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k1[j]
//				k[j] := (a[i+31:i] < b[i+31:i]) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VCMPPS'. Intrinsic: '_mm512_mask_cmplt_ps_mask'.
// Requires KNCNI.
func M512MaskCmpltPsMask(k1 x86.Mmask16, a x86.M512, b x86.M512) (dst x86.Mmask16) {
	panic("not implemented")
}


// M512CmpneqEpi32Mask: Compare packed 32-bit integers in 'a' and 'b' for
// not-equal, and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			k[j] := ( a[i+31:i] != b[i+31:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPD'. Intrinsic: '_mm512_cmpneq_epi32_mask'.
// Requires KNCNI.
func M512CmpneqEpi32Mask(a x86.M512i, b x86.M512i) (dst x86.Mmask16) {
	panic("not implemented")
}


// M512MaskCmpneqEpi32Mask: Compare packed 32-bit integers in 'a' and 'b' for
// not-equal, and store the results in mask vector 'k1' using zeromask 'k'
// (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k1[j]
//				k[j] := ( a[i+31:i] != b[i+31:i] ) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPD'. Intrinsic: '_mm512_mask_cmpneq_epi32_mask'.
// Requires KNCNI.
func M512MaskCmpneqEpi32Mask(k1 x86.Mmask16, a x86.M512i, b x86.M512i) (dst x86.Mmask16) {
	panic("not implemented")
}


// M512CmpneqEpu32Mask: Compare packed unsigned 32-bit integers in 'a' and 'b'
// for not-equal, and store the results in mask vector 'k'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			k[j] := ( a[i+31:i] != b[i+31:i] ) ? 1 : 0
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPUD'. Intrinsic: '_mm512_cmpneq_epu32_mask'.
// Requires KNCNI.
func M512CmpneqEpu32Mask(a x86.M512i, b x86.M512i) (dst x86.Mmask16) {
	panic("not implemented")
}


// M512MaskCmpneqEpu32Mask: Compare packed unsigned 32-bit integers in 'a' and
// 'b' for not-equal, and store the results in mask vector 'k1' using zeromask
// 'k' (elements are zeroed out when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k1[j]
//				k[j] := ( a[i+31:i] != b[i+31:i] ) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPCMPUD'. Intrinsic: '_mm512_mask_cmpneq_epu32_mask'.
// Requires KNCNI.
func M512MaskCmpneqEpu32Mask(k1 x86.Mmask16, a x86.M512i, b x86.M512i) (dst x86.Mmask16) {
	panic("not implemented")
}


// M512CmpneqPdMask: Compare packed double-precision (64-bit) floating-point
// elements in 'a' and 'b' for not-equal, and store the results in mask vector
// 'k'. 
//
//		FOR j := 0 to 7
//			i := j*64
//			k[j] := (a[i+63:i] != b[i+63:i]) ? 1 : 0
//		ENDFOR
//		k[MAX:8] := 0
//
// Instruction: 'VCMPPD'. Intrinsic: '_mm512_cmpneq_pd_mask'.
// Requires KNCNI.
func M512CmpneqPdMask(a x86.M512d, b x86.M512d) (dst x86.Mmask8) {
	panic("not implemented")
}


// M512MaskCmpneqPdMask: Compare packed double-precision (64-bit)
// floating-point elements in 'a' and 'b' for not-equal, and store the results
// in mask vector 'k' using zeromask 'k1' (elements are zeroed out when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k1[j]
//				k[j] := (a[i+63:i] != b[i+63:i]) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:8] := 0
//
// Instruction: 'VCMPPD'. Intrinsic: '_mm512_mask_cmpneq_pd_mask'.
// Requires KNCNI.
func M512MaskCmpneqPdMask(k1 x86.Mmask8, a x86.M512d, b x86.M512d) (dst x86.Mmask8) {
	panic("not implemented")
}


// M512CmpneqPsMask: Compare packed single-precision (32-bit) floating-point
// elements in 'a' and 'b' for not-equal, and store the results in mask vector
// 'k'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			k[j] := (a[i+31:i] != b[i+31:i]) ? 1 : 0
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VCMPPS'. Intrinsic: '_mm512_cmpneq_ps_mask'.
// Requires KNCNI.
func M512CmpneqPsMask(a x86.M512, b x86.M512) (dst x86.Mmask16) {
	panic("not implemented")
}


// M512MaskCmpneqPsMask: Compare packed single-precision (32-bit)
// floating-point elements in 'a' and 'b' for not-equal, and store the results
// in mask vector 'k' using zeromask 'k1' (elements are zeroed out when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k1[j]
//				k[j] := (a[i+31:i] != b[i+31:i]) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VCMPPS'. Intrinsic: '_mm512_mask_cmpneq_ps_mask'.
// Requires KNCNI.
func M512MaskCmpneqPsMask(k1 x86.Mmask16, a x86.M512, b x86.M512) (dst x86.Mmask16) {
	panic("not implemented")
}


// M512CmpnlePdMask: Compare packed double-precision (64-bit) floating-point
// elements in 'a' and 'b' for not-less-than-or-equal, and store the results in
// mask vector 'k'. 
//
//		FOR j := 0 to 7
//			i := j*64
//			k[j] := !(a[i+63:i] <= b[i+63:i]) ? 1 : 0
//		ENDFOR
//		k[MAX:8] := 0
//
// Instruction: 'VCMPPD'. Intrinsic: '_mm512_cmpnle_pd_mask'.
// Requires KNCNI.
func M512CmpnlePdMask(a x86.M512d, b x86.M512d) (dst x86.Mmask8) {
	panic("not implemented")
}


// M512MaskCmpnlePdMask: Compare packed double-precision (64-bit)
// floating-point elements in 'a' and 'b' for not-less-than-or-equal, and store
// the results in mask vector 'k' using zeromask 'k1' (elements are zeroed out
// when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k1[j]
//				k[j] := !(a[i+63:i] <= b[i+63:i]) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:8] := 0
//
// Instruction: 'VCMPPD'. Intrinsic: '_mm512_mask_cmpnle_pd_mask'.
// Requires KNCNI.
func M512MaskCmpnlePdMask(k1 x86.Mmask8, a x86.M512d, b x86.M512d) (dst x86.Mmask8) {
	panic("not implemented")
}


// M512CmpnlePsMask: Compare packed single-precision (32-bit) floating-point
// elements in 'a' and 'b' for not-less-than-or-equal, and store the results in
// mask vector 'k'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			k[j] := !(a[i+31:i] <= b[i+31:i]) ? 1 : 0
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VCMPPS'. Intrinsic: '_mm512_cmpnle_ps_mask'.
// Requires KNCNI.
func M512CmpnlePsMask(a x86.M512, b x86.M512) (dst x86.Mmask16) {
	panic("not implemented")
}


// M512MaskCmpnlePsMask: Compare packed single-precision (32-bit)
// floating-point elements in 'a' and 'b' for not-less-than-or-equal, and store
// the results in mask vector 'k' using zeromask 'k1' (elements are zeroed out
// when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k1[j]
//				k[j] := !(a[i+31:i] <= b[i+31:i]) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VCMPPS'. Intrinsic: '_mm512_mask_cmpnle_ps_mask'.
// Requires KNCNI.
func M512MaskCmpnlePsMask(k1 x86.Mmask16, a x86.M512, b x86.M512) (dst x86.Mmask16) {
	panic("not implemented")
}


// M512CmpnltPdMask: Compare packed double-precision (64-bit) floating-point
// elements in 'a' and 'b' for not-less-than, and store the results in mask
// vector 'k'. 
//
//		FOR j := 0 to 7
//			i := j*64
//			k[j] := !(a[i+63:i] < b[i+63:i]) ? 1 : 0
//		ENDFOR
//		k[MAX:8] := 0
//
// Instruction: 'VCMPPD'. Intrinsic: '_mm512_cmpnlt_pd_mask'.
// Requires KNCNI.
func M512CmpnltPdMask(a x86.M512d, b x86.M512d) (dst x86.Mmask8) {
	panic("not implemented")
}


// M512MaskCmpnltPdMask: Compare packed double-precision (64-bit)
// floating-point elements in 'a' and 'b' for not-less-than, and store the
// results in mask vector 'k' using zeromask 'k1' (elements are zeroed out when
// the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k1[j]
//				k[j] := !(a[i+63:i] < b[i+63:i]) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:8] := 0
//
// Instruction: 'VCMPPD'. Intrinsic: '_mm512_mask_cmpnlt_pd_mask'.
// Requires KNCNI.
func M512MaskCmpnltPdMask(k1 x86.Mmask8, a x86.M512d, b x86.M512d) (dst x86.Mmask8) {
	panic("not implemented")
}


// M512CmpnltPsMask: Compare packed single-precision (32-bit) floating-point
// elements in 'a' and 'b' for not-less-than, and store the results in mask
// vector 'k'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			k[j] := !(a[i+31:i] < b[i+31:i]) ? 1 : 0
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VCMPPS'. Intrinsic: '_mm512_cmpnlt_ps_mask'.
// Requires KNCNI.
func M512CmpnltPsMask(a x86.M512, b x86.M512) (dst x86.Mmask16) {
	panic("not implemented")
}


// M512MaskCmpnltPsMask: Compare packed single-precision (32-bit)
// floating-point elements in 'a' and 'b' for not-less-than, and store the
// results in mask vector 'k' using zeromask 'k1' (elements are zeroed out when
// the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k1[j]
//				k[j] := !(a[i+31:i] < b[i+31:i]) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VCMPPS'. Intrinsic: '_mm512_mask_cmpnlt_ps_mask'.
// Requires KNCNI.
func M512MaskCmpnltPsMask(k1 x86.Mmask16, a x86.M512, b x86.M512) (dst x86.Mmask16) {
	panic("not implemented")
}


// M512CmpordPdMask: Compare packed double-precision (64-bit) floating-point
// elements in 'a' and 'b' to see if neither is NaN, and store the results in
// mask vector 'k'. 
//
//		FOR j := 0 to 7
//			i := j*64
//			k[j] := (a[i+63:i] != NaN AND b[i+63:i] != NaN) ? 1 : 0 
//		ENDFOR
//		k[MAX:8] := 0
//
// Instruction: 'VCMPPD'. Intrinsic: '_mm512_cmpord_pd_mask'.
// Requires KNCNI.
func M512CmpordPdMask(a x86.M512d, b x86.M512d) (dst x86.Mmask8) {
	panic("not implemented")
}


// M512MaskCmpordPdMask: Compare packed double-precision (64-bit)
// floating-point elements in 'a' and 'b' to see if neither is NaN, and store
// the results in mask vector 'k' using zeromask 'k1' (elements are zeroed out
// when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k1[j]
//				k[j] := (a[i+63:i] != NaN AND b[i+63:i] != NaN) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:8] := 0
//
// Instruction: 'VCMPPD'. Intrinsic: '_mm512_mask_cmpord_pd_mask'.
// Requires KNCNI.
func M512MaskCmpordPdMask(k1 x86.Mmask8, a x86.M512d, b x86.M512d) (dst x86.Mmask8) {
	panic("not implemented")
}


// M512CmpordPsMask: Compare packed single-precision (32-bit) floating-point
// elements in 'a' and 'b' to see if neither is NaN, and store the results in
// mask vector 'k'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			k[j] := (a[i+31:i] != NaN AND b[i+31:i] != NaN) ? 1 : 0 
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VCMPPS'. Intrinsic: '_mm512_cmpord_ps_mask'.
// Requires KNCNI.
func M512CmpordPsMask(a x86.M512, b x86.M512) (dst x86.Mmask16) {
	panic("not implemented")
}


// M512MaskCmpordPsMask: Compare packed single-precision (32-bit)
// floating-point elements in 'a' and 'b' to see if neither is NaN, and store
// the results in mask vector 'k' using zeromask 'k1' (elements are zeroed out
// when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k1[j]
//				k[j] := (a[i+31:i] != NaN AND b[i+31:i] != NaN) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VCMPPS'. Intrinsic: '_mm512_mask_cmpord_ps_mask'.
// Requires KNCNI.
func M512MaskCmpordPsMask(k1 x86.Mmask16, a x86.M512, b x86.M512) (dst x86.Mmask16) {
	panic("not implemented")
}


// M512CmpunordPdMask: Compare packed double-precision (64-bit) floating-point
// elements in 'a' and 'b' to see if either is NaN, and store the results in
// mask vector 'k'. 
//
//		FOR j := 0 to 7
//			i := j*64
//			k[j] := (a[i+63:i] == NaN OR b[i+63:i] == NaN) ? 1 : 0 
//		ENDFOR
//		k[MAX:8] := 0
//
// Instruction: 'VCMPPD'. Intrinsic: '_mm512_cmpunord_pd_mask'.
// Requires KNCNI.
func M512CmpunordPdMask(a x86.M512d, b x86.M512d) (dst x86.Mmask8) {
	panic("not implemented")
}


// M512MaskCmpunordPdMask: Compare packed double-precision (64-bit)
// floating-point elements in 'a' and 'b' to see if either is NaN, and store
// the results in mask vector 'k' using zeromask 'k1' (elements are zeroed out
// when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k1[j]
//				k[j] := (a[i+63:i] == NaN OR b[i+63:i] == NaN) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:8] := 0
//
// Instruction: 'VCMPPD'. Intrinsic: '_mm512_mask_cmpunord_pd_mask'.
// Requires KNCNI.
func M512MaskCmpunordPdMask(k1 x86.Mmask8, a x86.M512d, b x86.M512d) (dst x86.Mmask8) {
	panic("not implemented")
}


// M512CmpunordPsMask: Compare packed single-precision (32-bit) floating-point
// elements in 'a' and 'b' to see if either is NaN, and store the results in
// mask vector 'k'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			k[j] := (a[i+31:i] == NaN OR b[i+31:i] == NaN) ? 1 : 0 
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VCMPPS'. Intrinsic: '_mm512_cmpunord_ps_mask'.
// Requires KNCNI.
func M512CmpunordPsMask(a x86.M512, b x86.M512) (dst x86.Mmask16) {
	panic("not implemented")
}


// M512MaskCmpunordPsMask: Compare packed single-precision (32-bit)
// floating-point elements in 'a' and 'b' to see if either is NaN, and store
// the results in mask vector 'k' using zeromask 'k1' (elements are zeroed out
// when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k1[j]
//				k[j] := (a[i+31:i] == NaN OR b[i+31:i] == NaN) ? 1 : 0
//			ELSE 
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VCMPPS'. Intrinsic: '_mm512_mask_cmpunord_ps_mask'.
// Requires KNCNI.
func M512MaskCmpunordPsMask(k1 x86.Mmask16, a x86.M512, b x86.M512) (dst x86.Mmask16) {
	panic("not implemented")
}


// Countbits32: Counts the number of set bits in 32-bit unsigned integer 'r1',
// returning the results in 'dst'. 
//
//		dst[31:0] := PopCount(r1[31:0])
//
// Instruction: 'POPCNT'. Intrinsic: '_mm_countbits_32'.
// Requires KNCNI.
func Countbits32(r1 uint32) uint32 {
	panic("not implemented")
}


// Countbits64: Counts the number of set bits in double-precision (32-bit)
// unsigned integer 'r1', returning the results in 'dst'. 
//
//		dst[63:0] := PopCount(r1[63:0])
//
// Instruction: 'POPCNT'. Intrinsic: '_mm_countbits_64'.
// Requires KNCNI.
func Countbits64(r1 uint64) uint64 {
	panic("not implemented")
}


// M512CvtRoundpdPslo: Performs element-by-element conversion of packed
// double-precision (64-bit) floating-point elements in 'v2' to packed
// single-precision (32-bit) floating-point elements, storing the results in
// 'dst'. Results are written to the lower half of 'dst', and the upper half
// locations are set to '0'.
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 7
//			i := j*64
//			k := j*32
//			dst[k+31:k] := Float64ToFloat32(v2[i+63:i])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VCVTPD2PS'. Intrinsic: '_mm512_cvt_roundpd_pslo'.
// Requires KNCNI.
func M512CvtRoundpdPslo(v2 x86.M512d, rounding int) (dst x86.M512) {
	panic("not implemented")
}


// M512MaskCvtRoundpdPslo: Performs element-by-element conversion of packed
// double-precision (64-bit) floating-point elements in 'v2' to packed
// single-precision (32-bit) floating-point elements, storing the results in
// 'dst' using writemask 'k' (elements are copied from 'src' when the
// corresponding mask bit is not set). Results are written to the lower half of
// 'dst', and the upper half locations are set to '0'.
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 7
//			i := j*64
//			l := j*32
//			IF k[j]
//				dst[l+31:l] := Float64ToFloat32(v2[i+63:i])
//			ELSE
//				dst[l+31:l] := src[l+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VCVTPD2PS'. Intrinsic: '_mm512_mask_cvt_roundpd_pslo'.
// Requires KNCNI.
func M512MaskCvtRoundpdPslo(src x86.M512, k x86.Mmask8, v2 x86.M512d, rounding int) (dst x86.M512) {
	panic("not implemented")
}


// M512Cvtepi32loPd: Performs element-by-element conversion of the lower half
// of packed 32-bit integer elements in 'v2' to packed double-precision
// (64-bit) floating-point elements, storing the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			l := j*64
//			dst[l+63:l] := Int32ToFloat64(v2[i+31:i])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VCVTDQ2PD'. Intrinsic: '_mm512_cvtepi32lo_pd'.
// Requires KNCNI.
func M512Cvtepi32loPd(v2 x86.M512i) (dst x86.M512d) {
	panic("not implemented")
}


// M512MaskCvtepi32loPd: Performs element-by-element conversion of the lower
// half of packed 32-bit integer elements in 'v2' to packed double-precision
// (64-bit) floating-point elements, storing the results in 'dst' using
// writemask 'k' (elements are copied from 'src' when the corresponding mask
// bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*32
//			n := j*64
//			IF k[j]
//				dst[k+63:k] := Int32ToFloat64(v2[i+31:i])
//			ELSE
//				dst[n+63:n] := src[n+63:n]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VCVTDQ2PD'. Intrinsic: '_mm512_mask_cvtepi32lo_pd'.
// Requires KNCNI.
func M512MaskCvtepi32loPd(src x86.M512d, k x86.Mmask8, v2 x86.M512i) (dst x86.M512d) {
	panic("not implemented")
}


// M512Cvtepu32loPd: Performs element-by-element conversion of the lower half
// of packed 32-bit unsigned integer elements in 'v2' to packed
// double-precision (64-bit) floating-point elements, storing the results in
// 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			k := j*64
//			dst[k+63:k] := UInt32ToFloat64(v2[i+31:i])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VCVTUDQ2PD'. Intrinsic: '_mm512_cvtepu32lo_pd'.
// Requires KNCNI.
func M512Cvtepu32loPd(v2 x86.M512i) (dst x86.M512d) {
	panic("not implemented")
}


// M512MaskCvtepu32loPd: Performs element-by-element conversion of the lower
// half of 32-bit unsigned integer elements in 'v2' to packed double-precision
// (64-bit) floating-point elements, storing the results in 'dst' using
// writemask 'k' (elements are copied from 'src' when the corresponding mask
// bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*32
//			l := j*64
//			IF k[j]
//				dst[l+63:l] := UInt32ToFloat64(v2[i+31:i])
//			ELSE
//				dst[l+63:l] := src[l+63:l]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VCVTUDQ2PD'. Intrinsic: '_mm512_mask_cvtepu32lo_pd'.
// Requires KNCNI.
func M512MaskCvtepu32loPd(src x86.M512d, k x86.Mmask8, v2 x86.M512i) (dst x86.M512d) {
	panic("not implemented")
}


// M512CvtfxpntRoundAdjustepi32Ps: Performs element-by-element conversion of
// packed 32-bit integer elements in 'v2' to packed single-precision (32-bit)
// floating-point elements and performing an optional exponent adjust using
// 'expadj', storing the results in 'dst'.
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := Int32ToFloat32(v2[i+31:i])
//			CASE expadj OF
//			_MM_EXPADJ_NONE: dst[i+31:i] = dst[i+31:i] * 2**0
//			_MM_EXPADJ_4:	dst[i+31:i] = dst[i+31:i] * 2**4
//			_MM_EXPADJ_5:	dst[i+31:i] = dst[i+31:i] * 2**5
//			_MM_EXPADJ_8:	dst[i+31:i] = dst[i+31:i] * 2**8
//			_MM_EXPADJ_16:   dst[i+31:i] = dst[i+31:i] * 2**16
//			_MM_EXPADJ_24:   dst[i+31:i] = dst[i+31:i] * 2**24
//			_MM_EXPADJ_31:   dst[i+31:i] = dst[i+31:i] * 2**31
//			_MM_EXPADJ_32:   dst[i+31:i] = dst[i+31:i] * 2**32
//			ESAC
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VCVTFXPNTDQ2PS'. Intrinsic: '_mm512_cvtfxpnt_round_adjustepi32_ps'.
// Requires KNCNI.
func M512CvtfxpntRoundAdjustepi32Ps(v2 x86.M512i, rounding int, expadj MMEXPADJENUM) (dst x86.M512) {
	panic("not implemented")
}


// M512CvtfxpntRoundAdjustepu32Ps: Performs element-by-element conversion of
// packed 32-bit unsigned integer elements in 'v2' to packed single-precision
// (32-bit) floating-point elements and performing an optional exponent adjust
// using 'expadj', storing the results in 'dst'.
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := UInt32ToFloat32(v2[i+31:i])
//			CASE expadj OF
//			_MM_EXPADJ_NONE: dst[i+31:i] = dst[i+31:i] * 2**0
//			_MM_EXPADJ_4:	 dst[i+31:i] = dst[i+31:i] * 2**4
//			_MM_EXPADJ_5:	 dst[i+31:i] = dst[i+31:i] * 2**5
//			_MM_EXPADJ_8:	 dst[i+31:i] = dst[i+31:i] * 2**8
//			_MM_EXPADJ_16:   dst[i+31:i] = dst[i+31:i] * 2**16
//			_MM_EXPADJ_24:   dst[i+31:i] = dst[i+31:i] * 2**24
//			_MM_EXPADJ_31:   dst[i+31:i] = dst[i+31:i] * 2**31
//			_MM_EXPADJ_32:   dst[i+31:i] = dst[i+31:i] * 2**32
//			ESAC
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VCVTFXPNTUDQ2PS'. Intrinsic: '_mm512_cvtfxpnt_round_adjustepu32_ps'.
// Requires KNCNI.
func M512CvtfxpntRoundAdjustepu32Ps(v2 x86.M512i, rounding int, expadj MMEXPADJENUM) (dst x86.M512) {
	panic("not implemented")
}


// M512MaskCvtfxpntRoundAdjustepu32Ps: Performs element-by-element conversion
// of packed 32-bit unsigned integer elements in 'v2' to packed
// single-precision (32-bit) floating-point elements and performing an optional
// exponent adjust using 'expadj', storing the results in 'dst' using writemask
// 'k' (elements are copied from 'src' when the corresponding mask bit is not
// set).
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := Int32ToFloat32(v2[i+31:i])
//				CASE expadj OF
//				_MM_EXPADJ_NONE: dst[i+31:i] = dst[i+31:i] * 2**0
//				_MM_EXPADJ_4:	 dst[i+31:i] = dst[i+31:i] * 2**4
//				_MM_EXPADJ_5:	 dst[i+31:i] = dst[i+31:i] * 2**5
//				_MM_EXPADJ_8:	 dst[i+31:i] = dst[i+31:i] * 2**8
//				_MM_EXPADJ_16:   dst[i+31:i] = dst[i+31:i] * 2**16
//				_MM_EXPADJ_24:   dst[i+31:i] = dst[i+31:i] * 2**24
//				_MM_EXPADJ_31:   dst[i+31:i] = dst[i+31:i] * 2**31
//				_MM_EXPADJ_32:   dst[i+31:i] = dst[i+31:i] * 2**32
//				ESAC
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VCVTFXPNTUDQ2PS'. Intrinsic: '_mm512_mask_cvtfxpnt_round_adjustepu32_ps'.
// Requires KNCNI.
func M512MaskCvtfxpntRoundAdjustepu32Ps(src x86.M512, k x86.Mmask16, v2 x86.M512i, rounding int, expadj MMEXPADJENUM) (dst x86.M512) {
	panic("not implemented")
}


// M512CvtfxpntRoundAdjustpsEpi32: Performs element-by-element conversion of
// packed single-precision (32-bit) floating-point elements in 'v2' to packed
// 32-bit integer elements and performs an optional exponent adjust using
// 'expadj', storing the results in 'dst'.
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := Float32ToInt32(v2[i+31:i])
//			CASE expadj OF
//			_MM_EXPADJ_NONE: dst[i+31:i] = dst[i+31:i] * 2**0
//			_MM_EXPADJ_4:	 dst[i+31:i] = dst[i+31:i] * 2**4
//			_MM_EXPADJ_5:	 dst[i+31:i] = dst[i+31:i] * 2**5
//			_MM_EXPADJ_8:	 dst[i+31:i] = dst[i+31:i] * 2**8
//			_MM_EXPADJ_16:   dst[i+31:i] = dst[i+31:i] * 2**16
//			_MM_EXPADJ_24:   dst[i+31:i] = dst[i+31:i] * 2**24
//			_MM_EXPADJ_31:   dst[i+31:i] = dst[i+31:i] * 2**31
//			_MM_EXPADJ_32:   dst[i+31:i] = dst[i+31:i] * 2**32
//			ESAC
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VCVTFXPNTPS2DQ'. Intrinsic: '_mm512_cvtfxpnt_round_adjustps_epi32'.
// Requires KNCNI.
func M512CvtfxpntRoundAdjustpsEpi32(v2 x86.M512, rounding int, expadj MMEXPADJENUM) (dst x86.M512i) {
	panic("not implemented")
}


// M512CvtfxpntRoundAdjustpsEpu32: Performs element-by-element conversion of
// packed single-precision (32-bit) floating-point elements in 'v2' to packed
// 32-bit unsigned integer elements and performing an optional exponent adjust
// using 'expadj', storing the results in 'dst'.
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := Float32ToUInt32(v2[i+31:i])
//			CASE expadj OF
//			_MM_EXPADJ_NONE: dst[i+31:i] = dst[i+31:i]  0
//			_MM_EXPADJ_4:	 dst[i+31:i] = dst[i+31:i]  4
//			_MM_EXPADJ_5:	 dst[i+31:i] = dst[i+31:i]  5
//			_MM_EXPADJ_8:	 dst[i+31:i] = dst[i+31:i]  8
//			_MM_EXPADJ_16:   dst[i+31:i] = dst[i+31:i]  16
//			_MM_EXPADJ_24:   dst[i+31:i] = dst[i+31:i]  24
//			_MM_EXPADJ_31:   dst[i+31:i] = dst[i+31:i]  31
//			_MM_EXPADJ_32:   dst[i+31:i] = dst[i+31:i]  32
//			ESAC
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VCVTFXPNTPS2UDQ'. Intrinsic: '_mm512_cvtfxpnt_round_adjustps_epu32'.
// Requires KNCNI.
func M512CvtfxpntRoundAdjustpsEpu32(v2 x86.M512, rounding int, expadj MMEXPADJENUM) (dst x86.M512i) {
	panic("not implemented")
}


// M512CvtfxpntRoundpdEpi32lo: Performs an element-by-element conversion of
// elements in packed double-precision (64-bit) floating-point vector 'v2' to
// 32-bit integer elements, storing them in the lower half of 'dst'. The
// elements in the upper half of 'dst' are set to 0.
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 7
//			i := j*64
//			k := j*32
//			dst[k+31:k] := Float64ToInt32(v2[i+63:i])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VCVTFXPNTPD2DQ'. Intrinsic: '_mm512_cvtfxpnt_roundpd_epi32lo'.
// Requires KNCNI.
func M512CvtfxpntRoundpdEpi32lo(v2 x86.M512d, rounding int) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskCvtfxpntRoundpdEpi32lo: Performs an element-by-element conversion of
// elements in packed double-precision (64-bit) floating-point vector 'v2' to
// 32-bit integer elements, storing them in the lower half of 'dst' using
// writemask 'k' (elements are copied from 'src' when the corresponding mask
// bit is not set). The elements in the upper half of 'dst' are set to 0.
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 7
//			i := j*64
//			l := j*32
//			IF k[j]
//				dst[l+31:l] := Float64ToInt32(v2[i+63:i])
//			ELSE
//				dst[l+31:l] := src[l+31:l]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VCVTFXPNTPD2DQ'. Intrinsic: '_mm512_mask_cvtfxpnt_roundpd_epi32lo'.
// Requires KNCNI.
func M512MaskCvtfxpntRoundpdEpi32lo(src x86.M512i, k x86.Mmask8, v2 x86.M512d, rounding int) (dst x86.M512i) {
	panic("not implemented")
}


// M512CvtfxpntRoundpdEpu32lo: Performs element-by-element conversion of packed
// double-precision (64-bit) floating-point elements in 'v2' to packed 32-bit
// unsigned integer elements, storing the results in 'dst'. Results are written
// to the lower half of 'dst', and the upper half locations are set to '0'.
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 7
//			i := j*64
//			k := j*32
//			dst[k+31:k] := Float64ToInt32(v2[i+63:i])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VCVTFXPNTPD2UDQ'. Intrinsic: '_mm512_cvtfxpnt_roundpd_epu32lo'.
// Requires KNCNI.
func M512CvtfxpntRoundpdEpu32lo(v2 x86.M512d, rounding int) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskCvtfxpntRoundpdEpu32lo: Performs element-by-element conversion of
// packed double-precision (64-bit) floating-point elements in 'v2' to packed
// 32-bit unsigned integer elements, storing the results in 'dst' using
// writemask 'k' (elements are copied from 'src' when the corresponding mask
// bit is not set). Results are written to the lower half of 'dst', and the
// upper half locations are set to '0'.
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 7
//			i := j*64
//			l := j*32
//			IF k[j]
//				dst[l+31:l] := Float64ToInt32(v2[i+63:i])
//			ELSE
//				dst[l+31:l] := src[l+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VCVTFXPNTPD2UDQ'. Intrinsic: '_mm512_mask_cvtfxpnt_roundpd_epu32lo'.
// Requires KNCNI.
func M512MaskCvtfxpntRoundpdEpu32lo(src x86.M512i, k x86.Mmask8, v2 x86.M512d, rounding int) (dst x86.M512i) {
	panic("not implemented")
}


// M512CvtpdPslo: Performs an element-by-element conversion of packed
// double-precision (64-bit) floating-point elements in 'v2' to
// single-precision (32-bit) floating-point elements and stores them in 'dst'.
// The elements are stored in the lower half of the results vector, while the
// remaining upper half locations are set to 0. 
//
//		FOR j := 0 to 7
//			i := j*64
//			k := j*32
//			dst[k+31:k] := Float64ToFloat32(v2[i+63:i])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VCVTPD2PS'. Intrinsic: '_mm512_cvtpd_pslo'.
// Requires KNCNI.
func M512CvtpdPslo(v2 x86.M512d) (dst x86.M512) {
	panic("not implemented")
}


// M512MaskCvtpdPslo: Performs an element-by-element conversion of packed
// double-precision (64-bit) floating-point elements in 'v2' to
// single-precision (32-bit) floating-point elements and stores them in 'dst'
// using writemask 'k' (elements are copied from 'src' when the corresponding
// mask bit is not set). The elements are stored in the lower half of the
// results vector, while the remaining upper half locations are set to 0. 
//
//		FOR j := 0 to 7
//			i := j*64
//			l := j*32
//			IF k[j]
//				dst[l+31:l] := Float64ToFloat32(v2[i+63:i])
//			ELSE
//				dst[l+31:l] := src[l+31:l]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VCVTPD2PS'. Intrinsic: '_mm512_mask_cvtpd_pslo'.
// Requires KNCNI.
func M512MaskCvtpdPslo(src x86.M512, k x86.Mmask8, v2 x86.M512d) (dst x86.M512) {
	panic("not implemented")
}


// M512CvtpsloPd: Performs element-by-element conversion of the lower half of
// packed single-precision (32-bit) floating-point elements in 'v2' to packed
// double-precision (64-bit) floating-point elements, storing the results in
// 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*32
//			k := j*64
//			dst[k+63:k] := Float32ToFloat64(v2[i+31:i])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VCVTPS2PD'. Intrinsic: '_mm512_cvtpslo_pd'.
// Requires KNCNI.
func M512CvtpsloPd(v2 x86.M512) (dst x86.M512d) {
	panic("not implemented")
}


// M512MaskCvtpsloPd: Performs element-by-element conversion of the lower half
// of packed single-precision (32-bit) floating-point elements in 'v2' to
// packed double-precision (64-bit) floating-point elements, storing the
// results in 'dst' using writemask 'k' (elements are copied from 'src' when
// the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*32
//			l := j*64
//			IF k[j]
//				dst[l+63:l] := Float32ToFloat64(v2[i+31:i])
//			ELSE
//				dst[l+63:l] := src[l+63:l]:
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VCVTPS2PD'. Intrinsic: '_mm512_mask_cvtpslo_pd'.
// Requires KNCNI.
func M512MaskCvtpsloPd(src x86.M512d, k x86.Mmask8, v2 x86.M512) (dst x86.M512d) {
	panic("not implemented")
}


// Delay32: Stalls a thread without blocking other threads for 32-bit unsigned
// integer 'r1' clock cycles. 
//
//		BlockThread(r1)
//
// Instruction: 'DELAY'. Intrinsic: '_mm_delay_32'.
// Requires KNCNI.
func Delay32(r1 uint32)  {
	panic("not implemented")
}


// Delay64: Stalls a thread without blocking other threads for 64-bit unsigned
// integer 'r1' clock cycles. 
//
//		BlockThread(r1)
//
// Instruction: 'DELAY'. Intrinsic: '_mm_delay_64'.
// Requires KNCNI.
func Delay64(r1 uint64)  {
	panic("not implemented")
}


// M512Exp223Ps: Approximates the base-2 exponent of the packed
// single-precision (32-bit) floating-point elements in 'v2' with eight bits
// for sign and magnitude and 24 bits for the fractional part. Results are
// stored in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := exp223(v2[i+31:i])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VEXP223PS'. Intrinsic: '_mm512_exp223_ps'.
// Requires KNCNI.
func M512Exp223Ps(v2 x86.M512i) (dst x86.M512) {
	panic("not implemented")
}


// M512MaskExp223Ps: Approximates the base-2 exponent of the packed
// single-precision (32-bit) floating-point elements in 'v2' with eight bits
// for sign and magnitude and 24 bits for the fractional part. Results are
// stored in 'dst' using writemask 'k' (elements are copied from 'src' when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := exp223(v2[i+31:i])
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VEXP223PS'. Intrinsic: '_mm512_mask_exp223_ps'.
// Requires KNCNI.
func M512MaskExp223Ps(src x86.M512, k x86.Mmask16, v2 x86.M512i) (dst x86.M512) {
	panic("not implemented")
}


// Skipped: _mm512_extload_epi32. Contains pointer parameter.


// Skipped: _mm512_mask_extload_epi32. Contains pointer parameter.


// Skipped: _mm512_extload_epi64. Contains pointer parameter.


// Skipped: _mm512_mask_extload_epi64. Contains pointer parameter.


// Skipped: _mm512_extload_pd. Contains pointer parameter.


// Skipped: _mm512_mask_extload_pd. Contains pointer parameter.


// Skipped: _mm512_extload_ps. Contains pointer parameter.


// Skipped: _mm512_mask_extload_ps. Contains pointer parameter.


// Skipped: _mm512_extloadunpackhi_epi32. Contains pointer parameter.


// Skipped: _mm512_mask_extloadunpackhi_epi32. Contains pointer parameter.


// Skipped: _mm512_extloadunpackhi_epi64. Contains pointer parameter.


// Skipped: _mm512_mask_extloadunpackhi_epi64. Contains pointer parameter.


// Skipped: _mm512_extloadunpackhi_pd. Contains pointer parameter.


// Skipped: _mm512_mask_extloadunpackhi_pd. Contains pointer parameter.


// Skipped: _mm512_extloadunpackhi_ps. Contains pointer parameter.


// Skipped: _mm512_mask_extloadunpackhi_ps. Contains pointer parameter.


// Skipped: _mm512_extloadunpacklo_epi32. Contains pointer parameter.


// Skipped: _mm512_mask_extloadunpacklo_epi32. Contains pointer parameter.


// Skipped: _mm512_extloadunpacklo_epi64. Contains pointer parameter.


// Skipped: _mm512_mask_extloadunpacklo_epi64. Contains pointer parameter.


// Skipped: _mm512_extloadunpacklo_pd. Contains pointer parameter.


// Skipped: _mm512_mask_extloadunpacklo_pd. Contains pointer parameter.


// Skipped: _mm512_extloadunpacklo_ps. Contains pointer parameter.


// Skipped: _mm512_mask_extloadunpacklo_ps. Contains pointer parameter.


// Skipped: _mm512_extpackstorehi_epi32. Contains pointer parameter.


// Skipped: _mm512_mask_extpackstorehi_epi32. Contains pointer parameter.


// Skipped: _mm512_extpackstorehi_epi64. Contains pointer parameter.


// Skipped: _mm512_mask_extpackstorehi_epi64. Contains pointer parameter.


// Skipped: _mm512_extpackstorehi_pd. Contains pointer parameter.


// Skipped: _mm512_mask_extpackstorehi_pd. Contains pointer parameter.


// Skipped: _mm512_extpackstorehi_ps. Contains pointer parameter.


// Skipped: _mm512_mask_extpackstorehi_ps. Contains pointer parameter.


// Skipped: _mm512_extpackstorelo_epi32. Contains pointer parameter.


// Skipped: _mm512_mask_extpackstorelo_epi32. Contains pointer parameter.


// Skipped: _mm512_extpackstorelo_epi64. Contains pointer parameter.


// Skipped: _mm512_mask_extpackstorelo_epi64. Contains pointer parameter.


// Skipped: _mm512_extpackstorelo_pd. Contains pointer parameter.


// Skipped: _mm512_mask_extpackstorelo_pd. Contains pointer parameter.


// Skipped: _mm512_extpackstorelo_ps. Contains pointer parameter.


// Skipped: _mm512_mask_extpackstorelo_ps. Contains pointer parameter.


// Skipped: _mm512_extstore_epi32. Contains pointer parameter.


// Skipped: _mm512_mask_extstore_epi32. Contains pointer parameter.


// Skipped: _mm512_extstore_epi64. Contains pointer parameter.


// Skipped: _mm512_mask_extstore_epi64. Contains pointer parameter.


// Skipped: _mm512_extstore_pd. Contains pointer parameter.


// Skipped: _mm512_mask_extstore_pd. Contains pointer parameter.


// Skipped: _mm512_extstore_ps. Contains pointer parameter.


// Skipped: _mm512_mask_extstore_ps. Contains pointer parameter.


// M512FixupnanPd: Fixes up NaN's from packed double-precision (64-bit)
// floating-point elements in 'v1' and 'v2', storing the results in 'dst' and
// storing the quietized NaN's from 'v1' in 'v3'. 
//
//		FOR j := 0 to 7
//			i := j*64
//			dst[i+63:i] := FixupNaNs(v1[i+63:i], v2[i+63:i])
//			v3[i+63:i] := QuietizeNaNs(v1[i+63:i])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VFIXUPNANPD'. Intrinsic: '_mm512_fixupnan_pd'.
// Requires KNCNI.
func M512FixupnanPd(v1 x86.M512d, v2 x86.M512d, v3 x86.M512i) (dst x86.M512d) {
	panic("not implemented")
}


// M512MaskFixupnanPd: Fixes up NaN's from packed double-precision (64-bit)
// floating-point elements in 'v1' and 'v2', storing the results in 'dst' using
// writemask 'k' (only elements whose corresponding mask bit is set are used in
// the computation). Quietized NaN's from 'v1' are stored in 'v3'. 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := FixupNaNs(v1[i+63:i], v2[i+63:i])
//				v3[i+63:i] := QuietizeNaNs(v1[i+63:i])
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VFIXUPNANPD'. Intrinsic: '_mm512_mask_fixupnan_pd'.
// Requires KNCNI.
func M512MaskFixupnanPd(v1 x86.M512d, k x86.Mmask8, v2 x86.M512d, v3 x86.M512i) (dst x86.M512d) {
	panic("not implemented")
}


// M512FixupnanPs: Fixes up NaN's from packed single-precision (32-bit)
// floating-point elements in 'v1' and 'v2', storing the results in 'dst' and
// storing the quietized NaN's from 'v1' in 'v3'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := FixupNaNs(v1[i+31:i], v2[i+31:i])
//			v3[i+31:i] := QuietizeNaNs(v1[i+31:i])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VFIXUPNANPS'. Intrinsic: '_mm512_fixupnan_ps'.
// Requires KNCNI.
func M512FixupnanPs(v1 x86.M512, v2 x86.M512, v3 x86.M512i) (dst x86.M512) {
	panic("not implemented")
}


// M512MaskFixupnanPs: Fixes up NaN's from packed single-precision (32-bit)
// floating-point elements in 'v1' and 'v2', storing the results in 'dst' using
// writemask 'k' (only elements whose corresponding mask bit is set are used in
// the computation). Quietized NaN's from 'v1' are stored in 'v3'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := FixupNaNs(v1[i+31:i], v2[i+31:i])
//				v3[i+31:i] := QuietizeNaNs(v1[i+31:i])
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VFIXUPNANPS'. Intrinsic: '_mm512_mask_fixupnan_ps'.
// Requires KNCNI.
func M512MaskFixupnanPs(v1 x86.M512, k x86.Mmask16, v2 x86.M512, v3 x86.M512i) (dst x86.M512) {
	panic("not implemented")
}


// M512FmaddEpi32: Multiply packed 32-bit integer elements in 'a' and 'b', add
// the intermediate result to packed elements in 'c' and store the results in
// 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := (a[i+31:i] * b[i+31:i]) + c[i+31:i]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPMADD231D'. Intrinsic: '_mm512_fmadd_epi32'.
// Requires KNCNI.
func M512FmaddEpi32(a x86.M512i, b x86.M512i, c x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskFmaddEpi32: Multiply packed 32-bit integer elements in 'a' and 'b',
// add the intermediate result to packed elements in 'c' and store the results
// in 'dst' using writemask 'k' (elements are copied from 'a' when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := (a[i+31:i] * b[i+31:i]) + c[i+31:i]
//			ELSE
//				dst[i+31:i] := a[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPMADD231D'. Intrinsic: '_mm512_mask_fmadd_epi32'.
// Requires KNCNI.
func M512MaskFmaddEpi32(a x86.M512i, k x86.Mmask16, b x86.M512i, c x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512Mask3FmaddEpi32: Multiply packed 32-bit integer elements in 'a' and 'b',
// add the intermediate result to packed elements in 'c' and store the results
// in 'dst' using writemask 'k' (elements are copied from 'c' when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := (a[i+31:i] * b[i+31:i]) + c[i+31:i]
//			ELSE
//				dst[i+31:i] := c[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPMADD231D'. Intrinsic: '_mm512_mask3_fmadd_epi32'.
// Requires KNCNI.
func M512Mask3FmaddEpi32(a x86.M512i, b x86.M512i, c x86.M512i, k x86.Mmask16) (dst x86.M512i) {
	panic("not implemented")
}


// M512FmaddPd: Multiply packed double-precision (64-bit) floating-point
// elements in 'a' and 'b', add the intermediate result to packed elements in
// 'c', and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*64
//			dst[i+63:i] := (a[i+63:i] * b[i+63:i]) + c[i+63:i]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VFMADD132PD, VFMADD213PD, VFMADD231PD'. Intrinsic: '_mm512_fmadd_pd'.
// Requires KNCNI.
func M512FmaddPd(a x86.M512d, b x86.M512d, c x86.M512d) (dst x86.M512d) {
	panic("not implemented")
}


// M512MaskFmaddPd: Multiply packed double-precision (64-bit) floating-point
// elements in 'a' and 'b', add the intermediate result to packed elements in
// 'c', and store the results in 'dst' using writemask 'k' (elements are copied
// from 'a' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := (a[i+63:i] * b[i+63:i]) + c[i+63:i]
//			ELSE
//				dst[i+63:i] := a[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VFMADD132PD, VFMADD213PD, VFMADD231PD'. Intrinsic: '_mm512_mask_fmadd_pd'.
// Requires KNCNI.
func M512MaskFmaddPd(a x86.M512d, k x86.Mmask8, b x86.M512d, c x86.M512d) (dst x86.M512d) {
	panic("not implemented")
}


// M512Mask3FmaddPd: Multiply packed double-precision (64-bit) floating-point
// elements in 'a' and 'b', add the intermediate result to packed elements in
// 'c', and store the results in 'dst' using writemask 'k' (elements are copied
// from 'c' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := (a[i+63:i] * b[i+63:i]) + c[i+63:i]
//			ELSE
//				dst[i+63:i] := c[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VFMADD132PD, VFMADD213PD, VFMADD231PD'. Intrinsic: '_mm512_mask3_fmadd_pd'.
// Requires KNCNI.
func M512Mask3FmaddPd(a x86.M512d, b x86.M512d, c x86.M512d, k x86.Mmask8) (dst x86.M512d) {
	panic("not implemented")
}


// M512FmaddPs: Multiply packed single-precision (32-bit) floating-point
// elements in 'a' and 'b', add the intermediate result to packed elements in
// 'c', and store the results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := (a[i+31:i] * b[i+31:i]) + c[i+31:i]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VFMADD132PS, VFMADD213PS, VFMADD231PS'. Intrinsic: '_mm512_fmadd_ps'.
// Requires KNCNI.
func M512FmaddPs(a x86.M512, b x86.M512, c x86.M512) (dst x86.M512) {
	panic("not implemented")
}


// M512MaskFmaddPs: Multiply packed single-precision (32-bit) floating-point
// elements in 'a' and 'b', add the intermediate result to packed elements in
// 'c', and store the results in 'dst' using writemask 'k' (elements are copied
// from 'a' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := (a[i+31:i] * b[i+31:i]) + c[i+31:i]
//			ELSE
//				dst[i+31:i] := a[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VFMADD132PS, VFMADD213PS, VFMADD231PS'. Intrinsic: '_mm512_mask_fmadd_ps'.
// Requires KNCNI.
func M512MaskFmaddPs(a x86.M512, k x86.Mmask16, b x86.M512, c x86.M512) (dst x86.M512) {
	panic("not implemented")
}


// M512Mask3FmaddPs: Multiply packed single-precision (32-bit) floating-point
// elements in 'a' and 'b', add the intermediate result to packed elements in
// 'c', and store the results in 'dst' using writemask 'k' (elements are copied
// from 'c' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := (a[i+31:i] * b[i+31:i]) + c[i+31:i]
//			ELSE
//				dst[i+31:i] := c[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VFMADD132PS, VFMADD213PS, VFMADD231PS'. Intrinsic: '_mm512_mask3_fmadd_ps'.
// Requires KNCNI.
func M512Mask3FmaddPs(a x86.M512, b x86.M512, c x86.M512, k x86.Mmask16) (dst x86.M512) {
	panic("not implemented")
}


// M512FmaddRoundPd: Multiply packed double-precision (64-bit) floating-point
// elements in 'a' and 'b', add the intermediate result to packed elements in
// 'c', and store the results in 'dst'. 
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 7
//			i := j*64
//			dst[i+63:i] := (a[i+63:i] * b[i+63:i]) + c[i+63:i]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VFMADD132PD, VFMADD213PD, VFMADD231PD'. Intrinsic: '_mm512_fmadd_round_pd'.
// Requires KNCNI.
func M512FmaddRoundPd(a x86.M512d, b x86.M512d, c x86.M512d, rounding int) (dst x86.M512d) {
	panic("not implemented")
}


// M512MaskFmaddRoundPd: Multiply packed double-precision (64-bit)
// floating-point elements in 'a' and 'b', add the intermediate result to
// packed elements in 'c', and store the results in 'dst' using writemask 'k'
// (elements are copied from 'a' when the corresponding mask bit is not set).
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := (a[i+63:i] * b[i+63:i]) + c[i+63:i]
//			ELSE
//				dst[i+63:i] := a[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VFMADD132PD, VFMADD213PD, VFMADD231PD'. Intrinsic: '_mm512_mask_fmadd_round_pd'.
// Requires KNCNI.
func M512MaskFmaddRoundPd(a x86.M512d, k x86.Mmask8, b x86.M512d, c x86.M512d, rounding int) (dst x86.M512d) {
	panic("not implemented")
}


// M512Mask3FmaddRoundPd: Multiply packed double-precision (64-bit)
// floating-point elements in 'a' and 'b', add the intermediate result to
// packed elements in 'c', and store the results in 'dst' using writemask 'k'
// (elements are copied from 'c' when the corresponding mask bit is not set). 
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := (a[i+63:i] * b[i+63:i]) + c[i+63:i]
//			ELSE 
//				dst[i+63:i] := c[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VFMADD132PD, VFMADD213PD, VFMADD231PD'. Intrinsic: '_mm512_mask3_fmadd_round_pd'.
// Requires KNCNI.
func M512Mask3FmaddRoundPd(a x86.M512d, b x86.M512d, c x86.M512d, k x86.Mmask8, rounding int) (dst x86.M512d) {
	panic("not implemented")
}


// M512FmaddRoundPs: Multiply packed single-precision (32-bit) floating-point
// elements in 'a' and 'b', add the intermediate result to packed elements in
// 'c', and store the results in 'dst'. 
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := (a[i+31:i] * b[i+31:i]) + c[i+31:i]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VFMADD132PS, VFMADD213PS, VFMADD231PS'. Intrinsic: '_mm512_fmadd_round_ps'.
// Requires KNCNI.
func M512FmaddRoundPs(a x86.M512, b x86.M512, c x86.M512, rounding int) (dst x86.M512) {
	panic("not implemented")
}


// M512MaskFmaddRoundPs: Multiply packed single-precision (32-bit)
// floating-point elements in 'a' and 'b', add the intermediate result to
// packed elements in 'c', and store the results in 'dst' using writemask 'k'
// (elements are copied from 'a' when the corresponding mask bit is not set). 
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := (a[i+31:i] * b[i+31:i]) + c[i+31:i]
//			ELSE
//				dst[i+31:i] := a[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VFMADD132PS, VFMADD213PS, VFMADD231PS'. Intrinsic: '_mm512_mask_fmadd_round_ps'.
// Requires KNCNI.
func M512MaskFmaddRoundPs(a x86.M512, k x86.Mmask16, b x86.M512, c x86.M512, rounding int) (dst x86.M512) {
	panic("not implemented")
}


// M512Mask3FmaddRoundPs: Multiply packed single-precision (32-bit)
// floating-point elements in 'a' and 'b', add the intermediate result to
// packed elements in 'c', and store the results in 'dst' using writemask 'k'
// (elements are copied from 'c' when the corresponding mask bit is not set). 
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := (a[i+31:i] * b[i+31:i]) + c[i+31:i]
//			ELSE
//				dst[i+31:i] := c[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VFMADD132PS, VFMADD213PS, VFMADD231PS'. Intrinsic: '_mm512_mask3_fmadd_round_ps'.
// Requires KNCNI.
func M512Mask3FmaddRoundPs(a x86.M512, b x86.M512, c x86.M512, k x86.Mmask16, rounding int) (dst x86.M512) {
	panic("not implemented")
}


// M512Fmadd233Epi32: Multiply packed 32-bit integer elements in each 4-element
// set of 'a' and by element 1 of the corresponding 4-element set from 'b', add
// the intermediate result to element 0 of the corresponding 4-element set from
// 'b', and store the results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			base := (j & ~0x3) * 32
//			scale[31:0] := b[base+63:base+32]
//			bias[31:0]  := b[base+31:base]
//			dst[i+31:i] := (a[i+31:i] * scale[31:0]) + bias[31:0]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPMADD233D'. Intrinsic: '_mm512_fmadd233_epi32'.
// Requires KNCNI.
func M512Fmadd233Epi32(a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskFmadd233Epi32: Multiply packed 32-bit integer elements in each
// 4-element set of 'a' and by element 1 of the corresponding 4-element set
// from 'b', add the intermediate result to element 0 of the corresponding
// 4-element set from 'b', and store the results in 'dst' using writemask 'k'
// (elements are copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				base := (j & ~0x3) * 32
//				scale[31:0] := b[base+63:base+32]
//				bias[31:0]  := b[base+31:base]
//				dst[i+31:i] := (a[i+31:i] * scale[31:0]) + bias[31:0]
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPMADD233D'. Intrinsic: '_mm512_mask_fmadd233_epi32'.
// Requires KNCNI.
func M512MaskFmadd233Epi32(src x86.M512i, k x86.Mmask16, a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512Fmadd233Ps: Performs multiplication between single-precision (32-bit)
// floating-point elements in 'a' and 'b' and adds the result to the elements
// in 'b', storing the results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := (a[i+31:i] * b[i+31:i]) + b[i+31:i]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VFMADD233PS'. Intrinsic: '_mm512_fmadd233_ps'.
// Requires KNCNI.
func M512Fmadd233Ps(a x86.M512, b x86.M512) (dst x86.M512) {
	panic("not implemented")
}


// M512MaskFmadd233Ps: Performs multiplication between single-precision
// (32-bit) floating-point elements in 'a' and 'b' and adds the result to the
// elements in 'b', storing the results in 'dst' using writemask 'k' (elements
// are copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := (a[i+31:i] * b[i+31:i]) + b[i+31:i]
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VFMADD233PS'. Intrinsic: '_mm512_mask_fmadd233_ps'.
// Requires KNCNI.
func M512MaskFmadd233Ps(src x86.M512, k x86.Mmask16, a x86.M512, b x86.M512) (dst x86.M512) {
	panic("not implemented")
}


// M512Fmadd233RoundPs: Multiply packed single-precision (32-bit)
// floating-point elements in each 4-element set of 'a' and by element 1 of the
// corresponding 4-element set from 'b', add the intermediate result to element
// 0 of the corresponding 4-element set from 'b', and store the results in
// 'dst'.
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 15
//			i := j*32
//			base := (j & ~0x3) * 32
//			scale[31:0] := b[base+63:base+32]
//			bias[31:0]  := b[base+31:base]
//			dst[i+31:i] := (a[i+31:i] * scale[31:0]) + bias[31:0]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VFMADD233PS'. Intrinsic: '_mm512_fmadd233_round_ps'.
// Requires KNCNI.
func M512Fmadd233RoundPs(a x86.M512, b x86.M512, rounding int) (dst x86.M512) {
	panic("not implemented")
}


// M512MaskFmadd233RoundPs: Multiply packed single-precision (32-bit)
// floating-point elements in each 4-element set of 'a' and by element 1 of the
// corresponding 4-element set from 'b', add the intermediate result to element
// 0 of the corresponding 4-element set from 'b', and store the results in
// 'dst' using writemask 'k' (elements are copied from 'src' when the
// corresponding mask bit is not set).
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				base := (j & ~0x3) * 32
//				scale[31:0] := b[base+63:base+32]
//				bias[31:0]  := b[base+31:base]
//				dst[i+31:i] := (a[i+31:i] * scale[31:0]) + bias[31:0]
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VFMADD233PS'. Intrinsic: '_mm512_mask_fmadd233_round_ps'.
// Requires KNCNI.
func M512MaskFmadd233RoundPs(src x86.M512, k x86.Mmask16, a x86.M512, b x86.M512, rounding int) (dst x86.M512) {
	panic("not implemented")
}


// M512FmsubPd: Multiply packed double-precision (64-bit) floating-point
// elements in 'a' and 'b', subtract packed elements in 'c' from the
// intermediate result, and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*64
//			dst[i+63:i] := (a[i+63:i] * b[i+63:i]) - c[i+63:i]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VFMSUB132PD, VFMSUB213PD, VFMSUB231PD'. Intrinsic: '_mm512_fmsub_pd'.
// Requires KNCNI.
func M512FmsubPd(a x86.M512d, b x86.M512d, c x86.M512d) (dst x86.M512d) {
	panic("not implemented")
}


// M512MaskFmsubPd: Multiply packed double-precision (64-bit) floating-point
// elements in 'a' and 'b', subtract packed elements in 'c' from the
// intermediate result, and store the results in 'dst' using writemask 'k'
// (elements are copied from 'a' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := (a[i+63:i] * b[i+63:i]) - c[i+63:i]
//			ELSE
//				dst[i+63:i] := a[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VFMSUB132PD, VFMSUB213PD, VFMSUB231PD'. Intrinsic: '_mm512_mask_fmsub_pd'.
// Requires KNCNI.
func M512MaskFmsubPd(a x86.M512d, k x86.Mmask8, b x86.M512d, c x86.M512d) (dst x86.M512d) {
	panic("not implemented")
}


// M512Mask3FmsubPd: Multiply packed double-precision (64-bit) floating-point
// elements in 'a' and 'b', subtract packed elements in 'c' from the
// intermediate result, and store the results in 'dst' using writemask 'k'
// (elements are copied from 'c' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := (a[i+63:i] * b[i+63:i]) - c[i+63:i]
//			ELSE
//				dst[i+63:i] := c[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VFMSUB132PD, VFMSUB213PD, VFMSUB231PD'. Intrinsic: '_mm512_mask3_fmsub_pd'.
// Requires KNCNI.
func M512Mask3FmsubPd(a x86.M512d, b x86.M512d, c x86.M512d, k x86.Mmask8) (dst x86.M512d) {
	panic("not implemented")
}


// M512FmsubPs: Multiply packed single-precision (32-bit) floating-point
// elements in 'a' and 'b', subtract packed elements in 'c' from the
// intermediate result, and store the results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := (a[i+31:i] * b[i+31:i]) - c[i+31:i]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VFMSUB132PS, VFMSUB213PS, VFMSUB231PS'. Intrinsic: '_mm512_fmsub_ps'.
// Requires KNCNI.
func M512FmsubPs(a x86.M512, b x86.M512, c x86.M512) (dst x86.M512) {
	panic("not implemented")
}


// M512MaskFmsubPs: Multiply packed single-precision (32-bit) floating-point
// elements in 'a' and 'b', subtract packed elements in 'c' from the
// intermediate result, and store the results in 'dst' using writemask 'k'
// (elements are copied from 'a' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := (a[i+31:i] * b[i+31:i]) - c[i+31:i]
//			ELSE
//				dst[i+31:i] := a[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VFMSUB132PS, VFMSUB213PS, VFMSUB231PS'. Intrinsic: '_mm512_mask_fmsub_ps'.
// Requires KNCNI.
func M512MaskFmsubPs(a x86.M512, k x86.Mmask16, b x86.M512, c x86.M512) (dst x86.M512) {
	panic("not implemented")
}


// M512Mask3FmsubPs: Multiply packed single-precision (32-bit) floating-point
// elements in 'a' and 'b', subtract packed elements in 'c' from the
// intermediate result, and store the results in 'dst' using writemask 'k'
// (elements are copied from 'c' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := (a[i+31:i] * b[i+31:i]) - c[i+31:i]
//			ELSE
//				dst[i+31:i] := c[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VFMSUB132PS, VFMSUB213PS, VFMSUB231PS'. Intrinsic: '_mm512_mask3_fmsub_ps'.
// Requires KNCNI.
func M512Mask3FmsubPs(a x86.M512, b x86.M512, c x86.M512, k x86.Mmask16) (dst x86.M512) {
	panic("not implemented")
}


// M512FmsubRoundPd: Multiply packed double-precision (64-bit) floating-point
// elements in 'a' and 'b', subtract packed elements in 'c' from the
// intermediate result, and store the results in 'dst'. 
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 7
//			i := j*64
//			dst[i+63:i] := (a[i+63:i] * b[i+63:i]) - c[i+63:i]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VFMSUB132PD, VFMSUB213PD, VFMSUB231PD'. Intrinsic: '_mm512_fmsub_round_pd'.
// Requires KNCNI.
func M512FmsubRoundPd(a x86.M512d, b x86.M512d, c x86.M512d, rounding int) (dst x86.M512d) {
	panic("not implemented")
}


// M512MaskFmsubRoundPd: Multiply packed double-precision (64-bit)
// floating-point elements in 'a' and 'b', subtract packed elements in 'c' from
// the intermediate result, and store the results in 'dst' using writemask 'k'
// (elements are copied from 'a' when the corresponding mask bit is not set).
// Rounding is done according to the 'rounding' parameter, which can be one of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := (a[i+63:i] * b[i+63:i]) - c[i+63:i]
//			ELSE
//				dst[i+63:i] := a[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VFMSUB132PD, VFMSUB213PD, VFMSUB231PD'. Intrinsic: '_mm512_mask_fmsub_round_pd'.
// Requires KNCNI.
func M512MaskFmsubRoundPd(a x86.M512d, k x86.Mmask8, b x86.M512d, c x86.M512d, rounding int) (dst x86.M512d) {
	panic("not implemented")
}


// M512Mask3FmsubRoundPd: Multiply packed double-precision (64-bit)
// floating-point elements in 'a' and 'b', subtract packed elements in 'c' from
// the intermediate result, and store the results in 'dst' using writemask 'k'
// (elements are copied from 'c' when the corresponding mask bit is not set).
// Rounding is done according to the 'rounding' parameter, which can be one of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := (a[i+63:i] * b[i+63:i]) - c[i+63:i]
//			ELSE
//				dst[i+63:i] := c[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VFMSUB132PD, VFMSUB213PD, VFMSUB231PD'. Intrinsic: '_mm512_mask3_fmsub_round_pd'.
// Requires KNCNI.
func M512Mask3FmsubRoundPd(a x86.M512d, b x86.M512d, c x86.M512d, k x86.Mmask8, rounding int) (dst x86.M512d) {
	panic("not implemented")
}


// M512FmsubRoundPs: Multiply packed single-precision (32-bit) floating-point
// elements in 'a' and 'b', subtract packed elements in 'c' from the
// intermediate result, and store the results in 'dst'. 
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := (a[i+31:i] * b[i+31:i]) - c[i+31:i]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VFMSUB132PS, VFMSUB213PS, VFMSUB231PS'. Intrinsic: '_mm512_fmsub_round_ps'.
// Requires KNCNI.
func M512FmsubRoundPs(a x86.M512, b x86.M512, c x86.M512, rounding int) (dst x86.M512) {
	panic("not implemented")
}


// M512MaskFmsubRoundPs: Multiply packed single-precision (32-bit)
// floating-point elements in 'a' and 'b', subtract packed elements in 'c' from
// the intermediate result, and store the results in 'dst' using writemask 'k'
// (elements are copied from 'a' when the corresponding mask bit is not set).
// Rounding is done according to the 'rounding' parameter, which can be one of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := (a[i+31:i] * b[i+31:i]) - c[i+31:i]
//			ELSE
//				dst[i+31:i] := a[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VFMSUB132PS, VFMSUB213PS, VFMSUB231PS'. Intrinsic: '_mm512_mask_fmsub_round_ps'.
// Requires KNCNI.
func M512MaskFmsubRoundPs(a x86.M512, k x86.Mmask16, b x86.M512, c x86.M512, rounding int) (dst x86.M512) {
	panic("not implemented")
}


// M512Mask3FmsubRoundPs: Multiply packed single-precision (32-bit)
// floating-point elements in 'a' and 'b', subtract packed elements in 'c' from
// the intermediate result, and store the results in 'dst' using writemask 'k'
// (elements are copied from 'c' when the corresponding mask bit is not set).
// Rounding is done according to the 'rounding' parameter, which can be one of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := (a[i+31:i] * b[i+31:i]) - c[i+31:i]
//			ELSE
//				dst[i+31:i] := c[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VFMSUB132PS, VFMSUB213PS, VFMSUB231PS'. Intrinsic: '_mm512_mask3_fmsub_round_ps'.
// Requires KNCNI.
func M512Mask3FmsubRoundPs(a x86.M512, b x86.M512, c x86.M512, k x86.Mmask16, rounding int) (dst x86.M512) {
	panic("not implemented")
}


// M512FnmaddPd: Multiply packed double-precision (64-bit) floating-point
// elements in 'a' and 'b', add the negated intermediate result to packed
// elements in 'c', and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*64
//			dst[i+63:i] := -(a[i+63:i] * b[i+63:i]) + c[i+63:i]
//		ENDFOR	
//		dst[MAX:512] := 0
//
// Instruction: 'VFNMADD132PD, VFNMADD213PD, VFNMADD231PD'. Intrinsic: '_mm512_fnmadd_pd'.
// Requires KNCNI.
func M512FnmaddPd(a x86.M512d, b x86.M512d, c x86.M512d) (dst x86.M512d) {
	panic("not implemented")
}


// M512MaskFnmaddPd: Multiply packed double-precision (64-bit) floating-point
// elements in 'a' and 'b', add the negated intermediate result to packed
// elements in 'c', and store the results in 'dst' using writemask 'k'
// (elements are copied from 'a' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := -(a[i+63:i] * b[i+63:i]) + c[i+63:i]
//			ELSE
//				dst[i+63:i] := a[i+63:i]
//			FI
//		ENDFOR	
//		dst[MAX:512] := 0
//
// Instruction: 'VFNMADD132PD, VFNMADD213PD, VFNMADD231PD'. Intrinsic: '_mm512_mask_fnmadd_pd'.
// Requires KNCNI.
func M512MaskFnmaddPd(a x86.M512d, k x86.Mmask8, b x86.M512d, c x86.M512d) (dst x86.M512d) {
	panic("not implemented")
}


// M512Mask3FnmaddPd: Multiply packed double-precision (64-bit) floating-point
// elements in 'a' and 'b', add the negated intermediate result to packed
// elements in 'c', and store the results in 'dst' using writemask 'k'
// (elements are copied from 'c' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := -(a[i+63:i] * b[i+63:i]) + c[i+63:i]
//			ELSE
//				dst[i+63:i] := c[i+63:i]
//			FI
//		ENDFOR	
//		dst[MAX:512] := 0
//
// Instruction: 'VFNMADD132PD, VFNMADD213PD, VFNMADD231PD'. Intrinsic: '_mm512_mask3_fnmadd_pd'.
// Requires KNCNI.
func M512Mask3FnmaddPd(a x86.M512d, b x86.M512d, c x86.M512d, k x86.Mmask8) (dst x86.M512d) {
	panic("not implemented")
}


// M512FnmaddPs: Multiply packed single-precision (32-bit) floating-point
// elements in 'a' and 'b', add the negated intermediate result to packed
// elements in 'c', and store the results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			a[i+31:i] := -(a[i+31:i] * b[i+31:i]) + c[i+31:i]
//		ENDFOR	
//		dst[MAX:512] := 0
//
// Instruction: 'VFNMADD132PS, VFNMADD213PS, VFNMADD231PS'. Intrinsic: '_mm512_fnmadd_ps'.
// Requires KNCNI.
func M512FnmaddPs(a x86.M512, b x86.M512, c x86.M512) (dst x86.M512) {
	panic("not implemented")
}


// M512MaskFnmaddPs: Multiply packed single-precision (32-bit) floating-point
// elements in 'a' and 'b', add the negated intermediate result to packed
// elements in 'c', and store the results in 'dst' using writemask 'k'
// (elements are copied from 'a' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := -(a[i+31:i] * b[i+31:i]) + c[i+31:i]
//			ELSE
//				dst[i+31:i] := a[i+31:i]
//			FI
//		ENDFOR	
//		dst[MAX:512] := 0
//
// Instruction: 'VFNMADD132PS, VFNMADD213PS, VFNMADD231PS'. Intrinsic: '_mm512_mask_fnmadd_ps'.
// Requires KNCNI.
func M512MaskFnmaddPs(a x86.M512, k x86.Mmask16, b x86.M512, c x86.M512) (dst x86.M512) {
	panic("not implemented")
}


// M512Mask3FnmaddPs: Multiply packed single-precision (32-bit) floating-point
// elements in 'a' and 'b', add the negated intermediate result to packed
// elements in 'c', and store the results in 'dst' using writemask 'k'
// (elements are copied from 'c' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := -(a[i+31:i] * b[i+31:i]) + c[i+31:i]
//			ELSE
//				dst[i+31:i] := c[i+31:i]
//			FI
//		ENDFOR	
//		dst[MAX:512] := 0
//
// Instruction: 'VFNMADD132PS, VFNMADD213PS, VFNMADD231PS'. Intrinsic: '_mm512_mask3_fnmadd_ps'.
// Requires KNCNI.
func M512Mask3FnmaddPs(a x86.M512, b x86.M512, c x86.M512, k x86.Mmask16) (dst x86.M512) {
	panic("not implemented")
}


// M512FnmaddRoundPd: Multiply packed double-precision (64-bit) floating-point
// elements in 'a' and 'b', add the negated intermediate result to packed
// elements in 'c', and store the results in 'dst'.
// 	 Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 7
//			i := j*64
//			dst[i+63:i] := -(a[i+63:i] * b[i+63:i]) + c[i+63:i]
//		ENDFOR	
//		dst[MAX:512] := 0
//
// Instruction: 'VFNMADD132PD, VFNMADD213PD, VFNMADD231PD'. Intrinsic: '_mm512_fnmadd_round_pd'.
// Requires KNCNI.
func M512FnmaddRoundPd(a x86.M512d, b x86.M512d, c x86.M512d, rounding int) (dst x86.M512d) {
	panic("not implemented")
}


// M512MaskFnmaddRoundPd: Multiply packed double-precision (64-bit)
// floating-point elements in 'a' and 'b', add the negated intermediate result
// to packed elements in 'c', and store the results in 'dst' using writemask
// 'k' (elements are copied from 'a' when the corresponding mask bit is not
// set). Rounding is done according to the 'rounding' parameter, which can be
// one of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := -(a[i+63:i] * b[i+63:i]) + c[i+63:i]
//			ELSE
//				dst[i+63:i] := a[i+63:i]
//			FI
//		ENDFOR	
//		dst[MAX:512] := 0
//
// Instruction: 'VFNMADD132PD, VFNMADD213PD, VFNMADD231PD'. Intrinsic: '_mm512_mask_fnmadd_round_pd'.
// Requires KNCNI.
func M512MaskFnmaddRoundPd(a x86.M512d, k x86.Mmask8, b x86.M512d, c x86.M512d, rounding int) (dst x86.M512d) {
	panic("not implemented")
}


// M512Mask3FnmaddRoundPd: Multiply packed double-precision (64-bit)
// floating-point elements in 'a' and 'b', add the negated intermediate result
// to packed elements in 'c', and store the results in 'dst' using writemask
// 'k' (elements are copied from 'c' when the corresponding mask bit is not
// set).  Rounding is done according to the 'rounding' parameter, which can be
// one of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := -(a[i+63:i] * b[i+63:i]) + c[i+63:i]
//			ELSE
//				dst[i+63:i] := c[i+63:i]
//			FI
//		ENDFOR	
//		dst[MAX:512] := 0
//
// Instruction: 'VFNMADD132PD, VFNMADD213PD, VFNMADD231PD'. Intrinsic: '_mm512_mask3_fnmadd_round_pd'.
// Requires KNCNI.
func M512Mask3FnmaddRoundPd(a x86.M512d, b x86.M512d, c x86.M512d, k x86.Mmask8, rounding int) (dst x86.M512d) {
	panic("not implemented")
}


// M512FnmaddRoundPs: Multiply packed single-precision (32-bit) floating-point
// elements in 'a' and 'b', add the negated intermediate result to packed
// elements in 'c', and store the results in 'dst'.  
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := -(a[i+31:i] * b[i+31:i]) + c[i+31:i]
//		ENDFOR	
//		dst[MAX:512] := 0
//
// Instruction: 'VFNMADD132PS, VFNMADD213PS, VFNMADD231PS'. Intrinsic: '_mm512_fnmadd_round_ps'.
// Requires KNCNI.
func M512FnmaddRoundPs(a x86.M512, b x86.M512, c x86.M512, rounding int) (dst x86.M512) {
	panic("not implemented")
}


// M512MaskFnmaddRoundPs: Multiply packed single-precision (32-bit)
// floating-point elements in 'a' and 'b', add the negated intermediate result
// to packed elements in 'c', and store the results in 'dst' using writemask
// 'k' (elements are copied from 'a' when the corresponding mask bit is not
// set). Rounding is done according to the 'rounding' parameter, which can be
// one of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := -(a[i+31:i] * b[i+31:i]) + c[i+31:i]
//			ELSE
//				dst[i+31:i] := a[i+31:i]
//			FI
//		ENDFOR	
//		dst[MAX:512] := 0
//
// Instruction: 'VFNMADD132PS, VFNMADD213PS, VFNMADD231PS'. Intrinsic: '_mm512_mask_fnmadd_round_ps'.
// Requires KNCNI.
func M512MaskFnmaddRoundPs(a x86.M512, k x86.Mmask16, b x86.M512, c x86.M512, rounding int) (dst x86.M512) {
	panic("not implemented")
}


// M512Mask3FnmaddRoundPs: Multiply packed single-precision (32-bit)
// floating-point elements in 'a' and 'b', add the negated intermediate result
// to packed elements in 'c', and store the results in 'dst' using writemask
// 'k' (elements are copied from 'c' when the corresponding mask bit is not
// set).  Rounding is done according to the 'rounding' parameter, which can be
// one of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := -(a[i+31:i] * b[i+31:i]) + c[i+31:i]
//			ELSE
//				dst[i+31:i] := c[i+31:i]
//			FI
//		ENDFOR	
//		dst[MAX:512] := 0
//
// Instruction: 'VFNMADD132PS, VFNMADD213PS, VFNMADD231PS'. Intrinsic: '_mm512_mask3_fnmadd_round_ps'.
// Requires KNCNI.
func M512Mask3FnmaddRoundPs(a x86.M512, b x86.M512, c x86.M512, k x86.Mmask16, rounding int) (dst x86.M512) {
	panic("not implemented")
}


// M512FnmsubPd: Multiply packed double-precision (64-bit) floating-point
// elements in 'a' and 'b', subtract packed elements in 'c' from the negated
// intermediate result, and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*64
//			dst[i+63:i] := -(a[i+63:i] * b[i+63:i]) - c[i+63:i]
//		ENDFOR	
//		dst[MAX:512] := 0
//
// Instruction: 'VFNMSUB132PD, VFNMSUB213PD, VFNMSUB231PD'. Intrinsic: '_mm512_fnmsub_pd'.
// Requires KNCNI.
func M512FnmsubPd(a x86.M512d, b x86.M512d, c x86.M512d) (dst x86.M512d) {
	panic("not implemented")
}


// M512MaskFnmsubPd: Multiply packed double-precision (64-bit) floating-point
// elements in 'a' and 'b', subtract packed elements in 'c' from the negated
// intermediate result, and store the results in 'dst' using writemask 'k'
// (elements are copied from 'a' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := -(a[i+63:i] * b[i+63:i]) - c[i+63:i]
//			ELSE
//				dst[i+63:i] := a[i+63:i]
//			FI
//		ENDFOR	
//		dst[MAX:512] := 0
//
// Instruction: 'VFNMSUB132PD, VFNMSUB213PD, VFNMSUB231PD'. Intrinsic: '_mm512_mask_fnmsub_pd'.
// Requires KNCNI.
func M512MaskFnmsubPd(a x86.M512d, k x86.Mmask8, b x86.M512d, c x86.M512d) (dst x86.M512d) {
	panic("not implemented")
}


// M512Mask3FnmsubPd: Multiply packed double-precision (64-bit) floating-point
// elements in 'a' and 'b', subtract packed elements in 'c' from the negated
// intermediate result, and store the results in 'dst' using writemask 'k'
// (elements are copied from 'c' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := -(a[i+63:i] * b[i+63:i]) - c[i+63:i]
//			ELSE
//				dst[i+63:i] := c[i+63:i]
//			FI
//		ENDFOR	
//		dst[MAX:512] := 0
//
// Instruction: 'VFNMSUB132PD, VFNMSUB213PD, VFNMSUB231PD'. Intrinsic: '_mm512_mask3_fnmsub_pd'.
// Requires KNCNI.
func M512Mask3FnmsubPd(a x86.M512d, b x86.M512d, c x86.M512d, k x86.Mmask8) (dst x86.M512d) {
	panic("not implemented")
}


// M512FnmsubPs: Multiply packed single-precision (32-bit) floating-point
// elements in 'a' and 'b', subtract packed elements in 'c' from the negated
// intermediate result, and store the results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := -(a[i+31:i] * b[i+31:i]) - c[i+31:i]
//		ENDFOR	
//		dst[MAX:512] := 0
//
// Instruction: 'VFNMSUB132PS, VFNMSUB213PS, VFNMSUB231PS'. Intrinsic: '_mm512_fnmsub_ps'.
// Requires KNCNI.
func M512FnmsubPs(a x86.M512, b x86.M512, c x86.M512) (dst x86.M512) {
	panic("not implemented")
}


// M512MaskFnmsubPs: Multiply packed single-precision (32-bit) floating-point
// elements in 'a' and 'b', subtract packed elements in 'c' from the negated
// intermediate result, and store the results in 'dst' using writemask 'k'
// (elements are copied from 'a' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := -(a[i+31:i] * b[i+31:i]) - c[i+31:i]
//			ELSE
//				dst[i+31:i] := a[i+31:i]
//			FI
//		ENDFOR	
//		dst[MAX:512] := 0
//
// Instruction: 'VFNMSUB132PS, VFNMSUB213PS, VFNMSUB231PS'. Intrinsic: '_mm512_mask_fnmsub_ps'.
// Requires KNCNI.
func M512MaskFnmsubPs(a x86.M512, k x86.Mmask16, b x86.M512, c x86.M512) (dst x86.M512) {
	panic("not implemented")
}


// M512Mask3FnmsubPs: Multiply packed single-precision (32-bit) floating-point
// elements in 'a' and 'b', subtract packed elements in 'c' from the negated
// intermediate result, and store the results in 'dst' using writemask 'k'
// (elements are copied from 'c' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := -(a[i+31:i] * b[i+31:i]) - c[i+31:i]
//			ELSE
//				dst[i+31:i] := c[i+31:i]
//			FI
//		ENDFOR	
//		dst[MAX:512] := 0
//
// Instruction: 'VFNMSUB132PS, VFNMSUB213PS, VFNMSUB231PS'. Intrinsic: '_mm512_mask3_fnmsub_ps'.
// Requires KNCNI.
func M512Mask3FnmsubPs(a x86.M512, b x86.M512, c x86.M512, k x86.Mmask16) (dst x86.M512) {
	panic("not implemented")
}


// M512FnmsubRoundPd: Multiply packed double-precision (64-bit) floating-point
// elements in 'a' and 'b', subtract packed elements in 'c' from the negated
// intermediate result, and store the results in 'dst'.  
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 7
//			i := j*64
//			dst[i+63:i] := -(a[i+63:i] * b[i+63:i]) - c[i+63:i]
//		ENDFOR	
//		dst[MAX:512] := 0
//
// Instruction: 'VFNMSUB132PD, VFNMSUB213PD, VFNMSUB231PD'. Intrinsic: '_mm512_fnmsub_round_pd'.
// Requires KNCNI.
func M512FnmsubRoundPd(a x86.M512d, b x86.M512d, c x86.M512d, rounding int) (dst x86.M512d) {
	panic("not implemented")
}


// M512MaskFnmsubRoundPd: Multiply packed double-precision (64-bit)
// floating-point elements in 'a' and 'b', subtract packed elements in 'c' from
// the negated intermediate result, and store the results in 'dst' using
// writemask 'k' (elements are copied from 'a' when the corresponding mask bit
// is not set).
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := -(a[i+63:i] * b[i+63:i]) - c[i+63:i]
//			ELSE
//				dst[i+63:i] := a[i+63:i]
//			FI
//		ENDFOR	
//		dst[MAX:512] := 0
//
// Instruction: 'VFNMSUB132PD, VFNMSUB213PD, VFNMSUB231PD'. Intrinsic: '_mm512_mask_fnmsub_round_pd'.
// Requires KNCNI.
func M512MaskFnmsubRoundPd(a x86.M512d, k x86.Mmask8, b x86.M512d, c x86.M512d, rounding int) (dst x86.M512d) {
	panic("not implemented")
}


// M512Mask3FnmsubRoundPd: Multiply packed double-precision (64-bit)
// floating-point elements in 'a' and 'b', subtract packed elements in 'c' from
// the negated intermediate result, and store the results in 'dst' using
// writemask 'k' (elements are copied from 'c' when the corresponding mask bit
// is not set). Rounding is done according to the 'rounding' parameter, which
// can be one of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := -(a[i+63:i] * b[i+63:i]) - c[i+63:i]
//			ELSE
//				dst[i+63:i] := c[i+63:i]
//			FI
//		ENDFOR	
//		dst[MAX:512] := 0
//
// Instruction: 'VFNMSUB132PD, VFNMSUB213PD, VFNMSUB231PD'. Intrinsic: '_mm512_mask3_fnmsub_round_pd'.
// Requires KNCNI.
func M512Mask3FnmsubRoundPd(a x86.M512d, b x86.M512d, c x86.M512d, k x86.Mmask8, rounding int) (dst x86.M512d) {
	panic("not implemented")
}


// M512FnmsubRoundPs: Multiply packed single-precision (32-bit) floating-point
// elements in 'a' and 'b', subtract packed elements in 'c' from the negated
// intermediate result, and store the results in 'dst'. 
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := -(a[i+31:i] * b[i+31:i]) - c[i+31:i]
//		ENDFOR	
//		dst[MAX:512] := 0
//
// Instruction: 'VFNMSUB132PS, VFNMSUB213PS, VFNMSUB231PS'. Intrinsic: '_mm512_fnmsub_round_ps'.
// Requires KNCNI.
func M512FnmsubRoundPs(a x86.M512, b x86.M512, c x86.M512, rounding int) (dst x86.M512) {
	panic("not implemented")
}


// M512MaskFnmsubRoundPs: Multiply packed single-precision (32-bit)
// floating-point elements in 'a' and 'b', subtract packed elements in 'c' from
// the negated intermediate result, and store the results in 'dst' using
// writemask 'k' (elements are copied from 'a' when the corresponding mask bit
// is not set). 
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := -(a[i+31:i] * b[i+31:i]) - c[i+31:i]
//			ELSE
//				dst[i+31:i] := a[i+31:i]
//			FI
//		ENDFOR	
//		dst[MAX:512] := 0
//
// Instruction: 'VFNMSUB132PS, VFNMSUB213PS, VFNMSUB231PS'. Intrinsic: '_mm512_mask_fnmsub_round_ps'.
// Requires KNCNI.
func M512MaskFnmsubRoundPs(a x86.M512, k x86.Mmask16, b x86.M512, c x86.M512, rounding int) (dst x86.M512) {
	panic("not implemented")
}


// M512Mask3FnmsubRoundPs: Multiply packed single-precision (32-bit)
// floating-point elements in 'a' and 'b', subtract packed elements in 'c' from
// the negated intermediate result, and store the results in 'dst' using
// writemask 'k' (elements are copied from 'c' when the corresponding mask bit
// is not set).  Rounding is done according to the 'rounding' parameter, which
// can be one of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := -(a[i+31:i] * b[i+31:i]) - c[i+31:i]
//			ELSE
//				dst[i+31:i] := c[i+31:i]
//			FI
//		ENDFOR	
//		dst[MAX:512] := 0
//
// Instruction: 'VFNMSUB132PS, VFNMSUB213PS, VFNMSUB231PS'. Intrinsic: '_mm512_mask3_fnmsub_round_ps'.
// Requires KNCNI.
func M512Mask3FnmsubRoundPs(a x86.M512, b x86.M512, c x86.M512, k x86.Mmask16, rounding int) (dst x86.M512) {
	panic("not implemented")
}


// M512GetexpPd: Convert the exponent of each packed double-precision (64-bit)
// floating-point element in 'a' to a double-precision (64-bit) floating-point
// number representing the integer exponent, and store the results in 'dst'.
// This intrinsic essentially calculates 'floor(log2(x))' for each element. 
//
//		FOR j := 0 to 7
//			i := j*64
//			dst[i+63:i] := ConvertExpFP64(a[i+63:i])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VGETEXPPD'. Intrinsic: '_mm512_getexp_pd'.
// Requires KNCNI.
func M512GetexpPd(a x86.M512d) (dst x86.M512d) {
	panic("not implemented")
}


// M512MaskGetexpPd: Convert the exponent of each packed double-precision
// (64-bit) floating-point element in 'a' to a double-precision (64-bit)
// floating-point number representing the integer exponent, and store the
// results in 'dst' using writemask 'k' (elements are copied from 'src' when
// the corresponding mask bit is not set). This intrinsic essentially
// calculates 'floor(log2(x))' for each element. 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := ConvertExpFP64(a[i+63:i])
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VGETEXPPD'. Intrinsic: '_mm512_mask_getexp_pd'.
// Requires KNCNI.
func M512MaskGetexpPd(src x86.M512d, k x86.Mmask8, a x86.M512d) (dst x86.M512d) {
	panic("not implemented")
}


// M512GetexpPs: Convert the exponent of each packed single-precision (32-bit)
// floating-point element in 'a' to a single-precision (32-bit) floating-point
// number representing the integer exponent, and store the results in 'dst'.
// This intrinsic essentially calculates 'floor(log2(x))' for each element. 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := ConvertExpFP32(a[i+31:i])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VGETEXPPS'. Intrinsic: '_mm512_getexp_ps'.
// Requires KNCNI.
func M512GetexpPs(a x86.M512) (dst x86.M512) {
	panic("not implemented")
}


// M512MaskGetexpPs: Convert the exponent of each packed single-precision
// (32-bit) floating-point element in 'a' to a single-precision (32-bit)
// floating-point number representing the integer exponent, and store the
// results in 'dst' using writemask 'k' (elements are copied from 'src' when
// the corresponding mask bit is not set). This intrinsic essentially
// calculates 'floor(log2(x))' for each element. 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := ConvertExpFP32(a[i+31:i])
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VGETEXPPS'. Intrinsic: '_mm512_mask_getexp_ps'.
// Requires KNCNI.
func M512MaskGetexpPs(src x86.M512, k x86.Mmask16, a x86.M512) (dst x86.M512) {
	panic("not implemented")
}


// M512GetexpRoundPd: Convert the exponent of each packed double-precision
// (64-bit) floating-point element in 'a' to a double-precision (64-bit)
// floating-point number representing the integer exponent, and store the
// results in 'dst'. This intrinsic essentially calculates 'floor(log2(x))' for
// each element.
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 7
//			i := j*64
//			dst[i+63:i] := ConvertExpFP64(a[i+63:i])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VGETEXPPD'. Intrinsic: '_mm512_getexp_round_pd'.
// Requires KNCNI.
func M512GetexpRoundPd(a x86.M512d, rounding int) (dst x86.M512d) {
	panic("not implemented")
}


// M512MaskGetexpRoundPd: Convert the exponent of each packed double-precision
// (64-bit) floating-point element in 'a' to a double-precision (64-bit)
// floating-point number representing the integer exponent, and store the
// results in 'dst' using writemask 'k' (elements are copied from 'src' when
// the corresponding mask bit is not set). This intrinsic essentially
// calculates 'floor(log2(x))' for each element.
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := ConvertExpFP64(a[i+63:i])
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VGETEXPPD'. Intrinsic: '_mm512_mask_getexp_round_pd'.
// Requires KNCNI.
func M512MaskGetexpRoundPd(src x86.M512d, k x86.Mmask8, a x86.M512d, rounding int) (dst x86.M512d) {
	panic("not implemented")
}


// M512GetexpRoundPs: Convert the exponent of each packed single-precision
// (32-bit) floating-point element in 'a' to a single-precision (32-bit)
// floating-point number representing the integer exponent, and store the
// results in 'dst'. This intrinsic essentially calculates 'floor(log2(x))' for
// each element.
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := ConvertExpFP32(a[i+31:i])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VGETEXPPS'. Intrinsic: '_mm512_getexp_round_ps'.
// Requires KNCNI.
func M512GetexpRoundPs(a x86.M512, rounding int) (dst x86.M512) {
	panic("not implemented")
}


// M512MaskGetexpRoundPs: Convert the exponent of each packed single-precision
// (32-bit) floating-point element in 'a' to a single-precision (32-bit)
// floating-point number representing the integer exponent, and store the
// results in 'dst' using writemask 'k' (elements are copied from 'src' when
// the corresponding mask bit is not set). This intrinsic essentially
// calculates 'floor(log2(x))' for each element.
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := ConvertExpFP32(a[i+31:i])
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VGETEXPPS'. Intrinsic: '_mm512_mask_getexp_round_ps'.
// Requires KNCNI.
func M512MaskGetexpRoundPs(src x86.M512, k x86.Mmask16, a x86.M512, rounding int) (dst x86.M512) {
	panic("not implemented")
}


// M512GetmantPd: Normalize the mantissas of packed double-precision (64-bit)
// floating-point elements in 'a', and store the results in 'dst'. This
// intrinsic essentially calculates '±(2^k)*|x.significand|', where 'k'
// depends on the interval range defined by 'interv' and the sign depends on
// 'sc' and the source sign.
// 	The mantissa is normalized to the interval specified by 'interv', which can
// take the following values:
//     _MM_MANT_NORM_1_2     // interval [1, 2)
//     _MM_MANT_NORM_p5_2    // interval [0.5, 2)
//     _MM_MANT_NORM_p5_1    // interval [0.5, 1)
//     _MM_MANT_NORM_p75_1p5 // interval [0.75, 1.5)The sign is determined by 'sc' which can take the following values:
//     _MM_MANT_SIGN_src     // sign = sign(src)
//     _MM_MANT_SIGN_zero    // sign = 0
//     _MM_MANT_SIGN_nan     // dst = NaN if sign(src) = 1 
//
//		FOR j := 0 to 7
//			i := j*64
//			dst[i+63:i] := GetNormalizedMantissa(a[i+63:i], sc, interv)
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VGETMANTPD'. Intrinsic: '_mm512_getmant_pd'.
// Requires KNCNI.
func M512GetmantPd(a x86.M512d, interv MMMANTISSANORMENUM, sc MMMANTISSASIGNENUM) (dst x86.M512d) {
	panic("not implemented")
}


// M512MaskGetmantPd: Normalize the mantissas of packed double-precision
// (64-bit) floating-point elements in 'a', and store the results in 'dst'
// using writemask 'k' (elements are copied from 'src' when the corresponding
// mask bit is not set). This intrinsic essentially calculates
// '±(2^k)*|x.significand|', where 'k' depends on the interval range defined
// by 'interv' and the sign depends on 'sc' and the source sign.
// 	The mantissa is normalized to the interval specified by 'interv', which can
// take the following values:
//     _MM_MANT_NORM_1_2     // interval [1, 2)
//     _MM_MANT_NORM_p5_2    // interval [0.5, 2)
//     _MM_MANT_NORM_p5_1    // interval [0.5, 1)
//     _MM_MANT_NORM_p75_1p5 // interval [0.75, 1.5)The sign is determined by 'sc' which can take the following values:
//     _MM_MANT_SIGN_src     // sign = sign(src)
//     _MM_MANT_SIGN_zero    // sign = 0
//     _MM_MANT_SIGN_nan     // dst = NaN if sign(src) = 1 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := GetNormalizedMantissa(a[i+63:i], sc, interv)
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VGETMANTPD'. Intrinsic: '_mm512_mask_getmant_pd'.
// Requires KNCNI.
func M512MaskGetmantPd(src x86.M512d, k x86.Mmask8, a x86.M512d, interv MMMANTISSANORMENUM, sc MMMANTISSASIGNENUM) (dst x86.M512d) {
	panic("not implemented")
}


// M512GetmantPs: Normalize the mantissas of packed single-precision (32-bit)
// floating-point elements in 'a', and store the results in 'dst'. This
// intrinsic essentially calculates '±(2^k)*|x.significand|', where 'k'
// depends on the interval range defined by 'interv' and the sign depends on
// 'sc' and the source sign.
// 	The mantissa is normalized to the interval specified by 'interv', which can
// take the following values:
//     _MM_MANT_NORM_1_2     // interval [1, 2)
//     _MM_MANT_NORM_p5_2    // interval [0.5, 2)
//     _MM_MANT_NORM_p5_1    // interval [0.5, 1)
//     _MM_MANT_NORM_p75_1p5 // interval [0.75, 1.5)The sign is determined by 'sc' which can take the following values:
//     _MM_MANT_SIGN_src     // sign = sign(src)
//     _MM_MANT_SIGN_zero    // sign = 0
//     _MM_MANT_SIGN_nan     // dst = NaN if sign(src) = 1 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := GetNormalizedMantissa(a[i+31:i], sc, interv)
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VGETMANTPS'. Intrinsic: '_mm512_getmant_ps'.
// Requires KNCNI.
func M512GetmantPs(a x86.M512, interv MMMANTISSANORMENUM, sc MMMANTISSASIGNENUM) (dst x86.M512) {
	panic("not implemented")
}


// M512MaskGetmantPs: Normalize the mantissas of packed single-precision
// (32-bit) floating-point elements in 'a', and store the results in 'dst'
// using writemask 'k' (elements are copied from 'src' when the corresponding
// mask bit is not set). This intrinsic essentially calculates
// '±(2^k)*|x.significand|', where 'k' depends on the interval range defined
// by 'interv' and the sign depends on 'sc' and the source sign.
// 	The mantissa is normalized to the interval specified by 'interv', which can
// take the following values:
//     _MM_MANT_NORM_1_2     // interval [1, 2)
//     _MM_MANT_NORM_p5_2    // interval [0.5, 2)
//     _MM_MANT_NORM_p5_1    // interval [0.5, 1)
//     _MM_MANT_NORM_p75_1p5 // interval [0.75, 1.5)The sign is determined by 'sc' which can take the following values:
//     _MM_MANT_SIGN_src     // sign = sign(src)
//     _MM_MANT_SIGN_zero    // sign = 0
//     _MM_MANT_SIGN_nan     // dst = NaN if sign(src) = 1 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := GetNormalizedMantissa(a[i+31:i], sc, interv)
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VGETMANTPS'. Intrinsic: '_mm512_mask_getmant_ps'.
// Requires KNCNI.
func M512MaskGetmantPs(src x86.M512, k x86.Mmask16, a x86.M512, interv MMMANTISSANORMENUM, sc MMMANTISSASIGNENUM) (dst x86.M512) {
	panic("not implemented")
}


// M512GetmantRoundPd: Normalize the mantissas of packed double-precision
// (64-bit) floating-point elements in 'a', and store the results in 'dst'.
// This intrinsic essentially calculates '±(2^k)*|x.significand|', where 'k'
// depends on the interval range defined by 'interv' and the sign depends on
// 'sc' and the source sign.
// 	The mantissa is normalized to the interval specified by 'interv', which can
// take the following values:
//     _MM_MANT_NORM_1_2     // interval [1, 2)
//     _MM_MANT_NORM_p5_2    // interval [0.5, 2)
//     _MM_MANT_NORM_p5_1    // interval [0.5, 1)
//     _MM_MANT_NORM_p75_1p5 // interval [0.75, 1.5)The sign is determined by 'sc' which can take the following values:
//     _MM_MANT_SIGN_src     // sign = sign(src)
//     _MM_MANT_SIGN_zero    // sign = 0
//     _MM_MANT_SIGN_nan     // dst = NaN if sign(src) = 1Rounding is done according to the 'rounding' parameter, which can be one of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 7
//			i := j*64
//			dst[i+63:i] := GetNormalizedMantissa(a[i+63:i], sc, interv)
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VGETMANTPD'. Intrinsic: '_mm512_getmant_round_pd'.
// Requires KNCNI.
func M512GetmantRoundPd(a x86.M512d, interv MMMANTISSANORMENUM, sc MMMANTISSASIGNENUM, rounding int) (dst x86.M512d) {
	panic("not implemented")
}


// M512MaskGetmantRoundPd: Normalize the mantissas of packed double-precision
// (64-bit) floating-point elements in 'a', and store the results in 'dst'
// using writemask 'k' (elements are copied from 'src' when the corresponding
// mask bit is not set). This intrinsic essentially calculates
// '±(2^k)*|x.significand|', where 'k' depends on the interval range defined
// by 'interv' and the sign depends on 'sc' and the source sign.
// 	The mantissa is normalized to the interval specified by 'interv', which can
// take the following values:
//     _MM_MANT_NORM_1_2     // interval [1, 2)
//     _MM_MANT_NORM_p5_2    // interval [0.5, 2)
//     _MM_MANT_NORM_p5_1    // interval [0.5, 1)
//     _MM_MANT_NORM_p75_1p5 // interval [0.75, 1.5)The sign is determined by 'sc' which can take the following values:
//     _MM_MANT_SIGN_src     // sign = sign(src)
//     _MM_MANT_SIGN_zero    // sign = 0
//     _MM_MANT_SIGN_nan     // dst = NaN if sign(src) = 1Rounding is done according to the 'rounding' parameter, which can be one of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := GetNormalizedMantissa(a[i+63:i], sc, interv)
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VGETMANTPD'. Intrinsic: '_mm512_mask_getmant_round_pd'.
// Requires KNCNI.
func M512MaskGetmantRoundPd(src x86.M512d, k x86.Mmask8, a x86.M512d, interv MMMANTISSANORMENUM, sc MMMANTISSASIGNENUM, rounding int) (dst x86.M512d) {
	panic("not implemented")
}


// M512GetmantRoundPs: Normalize the mantissas of packed single-precision
// (32-bit) floating-point elements in 'a', and store the results in 'dst'.
// This intrinsic essentially calculates '±(2^k)*|x.significand|', where 'k'
// depends on the interval range defined by 'interv' and the sign depends on
// 'sc' and the source sign.
// 	The mantissa is normalized to the interval specified by 'interv', which can
// take the following values:
//     _MM_MANT_NORM_1_2     // interval [1, 2)
//     _MM_MANT_NORM_p5_2    // interval [0.5, 2)
//     _MM_MANT_NORM_p5_1    // interval [0.5, 1)
//     _MM_MANT_NORM_p75_1p5 // interval [0.75, 1.5)The sign is determined by 'sc' which can take the following values:
//     _MM_MANT_SIGN_src     // sign = sign(src)
//     _MM_MANT_SIGN_zero    // sign = 0
//     _MM_MANT_SIGN_nan     // dst = NaN if sign(src) = 1Rounding is done according to the 'rounding' parameter, which can be one of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := GetNormalizedMantissa(a[i+31:i], sc, interv)
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VGETMANTPS'. Intrinsic: '_mm512_getmant_round_ps'.
// Requires KNCNI.
func M512GetmantRoundPs(a x86.M512, interv MMMANTISSANORMENUM, sc MMMANTISSASIGNENUM, rounding int) (dst x86.M512) {
	panic("not implemented")
}


// M512MaskGetmantRoundPs: Normalize the mantissas of packed single-precision
// (32-bit) floating-point elements in 'a', and store the results in 'dst'
// using writemask 'k' (elements are copied from 'src' when the corresponding
// mask bit is not set). This intrinsic essentially calculates
// '±(2^k)*|x.significand|', where 'k' depends on the interval range defined
// by 'interv' and the sign depends on 'sc' and the source sign.
// 	The mantissa is normalized to the interval specified by 'interv', which can
// take the following values:
//     _MM_MANT_NORM_1_2     // interval [1, 2)
//     _MM_MANT_NORM_p5_2    // interval [0.5, 2)
//     _MM_MANT_NORM_p5_1    // interval [0.5, 1)
//     _MM_MANT_NORM_p75_1p5 // interval [0.75, 1.5)The sign is determined by 'sc' which can take the following values:
//     _MM_MANT_SIGN_src     // sign = sign(src)
//     _MM_MANT_SIGN_zero    // sign = 0
//     _MM_MANT_SIGN_nan     // dst = NaN if sign(src) = 1Rounding is done according to the 'rounding' parameter, which can be one of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := GetNormalizedMantissa(a[i+31:i], sc, interv)
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VGETMANTPS'. Intrinsic: '_mm512_mask_getmant_round_ps'.
// Requires KNCNI.
func M512MaskGetmantRoundPs(src x86.M512, k x86.Mmask16, a x86.M512, interv MMMANTISSANORMENUM, sc MMMANTISSASIGNENUM, rounding int) (dst x86.M512) {
	panic("not implemented")
}


// M512GmaxPd: Determines the maximum of each pair of corresponding elements in
// packed double-precision (64-bit) floating-point elements in 'a' and 'b',
// storing the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*64
//			dst[i+63:i] := FpMax(a[i+63:i], b[i+63:i])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VGMAXPD'. Intrinsic: '_mm512_gmax_pd'.
// Requires KNCNI.
func M512GmaxPd(a x86.M512d, b x86.M512d) (dst x86.M512d) {
	panic("not implemented")
}


// M512MaskGmaxPd: Determines the maximum of each pair of corresponding
// elements of packed double-precision (64-bit) floating-point elements in 'a'
// and 'b', storing the results in 'dst' using writemask 'k' (elements are
// copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := FpMax(a[i+63:i], b[i+63:i])
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VGMAXPD'. Intrinsic: '_mm512_mask_gmax_pd'.
// Requires KNCNI.
func M512MaskGmaxPd(src x86.M512d, k x86.Mmask8, a x86.M512d, b x86.M512d) (dst x86.M512d) {
	panic("not implemented")
}


// M512GmaxPs: Determines the maximum of each pair of corresponding elements in
// packed single-precision (32-bit) floating-point elements in 'a' and 'b',
// storing the results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := FpMax(a[i+31:i], b[i+31:i])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VGMAXPS'. Intrinsic: '_mm512_gmax_ps'.
// Requires KNCNI.
func M512GmaxPs(a x86.M512, b x86.M512) (dst x86.M512) {
	panic("not implemented")
}


// M512MaskGmaxPs: Determines the maximum of each pair of corresponding
// elements of packed single-precision (32-bit) floating-point elements in 'a'
// and 'b', storing the results in 'dst' using writemask 'k' (elements are
// copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := FpMax(a[i+31:i], b[i+31:i])
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VGMAXPS'. Intrinsic: '_mm512_mask_gmax_ps'.
// Requires KNCNI.
func M512MaskGmaxPs(src x86.M512, k x86.Mmask16, a x86.M512, b x86.M512) (dst x86.M512) {
	panic("not implemented")
}


// M512GmaxabsPs: Determines the maximum of the absolute elements of each pair
// of corresponding elements of packed single-precision (32-bit) floating-point
// elements in 'a' and 'b', storing the results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := FpMax(Abs(a[i+31:i]), Abs(b[i+31:i]))
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VGMAXABSPS'. Intrinsic: '_mm512_gmaxabs_ps'.
// Requires KNCNI.
func M512GmaxabsPs(a x86.M512, b x86.M512) (dst x86.M512) {
	panic("not implemented")
}


// M512MaskGmaxabsPs: Determines the maximum of the absolute elements of each
// pair of corresponding elements of packed single-precision (32-bit)
// floating-point elements in 'a' and 'b', storing the results in 'dst' using
// writemask 'k' (elements are copied from 'src' when the corresponding mask
// bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := FpMax(Abs(a[i+31:i]), Abs(b[i+31:i]))
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VGMAXABSPS'. Intrinsic: '_mm512_mask_gmaxabs_ps'.
// Requires KNCNI.
func M512MaskGmaxabsPs(src x86.M512, k x86.Mmask16, a x86.M512, b x86.M512) (dst x86.M512) {
	panic("not implemented")
}


// M512GminPd: Determines the minimum of each pair of corresponding elements in
// packed double-precision (64-bit) floating-point elements in 'a' and 'b',
// storing the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*64
//			dst[i+63:i] := FpMin(a[i+63:i], b[i+63:i])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VGMINPD'. Intrinsic: '_mm512_gmin_pd'.
// Requires KNCNI.
func M512GminPd(a x86.M512d, b x86.M512d) (dst x86.M512d) {
	panic("not implemented")
}


// M512MaskGminPd: Determines the maximum of each pair of corresponding
// elements of packed double-precision (64-bit) floating-point elements in 'a'
// and 'b', storing the results in 'dst' using writemask 'k' (elements are
// copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := FpMin(a[i+63:i], b[i+63:i])
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VGMINPD'. Intrinsic: '_mm512_mask_gmin_pd'.
// Requires KNCNI.
func M512MaskGminPd(src x86.M512d, k x86.Mmask8, a x86.M512d, b x86.M512d) (dst x86.M512d) {
	panic("not implemented")
}


// M512GminPs: Determines the minimum of each pair of corresponding elements in
// packed single-precision (32-bit) floating-point elements in 'a' and 'b',
// storing the results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := FpMin(a[i+31:i], b[i+31:i])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VGMINPS'. Intrinsic: '_mm512_gmin_ps'.
// Requires KNCNI.
func M512GminPs(a x86.M512, b x86.M512) (dst x86.M512) {
	panic("not implemented")
}


// M512MaskGminPs: Determines the maximum of each pair of corresponding
// elements of packed single-precision (32-bit) floating-point elements in 'a'
// and 'b', storing the results in 'dst' using writemask 'k' (elements are
// copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := FpMin(a[i+31:i], b[i+31:i])
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VGMINPS'. Intrinsic: '_mm512_mask_gmin_ps'.
// Requires KNCNI.
func M512MaskGminPs(src x86.M512, k x86.Mmask16, a x86.M512, b x86.M512) (dst x86.M512) {
	panic("not implemented")
}


// Skipped: _mm512_i32extgather_epi32. Contains pointer parameter.


// Skipped: _mm512_mask_i32extgather_epi32. Contains pointer parameter.


// Skipped: _mm512_i32extgather_ps. Contains pointer parameter.


// Skipped: _mm512_mask_i32extgather_ps. Contains pointer parameter.


// Skipped: _mm512_i32extscatter_epi32. Contains pointer parameter.


// Skipped: _mm512_mask_i32extscatter_epi32. Contains pointer parameter.


// Skipped: _mm512_i32extscatter_ps. Contains pointer parameter.


// Skipped: _mm512_mask_i32extscatter_ps. Contains pointer parameter.


// Skipped: _mm512_i32gather_epi32. Contains pointer parameter.


// Skipped: _mm512_mask_i32gather_epi32. Contains pointer parameter.


// Skipped: _mm512_i32gather_ps. Contains pointer parameter.


// Skipped: _mm512_mask_i32gather_ps. Contains pointer parameter.


// Skipped: _mm512_i32loextgather_epi64. Contains pointer parameter.


// Skipped: _mm512_mask_i32loextgather_epi64. Contains pointer parameter.


// Skipped: _mm512_i32loextgather_pd. Contains pointer parameter.


// Skipped: _mm512_mask_i32loextgather_pd. Contains pointer parameter.


// Skipped: _mm512_i32loextscatter_epi64. Contains pointer parameter.


// Skipped: _mm512_mask_i32loextscatter_epi64. Contains pointer parameter.


// Skipped: _mm512_i32loextscatter_pd. Contains pointer parameter.


// Skipped: _mm512_mask_i32loextscatter_pd. Contains pointer parameter.


// Skipped: _mm512_i32logather_epi64. Contains pointer parameter.


// Skipped: _mm512_mask_i32logather_epi64. Contains pointer parameter.


// Skipped: _mm512_i32logather_pd. Contains pointer parameter.


// Skipped: _mm512_mask_i32logather_pd. Contains pointer parameter.


// Skipped: _mm512_i32loscatter_epi64. Contains pointer parameter.


// Skipped: _mm512_mask_i32loscatter_epi64. Contains pointer parameter.


// Skipped: _mm512_i32loscatter_pd. Contains pointer parameter.


// Skipped: _mm512_mask_i32loscatter_pd. Contains pointer parameter.


// Skipped: _mm512_i32scatter_epi32. Contains pointer parameter.


// Skipped: _mm512_mask_i32scatter_epi32. Contains pointer parameter.


// Skipped: _mm512_i32scatter_ps. Contains pointer parameter.


// Skipped: _mm512_mask_i32scatter_ps. Contains pointer parameter.


// Skipped: _mm512_i64extgather_epi32lo. Contains pointer parameter.


// Skipped: _mm512_mask_i64extgather_epi32lo. Contains pointer parameter.


// Skipped: _mm512_i64extgather_epi64. Contains pointer parameter.


// Skipped: _mm512_mask_i64extgather_epi64. Contains pointer parameter.


// Skipped: _mm512_i64extgather_pd. Contains pointer parameter.


// Skipped: _mm512_mask_i64extgather_pd. Contains pointer parameter.


// Skipped: _mm512_i64extgather_pslo. Contains pointer parameter.


// Skipped: _mm512_mask_i64extgather_pslo. Contains pointer parameter.


// Skipped: _mm512_i64extscatter_epi32lo. Contains pointer parameter.


// Skipped: _mm512_mask_i64extscatter_epi32lo. Contains pointer parameter.


// Skipped: _mm512_i64extscatter_epi64. Contains pointer parameter.


// Skipped: _mm512_mask_i64extscatter_epi64. Contains pointer parameter.


// Skipped: _mm512_i64extscatter_pd. Contains pointer parameter.


// Skipped: _mm512_mask_i64extscatter_pd. Contains pointer parameter.


// Skipped: _mm512_i64extscatter_pslo. Contains pointer parameter.


// Skipped: _mm512_mask_i64extscatter_pslo. Contains pointer parameter.


// Skipped: _mm512_i64gather_epi32lo. Contains pointer parameter.


// Skipped: _mm512_mask_i64gather_epi32lo. Contains pointer parameter.


// Skipped: _mm512_i64gather_pslo. Contains pointer parameter.


// Skipped: _mm512_mask_i64gather_pslo. Contains pointer parameter.


// Skipped: _mm512_i64scatter_epi32lo. Contains pointer parameter.


// Skipped: _mm512_mask_i64scatter_epi32lo. Contains pointer parameter.


// Skipped: _mm512_i64scatter_pslo. Contains pointer parameter.


// Skipped: _mm512_mask_i64scatter_pslo. Contains pointer parameter.


// M512Int2mask: Converts integer 'mask' into bitmask, storing the result in
// 'dst'. 
//
//		dst := mask[15:0]
//
// Instruction: 'KMOV'. Intrinsic: '_mm512_int2mask'.
// Requires KNCNI.
func M512Int2mask(mask int) (dst x86.Mmask16) {
	panic("not implemented")
}


// M512Kand: Compute the bitwise AND of 16-bit masks 'a' and 'b', and store the
// result in 'k'. 
//
//		k[15:0] := a[15:0] AND b[15:0]
//		k[MAX:16] := 0
//
// Instruction: 'KAND'. Intrinsic: '_mm512_kand'.
// Requires KNCNI.
func M512Kand(a x86.Mmask16, b x86.Mmask16) (dst x86.Mmask16) {
	panic("not implemented")
}


// M512Kandn: Compute the bitwise AND NOT of 16-bit masks 'a' and 'b', and
// store the result in 'k'. 
//
//		k[15:0] := (NOT a[15:0]) AND b[15:0]
//		k[MAX:16] := 0
//
// Instruction: 'KANDN'. Intrinsic: '_mm512_kandn'.
// Requires KNCNI.
func M512Kandn(a x86.Mmask16, b x86.Mmask16) (dst x86.Mmask16) {
	panic("not implemented")
}


// M512Kandnr: Performs a bitwise AND operation between NOT of 'k2' and 'k1',
// storing the result in 'dst'. 
//
//		dst[15:0] := NOT(k2[15:0]) & k1[15:0]
//
// Instruction: 'KANDNR'. Intrinsic: '_mm512_kandnr'.
// Requires KNCNI.
func M512Kandnr(k1 x86.Mmask16, k2 x86.Mmask16) (dst x86.Mmask16) {
	panic("not implemented")
}


// M512Kconcathi64: Packs masks 'k1' and 'k2' into the high 32 bits of 'dst'.
// The rest of 'dst' is set to 0. 
//
//		dst[63:48] := k1[15:0]
//		dst[47:32] := k2[15:0]
//		dst[31:0]  := 0
//
// Instruction: 'KCONCATH'. Intrinsic: '_mm512_kconcathi_64'.
// Requires KNCNI.
func M512Kconcathi64(k1 x86.Mmask16, k2 x86.Mmask16) int64 {
	panic("not implemented")
}


// M512Kconcatlo64: Packs masks 'k1' and 'k2' into the low 32 bits of 'dst'.
// The rest of 'dst' is set to 0. 
//
//		dst[31:16] := k1[15:0]
//		dst[15:0]  := k2[15:0]
//		dst[63:32] := 0
//
// Instruction: 'KCONCATL'. Intrinsic: '_mm512_kconcatlo_64'.
// Requires KNCNI.
func M512Kconcatlo64(k1 x86.Mmask16, k2 x86.Mmask16) int64 {
	panic("not implemented")
}


// M512Kextract64: Extracts 16-bit value 'b' from 64-bit integer 'a', storing
// the result in 'dst'. 
//
//		CASE b of
//		0: dst[15:0] := a[63:48]
//		1: dst[15:0] := a[47:32]
//		2: dst[15:0] := a[31:16]
//		3: dst[15:0] := a[15:0]
//		ESAC
//		dst[MAX:15] := 0
//
// Instruction: 'KEXTRACT'. Intrinsic: '_mm512_kextract_64'.
// Requires KNCNI.
func M512Kextract64(a int64, b int) (dst x86.Mmask16) {
	panic("not implemented")
}


// M512Kmerge2l1h: Move the high element from 'k1' to the low element of 'k1',
// and insert the low element of 'k2' into the high element of 'k1'. 
//
//		tmp[7:0] := k1[15:8]
//		k1[15:8] := k2[7:0]
//		k1[7:0]  := tmp[7:0]
//
// Instruction: 'KMERGE2L1H'. Intrinsic: '_mm512_kmerge2l1h'.
// Requires KNCNI.
func M512Kmerge2l1h(k1 x86.Mmask16, k2 x86.Mmask16) (dst x86.Mmask16) {
	panic("not implemented")
}


// M512Kmerge2l1l: Insert the low element of 'k2' into the high element of
// 'k1'. 
//
//		k1[15:8] := k2[7:0]
//
// Instruction: 'KMERGE2L1L'. Intrinsic: '_mm512_kmerge2l1l'.
// Requires KNCNI.
func M512Kmerge2l1l(k1 x86.Mmask16, k2 x86.Mmask16) (dst x86.Mmask16) {
	panic("not implemented")
}


// M512Kmov: Copy 16-bit mask 'a' to 'k'. 
//
//		k[15:0] := a[15:0]
//		k[MAX:16] := 0
//
// Instruction: 'KMOV'. Intrinsic: '_mm512_kmov'.
// Requires KNCNI.
func M512Kmov(a x86.Mmask16) (dst x86.Mmask16) {
	panic("not implemented")
}


// M512Kmovlhb: Inserts the low byte of mask 'k2' into the high byte of 'dst',
// and copies the low byte of 'k1' to the low byte of 'dst'. 
//
//		dst[7:0] := k1[7:0]
//		dst[15:8] := k2[7:0]
//
// Instruction: 'KMERGE2L1L'. Intrinsic: '_mm512_kmovlhb'.
// Requires KNCNI.
func M512Kmovlhb(k1 x86.Mmask16, k2 x86.Mmask16) (dst x86.Mmask16) {
	panic("not implemented")
}


// M512Knot: Compute the bitwise NOT of 16-bit mask 'a', and store the result
// in 'k'. 
//
//		k[15:0] := NOT a[15:0]
//		k[MAX:16] := 0
//
// Instruction: 'KNOT'. Intrinsic: '_mm512_knot'.
// Requires KNCNI.
func M512Knot(a x86.Mmask16) (dst x86.Mmask16) {
	panic("not implemented")
}


// M512Kor: Compute the bitwise OR of 16-bit masks 'a' and 'b', and store the
// result in 'k'. 
//
//		k[15:0] := a[15:0] OR b[15:0]
//		k[MAX:16] := 0
//
// Instruction: 'KOR'. Intrinsic: '_mm512_kor'.
// Requires KNCNI.
func M512Kor(a x86.Mmask16, b x86.Mmask16) (dst x86.Mmask16) {
	panic("not implemented")
}


// M512Kortestc: Performs bitwise OR between 'k1' and 'k2', storing the result
// in 'dst'. CF flag is set if 'dst' consists of all 1's. 
//
//		dst[15:0] := k1[15:0] | k2[15:0]
//		IF PopCount(dst[15:0]) = 16
//			SetCF()
//		FI
//
// Instruction: 'KORTEST'. Intrinsic: '_mm512_kortestc'.
// Requires KNCNI.
func M512Kortestc(k1 x86.Mmask16, k2 x86.Mmask16) int {
	panic("not implemented")
}


// M512Kortestz: Performs bitwise OR between 'k1' and 'k2', storing the result
// in 'dst'. ZF flag is set if 'dst' is 0. 
//
//		dst[15:0] := k1[15:0] | k2[15:0]
//		IF dst = 0
//			SetZF()
//		FI
//
// Instruction: 'KORTEST'. Intrinsic: '_mm512_kortestz'.
// Requires KNCNI.
func M512Kortestz(k1 x86.Mmask16, k2 x86.Mmask16) int {
	panic("not implemented")
}


// M512Kswapb: Moves high byte from 'k2' to low byte of 'k1', and moves low
// byte of 'k2' to high byte of 'k1'. 
//
//		tmp[7:0] := k2[15:8]
//		k2[15:8] := k1[7:0]
//		k1[7:0]  := tmp[7:0]
//		
//		tmp[7:0] := k2[7:0]
//		k2[7:0]  := k1[15:8]
//		k1[15:8] := tmp[7:0]
//
// Instruction: 'KMERGE2L1H'. Intrinsic: '_mm512_kswapb'.
// Requires KNCNI.
func M512Kswapb(k1 x86.Mmask16, k2 x86.Mmask16) (dst x86.Mmask16) {
	panic("not implemented")
}


// M512Kxnor: Compute the bitwise XNOR of 16-bit masks 'a' and 'b', and store
// the result in 'k'. 
//
//		k[15:0] := NOT (a[15:0] XOR b[15:0])
//		k[MAX:16] := 0
//
// Instruction: 'KXNOR'. Intrinsic: '_mm512_kxnor'.
// Requires KNCNI.
func M512Kxnor(a x86.Mmask16, b x86.Mmask16) (dst x86.Mmask16) {
	panic("not implemented")
}


// M512Kxor: Compute the bitwise XOR of 16-bit masks 'a' and 'b', and store the
// result in 'k'. 
//
//		k[15:0] := a[15:0] XOR b[15:0]
//		k[MAX:16] := 0
//
// Instruction: 'KXOR'. Intrinsic: '_mm512_kxor'.
// Requires KNCNI.
func M512Kxor(a x86.Mmask16, b x86.Mmask16) (dst x86.Mmask16) {
	panic("not implemented")
}


// Skipped: _mm512_load_epi32. Contains pointer parameter.


// Skipped: _mm512_mask_load_epi32. Contains pointer parameter.


// Skipped: _mm512_load_epi64. Contains pointer parameter.


// Skipped: _mm512_mask_load_epi64. Contains pointer parameter.


// Skipped: _mm512_load_pd. Contains pointer parameter.


// Skipped: _mm512_mask_load_pd. Contains pointer parameter.


// Skipped: _mm512_load_ps. Contains pointer parameter.


// Skipped: _mm512_mask_load_ps. Contains pointer parameter.


// Skipped: _mm512_load_si512. Contains pointer parameter.


// Skipped: _mm512_loadunpackhi_epi32. Contains pointer parameter.


// Skipped: _mm512_mask_loadunpackhi_epi32. Contains pointer parameter.


// Skipped: _mm512_loadunpackhi_epi64. Contains pointer parameter.


// Skipped: _mm512_mask_loadunpackhi_epi64. Contains pointer parameter.


// Skipped: _mm512_loadunpackhi_pd. Contains pointer parameter.


// Skipped: _mm512_mask_loadunpackhi_pd. Contains pointer parameter.


// Skipped: _mm512_loadunpackhi_ps. Contains pointer parameter.


// Skipped: _mm512_mask_loadunpackhi_ps. Contains pointer parameter.


// Skipped: _mm512_loadunpacklo_epi32. Contains pointer parameter.


// Skipped: _mm512_mask_loadunpacklo_epi32. Contains pointer parameter.


// Skipped: _mm512_loadunpacklo_epi64. Contains pointer parameter.


// Skipped: _mm512_mask_loadunpacklo_epi64. Contains pointer parameter.


// Skipped: _mm512_loadunpacklo_pd. Contains pointer parameter.


// Skipped: _mm512_mask_loadunpacklo_pd. Contains pointer parameter.


// Skipped: _mm512_loadunpacklo_ps. Contains pointer parameter.


// Skipped: _mm512_mask_loadunpacklo_ps. Contains pointer parameter.


// M512Log2Ps: Compute the base-2 logarithm of packed single-precision (32-bit)
// floating-point elements in 'a', and store the results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := log2(a[i+31:i])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VLOG2PS'. Intrinsic: '_mm512_log2_ps'.
// Requires KNCNI.
func M512Log2Ps(a x86.M512) (dst x86.M512) {
	panic("not implemented")
}


// M512MaskLog2Ps: Compute the base-2 logarithm of packed single-precision
// (32-bit) floating-point elements in 'a', and store the results in 'dst'
// using writemask 'k' (elements are copied from 'src' when the corresponding
// mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := log2(a[i+31:i])
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VLOG2PS'. Intrinsic: '_mm512_mask_log2_ps'.
// Requires KNCNI.
func M512MaskLog2Ps(src x86.M512, k x86.Mmask16, a x86.M512) (dst x86.M512) {
	panic("not implemented")
}


// M512Log2ae23Ps: Compute the base-2 logarithm of packed single-precision
// (32-bit) floating-point elements in 'a' with absolute error of 2^(-23) and
// store the results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := Log2ae23(a[i+31:i])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VLOG2PS'. Intrinsic: '_mm512_log2ae23_ps'.
// Requires KNCNI.
func M512Log2ae23Ps(a x86.M512) (dst x86.M512) {
	panic("not implemented")
}


// M512MaskLog2ae23Ps: Compute the base-2 logarithm of packed single-precision
// (32-bit) floating-point elements in 'a' with absolute error of 2^(-23) and
// store the results in 'dst' using writemask 'k' (elements are copied from
// 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := Log2ae23(a[i+31:i])
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VLOG2PS'. Intrinsic: '_mm512_mask_log2ae23_ps'.
// Requires KNCNI.
func M512MaskLog2ae23Ps(src x86.M512, k x86.Mmask16, a x86.M512) (dst x86.M512) {
	panic("not implemented")
}


// M512Mask2int: Converts bit mask 'k1' into an integer value, storing the
// results in 'dst'. 
//
//		dst := SignExtend(k1)
//
// Instruction: 'KMOV'. Intrinsic: '_mm512_mask2int'.
// Requires KNCNI.
func M512Mask2int(k1 x86.Mmask16) int {
	panic("not implemented")
}


// M512MaskMaxEpi32: Compare packed 32-bit integers in 'a' and 'b', and store
// packed maximum values in 'dst' using writemask 'k' (elements are copied from
// 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				IF a[i+31:i] > b[i+31:i]
//					dst[i+31:i] := a[i+31:i]
//				ELSE
//					dst[i+31:i] := b[i+31:i]
//				FI
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPMAXSD'. Intrinsic: '_mm512_mask_max_epi32'.
// Requires KNCNI.
func M512MaskMaxEpi32(src x86.M512i, k x86.Mmask16, a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaxEpi32: Compare packed 32-bit integers in 'a' and 'b', and store
// packed maximum values in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF a[i+31:i] > b[i+31:i]
//				dst[i+31:i] := a[i+31:i]
//			ELSE
//				dst[i+31:i] := b[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPMAXSD'. Intrinsic: '_mm512_max_epi32'.
// Requires KNCNI.
func M512MaxEpi32(a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskMaxEpu32: Compare packed unsigned 32-bit integers in 'a' and 'b',
// and store packed maximum values in 'dst' using writemask 'k' (elements are
// copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				IF a[i+31:i] > b[i+31:i]
//					dst[i+31:i] := a[i+31:i]
//				ELSE
//					dst[i+31:i] := b[i+31:i]
//				FI
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPMAXUD'. Intrinsic: '_mm512_mask_max_epu32'.
// Requires KNCNI.
func M512MaskMaxEpu32(src x86.M512i, k x86.Mmask16, a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaxEpu32: Compare packed unsigned 32-bit integers in 'a' and 'b', and
// store packed maximum values in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF a[i+31:i] > b[i+31:i]
//				dst[i+31:i] := a[i+31:i]
//			ELSE
//				dst[i+31:i] := b[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPMAXUD'. Intrinsic: '_mm512_max_epu32'.
// Requires KNCNI.
func M512MaxEpu32(a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskMaxabsPs: Determines the maximum of the absolute elements of each
// pair of corresponding elements of packed single-precision (32-bit)
// floating-point elements in 'a' and 'b', storing the results in 'dst' using
// writemask 'k' (elements are copied from 'src' when the corresponding mask
// bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := FpMax(Abs(a[i+31:i]), Abs(b[i+31:i]))
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VGMAXABSPS'. Intrinsic: '_mm512_mask_maxabs_ps'.
// Requires KNCNI.
func M512MaskMaxabsPs(src x86.M512, k x86.Mmask16, a x86.M512, b x86.M512) (dst x86.M512) {
	panic("not implemented")
}


// M512MaxabsPs: Determines the maximum of the absolute elements of each pair
// of corresponding elements of packed single-precision (32-bit) floating-point
// elements in 'a' and 'b', storing the results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := FpMax(Abs(a[i+31:i]), Abs(b[i+31:i]))
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VGMAXABSPS'. Intrinsic: '_mm512_maxabs_ps'.
// Requires KNCNI.
func M512MaxabsPs(a x86.M512, b x86.M512) (dst x86.M512) {
	panic("not implemented")
}


// M512MaskMinEpi32: Compare packed 32-bit integers in 'a' and 'b', and store
// packed minimum values in 'dst' using writemask 'k' (elements are copied from
// 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				IF a[i+31:i] < b[i+31:i]
//						dst[i+31:i] := a[i+31:i]
//				ELSE
//					dst[i+31:i] := b[i+31:i]
//				FI
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPMINSD'. Intrinsic: '_mm512_mask_min_epi32'.
// Requires KNCNI.
func M512MaskMinEpi32(src x86.M512i, k x86.Mmask16, a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MinEpi32: Compare packed 32-bit integers in 'a' and 'b', and store
// packed minimum values in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF a[i+31:i] < b[i+31:i]
//				dst[i+31:i] := a[i+31:i]
//			ELSE
//				dst[i+31:i] := b[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPMINSD'. Intrinsic: '_mm512_min_epi32'.
// Requires KNCNI.
func M512MinEpi32(a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskMinEpu32: Compare packed unsigned 32-bit integers in 'a' and 'b',
// and store packed minimum values in 'dst' using writemask 'k' (elements are
// copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				IF a[i+31:i] < b[i+31:i]
//					dst[i+31:i] := a[i+31:i]
//				ELSE
//					dst[i+31:i] := b[i+31:i]
//				FI
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPMINUD'. Intrinsic: '_mm512_mask_min_epu32'.
// Requires KNCNI.
func M512MaskMinEpu32(src x86.M512i, k x86.Mmask16, a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MinEpu32: Compare packed unsigned 32-bit integers in 'a' and 'b', and
// store packed minimum values in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF a[i+31:i] < b[i+31:i]
//				dst[i+31:i] := a[i+31:i]
//			ELSE
//				dst[i+31:i] := b[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPMINUD'. Intrinsic: '_mm512_min_epu32'.
// Requires KNCNI.
func M512MinEpu32(a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskMovEpi32: Move packed 32-bit integers from 'a' to 'dst' using
// writemask 'k' (elements are copied from 'src' when the corresponding mask
// bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := a[i+31:i]
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VMOVDQA32'. Intrinsic: '_mm512_mask_mov_epi32'.
// Requires KNCNI.
func M512MaskMovEpi32(src x86.M512i, k x86.Mmask16, a x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskMovEpi64: Move packed 64-bit integers from 'a' to 'dst' using
// writemask 'k' (elements are copied from 'src' when the corresponding mask
// bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := a[i+63:i]
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VMOVDQA64'. Intrinsic: '_mm512_mask_mov_epi64'.
// Requires KNCNI.
func M512MaskMovEpi64(src x86.M512i, k x86.Mmask8, a x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskMovPd: Move packed double-precision (64-bit) floating-point elements
// from 'a' to 'dst' using writemask 'k' (elements are copied from 'src' when
// the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := a[i+63:i]
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VMOVAPD'. Intrinsic: '_mm512_mask_mov_pd'.
// Requires KNCNI.
func M512MaskMovPd(src x86.M512d, k x86.Mmask8, a x86.M512d) (dst x86.M512d) {
	panic("not implemented")
}


// M512MaskMovPs: Move packed single-precision (32-bit) floating-point elements
// from 'a' to 'dst' using writemask 'k' (elements are copied from 'src' when
// the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := a[i+31:i]
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VMOVAPS'. Intrinsic: '_mm512_mask_mov_ps'.
// Requires KNCNI.
func M512MaskMovPs(src x86.M512, k x86.Mmask16, a x86.M512) (dst x86.M512) {
	panic("not implemented")
}


// M512MaskMulPd: Multiply packed double-precision (64-bit) floating-point
// elements in 'a' and 'b', and store the results in 'dst' using writemask 'k'
// (elements are copied from 'src' when the corresponding mask bit is not set).
// RM. 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := a[i+63:i] * b[i+63:i]
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VMULPD'. Intrinsic: '_mm512_mask_mul_pd'.
// Requires KNCNI.
func M512MaskMulPd(src x86.M512d, k x86.Mmask8, a x86.M512d, b x86.M512d) (dst x86.M512d) {
	panic("not implemented")
}


// M512MulPd: Multiply packed double-precision (64-bit) floating-point elements
// in 'a' and 'b', and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*64
//			dst[i+63:i] := a[i+63:i] * b[i+63:i]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VMULPD'. Intrinsic: '_mm512_mul_pd'.
// Requires KNCNI.
func M512MulPd(a x86.M512d, b x86.M512d) (dst x86.M512d) {
	panic("not implemented")
}


// M512MaskMulPs: Multiply packed single-precision (32-bit) floating-point
// elements in 'a' and 'b', and store the results in 'dst' using writemask 'k'
// (elements are copied from 'src' when the corresponding mask bit is not set).
// RM. 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := a[i+31:i] * b[i+31:i]
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VMULPS'. Intrinsic: '_mm512_mask_mul_ps'.
// Requires KNCNI.
func M512MaskMulPs(src x86.M512, k x86.Mmask16, a x86.M512, b x86.M512) (dst x86.M512) {
	panic("not implemented")
}


// M512MulPs: Multiply packed single-precision (32-bit) floating-point elements
// in 'a' and 'b', and store the results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := a[i+31:i] * b[i+31:i]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VMULPS'. Intrinsic: '_mm512_mul_ps'.
// Requires KNCNI.
func M512MulPs(a x86.M512, b x86.M512) (dst x86.M512) {
	panic("not implemented")
}


// M512MaskMulRoundPd: Multiply packed double-precision (64-bit) floating-point
// elements in 'a' and 'b', and store the results in 'dst' using writemask 'k'
// (elements are copied from 'src' when the corresponding mask bit is not set).
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := a[i+63:i] * b[i+63:i]
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VMULPD'. Intrinsic: '_mm512_mask_mul_round_pd'.
// Requires KNCNI.
func M512MaskMulRoundPd(src x86.M512d, k x86.Mmask8, a x86.M512d, b x86.M512d, rounding int) (dst x86.M512d) {
	panic("not implemented")
}


// M512MulRoundPd: Multiply packed double-precision (64-bit) floating-point
// elements in 'a' and 'b', and store the results in 'dst'. 
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 7
//			i := j*64
//			dst[i+63:i] := a[i+63:i] * b[i+63:i]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VMULPD'. Intrinsic: '_mm512_mul_round_pd'.
// Requires KNCNI.
func M512MulRoundPd(a x86.M512d, b x86.M512d, rounding int) (dst x86.M512d) {
	panic("not implemented")
}


// M512MaskMulRoundPs: Multiply packed single-precision (32-bit) floating-point
// elements in 'a' and 'b', and store the results in 'dst' using writemask 'k'
// (elements are copied from 'src' when the corresponding mask bit is not set).
// 	 Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := a[i+31:i] * b[i+31:i]
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VMULPS'. Intrinsic: '_mm512_mask_mul_round_ps'.
// Requires KNCNI.
func M512MaskMulRoundPs(src x86.M512, k x86.Mmask16, a x86.M512, b x86.M512, rounding int) (dst x86.M512) {
	panic("not implemented")
}


// M512MulRoundPs: Multiply packed single-precision (32-bit) floating-point
// elements in 'a' and 'b', and store the results in 'dst'. 
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := a[i+31:i] * b[i+31:i]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VMULPS'. Intrinsic: '_mm512_mul_round_ps'.
// Requires KNCNI.
func M512MulRoundPs(a x86.M512, b x86.M512, rounding int) (dst x86.M512) {
	panic("not implemented")
}


// M512MaskMulhiEpi32: Performs element-by-element multiplication between
// packed 32-bit integer elements in 'a' and 'b' and stores the high 32 bits of
// each result into 'dst' using writemask 'k' (elements are copied from 'src'
// when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := (a[i+31:i] * b[i+31:i]) >> 32
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPMULHD'. Intrinsic: '_mm512_mask_mulhi_epi32'.
// Requires KNCNI.
func M512MaskMulhiEpi32(src x86.M512i, k x86.Mmask16, a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MulhiEpi32: Performs element-by-element multiplication between packed
// 32-bit integer elements in 'a' and 'b' and stores the high 32 bits of each
// result into 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := (a[i+31:i] * b[i+31:i]) >> 32
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPMULHD'. Intrinsic: '_mm512_mulhi_epi32'.
// Requires KNCNI.
func M512MulhiEpi32(a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskMulhiEpu32: Performs element-by-element multiplication between
// packed unsigned 32-bit integer elements in 'a' and 'b' and stores the high
// 32 bits of each result into 'dst' using writemask 'k' (elements are copied
// from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := (a[i+31:i] * b[i+31:i]) >> 32
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPMULHUD'. Intrinsic: '_mm512_mask_mulhi_epu32'.
// Requires KNCNI.
func M512MaskMulhiEpu32(src x86.M512i, k x86.Mmask16, a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MulhiEpu32: Performs element-by-element multiplication between packed
// unsigned 32-bit integer elements in 'a' and 'b' and stores the high 32 bits
// of each result into 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := (a[i+31:i] * b[i+31:i]) >> 32
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPMULHUD'. Intrinsic: '_mm512_mulhi_epu32'.
// Requires KNCNI.
func M512MulhiEpu32(a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskMulloEpi32: Multiply the packed 32-bit integers in 'a' and 'b',
// producing intermediate 64-bit integers, and store the low 32 bits of the
// intermediate integers in 'dst' using writemask 'k' (elements are copied from
// 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				tmp[63:0] := a[i+31:i] * b[i+31:i]
//				dst[i+31:i] := tmp[31:0]
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPMULLD'. Intrinsic: '_mm512_mask_mullo_epi32'.
// Requires KNCNI.
func M512MaskMulloEpi32(src x86.M512i, k x86.Mmask16, a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MulloEpi32: Multiply the packed 32-bit integers in 'a' and 'b',
// producing intermediate 64-bit integers, and store the low 32 bits of the
// intermediate integers in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			tmp[63:0] := a[i+31:i] * b[i+31:i]
//			dst[i+31:i] := tmp[31:0]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPMULLD'. Intrinsic: '_mm512_mullo_epi32'.
// Requires KNCNI.
func M512MulloEpi32(a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskOrEpi32: Compute the bitwise OR of packed 32-bit integers in 'a' and
// 'b', and store the results in 'dst' using writemask 'k' (elements are copied
// from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := a[i+31:i] OR b[i+31:i]
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPORD'. Intrinsic: '_mm512_mask_or_epi32'.
// Requires KNCNI.
func M512MaskOrEpi32(src x86.M512i, k x86.Mmask16, a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512OrEpi32: Compute the bitwise OR of packed 32-bit integers in 'a' and
// 'b', and store the results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := a[i+31:i] OR b[i+31:i]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPORD'. Intrinsic: '_mm512_or_epi32'.
// Requires KNCNI.
func M512OrEpi32(a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskOrEpi64: Compute the bitwise OR of packed 64-bit integers in 'a' and
// 'b', and store the results in 'dst' using writemask 'k' (elements are copied
// from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := a[i+63:i] OR b[i+63:i]
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPORQ'. Intrinsic: '_mm512_mask_or_epi64'.
// Requires KNCNI.
func M512MaskOrEpi64(src x86.M512i, k x86.Mmask8, a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512OrEpi64: Compute the bitwise OR of packed 64-bit integers in 'a' and
// 'b', and store the resut in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*64
//			dst[i+63:i] := a[i+63:i] OR b[i+63:i]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPORQ'. Intrinsic: '_mm512_or_epi64'.
// Requires KNCNI.
func M512OrEpi64(a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512OrSi512: Compute the bitwise OR of 512 bits (representing integer data)
// in 'a' and 'b', and store the result in 'dst'. 
//
//		dst[511:0] := (a[511:0] OR b[511:0])
//		dst[MAX:512] := 0
//
// Instruction: 'VPORD'. Intrinsic: '_mm512_or_si512'.
// Requires KNCNI.
func M512OrSi512(a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// Skipped: _mm512_mask_packstorehi_epi32. Contains pointer parameter.


// Skipped: _mm512_packstorehi_epi32. Contains pointer parameter.


// Skipped: _mm512_mask_packstorehi_epi64. Contains pointer parameter.


// Skipped: _mm512_packstorehi_epi64. Contains pointer parameter.


// Skipped: _mm512_mask_packstorehi_pd. Contains pointer parameter.


// Skipped: _mm512_packstorehi_pd. Contains pointer parameter.


// Skipped: _mm512_mask_packstorehi_ps. Contains pointer parameter.


// Skipped: _mm512_packstorehi_ps. Contains pointer parameter.


// Skipped: _mm512_mask_packstorelo_epi32. Contains pointer parameter.


// Skipped: _mm512_packstorelo_epi32. Contains pointer parameter.


// Skipped: _mm512_mask_packstorelo_epi64. Contains pointer parameter.


// Skipped: _mm512_packstorelo_epi64. Contains pointer parameter.


// Skipped: _mm512_mask_packstorelo_pd. Contains pointer parameter.


// Skipped: _mm512_packstorelo_pd. Contains pointer parameter.


// Skipped: _mm512_mask_packstorelo_ps. Contains pointer parameter.


// Skipped: _mm512_packstorelo_ps. Contains pointer parameter.


// M512MaskPermute4f128Epi32: Permutes 128-bit blocks of the packed 32-bit
// integer vector 'a' using constant 'imm8'. The results are stored in 'dst'
// using writemask 'k' (elements are copied from 'src' when the corresponding
// mask bit is not set). 
//
//		SELECT4(src, control) {
//			CASE control[1:0] OF
//			0: tmp[127:0] := src[127:0]
//			1: tmp[127:0] := src[255:128]
//			2: tmp[127:0] := src[383:256]
//			3: tmp[127:0] := src[511:384]
//			ESAC
//			RETURN tmp[127:0]
//		}
//		
//		tmp[511:0] := 0
//		FOR j := 0 to 4
//			i := j*128
//			n := j*2
//			tmp[i+127:i] := SELECT4(a[511:0], imm8[n+1:n])
//		ENDFOR
//		FOR j := 0 to 15
//			IF k[j]
//				dst[i+31:i] := tmp[i+31:i]
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPERMF32X4'. Intrinsic: '_mm512_mask_permute4f128_epi32'.
// Requires KNCNI.
//
// FIXME: Requires compiler support (has immediate)
func M512MaskPermute4f128Epi32(src x86.M512i, k x86.Mmask16, a x86.M512i, imm8 byte) (dst x86.M512i) {
	panic("not implemented")
}


// M512Permute4f128Epi32: Permutes 128-bit blocks of the packed 32-bit integer
// vector 'a' using constant 'imm8'. The results are stored in 'dst'. 
//
//		SELECT4(src, control) {
//			CASE control[1:0] OF
//			0: tmp[127:0] := src[127:0]
//			1: tmp[127:0] := src[255:128]
//			2: tmp[127:0] := src[383:256]
//			3: tmp[127:0] := src[511:384]
//			ESAC
//			RETURN tmp[127:0]
//		}
//		
//		FOR j := 0 to 3
//			i := j*128
//			n := j*2
//			dst[i+127:i] := SELECT4(a[511:0], imm8[n+1:n])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPERMF32X4'. Intrinsic: '_mm512_permute4f128_epi32'.
// Requires KNCNI.
//
// FIXME: Requires compiler support (has immediate)
func M512Permute4f128Epi32(a x86.M512i, imm8 byte) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskPermute4f128Ps: Permutes 128-bit blocks of the packed
// single-precision (32-bit) floating-point elements in 'a' using constant
// 'imm8'. The results are stored in 'dst' using writemask 'k' (elements are
// copied from 'src' when the corresponding mask bit is not set). 
//
//		SELECT4(src, control) {
//			CASE control[1:0] OF
//			0: tmp[127:0] := src[127:0]
//			1: tmp[127:0] := src[255:128]
//			2: tmp[127:0] := src[383:256]
//			3: tmp[127:0] := src[511:384]
//			ESAC
//			RETURN tmp[127:0]
//		}
//		
//		tmp[511:0] := 0
//		FOR j := 0 to 4
//			i := j*128
//			n := j*2
//			tmp[i+127:i] := SELECT4(a[511:0], imm8[n+1:n])
//		ENDFOR
//		FOR j := 0 to 15
//			IF k[j]
//				dst[i+31:i] := tmp[i+31:i]
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPERMF32X4'. Intrinsic: '_mm512_mask_permute4f128_ps'.
// Requires KNCNI.
//
// FIXME: Requires compiler support (has immediate)
func M512MaskPermute4f128Ps(src x86.M512, k x86.Mmask16, a x86.M512, imm8 byte) (dst x86.M512) {
	panic("not implemented")
}


// M512Permute4f128Ps: Permutes 128-bit blocks of the packed single-precision
// (32-bit) floating-point elements in 'a' using constant 'imm8'. The results
// are stored in 'dst'. 
//
//		SELECT4(src, control) {
//			CASE control[1:0] OF
//			0: tmp[127:0] := src[127:0]
//			1: tmp[127:0] := src[255:128]
//			2: tmp[127:0] := src[383:256]
//			3: tmp[127:0] := src[511:384]
//			ESAC
//			RETURN tmp[127:0]
//		}
//		
//		FOR j := 0 to 3
//			i := j*128
//			n := j*2
//			dst[i+127:i] := SELECT4(a[511:0], imm8[n+1:n])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPERMF32X4'. Intrinsic: '_mm512_permute4f128_ps'.
// Requires KNCNI.
//
// FIXME: Requires compiler support (has immediate)
func M512Permute4f128Ps(a x86.M512, imm8 byte) (dst x86.M512) {
	panic("not implemented")
}


// M512MaskPermutevarEpi32: Shuffle 32-bit integers in 'a' across lanes using
// the corresponding index in 'idx', and store the results in 'dst' using
// writemask 'k' (elements are copied from 'src' when the corresponding mask
// bit is not set). Note that this intrinsic shuffles across 128-bit lanes,
// unlike past intrinsics that use the 'permutevar' name. This intrinsic is
// identical to '_mm512_mask_permutexvar_epi32', and it is recommended that you
// use that intrinsic name. 
//
//		FOR j := 0 to 15
//			i := j*32
//			id := idx[i+3:i]*32
//			IF k[j]
//				dst[i+31:i] := a[id+31:id]
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPERMD'. Intrinsic: '_mm512_mask_permutevar_epi32'.
// Requires KNCNI.
func M512MaskPermutevarEpi32(src x86.M512i, k x86.Mmask16, idx x86.M512i, a x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512PermutevarEpi32: Shuffle 32-bit integers in 'a' across lanes using the
// corresponding index in 'idx', and store the results in 'dst'. Note that this
// intrinsic shuffles across 128-bit lanes, unlike past intrinsics that use the
// 'permutevar' name. This intrinsic is identical to
// '_mm512_permutexvar_epi32', and it is recommended that you use that
// intrinsic name. 
//
//		FOR j := 0 to 15
//			i := j*32
//			id := idx[i+3:i]*32
//			dst[i+31:i] := a[id+31:id]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPERMD'. Intrinsic: '_mm512_permutevar_epi32'.
// Requires KNCNI.
func M512PermutevarEpi32(idx x86.M512i, a x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// Prefetch: Fetch the line of data from memory that contains address 'p' to a
// location in the cache heirarchy specified by the locality hint 'i'. 
//
//		
//
// Instruction: 'VPREFETCH0, VPREFETCH1, VPREFETCH2, VPREFETCHNTA, VPREFETCHE0, VPREFETCHE1, VPREFETCHE2, VPREFETCHENTA'. Intrinsic: '_mm_prefetch'.
// Requires KNCNI.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func Prefetch(p *byte, i int)  {
	panic("not implemented")
}


// Skipped: _mm512_mask_prefetch_i32extgather_ps. Contains pointer parameter.


// Skipped: _mm512_prefetch_i32extgather_ps. Contains pointer parameter.


// Skipped: _mm512_mask_prefetch_i32extscatter_ps. Contains pointer parameter.


// Skipped: _mm512_prefetch_i32extscatter_ps. Contains pointer parameter.


// Skipped: _mm512_mask_prefetch_i32gather_ps. Contains pointer parameter.


// Skipped: _mm512_prefetch_i32gather_ps. Contains pointer parameter.


// Skipped: _mm512_mask_prefetch_i32scatter_ps. Contains pointer parameter.


// Skipped: _mm512_prefetch_i32scatter_ps. Contains pointer parameter.


// M512MaskRcp23Ps: Approximates the reciprocals of packed single-precision
// (32-bit) floating-point elements in 'a' to 23 bits of precision, storing the
// results in 'dst' using writemask 'k' (elements are copied from 'src' when
// the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := APPROXIMATE(1.0/a[i+31:i])
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VRCP23PS'. Intrinsic: '_mm512_mask_rcp23_ps'.
// Requires KNCNI.
func M512MaskRcp23Ps(src x86.M512, k x86.Mmask16, a x86.M512) (dst x86.M512) {
	panic("not implemented")
}


// M512Rcp23Ps: Approximates the reciprocals of packed single-precision
// (32-bit) floating-point elements in 'a' to 23 bits of precision, storing the
// results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := APPROXIMATE(1.0/a[i+31:i])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VRCP23PS'. Intrinsic: '_mm512_rcp23_ps'.
// Requires KNCNI.
func M512Rcp23Ps(a x86.M512) (dst x86.M512) {
	panic("not implemented")
}


// M512MaskReduceAddEpi32: Reduce the packed 32-bit integers in 'a' by addition
// using mask 'k'. Returns the sum of all active elements in 'a'. 
//
//		sum[31:0] := 0
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				sum[31:0] := sum[31:0] + a[i+31:i]
//			FI
//		ENDFOR
//		RETURN sum[31:0]
//
// Instruction: '...'. Intrinsic: '_mm512_mask_reduce_add_epi32'.
// Requires KNCNI.
func M512MaskReduceAddEpi32(k x86.Mmask16, a x86.M512i) int {
	panic("not implemented")
}


// M512ReduceAddEpi32: Reduce the packed 32-bit integers in 'a' by addition.
// Returns the sum of all elements in 'a'. 
//
//		sum[31:0] := 0
//		FOR j := 0 to 15
//			i := j*32
//			sum[31:0] := sum[31:0] + a[i+31:i]
//		ENDFOR
//		RETURN sum[63:0]
//
// Instruction: '...'. Intrinsic: '_mm512_reduce_add_epi32'.
// Requires KNCNI.
func M512ReduceAddEpi32(a x86.M512i) int {
	panic("not implemented")
}


// M512MaskReduceAddEpi64: Reduce the packed 64-bit integers in 'a' by addition
// using mask 'k'. Returns the sum of all active elements in 'a'. 
//
//		sum[63:0] := 0
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				sum[63:0] := sum[63:0] + a[i+63:i]
//			FI
//		ENDFOR
//		RETURN sum[63:0]
//
// Instruction: '...'. Intrinsic: '_mm512_mask_reduce_add_epi64'.
// Requires KNCNI.
func M512MaskReduceAddEpi64(k x86.Mmask8, a x86.M512i) int64 {
	panic("not implemented")
}


// M512ReduceAddEpi64: Reduce the packed 64-bit integers in 'a' by addition.
// Returns the sum of all elements in 'a'. 
//
//		sum[63:0] := 0
//		FOR j := 0 to 7
//			i := j*64
//			sum[63:0] := sum[63:0] + a[i+63:i]
//		ENDFOR
//		RETURN sum[63:0]
//
// Instruction: '...'. Intrinsic: '_mm512_reduce_add_epi64'.
// Requires KNCNI.
func M512ReduceAddEpi64(a x86.M512i) int64 {
	panic("not implemented")
}


// M512MaskReduceAddPd: Reduce the packed double-precision (64-bit)
// floating-point elements in 'a' by addition using mask 'k'. Returns the sum
// of all active elements in 'a'. 
//
//		sum[63:0] := 0
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				sum[63:0] := sum[63:0] + a[i+63:i]
//			FI
//		ENDFOR
//		RETURN sum[63:0]
//
// Instruction: '...'. Intrinsic: '_mm512_mask_reduce_add_pd'.
// Requires KNCNI.
func M512MaskReduceAddPd(k x86.Mmask8, a x86.M512d) float64 {
	panic("not implemented")
}


// M512ReduceAddPd: Reduce the packed double-precision (64-bit) floating-point
// elements in 'a' by addition. Returns the sum of all elements in 'a'. 
//
//		sum[63:0] := 0
//		FOR j := 0 to 7
//			i := j*64
//			sum[63:0] := sum[63:0] + a[i+63:i]
//		ENDFOR
//		RETURN sum[63:0]
//
// Instruction: '...'. Intrinsic: '_mm512_reduce_add_pd'.
// Requires KNCNI.
func M512ReduceAddPd(a x86.M512d) float64 {
	panic("not implemented")
}


// M512MaskReduceAddPs: Reduce the packed single-precision (32-bit)
// floating-point elements in 'a' by addition using mask 'k'. Returns the sum
// of all active elements in 'a'. 
//
//		sum[31:0] := 0
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				sum[31:0] := sum[31:0] + a[i+31:i]
//			FI
//		ENDFOR
//		RETURN sum[31:0]
//
// Instruction: '...'. Intrinsic: '_mm512_mask_reduce_add_ps'.
// Requires KNCNI.
func M512MaskReduceAddPs(k x86.Mmask16, a x86.M512) float32 {
	panic("not implemented")
}


// M512ReduceAddPs: Reduce the packed single-precision (32-bit) floating-point
// elements in 'a' by addition. Returns the sum of all elements in 'a'. 
//
//		sum[31:0] := 0
//		FOR j := 0 to 15
//			i := j*32
//			sum[31:0] := sum[31:0] + a[i+31:i]
//		ENDFOR
//		RETURN sum[63:0]
//
// Instruction: '...'. Intrinsic: '_mm512_reduce_add_ps'.
// Requires KNCNI.
func M512ReduceAddPs(a x86.M512) float32 {
	panic("not implemented")
}


// M512MaskReduceAndEpi32: Reduce the packed 32-bit integers in 'a' by bitwise
// AND using mask 'k'. Returns the bitwise AND of all active elements in 'a'. 
//
//		reduced[31:0] := 0xFFFFFFFF
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				reduced[31:0] := reduced[31:0] AND a[i+31:i]
//			FI
//		ENDFOR
//		RETURN reduced[31:0]
//
// Instruction: '...'. Intrinsic: '_mm512_mask_reduce_and_epi32'.
// Requires KNCNI.
func M512MaskReduceAndEpi32(k x86.Mmask16, a x86.M512i) int {
	panic("not implemented")
}


// M512ReduceAndEpi32: Reduce the packed 32-bit integers in 'a' by bitwise AND.
// Returns the bitwise AND of all elements in 'a'. 
//
//		reduced[31:0] := 0xFFFFFFFF
//		FOR j := 0 to 15
//			i := j*32
//			reduced[31:0] := reduced[31:0] AND a[i+31:i]
//		ENDFOR
//		RETURN reduced[31:0]
//
// Instruction: '...'. Intrinsic: '_mm512_reduce_and_epi32'.
// Requires KNCNI.
func M512ReduceAndEpi32(a x86.M512i) int {
	panic("not implemented")
}


// M512MaskReduceAndEpi64: Reduce the packed 64-bit integers in 'a' by bitwise
// AND using mask 'k'. Returns the bitwise AND of all active elements in 'a'. 
//
//		reduced[63:0] := 0xFFFFFFFFFFFFFFFF
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				reduced[63:0] := reduced[63:0] AND a[i+63:i]
//			FI
//		ENDFOR
//		RETURN reduced[63:0]
//
// Instruction: '...'. Intrinsic: '_mm512_mask_reduce_and_epi64'.
// Requires KNCNI.
func M512MaskReduceAndEpi64(k x86.Mmask8, a x86.M512i) int64 {
	panic("not implemented")
}


// M512ReduceAndEpi64: Reduce the packed 64-bit integers in 'a' by bitwise AND.
// Returns the bitwise AND of all elements in 'a'. 
//
//		reduced[63:0] := 0xFFFFFFFFFFFFFFFF
//		FOR j := 0 to 7
//			i := j*64
//			reduced[63:0] := reduced[63:0] AND a[i+63:i]
//		ENDFOR
//		RETURN reduced[63:0]
//
// Instruction: '...'. Intrinsic: '_mm512_reduce_and_epi64'.
// Requires KNCNI.
func M512ReduceAndEpi64(a x86.M512i) int64 {
	panic("not implemented")
}


// M512MaskReduceGmaxPd: Determines the maximum element of the packed
// double-precision (64-bit) floating-point elements stored in 'a' and stores
// the result in 'dst'. Bitmask 'k' is used to exclude certain elements
// (elements are ignored when the corresponding mask bit is not set). 
//
//		max = a[63:0]
//		FOR j := 1 to 7
//			i := j*64
//			IF k[j]
//				CONTINUE
//			ELSE
//				dst = FpMax(max, a[i+63:i])
//			FI
//		ENDFOR
//		dst := max
//
// Instruction: '...'. Intrinsic: '_mm512_mask_reduce_gmax_pd'.
// Requires KNCNI.
func M512MaskReduceGmaxPd(k x86.Mmask8, a x86.M512d) float64 {
	panic("not implemented")
}


// M512ReduceGmaxPd: Determines the maximum element of the packed
// double-precision (64-bit) floating-point elements stored in 'a' and stores
// the result in 'dst'. 
//
//		max = a[63:0]
//		FOR j := 1 to 7
//			i := j*64
//			dst = FpMax(max, a[i+63:i])
//		ENDFOR
//		dst := max
//
// Instruction: '...'. Intrinsic: '_mm512_reduce_gmax_pd'.
// Requires KNCNI.
func M512ReduceGmaxPd(a x86.M512d) float64 {
	panic("not implemented")
}


// M512MaskReduceGmaxPs: Determines the maximum element of the packed
// single-precision (32-bit) floating-point elements stored in 'a' and stores
// the result in 'dst'. Bitmask 'k' is used to exclude certain elements
// (elements are ignored when the corresponding mask bit is not set). 
//
//		max = a[31:0]
//		FOR j := 1 to 15
//			i := j*32
//			IF k[j]
//				CONTINUE
//			ELSE
//				dst = FpMax(max, a[i+31:i])
//			FI
//		ENDFOR
//		dst := max
//
// Instruction: '...'. Intrinsic: '_mm512_mask_reduce_gmax_ps'.
// Requires KNCNI.
func M512MaskReduceGmaxPs(k x86.Mmask16, a x86.M512) float32 {
	panic("not implemented")
}


// M512ReduceGmaxPs: Determines the maximum element of the packed
// single-precision (32-bit) floating-point elements stored in 'a' and stores
// the result in 'dst'. 
//
//		max = a[31:0]
//		FOR j := 1 to 15
//			i := j*32
//			dst = FpMax(max, a[i+31:i])
//		ENDFOR
//		dst := max
//
// Instruction: '...'. Intrinsic: '_mm512_reduce_gmax_ps'.
// Requires KNCNI.
func M512ReduceGmaxPs(a x86.M512) float32 {
	panic("not implemented")
}


// M512MaskReduceGminPd: Determines the minimum element of the packed
// double-precision (64-bit) floating-point elements stored in 'a' and stores
// the result in 'dst'. Bitmask 'k' is used to exclude certain elements
// (elements are ignored when the corresponding mask bit is not set). 
//
//		min = a[63:0]
//		FOR j := 1 to 7
//			i := j*64
//			IF k[j]
//				CONTINUE
//			ELSE
//				dst = FpMin(min, a[i+63:i])
//			FI
//		ENDFOR
//		dst := min
//
// Instruction: '...'. Intrinsic: '_mm512_mask_reduce_gmin_pd'.
// Requires KNCNI.
func M512MaskReduceGminPd(k x86.Mmask8, a x86.M512d) float64 {
	panic("not implemented")
}


// M512ReduceGminPd: Determines the minimum element of the packed
// double-precision (64-bit) floating-point elements stored in 'a' and stores
// the result in 'dst'. 
//
//		min = a[63:0]
//		FOR j := 1 to 7
//			i := j*64
//			dst = FpMin(min, a[i+63:i])
//		ENDFOR
//		dst := min
//
// Instruction: '...'. Intrinsic: '_mm512_reduce_gmin_pd'.
// Requires KNCNI.
func M512ReduceGminPd(a x86.M512d) float64 {
	panic("not implemented")
}


// M512MaskReduceGminPs: Determines the minimum element of the packed
// single-precision (32-bit) floating-point elements stored in 'a' and stores
// the result in 'dst' using writemask 'k' (elements are ignored when the
// corresponding mask bit is not set). 
//
//		min = a[31:0]
//		FOR j := 1 to 15
//			i := j*32
//			IF k[j]
//				CONTINUE
//			ELSE
//				dst = FpMin(min, a[i+31:i])
//			FI
//		ENDFOR
//		dst := min
//
// Instruction: '...'. Intrinsic: '_mm512_mask_reduce_gmin_ps'.
// Requires KNCNI.
func M512MaskReduceGminPs(k x86.Mmask16, a x86.M512) float32 {
	panic("not implemented")
}


// M512ReduceGminPs: Determines the minimum element of the packed
// single-precision (32-bit) floating-point elements stored in 'a' and stores
// the result in 'dst'. 
//
//		min = a[31:0]
//		FOR j := 1 to 15
//			i := j*32
//			dst = FpMin(min, a[i+31:i])
//		ENDFOR
//		dst := min
//
// Instruction: '...'. Intrinsic: '_mm512_reduce_gmin_ps'.
// Requires KNCNI.
func M512ReduceGminPs(a x86.M512) float32 {
	panic("not implemented")
}


// M512MaskReduceMaxEpi32: Reduce the packed 32-bit integers in 'a' by maximum
// using mask 'k'. Returns the maximum of all active elements in 'a'. 
//
//		max[31:0] := MIN_INT
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				max[31:0] := MAXIMUM(max[31:0], a[i+31:i])
//			FI
//		ENDFOR
//		RETURN max[31:0]
//
// Instruction: '...'. Intrinsic: '_mm512_mask_reduce_max_epi32'.
// Requires KNCNI.
func M512MaskReduceMaxEpi32(k x86.Mmask16, a x86.M512i) int {
	panic("not implemented")
}


// M512ReduceMaxEpi32: Reduce the packed 32-bit integers in 'a' by maximum.
// Returns the maximum of all elements in 'a'. 
//
//		max[31:0] := MIN_INT
//		FOR j := 0 to 15
//			i := j*32
//			max[31:0] := MAXIMUM(max[31:0], a[i+31:i])
//		ENDFOR
//		RETURN max[31:0]
//
// Instruction: '...'. Intrinsic: '_mm512_reduce_max_epi32'.
// Requires KNCNI.
func M512ReduceMaxEpi32(a x86.M512i) int {
	panic("not implemented")
}


// M512MaskReduceMaxEpi64: Reduce the packed 64-bit integers in 'a' by maximum
// using mask 'k'. Returns the maximum of all active elements in 'a'. 
//
//		max[63:0] := MIN_INT
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				max[63:0] := MAXIMUM(max[63:0], a[i+63:i])
//			FI
//		ENDFOR
//		RETURN max[63:0]
//
// Instruction: '...'. Intrinsic: '_mm512_mask_reduce_max_epi64'.
// Requires KNCNI.
func M512MaskReduceMaxEpi64(k x86.Mmask8, a x86.M512i) int64 {
	panic("not implemented")
}


// M512ReduceMaxEpi64: Reduce the packed 64-bit integers in 'a' by maximum.
// Returns the maximum of all elements in 'a'. 
//
//		max[63:0] := MIN_INT
//		FOR j := 0 to 7
//			i := j*64
//			max[63:0] := MAXIMUM(max[63:0], a[i+63:i])
//		ENDFOR
//		RETURN max[63:0]
//
// Instruction: '...'. Intrinsic: '_mm512_reduce_max_epi64'.
// Requires KNCNI.
func M512ReduceMaxEpi64(a x86.M512i) int64 {
	panic("not implemented")
}


// M512MaskReduceMaxEpu32: Reduce the packed unsigned 32-bit integers in 'a' by
// maximum using mask 'k'. Returns the maximum of all active elements in 'a'. 
//
//		max[31:0] := 0
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				max[31:0] := MAXIMUM(max[31:0], a[i+31:i])
//			FI
//		ENDFOR
//		RETURN max[31:0]
//
// Instruction: '...'. Intrinsic: '_mm512_mask_reduce_max_epu32'.
// Requires KNCNI.
func M512MaskReduceMaxEpu32(k x86.Mmask16, a x86.M512i) uint32 {
	panic("not implemented")
}


// M512ReduceMaxEpu32: Reduce the packed unsigned 32-bit integers in 'a' by
// maximum. Returns the maximum of all elements in 'a'. 
//
//		max[31:0] := 0
//		FOR j := 0 to 15
//			i := j*32
//			max[31:0] := MAXIMUM(max[31:0], a[i+31:i])
//		ENDFOR
//		RETURN max[31:0]
//
// Instruction: '...'. Intrinsic: '_mm512_reduce_max_epu32'.
// Requires KNCNI.
func M512ReduceMaxEpu32(a x86.M512i) uint32 {
	panic("not implemented")
}


// M512MaskReduceMaxEpu64: Reduce the packed unsigned 64-bit integers in 'a' by
// maximum using mask 'k'. Returns the maximum of all active elements in 'a'. 
//
//		max[63:0] := 0
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				max[63:0] := MAXIMUM(max[63:0], a[i+63:i])
//			FI
//		ENDFOR
//		RETURN max[63:0]
//
// Instruction: '...'. Intrinsic: '_mm512_mask_reduce_max_epu64'.
// Requires KNCNI.
func M512MaskReduceMaxEpu64(k x86.Mmask8, a x86.M512i) uint64 {
	panic("not implemented")
}


// M512ReduceMaxEpu64: Reduce the packed unsigned 64-bit integers in 'a' by
// maximum. Returns the maximum of all elements in 'a'. 
//
//		max[63:0] := 0
//		FOR j := 0 to 7
//			i := j*64
//			max[63:0] := MAXIMUM(max[63:0], a[i+63:i])
//		ENDFOR
//		RETURN max[63:0]
//
// Instruction: '...'. Intrinsic: '_mm512_reduce_max_epu64'.
// Requires KNCNI.
func M512ReduceMaxEpu64(a x86.M512i) uint64 {
	panic("not implemented")
}


// M512MaskReduceMaxPd: Reduce the packed double-precision (64-bit)
// floating-point elements in 'a' by maximum using mask 'k'. Returns the
// maximum of all active elements in 'a'. 
//
//		max[63:0] := MIN_DOUBLE
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				max[63:0] := MAXIMUM(max[63:0], a[i+63:i])
//			FI
//		ENDFOR
//		RETURN max[63:0]
//
// Instruction: '...'. Intrinsic: '_mm512_mask_reduce_max_pd'.
// Requires KNCNI.
func M512MaskReduceMaxPd(k x86.Mmask8, a x86.M512d) float64 {
	panic("not implemented")
}


// M512ReduceMaxPd: Reduce the packed double-precision (64-bit) floating-point
// elements in 'a' by maximum. Returns the maximum of all elements in 'a'. 
//
//		max[63:0] := MIN_DOUBLE
//		FOR j := 0 to 7
//			i := j*64
//			max[63:0] := MAXIMUM(max[63:0], a[i+63:i])
//		ENDFOR
//		RETURN max[63:0]
//
// Instruction: '...'. Intrinsic: '_mm512_reduce_max_pd'.
// Requires KNCNI.
func M512ReduceMaxPd(a x86.M512d) float64 {
	panic("not implemented")
}


// M512MaskReduceMaxPs: Reduce the packed single-precision (32-bit)
// floating-point elements in 'a' by maximum using mask 'k'. Returns the
// maximum of all active elements in 'a'. 
//
//		max[31:0] := MIN_FLOAT
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				max[31:0] := MAXIMUM(max[31:0], a[i+31:i])
//			FI
//		ENDFOR
//		RETURN max[31:0]
//
// Instruction: '...'. Intrinsic: '_mm512_mask_reduce_max_ps'.
// Requires KNCNI.
func M512MaskReduceMaxPs(k x86.Mmask16, a x86.M512) float32 {
	panic("not implemented")
}


// M512ReduceMaxPs: Reduce the packed single-precision (32-bit) floating-point
// elements in 'a' by maximum. Returns the maximum of all elements in 'a'. 
//
//		max[31:0] := MIN_FLOAT
//		FOR j := 0 to 15
//			i := j*32
//			max[31:0] := MAXIMUM(max[31:0], a[i+31:i])
//		ENDFOR
//		RETURN max[31:0]
//
// Instruction: '...'. Intrinsic: '_mm512_reduce_max_ps'.
// Requires KNCNI.
func M512ReduceMaxPs(a x86.M512) float32 {
	panic("not implemented")
}


// M512MaskReduceMinEpi32: Reduce the packed 32-bit integers in 'a' by maximum
// using mask 'k'. Returns the minimum of all active elements in 'a'. 
//
//		min[31:0] := MAX_INT
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				min[31:0] := MINIMUM(min[31:0], a[i+31:i])
//			FI
//		ENDFOR
//		RETURN min[31:0]
//
// Instruction: '...'. Intrinsic: '_mm512_mask_reduce_min_epi32'.
// Requires KNCNI.
func M512MaskReduceMinEpi32(k x86.Mmask16, a x86.M512i) int {
	panic("not implemented")
}


// M512ReduceMinEpi32: Reduce the packed 32-bit integers in 'a' by minimum.
// Returns the minimum of all elements in 'a'. 
//
//		min[31:0] := MAX_INT
//		FOR j := 0 to 15
//			i := j*32
//			min[31:0] := MINIMUM(min[31:0], a[i+31:i])
//		ENDFOR
//		RETURN min[31:0]
//
// Instruction: '...'. Intrinsic: '_mm512_reduce_min_epi32'.
// Requires KNCNI.
func M512ReduceMinEpi32(a x86.M512i) int {
	panic("not implemented")
}


// M512MaskReduceMinEpi64: Reduce the packed 64-bit integers in 'a' by maximum
// using mask 'k'. Returns the minimum of all active elements in 'a'. 
//
//		min[63:0] := MAX_INT
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				min[63:0] := MINIMUM(min[63:0], a[i+63:i])
//			FI
//		ENDFOR
//		RETURN min[63:0]
//
// Instruction: '...'. Intrinsic: '_mm512_mask_reduce_min_epi64'.
// Requires KNCNI.
func M512MaskReduceMinEpi64(k x86.Mmask8, a x86.M512i) int64 {
	panic("not implemented")
}


// M512ReduceMinEpi64: Reduce the packed 64-bit integers in 'a' by minimum.
// Returns the minimum of all elements in 'a'. 
//
//		min[63:0] := MAX_INT
//		FOR j := 0 to 7
//			i := j*64
//			min[63:0] := MINIMUM(min[63:0], a[i+63:i])
//		ENDFOR
//		RETURN min[63:0]
//
// Instruction: '...'. Intrinsic: '_mm512_reduce_min_epi64'.
// Requires KNCNI.
func M512ReduceMinEpi64(a x86.M512i) int64 {
	panic("not implemented")
}


// M512MaskReduceMinEpu32: Reduce the packed unsigned 32-bit integers in 'a' by
// maximum using mask 'k'. Returns the minimum of all active elements in 'a'. 
//
//		min[31:0] := MAX_UINT
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				min[31:0] := MINIMUM(min[31:0], a[i+31:i])
//			FI
//		ENDFOR
//		RETURN min[31:0]
//
// Instruction: '...'. Intrinsic: '_mm512_mask_reduce_min_epu32'.
// Requires KNCNI.
func M512MaskReduceMinEpu32(k x86.Mmask16, a x86.M512i) uint32 {
	panic("not implemented")
}


// M512ReduceMinEpu32: Reduce the packed unsigned 32-bit integers in 'a' by
// minimum. Returns the minimum of all elements in 'a'. 
//
//		min[31:0] := MAX_UINT
//		FOR j := 0 to 15
//			i := j*32
//			min[31:0] := MINIMUM(min[31:0], a[i+31:i])
//		ENDFOR
//		RETURN min[31:0]
//
// Instruction: '...'. Intrinsic: '_mm512_reduce_min_epu32'.
// Requires KNCNI.
func M512ReduceMinEpu32(a x86.M512i) uint32 {
	panic("not implemented")
}


// M512MaskReduceMinEpu64: Reduce the packed unsigned 64-bit integers in 'a' by
// minimum using mask 'k'. Returns the minimum of all active elements in 'a'. 
//
//		min[63:0] := MAX_UINT
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				min[63:0] := MINIMUM(min[63:0], a[i+63:i])
//			FI
//		ENDFOR
//		RETURN min[63:0]
//
// Instruction: '...'. Intrinsic: '_mm512_mask_reduce_min_epu64'.
// Requires KNCNI.
func M512MaskReduceMinEpu64(k x86.Mmask8, a x86.M512i) uint64 {
	panic("not implemented")
}


// M512ReduceMinEpu64: Reduce the packed unsigned 64-bit integers in 'a' by
// minimum. Returns the minimum of all elements in 'a'. 
//
//		min[63:0] := MAX_UINT
//		FOR j := 0 to 7
//			i := j*64
//			min[63:0] := MINIMUM(min[63:0], a[i+63:i])
//		ENDFOR
//		RETURN min[63:0]
//
// Instruction: '...'. Intrinsic: '_mm512_reduce_min_epu64'.
// Requires KNCNI.
func M512ReduceMinEpu64(a x86.M512i) uint64 {
	panic("not implemented")
}


// M512MaskReduceMinPd: Reduce the packed double-precision (64-bit)
// floating-point elements in 'a' by maximum using mask 'k'. Returns the
// minimum of all active elements in 'a'. 
//
//		min[63:0] := MAX_DOUBLE
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				min[63:0] := MINIMUM(min[63:0], a[i+63:i])
//			FI
//		ENDFOR
//		RETURN min[63:0]
//
// Instruction: '...'. Intrinsic: '_mm512_mask_reduce_min_pd'.
// Requires KNCNI.
func M512MaskReduceMinPd(k x86.Mmask8, a x86.M512d) float64 {
	panic("not implemented")
}


// M512ReduceMinPd: Reduce the packed double-precision (64-bit) floating-point
// elements in 'a' by minimum. Returns the minimum of all elements in 'a'. 
//
//		min[63:0] := MAX_DOUBLE
//		FOR j := 0 to 7
//			i := j*64
//			min[63:0] := MINIMUM(min[63:0], a[i+63:i])
//		ENDFOR
//		RETURN min[63:0]
//
// Instruction: '...'. Intrinsic: '_mm512_reduce_min_pd'.
// Requires KNCNI.
func M512ReduceMinPd(a x86.M512d) float64 {
	panic("not implemented")
}


// M512MaskReduceMinPs: Reduce the packed single-precision (32-bit)
// floating-point elements in 'a' by maximum using mask 'k'. Returns the
// minimum of all active elements in 'a'. 
//
//		min[31:0] := MAX_FLOAT
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				min[31:0] := MINIMUM(min[31:0], a[i+31:i])
//			FI
//		ENDFOR
//		RETURN min[31:0]
//
// Instruction: '...'. Intrinsic: '_mm512_mask_reduce_min_ps'.
// Requires KNCNI.
func M512MaskReduceMinPs(k x86.Mmask16, a x86.M512) float32 {
	panic("not implemented")
}


// M512ReduceMinPs: Reduce the packed single-precision (32-bit) floating-point
// elements in 'a' by minimum. Returns the minimum of all elements in 'a'. 
//
//		min[31:0] := MAX_INT
//		FOR j := 0 to 15
//			i := j*32
//			min[31:0] := MINIMUM(min[31:0], a[i+31:i])
//		ENDFOR
//		RETURN min[31:0]
//
// Instruction: '...'. Intrinsic: '_mm512_reduce_min_ps'.
// Requires KNCNI.
func M512ReduceMinPs(a x86.M512) float32 {
	panic("not implemented")
}


// M512MaskReduceMulEpi32: Reduce the packed 32-bit integers in 'a' by
// multiplication using mask 'k'. Returns the product of all active elements in
// 'a'. 
//
//		prod[31:0] := 1
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				prod[31:0] := prod[31:0] * a[i+31:i]
//			FI
//		ENDFOR
//		RETURN prod[31:0]
//
// Instruction: '...'. Intrinsic: '_mm512_mask_reduce_mul_epi32'.
// Requires KNCNI.
func M512MaskReduceMulEpi32(k x86.Mmask16, a x86.M512i) int {
	panic("not implemented")
}


// M512ReduceMulEpi32: Reduce the packed 32-bit integers in 'a' by
// multiplication. Returns the product of all elements in 'a'. 
//
//		prod[31:0] := 1
//		FOR j := 0 to 15
//			i := j*32
//			prod[31:0] := prod[31:0] * a[i+31:i]
//		ENDFOR
//		RETURN prod[31:0]
//
// Instruction: '...'. Intrinsic: '_mm512_reduce_mul_epi32'.
// Requires KNCNI.
func M512ReduceMulEpi32(a x86.M512i) int {
	panic("not implemented")
}


// M512MaskReduceMulEpi64: Reduce the packed 64-bit integers in 'a' by
// multiplication using mask 'k'. Returns the product of all active elements in
// 'a'. 
//
//		prod[63:0] := 1
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				prod[63:0] := prod[63:0] * a[i+63:i]
//			FI
//		ENDFOR
//		RETURN prod[63:0]
//
// Instruction: '...'. Intrinsic: '_mm512_mask_reduce_mul_epi64'.
// Requires KNCNI.
func M512MaskReduceMulEpi64(k x86.Mmask8, a x86.M512i) int64 {
	panic("not implemented")
}


// M512ReduceMulEpi64: Reduce the packed 64-bit integers in 'a' by
// multiplication. Returns the product of all elements in 'a'. 
//
//		prod[63:0] := 1
//		FOR j := 0 to 7
//			i := j*64
//			prod[63:0] := prod[63:0] * a[i+63:i]
//		ENDFOR
//		RETURN prod[63:0]
//
// Instruction: '...'. Intrinsic: '_mm512_reduce_mul_epi64'.
// Requires KNCNI.
func M512ReduceMulEpi64(a x86.M512i) int64 {
	panic("not implemented")
}


// M512MaskReduceMulPd: Reduce the packed double-precision (64-bit)
// floating-point elements in 'a' by multiplication using mask 'k'. Returns the
// product of all active elements in 'a'. 
//
//		prod[63:0] := 1
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				prod[63:0] := prod[63:0] * a[i+63:i]
//			FI
//		ENDFOR
//		RETURN prod[63:0]
//
// Instruction: '...'. Intrinsic: '_mm512_mask_reduce_mul_pd'.
// Requires KNCNI.
func M512MaskReduceMulPd(k x86.Mmask8, a x86.M512d) float64 {
	panic("not implemented")
}


// M512ReduceMulPd: Reduce the packed double-precision (64-bit) floating-point
// elements in 'a' by multiplication. Returns the product of all elements in
// 'a'. 
//
//		prod[63:0] := 1
//		FOR j := 0 to 7
//			i := j*64
//			prod[63:0] := prod[63:0] * a[i+63:i]
//		ENDFOR
//		RETURN prod[63:0]
//
// Instruction: '...'. Intrinsic: '_mm512_reduce_mul_pd'.
// Requires KNCNI.
func M512ReduceMulPd(a x86.M512d) float64 {
	panic("not implemented")
}


// M512MaskReduceMulPs: Reduce the packed single-precision (32-bit)
// floating-point elements in 'a' by multiplication using mask 'k'. Returns the
// product of all active elements in 'a'. 
//
//		prod[31:0] := 1
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				prod[31:0] := prod[31:0] * a[i+31:i]
//			FI
//		ENDFOR
//		RETURN prod[31:0]
//
// Instruction: '...'. Intrinsic: '_mm512_mask_reduce_mul_ps'.
// Requires KNCNI.
func M512MaskReduceMulPs(k x86.Mmask16, a x86.M512) float32 {
	panic("not implemented")
}


// M512ReduceMulPs: Reduce the packed single-precision (32-bit) floating-point
// elements in 'a' by multiplication. Returns the product of all elements in
// 'a'. 
//
//		prod[31:0] := 1
//		FOR j := 0 to 15
//			i := j*32
//			prod[31:0] := prod[31:0] * a[i+31:i]
//		ENDFOR
//		RETURN prod[31:0]
//
// Instruction: '...'. Intrinsic: '_mm512_reduce_mul_ps'.
// Requires KNCNI.
func M512ReduceMulPs(a x86.M512) float32 {
	panic("not implemented")
}


// M512MaskReduceOrEpi32: Reduce the packed 32-bit integers in 'a' by bitwise
// OR using mask 'k'. Returns the bitwise OR of all active elements in 'a'. 
//
//		reduced[31:0] := 0
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				reduced[31:0] := reduced[31:0] OR a[i+31:i]
//			FI
//		ENDFOR
//		RETURN reduced[31:0]
//
// Instruction: '...'. Intrinsic: '_mm512_mask_reduce_or_epi32'.
// Requires KNCNI.
func M512MaskReduceOrEpi32(k x86.Mmask16, a x86.M512i) int {
	panic("not implemented")
}


// M512ReduceOrEpi32: Reduce the packed 32-bit integers in 'a' by bitwise OR.
// Returns the bitwise OR of all elements in 'a'. 
//
//		reduced[31:0] := 0
//		FOR j := 0 to 15
//			i := j*32
//			reduced[31:0] := reduced[31:0] OR a[i+31:i]
//		ENDFOR
//		RETURN reduced[31:0]
//
// Instruction: '...'. Intrinsic: '_mm512_reduce_or_epi32'.
// Requires KNCNI.
func M512ReduceOrEpi32(a x86.M512i) int {
	panic("not implemented")
}


// M512MaskReduceOrEpi64: Reduce the packed 64-bit integers in 'a' by bitwise
// OR using mask 'k'. Returns the bitwise OR of all active elements in 'a'. 
//
//		reduced[63:0] := 0
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				reduced[63:0] := reduced[63:0] OR a[i+63:i]
//			FI
//		ENDFOR
//		RETURN reduced[63:0]
//
// Instruction: '...'. Intrinsic: '_mm512_mask_reduce_or_epi64'.
// Requires KNCNI.
func M512MaskReduceOrEpi64(k x86.Mmask8, a x86.M512i) int64 {
	panic("not implemented")
}


// M512ReduceOrEpi64: Reduce the packed 64-bit integers in 'a' by bitwise OR.
// Returns the bitwise OR of all elements in 'a'. 
//
//		reduced[63:0] := 0
//		FOR j := 0 to 7
//			i := j*64
//			reduced[63:0] := reduced[63:0] OR a[i+63:i]
//		ENDFOR
//		RETURN reduced[63:0]
//
// Instruction: '...'. Intrinsic: '_mm512_reduce_or_epi64'.
// Requires KNCNI.
func M512ReduceOrEpi64(a x86.M512i) int64 {
	panic("not implemented")
}


// M512MaskRoundPs: Round the packed single-precision (32-bit) floating-point
// elements in 'a' to the nearest integer value using 'expadj' and in the
// direction of 'rounding', and store the results as packed single-precision
// floating-point elements in 'dst' using writemask 'k' (elements are copied
// from 'src' when the corresponding mask bit is not set).
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := ROUND(a[i+31:i])
//				CASE expadj OF
//				_MM_EXPADJ_NONE: dst[i+31:i] = dst[i+31:i] * 2**0
//				_MM_EXPADJ_4:	dst[i+31:i] = dst[i+31:i] * 2**4
//				_MM_EXPADJ_5:	dst[i+31:i] = dst[i+31:i] * 2**5
//				_MM_EXPADJ_8:	dst[i+31:i] = dst[i+31:i] * 2**8
//				_MM_EXPADJ_16:   dst[i+31:i] = dst[i+31:i] * 2**16
//				_MM_EXPADJ_24:   dst[i+31:i] = dst[i+31:i] * 2**24
//				_MM_EXPADJ_31:   dst[i+31:i] = dst[i+31:i] * 2**31
//				_MM_EXPADJ_32:   dst[i+31:i] = dst[i+31:i] * 2**32
//				ESAC
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VROUNDPS'. Intrinsic: '_mm512_mask_round_ps'.
// Requires KNCNI.
func M512MaskRoundPs(src x86.M512, k x86.Mmask16, a x86.M512, rounding int, expadj MMEXPADJENUM) (dst x86.M512) {
	panic("not implemented")
}


// M512RoundPs: Round the packed single-precision (32-bit) floating-point
// elements in 'a' to the nearest integer value using 'expadj' and in the
// direction of 'rounding', and store the results as packed single-precision
// floating-point elements in 'dst'.
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := ROUND(a[i+31:i])
//			CASE expadj OF
//			_MM_EXPADJ_NONE: dst[i+31:i] = dst[i+31:i] * 2**0
//			_MM_EXPADJ_4:	dst[i+31:i] = dst[i+31:i] * 2**4
//			_MM_EXPADJ_5:	dst[i+31:i] = dst[i+31:i] * 2**5
//			_MM_EXPADJ_8:	dst[i+31:i] = dst[i+31:i] * 2**8
//			_MM_EXPADJ_16:   dst[i+31:i] = dst[i+31:i] * 2**16
//			_MM_EXPADJ_24:   dst[i+31:i] = dst[i+31:i] * 2**24
//			_MM_EXPADJ_31:   dst[i+31:i] = dst[i+31:i] * 2**31
//			_MM_EXPADJ_32:   dst[i+31:i] = dst[i+31:i] * 2**32
//			ESAC
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VROUNDPS'. Intrinsic: '_mm512_round_ps'.
// Requires KNCNI.
func M512RoundPs(a x86.M512, rounding int, expadj MMEXPADJENUM) (dst x86.M512) {
	panic("not implemented")
}


// M512MaskRoundfxpntAdjustPd: Performs element-by-element rounding of packed
// double-precision (64-bit) floating-point elements in 'a' using 'expadj' and
// in the direction of 'rounding' and stores results in 'dst' using writemask
// 'k' (elements are copied from 'src' when the corresponding mask bit is not
// set).
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := ROUND(a[i+63:i])
//				CASE expadj OF
//				_MM_EXPADJ_NONE: dst[i+31:i] = dst[i+31:i] * 2**0
//				_MM_EXPADJ_4:	dst[i+31:i] = dst[i+31:i] * 2**4
//				_MM_EXPADJ_5:	dst[i+31:i] = dst[i+31:i] * 2**5
//				_MM_EXPADJ_8:	dst[i+31:i] = dst[i+31:i] * 2**8
//				_MM_EXPADJ_16:   dst[i+31:i] = dst[i+31:i] * 2**16
//				_MM_EXPADJ_24:   dst[i+31:i] = dst[i+31:i] * 2**24
//				_MM_EXPADJ_31:   dst[i+31:i] = dst[i+31:i] * 2**31
//				_MM_EXPADJ_32:   dst[i+31:i] = dst[i+31:i] * 2**32
//				ESAC
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VRNDFXPNTPD'. Intrinsic: '_mm512_mask_roundfxpnt_adjust_pd'.
// Requires KNCNI.
func M512MaskRoundfxpntAdjustPd(src x86.M512d, k x86.Mmask8, a x86.M512d, rounding int, expadj MMEXPADJENUM) (dst x86.M512d) {
	panic("not implemented")
}


// M512RoundfxpntAdjustPd: Performs element-by-element rounding of packed
// double-precision (64-bit) floating-point elements in 'a' using 'expadj' and
// in the direction of 'rounding' and stores results in 'dst'.
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 7
//			i := j*64
//			dst[i+63:i] := ROUND(a[i+63:i])
//			CASE expadj OF
//			_MM_EXPADJ_NONE: dst[i+31:i] = dst[i+31:i] * 2**0
//			_MM_EXPADJ_4:	dst[i+31:i] = dst[i+31:i] * 2**4
//			_MM_EXPADJ_5:	dst[i+31:i] = dst[i+31:i] * 2**5
//			_MM_EXPADJ_8:	dst[i+31:i] = dst[i+31:i] * 2**8
//			_MM_EXPADJ_16:   dst[i+31:i] = dst[i+31:i] * 2**16
//			_MM_EXPADJ_24:   dst[i+31:i] = dst[i+31:i] * 2**24
//			_MM_EXPADJ_31:   dst[i+31:i] = dst[i+31:i] * 2**31
//			_MM_EXPADJ_32:   dst[i+31:i] = dst[i+31:i] * 2**32
//			ESAC
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VRNDFXPNTPD'. Intrinsic: '_mm512_roundfxpnt_adjust_pd'.
// Requires KNCNI.
func M512RoundfxpntAdjustPd(a x86.M512d, rounding int, expadj MMEXPADJENUM) (dst x86.M512d) {
	panic("not implemented")
}


// M512MaskRoundfxpntAdjustPs: Performs element-by-element rounding of packed
// single-precision (32-bit) floating-point elements in 'a' using 'expadj' and
// in the direction of 'rounding' and stores results in 'dst' using writemask
// 'k' (elements are copied from 'src' when the corresponding mask bit is not
// set).
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := ROUND(a[i+31:i])
//				CASE expadj OF
//				_MM_EXPADJ_NONE: dst[i+31:i] = dst[i+31:i] * 2**0
//				_MM_EXPADJ_4:	dst[i+31:i] = dst[i+31:i] * 2**4
//				_MM_EXPADJ_5:	dst[i+31:i] = dst[i+31:i] * 2**5
//				_MM_EXPADJ_8:	dst[i+31:i] = dst[i+31:i] * 2**8
//				_MM_EXPADJ_16:   dst[i+31:i] = dst[i+31:i] * 2**16
//				_MM_EXPADJ_24:   dst[i+31:i] = dst[i+31:i] * 2**24
//				_MM_EXPADJ_31:   dst[i+31:i] = dst[i+31:i] * 2**31
//				_MM_EXPADJ_32:   dst[i+31:i] = dst[i+31:i] * 2**32
//				ESAC
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VRNDFXPNTPS'. Intrinsic: '_mm512_mask_roundfxpnt_adjust_ps'.
// Requires KNCNI.
func M512MaskRoundfxpntAdjustPs(src x86.M512, k x86.Mmask16, a x86.M512, rounding int, expadj MMEXPADJENUM) (dst x86.M512) {
	panic("not implemented")
}


// M512RoundfxpntAdjustPs: Performs element-by-element rounding of packed
// single-precision (32-bit) floating-point elements in 'a' using 'expadj' and
// in the direction of 'rounding' and stores results in 'dst'.
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := ROUND(a[i+31:i])
//			CASE expadj OF
//			_MM_EXPADJ_NONE: dst[i+31:i] = dst[i+31:i] * 2**0
//			_MM_EXPADJ_4:	dst[i+31:i] = dst[i+31:i] * 2**4
//			_MM_EXPADJ_5:	dst[i+31:i] = dst[i+31:i] * 2**5
//			_MM_EXPADJ_8:	dst[i+31:i] = dst[i+31:i] * 2**8
//			_MM_EXPADJ_16:   dst[i+31:i] = dst[i+31:i] * 2**16
//			_MM_EXPADJ_24:   dst[i+31:i] = dst[i+31:i] * 2**24
//			_MM_EXPADJ_31:   dst[i+31:i] = dst[i+31:i] * 2**31
//			_MM_EXPADJ_32:   dst[i+31:i] = dst[i+31:i] * 2**32
//			ESAC
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VRNDFXPNTPS'. Intrinsic: '_mm512_roundfxpnt_adjust_ps'.
// Requires KNCNI.
func M512RoundfxpntAdjustPs(a x86.M512, rounding int, expadj MMEXPADJENUM) (dst x86.M512) {
	panic("not implemented")
}


// M512MaskRsqrt23Ps: Calculates the reciprocal square root of packed
// single-precision (32-bit) floating-point elements in 'a' to 23 bits of
// accuracy and stores the result in 'dst' using writemask 'k' (elements are
// copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := Sqrt(1.0 / a[i+31:i])
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VRSQRT23PS'. Intrinsic: '_mm512_mask_rsqrt23_ps'.
// Requires KNCNI.
func M512MaskRsqrt23Ps(src x86.M512, k x86.Mmask16, a x86.M512) (dst x86.M512) {
	panic("not implemented")
}


// M512Rsqrt23Ps: Calculates the reciprocal square root of packed
// single-precision (32-bit) floating-point elements in 'a' to 23 bits of
// accuracy and stores the result in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := Sqrt(1.0 / a[i+31:i])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VRSQRT23PS'. Intrinsic: '_mm512_rsqrt23_ps'.
// Requires KNCNI.
func M512Rsqrt23Ps(a x86.M512) (dst x86.M512) {
	panic("not implemented")
}


// M512MaskSbbEpi32: Performs element-by-element three-input subtraction of
// packed 32-bit integer elements of 'v3' as well as the corresponding bit from
// 'k2' from 'v2'. The borrowed value from the subtraction difference for the
// nth element is written to the nth bit of 'borrow' (borrow flag). Results are
// stored in 'dst' using writemask 'k1' (elements are copied from 'v2' when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k1[j]
//				dst[i+31:i] := v2[i+31:i] - v3[i+31:i] - k2[j]
//				borrow[j] := Borrow(v2[i+31:i] - v3[i+31:i] - k2[j])
//			ELSE
//				dst[i+31:i] := v2[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSBBD'. Intrinsic: '_mm512_mask_sbb_epi32'.
// Requires KNCNI.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func M512MaskSbbEpi32(v2 x86.M512i, k1 x86.Mmask16, k2 x86.Mmask16, v3 x86.M512i, borrow *x86.Mmask16) (dst x86.M512i) {
	panic("not implemented")
}


// M512SbbEpi32: Performs element-by-element three-input subtraction of packed
// 32-bit integer elements of 'v3' as well as the corresponding bit from 'k'
// from 'v2'. The borrowed value from the subtraction difference for the nth
// element is written to the nth bit of 'borrow' (borrow flag). Results are
// stored in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := v2[i+31:i] - v3[i+31:i] - k[j]
//			borrow[j] := Borrow(v2[i+31:i] - v3[i+31:i] - k[j])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSBBD'. Intrinsic: '_mm512_sbb_epi32'.
// Requires KNCNI.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func M512SbbEpi32(v2 x86.M512i, k x86.Mmask16, v3 x86.M512i, borrow *x86.Mmask16) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskSbbrEpi32: Performs element-by-element three-input subtraction of
// packed 32-bit integer elements of 'v2' as well as the corresponding bit from
// 'k2' from 'v3'. The borrowed value from the subtraction difference for the
// nth element is written to the nth bit of 'borrow' (borrow flag). Results are
// stored in 'dst' using writemask 'k1' (elements are copied from 'v2' when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k1[j]
//				dst[i+31:i] := v3[i+31:i] - v2[i+31:i] - k2[j]
//				borrow[j] := Borrow(v2[i+31:i] - v3[i+31:i] - k[j])
//			ELSE
//				dst[i+31:i] := v2[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSBBRD'. Intrinsic: '_mm512_mask_sbbr_epi32'.
// Requires KNCNI.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func M512MaskSbbrEpi32(v2 x86.M512i, k1 x86.Mmask16, k2 x86.Mmask16, v3 x86.M512i, borrow *x86.Mmask16) (dst x86.M512i) {
	panic("not implemented")
}


// M512SbbrEpi32: Performs element-by-element three-input subtraction of packed
// 32-bit integer elements of 'v2' as well as the corresponding bit from 'k'
// from 'v3'. The borrowed value from the subtraction difference for the nth
// element is written to the nth bit of 'borrow' (borrow flag). Results are
// stored in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := v3[i+31:i] - v2[i+31:i] - k[j]
//			borrow[j] := Borrow(v2[i+31:i] - v3[i+31:i] - k[j])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSBBRD'. Intrinsic: '_mm512_sbbr_epi32'.
// Requires KNCNI.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func M512SbbrEpi32(v2 x86.M512i, k x86.Mmask16, v3 x86.M512i, borrow *x86.Mmask16) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskScalePs: Scales each single-precision (32-bit) floating-point
// element in 'a' by multiplying it by 2**exponent, where the exponenet is the
// corresponding 32-bit integer element in 'b', storing results in 'dst' using
// writemask 'k' (elements are copied from 'src' when the corresponding mask
// bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := a[i+31:i] * Pow(2, b[i+31:i])
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VSCALEPS'. Intrinsic: '_mm512_mask_scale_ps'.
// Requires KNCNI.
func M512MaskScalePs(src x86.M512, k x86.Mmask16, a x86.M512, b x86.M512i) (dst x86.M512) {
	panic("not implemented")
}


// M512ScalePs: Scales each single-precision (32-bit) floating-point element in
// 'a' by multiplying it by 2**exponent, where the exponent is the
// corresponding 32-bit integer element in 'b', storing results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := a[i+31:i] * Pow(2, b[i+31:i])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VSCALEPS'. Intrinsic: '_mm512_scale_ps'.
// Requires KNCNI.
func M512ScalePs(a x86.M512, b x86.M512i) (dst x86.M512) {
	panic("not implemented")
}


// M512MaskScaleRoundPs: Scales each single-precision (32-bit) floating-point
// element in 'a' by multiplying it by 2**exp, where the exp is the
// corresponding 32-bit integer element in 'b', storing results in 'dst' using
// writemask 'k' (elements are copied from 'src' when the corresponding mask
// bit is not set). Results are rounded using constant 'rounding'.
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := a[i+31:i] * Pow(2, b[i+31:i])
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VSCALEPS'. Intrinsic: '_mm512_mask_scale_round_ps'.
// Requires KNCNI.
func M512MaskScaleRoundPs(src x86.M512, k x86.Mmask16, a x86.M512, b x86.M512i, rounding int) (dst x86.M512) {
	panic("not implemented")
}


// M512ScaleRoundPs: Scales each single-precision (32-bit) floating-point
// element in 'a' by multiplying it by 2**exponent, where the exponenet is the
// corresponding 32-bit integer element in 'b', storing results in 'dst'.
// Intermediate elements are rounded using 'rounding'.
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := a[i+31:i] * Pow(2, b[i+31:i])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: '...'. Intrinsic: '_mm512_scale_round_ps'.
// Requires KNCNI.
func M512ScaleRoundPs(a x86.M512, b x86.M512i, rounding int) (dst x86.M512) {
	panic("not implemented")
}


// M512MaskShuffleEpi32: Shuffle 32-bit integers in 'a' within 128-bit lanes
// using the control in 'imm8', and store the results in 'dst' using writemask
// 'k' (elements are copied from 'src' when the corresponding mask bit is not
// set). 
//
//		SELECT4(src, control){
//			CASE(control[1:0])
//			0:	tmp[31:0] := src[31:0]
//			1:	tmp[31:0] := src[63:32]
//			2:	tmp[31:0] := src[95:64]
//			3:	tmp[31:0] := src[127:96]
//			ESAC
//			RETURN tmp[31:0]
//		}
//		
//		tmp_dst[31:0] := SELECT4(a[127:0], imm8[1:0])
//		tmp_dst[63:32] := SELECT4(a[127:0], imm8[3:2])
//		tmp_dst[95:64] := SELECT4(a[127:0], imm8[5:4])
//		tmp_dst[127:96] := SELECT4(a[127:0], imm8[7:6])
//		tmp_dst[159:128] := SELECT4(a[255:128], imm8[1:0])
//		tmp_dst[191:160] := SELECT4(a[255:128], imm8[3:2])
//		tmp_dst[223:192] := SELECT4(a[255:128], imm8[5:4])
//		tmp_dst[255:224] := SELECT4(a[255:128], imm8[7:6])
//		tmp_dst[287:256] := SELECT4(a[383:256], imm8[1:0])
//		tmp_dst[319:288] := SELECT4(a[383:256], imm8[3:2])
//		tmp_dst[351:320] := SELECT4(a[383:256], imm8[5:4])
//		tmp_dst[383:352] := SELECT4(a[383:256], imm8[7:6])
//		tmp_dst[415:384] := SELECT4(a[511:384], imm8[1:0])
//		tmp_dst[447:416] := SELECT4(a[511:384], imm8[3:2])
//		tmp_dst[479:448] := SELECT4(a[511:384], imm8[5:4])
//		tmp_dst[511:480] := SELECT4(a[511:384], imm8[7:6])
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := tmp_dst[i+31:i]
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSHUFD'. Intrinsic: '_mm512_mask_shuffle_epi32'.
// Requires KNCNI.
//
// FIXME: Requires compiler support (has immediate)
func M512MaskShuffleEpi32(src x86.M512i, k x86.Mmask16, a x86.M512i, imm8 byte) (dst x86.M512i) {
	panic("not implemented")
}


// M512ShuffleEpi32: Shuffle 32-bit integers in 'a' within 128-bit lanes using
// the control in 'imm8', and store the results in 'dst'. 
//
//		SELECT4(src, control){
//			CASE(control[1:0])
//			0:	tmp[31:0] := src[31:0]
//			1:	tmp[31:0] := src[63:32]
//			2:	tmp[31:0] := src[95:64]
//			3:	tmp[31:0] := src[127:96]
//			ESAC
//			RETURN tmp[31:0]
//		}
//		
//		dst[31:0] := SELECT4(a[127:0], imm8[1:0])
//		dst[63:32] := SELECT4(a[127:0], imm8[3:2])
//		dst[95:64] := SELECT4(a[127:0], imm8[5:4])
//		dst[127:96] := SELECT4(a[127:0], imm8[7:6])
//		dst[159:128] := SELECT4(a[255:128], imm8[1:0])
//		dst[191:160] := SELECT4(a[255:128], imm8[3:2])
//		dst[223:192] := SELECT4(a[255:128], imm8[5:4])
//		dst[255:224] := SELECT4(a[255:128], imm8[7:6])
//		dst[287:256] := SELECT4(a[383:256], imm8[1:0])
//		dst[319:288] := SELECT4(a[383:256], imm8[3:2])
//		dst[351:320] := SELECT4(a[383:256], imm8[5:4])
//		dst[383:352] := SELECT4(a[383:256], imm8[7:6])
//		dst[415:384] := SELECT4(a[511:384], imm8[1:0])
//		dst[447:416] := SELECT4(a[511:384], imm8[3:2])
//		dst[479:448] := SELECT4(a[511:384], imm8[5:4])
//		dst[511:480] := SELECT4(a[511:384], imm8[7:6])
//		dst[MAX:512] := 0
//
// Instruction: 'VPSHUFD'. Intrinsic: '_mm512_shuffle_epi32'.
// Requires KNCNI.
//
// FIXME: Requires compiler support (has immediate)
func M512ShuffleEpi32(a x86.M512i, imm8 byte) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskSlliEpi32: Shift packed 32-bit integers in 'a' left by 'imm8' while
// shifting in zeros, and store the results in 'dst' using writemask 'k'
// (elements are copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				IF imm8[7:0] > 31
//					dst[i+31:i] := 0
//				ELSE
//					dst[i+31:i] := ZeroExtend(a[i+31:i] << imm8[7:0])
//				FI
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSLLD'. Intrinsic: '_mm512_mask_slli_epi32'.
// Requires KNCNI.
//
// FIXME: Requires compiler support (has immediate)
func M512MaskSlliEpi32(src x86.M512i, k x86.Mmask16, a x86.M512i, imm8 byte) (dst x86.M512i) {
	panic("not implemented")
}


// M512SlliEpi32: Shift packed 32-bit integers in 'a' left by 'imm8' while
// shifting in zeros, and store the results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF imm8[7:0] > 31
//				dst[i+31:i] := 0
//			ELSE
//				dst[i+31:i] := ZeroExtend(a[i+31:i] << imm8[7:0])
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSLLD'. Intrinsic: '_mm512_slli_epi32'.
// Requires KNCNI.
//
// FIXME: Requires compiler support (has immediate)
func M512SlliEpi32(a x86.M512i, imm8 byte) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskSllvEpi32: Shift packed 32-bit integers in 'a' left by the amount
// specified by the corresponding element in 'count' while shifting in zeros,
// and store the results in 'dst' using writemask 'k' (elements are copied from
// 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := ZeroExtend(a[i+31:i] << count[i+31:i])
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI	
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSLLVD'. Intrinsic: '_mm512_mask_sllv_epi32'.
// Requires KNCNI.
func M512MaskSllvEpi32(src x86.M512i, k x86.Mmask16, a x86.M512i, count x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512SllvEpi32: Shift packed 32-bit integers in 'a' left by the amount
// specified by the corresponding element in 'count' while shifting in zeros,
// and store the results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := ZeroExtend(a[i+31:i] << count[i+31:i])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSLLVD'. Intrinsic: '_mm512_sllv_epi32'.
// Requires KNCNI.
func M512SllvEpi32(a x86.M512i, count x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// Spflt32: Set performance monitoring filtering mask to 32-bit unsigned
// integer 'r1'. 
//
//		SetPerfMonMask(r1[31:0])
//
// Instruction: 'SPFLT'. Intrinsic: '_mm_spflt_32'.
// Requires KNCNI.
func Spflt32(r1 uint32)  {
	panic("not implemented")
}


// Spflt64: Set performance monitoring filtering mask to 64-bit unsigned
// integer 'r1'. 
//
//		SetPerfMonMask(r1[63:0])
//
// Instruction: 'SPFLT'. Intrinsic: '_mm_spflt_64'.
// Requires KNCNI.
func Spflt64(r1 uint64)  {
	panic("not implemented")
}


// M512MaskSraiEpi32: Shift packed 32-bit integers in 'a' right by 'imm8' while
// shifting in sign bits, and store the results in 'dst' using writemask 'k'
// (elements are copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				IF imm8[7:0] > 31
//					dst[i+31:i] := SignBit
//				ELSE
//					dst[i+31:i] := SignExtend(a[i+31:i] >> imm8[7:0])
//				FI
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSRAD'. Intrinsic: '_mm512_mask_srai_epi32'.
// Requires KNCNI.
//
// FIXME: Requires compiler support (has immediate)
func M512MaskSraiEpi32(src x86.M512i, k x86.Mmask16, a x86.M512i, imm8 byte) (dst x86.M512i) {
	panic("not implemented")
}


// M512SraiEpi32: Shift packed 32-bit integers in 'a' right by 'imm8' while
// shifting in sign bits, and store the results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF imm8[7:0] > 31
//				dst[i+31:i] := SignBit
//			ELSE
//				dst[i+31:i] := SignExtend(a[i+31:i] >> imm8[7:0])
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSRAD'. Intrinsic: '_mm512_srai_epi32'.
// Requires KNCNI.
//
// FIXME: Requires compiler support (has immediate)
func M512SraiEpi32(a x86.M512i, imm8 byte) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskSravEpi32: Shift packed 32-bit integers in 'a' right by the amount
// specified by the corresponding element in 'count' while shifting in sign
// bits, and store the results in 'dst' using writemask 'k' (elements are
// copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := SignExtend(a[i+31:i] >> count[i+31:i])
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI	
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSRAVD'. Intrinsic: '_mm512_mask_srav_epi32'.
// Requires KNCNI.
func M512MaskSravEpi32(src x86.M512i, k x86.Mmask16, a x86.M512i, count x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512SravEpi32: Shift packed 32-bit integers in 'a' right by the amount
// specified by the corresponding element in 'count' while shifting in sign
// bits, and store the results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := SignExtend(a[i+31:i] >> count[i+31:i])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSRAVD'. Intrinsic: '_mm512_srav_epi32'.
// Requires KNCNI.
func M512SravEpi32(a x86.M512i, count x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskSrliEpi32: Shift packed 32-bit integers in 'a' right by 'imm8' while
// shifting in zeros, and store the results in 'dst' using writemask 'k'
// (elements are copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				IF imm8[7:0] > 31
//					dst[i+31:i] := 0
//				ELSE
//					dst[i+31:i] := ZeroExtend(a[i+31:i] >> imm8[7:0])
//				FI
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSRLD'. Intrinsic: '_mm512_mask_srli_epi32'.
// Requires KNCNI.
//
// FIXME: Requires compiler support (has immediate)
func M512MaskSrliEpi32(src x86.M512i, k x86.Mmask16, a x86.M512i, imm8 byte) (dst x86.M512i) {
	panic("not implemented")
}


// M512SrliEpi32: Shift packed 32-bit integers in 'a' right by 'imm8' while
// shifting in zeros, and store the results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF imm8[7:0] > 31
//				dst[i+31:i] := 0
//			ELSE
//				dst[i+31:i] := ZeroExtend(a[i+31:i] >> imm8[7:0])
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSRLD'. Intrinsic: '_mm512_srli_epi32'.
// Requires KNCNI.
//
// FIXME: Requires compiler support (has immediate)
func M512SrliEpi32(a x86.M512i, imm8 byte) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskSrlvEpi32: Shift packed 32-bit integers in 'a' right by the amount
// specified by the corresponding element in 'count' while shifting in zeros,
// and store the results in 'dst' using writemask 'k' (elements are copied from
// 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := ZeroExtend(a[i+31:i] >> count[i+31:i])
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI	
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSRLVD'. Intrinsic: '_mm512_mask_srlv_epi32'.
// Requires KNCNI.
func M512MaskSrlvEpi32(src x86.M512i, k x86.Mmask16, a x86.M512i, count x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512SrlvEpi32: Shift packed 32-bit integers in 'a' right by the amount
// specified by the corresponding element in 'count' while shifting in zeros,
// and store the results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := ZeroExtend(a[i+31:i] >> count[i+31:i])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSRLVD'. Intrinsic: '_mm512_srlv_epi32'.
// Requires KNCNI.
func M512SrlvEpi32(a x86.M512i, count x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// Skipped: _mm512_mask_store_epi32. Contains pointer parameter.


// Skipped: _mm512_store_epi32. Contains pointer parameter.


// Skipped: _mm512_mask_store_epi64. Contains pointer parameter.


// Skipped: _mm512_store_epi64. Contains pointer parameter.


// Skipped: _mm512_mask_store_pd. Contains pointer parameter.


// Skipped: _mm512_store_pd. Contains pointer parameter.


// Skipped: _mm512_mask_store_ps. Contains pointer parameter.


// Skipped: _mm512_store_ps. Contains pointer parameter.


// Skipped: _mm512_store_si512. Contains pointer parameter.


// Skipped: _mm512_storenr_pd. Contains pointer parameter.


// Skipped: _mm512_storenr_ps. Contains pointer parameter.


// Skipped: _mm512_storenrngo_pd. Contains pointer parameter.


// Skipped: _mm512_storenrngo_ps. Contains pointer parameter.


// M512MaskSubEpi32: Subtract packed 32-bit integers in 'b' from packed 32-bit
// integers in 'a', and store the results in 'dst' using writemask 'k'
// (elements are copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := a[i+31:i] - b[i+31:i]
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI	
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSUBD'. Intrinsic: '_mm512_mask_sub_epi32'.
// Requires KNCNI.
func M512MaskSubEpi32(src x86.M512i, k x86.Mmask16, a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512SubEpi32: Subtract packed 32-bit integers in 'b' from packed 32-bit
// integers in 'a', and store the results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := a[i+31:i] - b[i+31:i]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSUBD'. Intrinsic: '_mm512_sub_epi32'.
// Requires KNCNI.
func M512SubEpi32(a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskSubPd: Subtract packed double-precision (64-bit) floating-point
// elements in 'b' from packed double-precision (64-bit) floating-point
// elements in 'a', and store the results in 'dst' using writemask 'k'
// (elements are copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := a[i+63:i] - b[i+63:i]
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VSUBPD'. Intrinsic: '_mm512_mask_sub_pd'.
// Requires KNCNI.
func M512MaskSubPd(src x86.M512d, k x86.Mmask8, a x86.M512d, b x86.M512d) (dst x86.M512d) {
	panic("not implemented")
}


// M512SubPd: Subtract packed double-precision (64-bit) floating-point elements
// in 'b' from packed double-precision (64-bit) floating-point elements in 'a',
// and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*64
//			dst[i+63:i] := a[i+63:i] - b[i+63:i]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VSUBPD'. Intrinsic: '_mm512_sub_pd'.
// Requires KNCNI.
func M512SubPd(a x86.M512d, b x86.M512d) (dst x86.M512d) {
	panic("not implemented")
}


// M512MaskSubPs: Subtract packed single-precision (32-bit) floating-point
// elements in 'b' from packed single-precision (32-bit) floating-point
// elements in 'a', and store the results in 'dst' using writemask 'k'
// (elements are copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := a[i+31:i] - b[i+31:i]
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI	
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VSUBPS'. Intrinsic: '_mm512_mask_sub_ps'.
// Requires KNCNI.
func M512MaskSubPs(src x86.M512, k x86.Mmask16, a x86.M512, b x86.M512) (dst x86.M512) {
	panic("not implemented")
}


// M512SubPs: Subtract packed single-precision (32-bit) floating-point elements
// in 'b' from packed single-precision (32-bit) floating-point elements in 'a',
// and store the results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := a[i+31:i] - b[i+31:i]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VSUBPS'. Intrinsic: '_mm512_sub_ps'.
// Requires KNCNI.
func M512SubPs(a x86.M512, b x86.M512) (dst x86.M512) {
	panic("not implemented")
}


// M512MaskSubRoundPd: Subtract packed double-precision (64-bit) floating-point
// elements in 'b' from packed double-precision (64-bit) floating-point
// elements in 'a', and store the results in 'dst' using writemask 'k'
// (elements are copied from 'src' when the corresponding mask bit is not set).
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := a[i+63:i] - b[i+63:i]
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VSUBPD'. Intrinsic: '_mm512_mask_sub_round_pd'.
// Requires KNCNI.
func M512MaskSubRoundPd(src x86.M512d, k x86.Mmask8, a x86.M512d, b x86.M512d, rounding int) (dst x86.M512d) {
	panic("not implemented")
}


// M512SubRoundPd: Subtract packed double-precision (64-bit) floating-point
// elements in 'b' from packed double-precision (64-bit) floating-point
// elements in 'a', and store the results in 'dst'.
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 7
//			i := j*64
//			dst[i+63:i] := a[i+63:i] - b[i+63:i]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VSUBPD'. Intrinsic: '_mm512_sub_round_pd'.
// Requires KNCNI.
func M512SubRoundPd(a x86.M512d, b x86.M512d, rounding int) (dst x86.M512d) {
	panic("not implemented")
}


// M512MaskSubRoundPs: Subtract packed single-precision (32-bit) floating-point
// elements in 'b' from packed single-precision (32-bit) floating-point
// elements in 'a', and store the results in 'dst' using writemask 'k'
// (elements are copied from 'src' when the corresponding mask bit is not set).
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := a[i+31:i] - b[i+31:i]
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI	
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VSUBPS'. Intrinsic: '_mm512_mask_sub_round_ps'.
// Requires KNCNI.
func M512MaskSubRoundPs(src x86.M512, k x86.Mmask16, a x86.M512, b x86.M512, rounding int) (dst x86.M512) {
	panic("not implemented")
}


// M512SubRoundPs: Subtract packed single-precision (32-bit) floating-point
// elements in 'b' from packed single-precision (32-bit) floating-point
// elements in 'a', and store the results in 'dst'.
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := a[i+31:i] - b[i+31:i]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VSUBPS'. Intrinsic: '_mm512_sub_round_ps'.
// Requires KNCNI.
func M512SubRoundPs(a x86.M512, b x86.M512, rounding int) (dst x86.M512) {
	panic("not implemented")
}


// M512MaskSubrEpi32: Performs element-by-element subtraction of packed 32-bit
// integer elements in 'v2' from 'v3' storing the results in 'dst' using
// writemask 'k' (elements are copied from 'src' when the corresponding mask
// bit is not set) 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := v3[i+31:i] - v2[i+31:i]
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSUBRD'. Intrinsic: '_mm512_mask_subr_epi32'.
// Requires KNCNI.
func M512MaskSubrEpi32(src x86.M512i, k x86.Mmask16, v2 x86.M512i, v3 x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512SubrEpi32: Performs element-by-element subtraction of packed 32-bit
// integer elements in 'v2' from 'v3' storing the results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := v3[i+31:i] - v2[i+31:i]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSUBRD'. Intrinsic: '_mm512_subr_epi32'.
// Requires KNCNI.
func M512SubrEpi32(v2 x86.M512i, v3 x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskSubrPd: Performs element-by-element subtraction of packed
// double-precision (64-bit) floating-point elements in 'v2' from 'v3' storing
// the results in 'dst' using writemask 'k' (elements are copied from 'src'
// when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := v3[i+63:i] - v2[i+63:i]
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VSUBRPD'. Intrinsic: '_mm512_mask_subr_pd'.
// Requires KNCNI.
func M512MaskSubrPd(src x86.M512d, k x86.Mmask8, v2 x86.M512d, v3 x86.M512d) (dst x86.M512d) {
	panic("not implemented")
}


// M512SubrPd: Performs element-by-element subtraction of packed
// double-precision (64-bit) floating-point elements in 'v2' from 'v3' storing
// the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*64
//			dst[i+63:i] := v3[i+63:i] - v2[i+63:i]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VSUBRPD'. Intrinsic: '_mm512_subr_pd'.
// Requires KNCNI.
func M512SubrPd(v2 x86.M512d, v3 x86.M512d) (dst x86.M512d) {
	panic("not implemented")
}


// M512MaskSubrPs: Performs element-by-element subtraction of packed
// single-precision (32-bit) floating-point elements in 'v2' from 'v3' storing
// the results in 'dst' using writemask 'k' (elements are copied from 'src'
// when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := v3[i+31:i] - v2[i+31:i]
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VSUBRPS'. Intrinsic: '_mm512_mask_subr_ps'.
// Requires KNCNI.
func M512MaskSubrPs(src x86.M512, k x86.Mmask16, v2 x86.M512, v3 x86.M512) (dst x86.M512) {
	panic("not implemented")
}


// M512SubrPs: Performs element-by-element subtraction of packed
// single-precision (32-bit) floating-point elements in 'v2' from 'v3' storing
// the results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := v3[i+31:i] - v2[i+31:i]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VSUBRPS'. Intrinsic: '_mm512_subr_ps'.
// Requires KNCNI.
func M512SubrPs(v2 x86.M512, v3 x86.M512) (dst x86.M512) {
	panic("not implemented")
}


// M512MaskSubrRoundPd: Performs element-by-element subtraction of packed
// double-precision (64-bit) floating-point elements in 'v2' from 'v3' storing
// the results in 'dst' using writemask 'k' (elements are copied from 'src'
// when the corresponding mask bit is not set).
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := v3[i+63:i] - v2[i+63:i]
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VSUBRPD'. Intrinsic: '_mm512_mask_subr_round_pd'.
// Requires KNCNI.
func M512MaskSubrRoundPd(src x86.M512d, k x86.Mmask8, v2 x86.M512d, v3 x86.M512d, rounding int) (dst x86.M512d) {
	panic("not implemented")
}


// M512SubrRoundPd: Performs element-by-element subtraction of packed
// double-precision (64-bit) floating-point elements in 'v2' from 'v3' storing
// the results in 'dst'.
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 7
//			i := j*64
//			dst[i+63:i] := v3[i+63:i] - v2[i+63:i]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VSUBRPD'. Intrinsic: '_mm512_subr_round_pd'.
// Requires KNCNI.
func M512SubrRoundPd(v2 x86.M512d, v3 x86.M512d, rounding int) (dst x86.M512d) {
	panic("not implemented")
}


// M512MaskSubrRoundPs: Performs element-by-element subtraction of packed
// single-precision (32-bit) floating-point elements in 'v2' from 'v3' storing
// the results in 'dst' using writemask 'k' (elements are copied from 'src'
// when the corresponding mask bit is not set).
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := v3[i+31:i] - v2[i+31:i]
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VSUBRPS'. Intrinsic: '_mm512_mask_subr_round_ps'.
// Requires KNCNI.
func M512MaskSubrRoundPs(src x86.M512, k x86.Mmask16, v2 x86.M512, v3 x86.M512, rounding int) (dst x86.M512) {
	panic("not implemented")
}


// M512SubrRoundPs: Performs element-by-element subtraction of packed
// single-precision (32-bit) floating-point elements in 'v2' from 'v3' storing
// the results in 'dst'.
// 	Rounding is done according to the 'rounding' parameter, which can be one
// of:
//     (_MM_FROUND_TO_NEAREST_INT |_MM_FROUND_NO_EXC) // round to nearest, and suppress exceptions
//     (_MM_FROUND_TO_NEG_INF |_MM_FROUND_NO_EXC)     // round down, and suppress exceptions
//     (_MM_FROUND_TO_POS_INF |_MM_FROUND_NO_EXC)     // round up, and suppress exceptions
//     (_MM_FROUND_TO_ZERO |_MM_FROUND_NO_EXC)        // truncate, and suppress exceptions
//     _MM_FROUND_CUR_DIRECTION // use MXCSR.RC; see _MM_SET_ROUNDING_MODE 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := v3[i+31:i] - v2[i+31:i])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VSUBRPS'. Intrinsic: '_mm512_subr_round_ps'.
// Requires KNCNI.
func M512SubrRoundPs(v2 x86.M512, v3 x86.M512, rounding int) (dst x86.M512) {
	panic("not implemented")
}


// M512MaskSubrsetbEpi32: Performs element-by-element subtraction of packed
// 32-bit integer elements in 'v2' from 'v3', storing the results in 'dst' and
// 'v2'. The borrowed value from the subtraction difference for the nth element
// is written to the nth bit of 'borrow' (borrow flag). Results are written
// using writemask 'k' (elements are copied from 'k' to 'k_old' when the
// corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				diff := v3[i+31:i] - v2[i+31:i]
//				borrow[j] := Borrow(v3[i+31:i] - v2[i+31:i])
//				dst[i+31:i] := diff
//				v2[i+31:i] := diff
//			ELSE
//				borrow[j] := k_old[j]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSUBRSETBD'. Intrinsic: '_mm512_mask_subrsetb_epi32'.
// Requires KNCNI.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func M512MaskSubrsetbEpi32(v2 x86.M512i, k x86.Mmask16, k_old x86.Mmask16, v3 x86.M512i, borrow *x86.Mmask16) (dst x86.M512i) {
	panic("not implemented")
}


// M512SubrsetbEpi32: Performs element-by-element subtraction of packed 32-bit
// integer elements in 'v2' from 'v3', storing the results in 'dst' and 'v2'.
// The borrowed value from the subtraction difference for the nth element is
// written to the nth bit of 'borrow' (borrow flag). 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := v3[i+31:i] - v2[i+31:i]
//			borrow[j] := Borrow(v3[i+31:i] - v2[i+31:i])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSUBRSETBD'. Intrinsic: '_mm512_subrsetb_epi32'.
// Requires KNCNI.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func M512SubrsetbEpi32(v2 x86.M512i, v3 x86.M512i, borrow *x86.Mmask16) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskSubsetbEpi32: Performs element-by-element subtraction of packed
// 32-bit integer elements in 'v3' from 'v2', storing the results in 'dst' and
// the nth borrow bit in the nth position of 'borrow' (borrow flag). Results
// are stored using writemask 'k' (elements are copied from 'v2' and 'k_old'
// when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := v2[i+31:i] - v3[i+31:i]
//				borrow[j] := Borrow(v2[i+31:i] - v3[i+31:i])
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//				borrow[j] := k_old[j]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSUBSETBD'. Intrinsic: '_mm512_mask_subsetb_epi32'.
// Requires KNCNI.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func M512MaskSubsetbEpi32(v2 x86.M512i, k x86.Mmask16, k_old x86.Mmask16, v3 x86.M512i, borrow *x86.Mmask16) (dst x86.M512i) {
	panic("not implemented")
}


// M512SubsetbEpi32: Performs element-by-element subtraction of packed 32-bit
// integer elements in 'v3' from 'v2', storing the results in 'dst' and the nth
// borrow bit in the nth position of 'borrow' (borrow flag). 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := v2[i+31:i] - v3[i+31:i]
//			borrow[j] := Borrow(v2[i+31:i] - v3[i+31:i])
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPSUBSETBD'. Intrinsic: '_mm512_subsetb_epi32'.
// Requires KNCNI.
//
// FIXME: Will likely need to be reworked (has pointer parameter).
func M512SubsetbEpi32(v2 x86.M512i, v3 x86.M512i, borrow *x86.Mmask16) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskSwizzleEpi32: Performs a swizzle transformation of each of the four
// groups of packed 4x32-bit integer elements in 'v' using swizzle parameter
// 's', storing the results in 'dst' using writemask 'k' (elements are copied
// from 'src' when the corresponding mask bit is not set). 
//
//		CASE s OF
//		_MM_SWIZ_REG_NONE:
//			dst[511:0] := v[511:0]
//		_MM_SWIZ_REG_DCBA:
//			dst[511:0] := v[511:0]
//		_MM_SWIZ_REG_CDAB:
//			FOR j := 0 to 7
//				i := j*64
//				IF k[j*2]
//					dst[i+31:i]	:= v[i+63:i+32]
//				ELSE
//					dst[i+31:i]	:= src[i+31:i]
//				FI
//				IF k[j*2+1]
//					dst[i+63:i+32] := v[i+31:i]
//				ELSE
//					dst[i+63:i+32] := src[i+63:i+32]
//				FI
//			ENDFOR
//		_MM_SWIZ_REG_BADC:
//			FOR j := 0 to 3
//				i := j*128
//				IF k[j*4]
//					dst[i+31:i]	 := v[i+95:i+64]
//				ELSE
//					dst[i+31:i]	 := src[i+31:i]
//				FI
//				IF k[j*4+1]
//					dst[i+63:i+32]  := v[i+127:i+96]
//				ELSE
//					dst[i+63:i+32]  := src[i+63:i+32]
//				FI
//				IF k[j*4+2]
//					dst[i+95:i+64]  := v[i+31:i]
//				ELSE
//					dst[i+95:i+64]  := src[i+95:i+64]
//				FI
//				IF k[j*4+3]
//					dst[i+127:i+96] := v[i+63:i+32]
//				ELSE
//					dst[i+127:i+96] := src[i+127:i+96]
//				FI
//			ENDFOR
//		_MM_SWIZ_REG_AAAA:
//			FOR j := 0 to 3
//				i := j*128
//				IF k[j*4]
//					dst[i+31:i]	 := v[i+31:i]
//				ELSE
//					dst[i+31:i]	 := src[i+31:i]
//				FI
//				IF k[j*4+1]
//					dst[i+63:i+32]  := v[i+31:i]
//				ELSE
//					dst[i+63:i+32]  := src[i+63:i+32]
//				FI
//				IF k[j*4+2]
//					dst[i+95:i+64]  := v[i+31:i]
//				ELSE
//					dst[i+95:i+64]  := src[i+95:i+64]
//				FI
//				IF k[j*4+3]
//					dst[i+127:i+96] := v[i+31:i]
//				ELSE
//					dst[i+127:i+96] := src[i+127:i+96]
//				FI
//			ENDFOR
//		_MM_SWIZ_REG_BBBB:
//			FOR j := 0 to 3
//				i := j*128
//				IF k[j*4]
//					dst[i+31:i]	 := v[i+63:i+32]
//				ELSE
//					dst[i+31:i]	 := src[i+31:i]
//				FI
//				IF k[j*4+1]
//					dst[i+63:i+32]  := v[i+63:i+32]
//				ELSE
//					dst[i+63:i+32]  := src[i+63:i+32]
//				FI
//				IF k[j*4+2]
//					dst[i+95:i+64]  := v[i+63:i+32]
//				ELSE
//					dst[i+95:i+64]  := src[i+95:i+64]
//				FI
//				IF k[j*4+3]
//					dst[i+127:i+96] := v[i+63:i+32]
//				ELSE
//					dst[i+127:i+96] := src[i+127:i+96]
//				FI
//			ENDFOR
//		_MM_SWIZ_REG_CCCC:
//			FOR j := 0 to 3
//				i := j*128
//				IF k[j*4]
//					dst[i+31:i]	 := v[i+95:i+64]
//				ELSE
//					dst[i+31:i]	 := src[i+31:i]
//				FI
//				IF k[j*4+1]
//					dst[i+63:i+32]  := v[i+95:i+64]
//				ELSE
//					dst[i+63:i+32]  := src[i+63:i+32]
//				FI
//				IF k[j*4+2]
//					dst[i+95:i+64]  := v[i+95:i+64]
//				ELSE
//					dst[i+95:i+64]  := src[i+95:i+64]
//				FI
//				IF k[j*4+3]
//					dst[i+127:i+96] := v[i+95:i+64]
//				ELSE
//					dst[i+127:i+96] := src[i+127:i+96]
//				FI
//			ENDFOR
//		_MM_SWIZ_REG_DDDD:
//			FOR j := 0 to 3
//				i := j*128
//				IF k[j*4]
//					dst[i+31:i]	 := v[i+127:i+96]
//				ELSE
//					dst[i+31:i]	 := src[i+31:i]
//				FI
//				IF k[j*4+1]
//					dst[i+63:i+32]  := v[i+127:i+96]
//				ELSE
//					dst[i+63:i+32]  := src[i+63:i+32]
//				FI
//				IF k[j*4+2]
//					dst[i+95:i+64]  := v[i+127:i+96]
//				ELSE
//					dst[i+95:i+64]  := src[i+95:i+64]
//				FI
//				IF k[j*4+3]
//					dst[i+127:i+96] := v[i+127:i+96]
//				ELSE
//					dst[i+127:i+96] := src[i+127:i+96]
//				FI
//			ENDFOR
//		_MM_SWIZ_REG_DACB:
//			FOR j := 0 to 3
//				i := j*128
//				IF k[j*4]
//					dst[i+31:i]	 := v[i+63:i+32]
//				ELSE
//					dst[i+31:i]	 := src[i+31:i]
//				FI
//				IF k[j*4+1]
//					dst[i+63:i+32]  := v[i+95:i+64]
//				ELSE
//					dst[i+63:i+32]  := src[i+63:i+32]
//				FI
//				IF k[j*4+2]
//					dst[i+95:i+64]  := v[i+31:i]
//				ELSE
//					dst[i+95:i+64]  := src[i+95:i+64]
//				FI
//				IF k[j*4+3]
//					dst[i+127:i+96] := v[i+127:i+96]
//				ELSE
//					dst[i+127:i+96] := src[i+127:i+96]
//				FI
//			ENDFOR
//		ESAC
//		dst[MAX:512] := 0
//
// Instruction: '...'. Intrinsic: '_mm512_mask_swizzle_epi32'.
// Requires KNCNI.
func M512MaskSwizzleEpi32(src x86.M512i, k x86.Mmask16, v x86.M512i, s MMSWIZZLEENUM) (dst x86.M512i) {
	panic("not implemented")
}


// M512SwizzleEpi32: Performs a swizzle transformation of each of the four
// groups of packed 4x 32-bit integer elements in 'v' using swizzle parameter
// 's', storing the results in 'dst'. 
//
//		CASE s OF
//		_MM_SWIZ_REG_NONE:
//			dst[511:0] := v[511:0]
//		_MM_SWIZ_REG_DCBA:
//			dst[511:0] := v[511:0]
//		_MM_SWIZ_REG_CDAB:
//			FOR j := 0 to 7
//				i := j*64
//				dst[i+31:i]    := v[i+63:i+32]
//				dst[i+63:i+32] := v[i+31:i]
//			ENDFOR
//		_MM_SWIZ_REG_BADC:
//			FOR j := 0 to 3
//				i := j*128
//				dst[i+31:i]	    := v[i+95:i+64]
//				dst[i+63:i+32]  := v[i+127:i+96]
//				dst[i+95:i+64]  := v[i+31:i]
//				dst[i+127:i+96] := v[i+63:i+32]
//			ENDFOR
//		_MM_SWIZ_REG_AAAA:
//			FOR j := 0 to 3
//				i := j*128
//				dst[i+31:i]	    := v[i+31:i]
//				dst[i+63:i+32]  := v[i+31:i]
//				dst[i+95:i+64]  := v[i+31:i]
//				dst[i+127:i+96] := v[i+31:i]
//			ENDFOR
//		_MM_SWIZ_REG_BBBB:
//			FOR j := 0 to 3
//				i := j*128
//				dst[i+31:i]	    := v[i+63:i+32]
//				dst[i+63:i+32]  := v[i+63:i+32]
//				dst[i+95:i+64]  := v[i+63:i+32]
//				dst[i+127:i+96] := v[i+63:i+32]
//			ENDFOR
//		_MM_SWIZ_REG_CCCC:
//			FOR j := 0 to 3
//				i := j*128
//				dst[i+31:i]	    := v[i+95:i+64]
//				dst[i+63:i+32]  := v[i+95:i+64]
//				dst[i+95:i+64]  := v[i+95:i+64]
//				dst[i+127:i+96] := v[i+95:i+64]
//			ENDFOR
//		_MM_SWIZ_REG_DDDD:
//			FOR j := 0 to 3
//				i := j*128
//				dst[i+31:i]	    := v[i+127:i+96]
//				dst[i+63:i+32]  := v[i+127:i+96]
//				dst[i+95:i+64]  := v[i+127:i+96]
//				dst[i+127:i+96] := v[i+127:i+96]
//			ENDFOR
//		_MM_SWIZ_REG_DACB:
//			FOR j := 0 to 3
//				i := j*128
//				dst[i+31:i]	    := v[i+63:i+32]
//				dst[i+63:i+32]  := v[i+95:i+64]
//				dst[i+95:i+64]  := v[i+31:i]
//				dst[i+127:i+96] := v[i+127:i+96]
//			ENDFOR
//		ESAC
//		dst[MAX:512] := 0
//
// Instruction: '...'. Intrinsic: '_mm512_swizzle_epi32'.
// Requires KNCNI.
func M512SwizzleEpi32(v x86.M512i, s MMSWIZZLEENUM) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskSwizzleEpi64: Performs a swizzle transformation of each of the four
// groups of packed 4x64-bit integer elements in 'v' using swizzle parameter
// 's', storing the results in 'dst' using writemask 'k' (elements are copied
// from 'src' when the corresponding mask bit is not set). 
//
//		CASE s OF
//		_MM_SWIZ_REG_NONE:
//			dst[511:0] := v[511:0]
//		_MM_SWIZ_REG_DCBA:
//			dst[511:0] := v[511:0]
//		_MM_SWIZ_REG_CDAB:
//			FOR j := 0 to 3
//				i := j*64
//				IF k[j*2]
//					dst[i+63:i]	 := v[i+127:i+64]
//				ELSE
//					dst[i+63:i]	 := src[i+63:i]
//				FI
//				IF k[j*2+1]
//					dst[i+127:i+64] := v[i+63:i]
//				ELSE
//					dst[i+127:i+64] := src[i+127:i+64]
//				FI
//			ENDFOR
//		_MM_SWIZ_REG_BADC:
//			FOR j := 0 to 1
//				i := j*256
//				IF k[j*4]
//					dst[i+63:i]	  := v[i+191:i+128]
//				ELSE
//					dst[i+63:i]	  := src[i+63:i]
//				FI
//				IF k[j*4+1]
//					dst[i+127:i+64]  := v[i+255:i+192]
//				ELSE
//					dst[i+127:i+64]  := src[i+127:i+64]
//				FI
//				IF k[j*4+2]
//					dst[i+191:i+128] := v[i+63:i]
//				ELSE
//					dst[i+191:i+128] := src[i+191:i+128]
//				FI
//				IF k[j*4+3]
//					dst[i+255:i+192] := v[i+127:i+64]
//				ELSE
//					dst[i+255:i+192] := src[i+255:i+192]
//				FI
//			ENDFOR
//		_MM_SWIZ_REG_AAAA:
//			FOR j := 0 to 1
//				i := j*256
//				IF k[j*4]
//					dst[i+63:i]	  := v[i+63:i]
//				ELSE
//					dst[i+63:i]	  := src[i+63:i]
//				FI
//				IF k[j*4+1]
//					dst[i+127:i+64]  := v[i+63:i]
//				ELSE
//					dst[i+127:i+64]  := src[i+127:i+64]
//				FI
//				IF k[j*4+2]
//					dst[i+191:i+128] := v[i+63:i]
//				ELSE
//					dst[i+191:i+128] := src[i+191:i+128]
//				FI
//				IF k[j*4+3]
//					dst[i+255:i+192] := v[i+63:i]
//				ELSE
//					dst[i+255:i+192] := src[i+255:i+192]
//				FI
//			ENDFOR
//		_MM_SWIZ_REG_BBBB:
//			FOR j := 0 to 1
//				i := j*256
//				IF k[j*4]
//					dst[i+63:i]	  := v[i+127:i+63]
//				ELSE
//					dst[i+63:i]	  := src[i+63:i]
//				FI
//				IF k[j*4+1]
//					dst[i+127:i+64]  := v[i+127:i+63]
//				ELSE
//					dst[i+127:i+64]  := src[i+127:i+64]
//				FI
//				IF k[j*4+2]
//					dst[i+191:i+128] := v[i+127:i+63]
//				ELSE
//					dst[i+191:i+128] := src[i+191:i+128]
//				FI
//				IF k[j*4+3]
//					dst[i+255:i+192] := v[i+127:i+63]
//				ELSE
//					dst[i+255:i+192] := src[i+255:i+192]
//				FI
//			ENDFOR
//		_MM_SWIZ_REG_CCCC:
//			FOR j := 0 to 1
//				i := j*256
//				IF k[j*4]
//					dst[i+63:i]	  := v[i+191:i+128]
//				ELSE
//					dst[i+63:i]	  := src[i+63:i]
//				FI
//				IF k[j*4+1]
//					dst[i+127:i+64]  := v[i+191:i+128]
//				ELSE
//					dst[i+127:i+64]  := src[i+127:i+64]
//				FI
//				IF k[j*4+2]
//					dst[i+191:i+128] := v[i+191:i+128]
//				ELSE
//					dst[i+191:i+128] := src[i+191:i+128]
//				FI
//				IF k[j*4+3]
//					dst[i+255:i+192] := v[i+191:i+128]
//				ELSE
//					dst[i+255:i+192] := src[i+255:i+192]
//				FI
//			ENDFOR
//		_MM_SWIZ_REG_DDDD:
//			FOR j := 0 to 1
//				i := j*256
//				IF k[j*4]
//					dst[i+63:i]	  := v[i+255:i+192]
//				ELSE
//					dst[i+63:i]	  := src[i+63:i]
//				FI
//				IF k[j*4+1]
//					dst[i+127:i+64]  := v[i+255:i+192]
//				ELSE
//					dst[i+127:i+64]  := src[i+127:i+64]
//				FI
//				IF k[j*4+2]
//					dst[i+191:i+128] := v[i+255:i+192]
//				ELSE
//					dst[i+191:i+128] := src[i+191:i+128]
//				FI
//				IF k[j*4+3]
//					dst[i+255:i+192] := v[i+255:i+192]
//				ELSE
//					dst[i+255:i+192] := src[i+255:i+192]
//				FI
//			ENDFOR
//		_MM_SWIZ_REG_DACB:
//			FOR j := 0 to 1
//				i := j*256
//				IF k[j*4]
//					dst[i+63:i]	  := v[i+127:i+64]
//				ELSE
//					dst[i+63:i]	  := src[i+63:i]
//				FI
//				IF k[j*4+1]
//					dst[i+127:i+64]  := v[i+191:i+128]
//				ELSE
//					dst[i+127:i+64]  := src[i+127:i+64]
//				FI
//				IF k[j*4+2]
//					dst[i+191:i+128] := v[i+63:i]
//				ELSE
//					dst[i+191:i+128] := src[i+191:i+128]
//				FI
//				IF k[j*4+3]
//					dst[i+255:i+192] := v[i+255:i+192]
//				ELSE
//					dst[i+255:i+192] := src[i+255:i+192]
//				FI
//			ENDFOR
//		ESAC
//		dst[MAX:512] := 0
//
// Instruction: '...'. Intrinsic: '_mm512_mask_swizzle_epi64'.
// Requires KNCNI.
func M512MaskSwizzleEpi64(src x86.M512i, k x86.Mmask8, v x86.M512i, s MMSWIZZLEENUM) (dst x86.M512i) {
	panic("not implemented")
}


// M512SwizzleEpi64: Performs a swizzle transformation of each of the two
// groups of packed 4x64-bit integer elements in 'v' using swizzle parameter
// 's', storing the results in 'dst'. 
//
//		CASE s OF
//		_MM_SWIZ_REG_NONE:
//			dst[511:0] := v[511:0]
//		_MM_SWIZ_REG_DCBA:
//			dst[511:0] := v[511:0]
//		_MM_SWIZ_REG_CDAB:
//			FOR j := 0 to 3
//				i := j*64
//				dst[i+63:i]	    := v[i+127:i+64]
//				dst[i+127:i+64] := v[i+63:i]
//			ENDFOR
//		_MM_SWIZ_REG_BADC:
//			FOR j := 0 to 1
//				i := j*256
//				dst[i+63:i]	     := v[i+191:i+128]
//				dst[i+127:i+64]  := v[i+255:i+192]
//				dst[i+191:i+128] := v[i+63:i]
//				dst[i+255:i+192] := v[i+127:i+64]
//			ENDFOR
//		_MM_SWIZ_REG_AAAA:
//			FOR j := 0 to 1
//				i := j*256
//				dst[i+63:i]	     := v[i+63:i]
//				dst[i+127:i+64]  := v[i+63:i]
//				dst[i+191:i+128] := v[i+63:i]
//				dst[i+255:i+192] := v[i+63:i]
//			ENDFOR
//		_MM_SWIZ_REG_BBBB:
//			FOR j := 0 to 1
//				i := j*256
//				dst[i+63:i]	     := v[i+127:i+63]
//				dst[i+127:i+64]  := v[i+127:i+63]
//				dst[i+191:i+128] := v[i+127:i+63]
//				dst[i+255:i+192] := v[i+127:i+63]
//			ENDFOR
//		_MM_SWIZ_REG_CCCC:
//			FOR j := 0 to 1
//				i := j*256
//				dst[i+63:i]	     := v[i+191:i+128]
//				dst[i+127:i+64]  := v[i+191:i+128]
//				dst[i+191:i+128] := v[i+191:i+128]
//				dst[i+255:i+192] := v[i+191:i+128]
//			ENDFOR
//		_MM_SWIZ_REG_DDDD:
//			FOR j := 0 to 1
//				i := j*256
//				dst[i+63:i]	     := v[i+255:i+192]
//				dst[i+127:i+64]  := v[i+255:i+192]
//				dst[i+191:i+128] := v[i+255:i+192]
//				dst[i+255:i+192] := v[i+255:i+192]
//			ENDFOR
//		_MM_SWIZ_REG_DACB:
//			FOR j := 0 to 1
//				i := j*256
//				dst[i+63:i]	     := v[i+127:i+64]
//				dst[i+127:i+64]  := v[i+191:i+128]
//				dst[i+191:i+128] := v[i+63:i]
//				dst[i+255:i+192] := v[i+255:i+192]
//			ENDFOR
//		ESAC
//		dst[MAX:512] := 0
//
// Instruction: '...'. Intrinsic: '_mm512_swizzle_epi64'.
// Requires KNCNI.
func M512SwizzleEpi64(v x86.M512i, s MMSWIZZLEENUM) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskSwizzlePd: Performs a swizzle transformation of each of the two
// groups of packed 4x double-precision (64-bit) floating-point elements in 'v'
// using swizzle parameter 's', storing the results in 'dst' using writemask
// 'k' (elements are copied from 'src' when the corresponding mask bit is not
// set). 
//
//		CASE s OF
//		_MM_SWIZ_REG_NONE:
//			dst[511:0] := v[511:0]
//		_MM_SWIZ_REG_DCBA:
//			dst[511:0] := v[511:0]
//		_MM_SWIZ_REG_CDAB:
//			FOR j := 0 to 3
//				i := j*64
//				IF k[j*2]
//					dst[i+63:i]	 := v[i+127:i+64]
//				ELSE
//					dst[i+63:i]	 := src[i+63:i]
//				FI
//				IF k[j*2+1]
//					dst[i+127:i+64] := v[i+63:i]
//				ELSE
//					dst[i+127:i+64] := src[i+127:i+64]
//				FI
//			ENDFOR
//		_MM_SWIZ_REG_BADC:
//			FOR j := 0 to 1
//				i := j*256
//				IF k[j*4]
//					dst[i+63:i]	  := v[i+191:i+128]
//				ELSE
//					dst[i+63:i]	  := src[i+63:i]
//				FI
//				IF k[j*4+1]
//					dst[i+127:i+64]  := v[i+255:i+192]
//				ELSE
//					dst[i+127:i+64]  := src[i+127:i+64]
//				FI
//				IF k[j*4+2]
//					dst[i+191:i+128] := v[i+63:i]
//				ELSE
//					dst[i+191:i+128] := src[i+191:i+128]
//				FI
//				IF k[j*4+3]
//					dst[i+255:i+192] := v[i+127:i+64]
//				ELSE
//					dst[i+255:i+192] := src[i+255:i+192]
//				FI
//			ENDFOR
//		_MM_SWIZ_REG_AAAA:
//			FOR j := 0 to 1
//				i := j*256
//				IF k[j*4]
//					dst[i+63:i]	  := v[i+63:i]
//				ELSE
//					dst[i+63:i]	  := src[i+63:i]
//				FI
//				IF k[j*4+1]
//					dst[i+127:i+64]  := v[i+63:i]
//				ELSE
//					dst[i+127:i+64]  := src[i+127:i+64]
//				FI
//				IF k[j*4+2]
//					dst[i+191:i+128] := v[i+63:i]
//				ELSE
//					dst[i+191:i+128] := src[i+191:i+128]
//				FI
//				IF k[j*4+3]
//					dst[i+255:i+192] := v[i+63:i]
//				ELSE
//					dst[i+255:i+192] := src[i+255:i+192]
//				FI
//			ENDFOR
//		_MM_SWIZ_REG_BBBB:
//			FOR j := 0 to 1
//				i := j*256
//				IF k[j*4]
//					dst[i+63:i]	  := v[i+127:i+63]
//				ELSE
//					dst[i+63:i]	  := src[i+63:i]
//				FI
//				IF k[j*4+1]
//					dst[i+127:i+64]  := v[i+127:i+63]
//				ELSE
//					dst[i+127:i+64]  := src[i+127:i+64]
//				FI
//				IF k[j*4+2]
//					dst[i+191:i+128] := v[i+127:i+63]
//				ELSE
//					dst[i+191:i+128] := src[i+191:i+128]
//				FI
//				IF k[j*4+3]
//					dst[i+255:i+192] := v[i+127:i+63]
//				ELSE
//					dst[i+255:i+192] := src[i+255:i+192]
//				FI
//			ENDFOR
//		_MM_SWIZ_REG_CCCC:
//			FOR j := 0 to 1
//				i := j*256
//				IF k[j*4]
//					dst[i+63:i]	  := v[i+191:i+128]
//				ELSE
//					dst[i+63:i]	  := src[i+63:i]
//				FI
//				IF k[j*4+1]
//					dst[i+127:i+64]  := v[i+191:i+128]
//				ELSE
//					dst[i+127:i+64]  := src[i+127:i+64]
//				FI
//				IF k[j*4+2]
//					dst[i+191:i+128] := v[i+191:i+128]
//				ELSE
//					dst[i+191:i+128] := src[i+191:i+128]
//				FI
//				IF k[j*4+3]
//					dst[i+255:i+192] := v[i+191:i+128]
//				ELSE
//					dst[i+255:i+192] := src[i+255:i+192]
//				FI
//			ENDFOR
//		_MM_SWIZ_REG_DDDD:
//			FOR j := 0 to 1
//				i := j*256
//				IF k[j*4]
//					dst[i+63:i]	  := v[i+255:i+192]
//				ELSE
//					dst[i+63:i]	  := src[i+63:i]
//				FI
//				IF k[j*4+1]
//					dst[i+127:i+64]  := v[i+255:i+192]
//				ELSE
//					dst[i+127:i+64]  := src[i+127:i+64]
//				FI
//				IF k[j*4+2]
//					dst[i+191:i+128] := v[i+255:i+192]
//				ELSE
//					dst[i+191:i+128] := src[i+191:i+128]
//				FI
//				IF k[j*4+3]
//					dst[i+255:i+192] := v[i+255:i+192]
//				ELSE
//					dst[i+255:i+192] := src[i+255:i+192]
//				FI
//			ENDFOR
//		_MM_SWIZ_REG_DACB:
//			FOR j := 0 to 1
//				i := j*256
//				IF k[j*4]
//					dst[i+63:i]	  := v[i+127:i+64]
//				ELSE
//					dst[i+63:i]	  := src[i+63:i]
//				FI
//				IF k[j*4+1]
//					dst[i+127:i+64]  := v[i+191:i+128]
//				ELSE
//					dst[i+127:i+64]  := src[i+127:i+64]
//				FI
//				IF k[j*4+2]
//					dst[i+191:i+128] := v[i+63:i]
//				ELSE
//					dst[i+191:i+128] := src[i+191:i+128]
//				FI
//				IF k[j*4+3]
//					dst[i+255:i+192] := v[i+255:i+192]
//				ELSE
//					dst[i+255:i+192] := src[i+255:i+192]
//				FI
//			ENDFOR
//		ESAC
//		dst[MAX:512] := 0
//
// Instruction: '...'. Intrinsic: '_mm512_mask_swizzle_pd'.
// Requires KNCNI.
func M512MaskSwizzlePd(src x86.M512d, k x86.Mmask8, v x86.M512d, s MMSWIZZLEENUM) (dst x86.M512d) {
	panic("not implemented")
}


// M512SwizzlePd: Performs a swizzle transformation of each of the two groups
// of packed 4x double-precision (64-bit) floating-point elements in 'v' using
// swizzle parameter 's', storing the results in 'dst'. 
//
//		CASE s OF
//		_MM_SWIZ_REG_NONE:
//			dst[511:0] := v[511:0]
//		_MM_SWIZ_REG_DCBA:
//			dst[511:0] := v[511:0]
//		_MM_SWIZ_REG_CDAB:
//			FOR j := 0 to 3
//				i := j*64
//				dst[i+63:i]     := v[i+127:i+64]
//				dst[i+127:i+64] := v[i+63:i]
//			ENDFOR
//		_MM_SWIZ_REG_BADC:
//			FOR j := 0 to 1
//				i := j*256
//				dst[i+63:i]      := v[i+191:i+128]
//				dst[i+127:i+64]  := v[i+255:i+192]
//				dst[i+191:i+128] := v[i+63:i]
//				dst[i+255:i+192] := v[i+127:i+64]
//			ENDFOR
//		_MM_SWIZ_REG_AAAA:
//			FOR j := 0 to 1
//				i := j*256
//				dst[i+63:i]      := v[i+63:i]
//				dst[i+127:i+64]  := v[i+63:i]
//				dst[i+191:i+128] := v[i+63:i]
//				dst[i+255:i+192] := v[i+63:i]
//			ENDFOR
//		_MM_SWIZ_REG_BBBB:
//			FOR j := 0 to 1
//				i := j*256
//				dst[i+63:i]      := v[i+127:i+63]
//				dst[i+127:i+64]  := v[i+127:i+63]
//				dst[i+191:i+128] := v[i+127:i+63]
//				dst[i+255:i+192] := v[i+127:i+63]
//			ENDFOR
//		_MM_SWIZ_REG_CCCC:
//			FOR j := 0 to 1
//				i := j*256
//				dst[i+63:i]      := v[i+191:i+128]
//				dst[i+127:i+64]  := v[i+191:i+128]
//				dst[i+191:i+128] := v[i+191:i+128]
//				dst[i+255:i+192] := v[i+191:i+128]
//			ENDFOR
//		_MM_SWIZ_REG_DDDD:
//			FOR j := 0 to 1
//				i := j*256
//				dst[i+63:i]	     := v[i+255:i+192]
//				dst[i+127:i+64]  := v[i+255:i+192]
//				dst[i+191:i+128] := v[i+255:i+192]
//				dst[i+255:i+192] := v[i+255:i+192]
//			ENDFOR
//		_MM_SWIZ_REG_DACB:
//			FOR j := 0 to 1
//				i := j*256
//				dst[i+63:i]	     := v[i+127:i+64]
//				dst[i+127:i+64]  := v[i+191:i+128]
//				dst[i+191:i+128] := v[i+63:i]
//				dst[i+255:i+192] := v[i+255:i+192]
//			ENDFOR
//		ESAC
//		dst[MAX:512] := 0
//
// Instruction: '...'. Intrinsic: '_mm512_swizzle_pd'.
// Requires KNCNI.
func M512SwizzlePd(v x86.M512d, s MMSWIZZLEENUM) (dst x86.M512d) {
	panic("not implemented")
}


// M512MaskSwizzlePs: Performs a swizzle transformation of each of the four
// groups of packed 4x single-precision (32-bit) floating-point elements in 'v'
// using swizzle parameter 's', storing the results in 'dst' using writemask
// 'k' (elements are copied from 'src' when the corresponding mask bit is not
// set). 
//
//		CASE s OF
//		_MM_SWIZ_REG_NONE:
//			dst[511:0] := v[511:0]
//		_MM_SWIZ_REG_DCBA:
//			dst[511:0] := v[511:0]
//		_MM_SWIZ_REG_CDAB:
//			FOR j := 0 to 7
//				i := j*64
//				IF k[j*2]
//					dst[i+31:i]	:= v[i+63:i+32]
//				ELSE
//					dst[i+31:i]	:= src[i+31:i]
//				FI
//				IF k[j*2+1]
//					dst[i+63:i+32] := v[i+31:i]
//				ELSE
//					dst[i+63:i+32] := src[i+63:i+32]
//				FI
//			ENDFOR
//		_MM_SWIZ_REG_BADC:
//			FOR j := 0 to 3
//				i := j*128
//				IF k[j*4]
//					dst[i+31:i]	 := v[i+95:i+64]
//				ELSE
//					dst[i+31:i]	 := src[i+31:i]
//				FI
//				IF k[j*4+1]
//					dst[i+63:i+32]  := v[i+127:i+96]
//				ELSE
//					dst[i+63:i+32]  := src[i+63:i+32]
//				FI
//				IF k[j*4+2]
//					dst[i+95:i+64]  := v[i+31:i]
//				ELSE
//					dst[i+95:i+64]  := src[i+95:i+64]
//				FI
//				IF k[j*4+3]
//					dst[i+127:i+96] := v[i+63:i+32]
//				ELSE
//					dst[i+127:i+96] := src[i+127:i+96]
//				FI
//			ENDFOR
//		_MM_SWIZ_REG_AAAA:
//			FOR j := 0 to 3
//				i := j*128
//				IF k[j*4]
//					dst[i+31:i]	 := v[i+31:i]
//				ELSE
//					dst[i+31:i]	 := src[i+31:i]
//				FI
//				IF k[j*4+1]
//					dst[i+63:i+32]  := v[i+31:i]
//				ELSE
//					dst[i+63:i+32]  := src[i+63:i+32]
//				FI
//				IF k[j*4+2]
//					dst[i+95:i+64]  := v[i+31:i]
//				ELSE
//					dst[i+95:i+64]  := src[i+95:i+64]
//				FI
//				IF k[j*4+3]
//					dst[i+127:i+96] := v[i+31:i]
//				ELSE
//					dst[i+127:i+96] := src[i+127:i+96]
//				FI
//			ENDFOR
//		_MM_SWIZ_REG_BBBB:
//			FOR j := 0 to 3
//				i := j*128
//				IF k[j*4]
//					dst[i+31:i]	 := v[i+63:i+32]
//				ELSE
//					dst[i+31:i]	 := src[i+31:i]
//				FI
//				IF k[j*4+1]
//					dst[i+63:i+32]  := v[i+63:i+32]
//				ELSE
//					dst[i+63:i+32]  := src[i+63:i+32]
//				FI
//				IF k[j*4+2]
//					dst[i+95:i+64]  := v[i+63:i+32]
//				ELSE
//					dst[i+95:i+64]  := src[i+95:i+64]
//				FI
//				IF k[j*4+3]
//					dst[i+127:i+96] := v[i+63:i+32]
//				ELSE
//					dst[i+127:i+96] := src[i+127:i+96]
//				FI
//			ENDFOR
//		_MM_SWIZ_REG_CCCC:
//			FOR j := 0 to 3
//				i := j*128
//				IF k[j*4]
//					dst[i+31:i]	 := v[i+95:i+64]
//				ELSE
//					dst[i+31:i]	 := src[i+31:i]
//				FI
//				IF k[j*4+1]
//					dst[i+63:i+32]  := v[i+95:i+64]
//				ELSE
//					dst[i+63:i+32]  := src[i+63:i+32]
//				FI
//				IF k[j*4+2]
//					dst[i+95:i+64]  := v[i+95:i+64]
//				ELSE
//					dst[i+95:i+64]  := src[i+95:i+64]
//				FI
//				IF k[j*4+3]
//					dst[i+127:i+96] := v[i+95:i+64]
//				ELSE
//					dst[i+127:i+96] := src[i+127:i+96]
//				FI
//			ENDFOR
//		_MM_SWIZ_REG_DDDD:
//			FOR j := 0 to 3
//				i := j*128
//				IF k[j*4]
//					dst[i+31:i]	 := v[i+127:i+96]
//				ELSE
//					dst[i+31:i]	 := src[i+31:i]
//				FI
//				IF k[j*4+1]
//					dst[i+63:i+32]  := v[i+127:i+96]
//				ELSE
//					dst[i+63:i+32]  := src[i+63:i+32]
//				FI
//				IF k[j*4+2]
//					dst[i+95:i+64]  := v[i+127:i+96]
//				ELSE
//					dst[i+95:i+64]  := src[i+95:i+64]
//				FI
//				IF k[j*4+3]
//					dst[i+127:i+96] := v[i+127:i+96]
//				ELSE
//					dst[i+127:i+96] := src[i+127:i+96]
//				FI
//			ENDFOR
//		_MM_SWIZ_REG_DACB:
//			FOR j := 0 to 3
//				i := j*128
//				IF k[j*4]
//					dst[i+31:i]	 := v[i+63:i+32]
//				ELSE
//					dst[i+31:i]	 := src[i+31:i]
//				FI
//				IF k[j*4+1]
//					dst[i+63:i+32]  := v[i+95:i+64]
//				ELSE
//					dst[i+63:i+32]  := src[i+63:i+32]
//				FI
//				IF k[j*4+2]
//					dst[i+95:i+64]  := v[i+31:i]
//				ELSE
//					dst[i+95:i+64]  := src[i+95:i+64]
//				FI
//				IF k[j*4+3]
//					dst[i+127:i+96] := v[i+127:i+96]
//				ELSE
//					dst[i+127:i+96] := src[i+127:i+96]
//				FI
//			ENDFOR
//		ESAC
//		dst[MAX:512] := 0
//
// Instruction: '...'. Intrinsic: '_mm512_mask_swizzle_ps'.
// Requires KNCNI.
func M512MaskSwizzlePs(src x86.M512, k x86.Mmask16, v x86.M512, s MMSWIZZLEENUM) (dst x86.M512) {
	panic("not implemented")
}


// M512SwizzlePs: Performs a swizzle transformation of each of the four groups
// of packed 4xsingle-precision (32-bit) floating-point elements in 'v' using
// swizzle parameter 's', storing the results in 'dst'. 
//
//		CASE s OF
//		_MM_SWIZ_REG_NONE:
//			dst[511:0] := v[511:0]
//		_MM_SWIZ_REG_DCBA:
//			dst[511:0] := v[511:0]
//		_MM_SWIZ_REG_CDAB:
//			FOR j := 0 to 7
//				i := j*64
//				dst[i+31:i]    := v[i+63:i+32]
//				dst[i+63:i+32] := v[i+31:i]
//			ENDFOR
//		_MM_SWIZ_REG_BADC:
//			FOR j := 0 to 3
//				i := j*128
//				dst[i+31:i]     := v[i+95:i+64]
//				dst[i+63:i+32]  := v[i+127:i+96]
//				dst[i+95:i+64]  := v[i+31:i]
//				dst[i+127:i+96] := v[i+63:i+32]
//			ENDFOR
//		_MM_SWIZ_REG_AAAA:
//			FOR j := 0 to 3
//				i := j*128
//				dst[i+31:i]     := v[i+31:i]
//				dst[i+63:i+32]  := v[i+31:i]
//				dst[i+95:i+64]  := v[i+31:i]
//				dst[i+127:i+96] := v[i+31:i]
//			ENDFOR
//		_MM_SWIZ_REG_BBBB:
//			FOR j := 0 to 3
//				i := j*128
//				dst[i+31:i]     := v[i+63:i+32]
//				dst[i+63:i+32]  := v[i+63:i+32]
//				dst[i+95:i+64]  := v[i+63:i+32]
//				dst[i+127:i+96] := v[i+63:i+32]
//			ENDFOR
//		_MM_SWIZ_REG_CCCC:
//			FOR j := 0 to 3
//				i := j*128
//				dst[i+31:i]     := v[i+95:i+64]
//				dst[i+63:i+32]  := v[i+95:i+64]
//				dst[i+95:i+64]  := v[i+95:i+64]
//				dst[i+127:i+96] := v[i+95:i+64]
//			ENDFOR
//		_MM_SWIZ_REG_DDDD:
//			FOR j := 0 to 3
//				i := j*128
//				dst[i+31:i]     := v[i+127:i+96]
//				dst[i+63:i+32]  := v[i+127:i+96]
//				dst[i+95:i+64]  := v[i+127:i+96]
//				dst[i+127:i+96] := v[i+127:i+96]
//			ENDFOR
//		_MM_SWIZ_REG_DACB:
//			FOR j := 0 to 3
//				i := j*128
//				dst[i+31:i]     := v[i+63:i+32]
//				dst[i+63:i+32]  := v[i+95:i+64]
//				dst[i+95:i+64]  := v[i+31:i]
//				dst[i+127:i+96] := v[i+127:i+96]
//			ENDFOR
//		ESAC
//		dst[MAX:512] := 0
//
// Instruction: '...'. Intrinsic: '_mm512_swizzle_ps'.
// Requires KNCNI.
func M512SwizzlePs(v x86.M512, s MMSWIZZLEENUM) (dst x86.M512) {
	panic("not implemented")
}


// M512MaskTestEpi32Mask: Compute the bitwise AND of packed 32-bit integers in
// 'a' and 'b', producing intermediate 32-bit values, and set the corresponding
// bit in result mask 'k' (subject to writemask 'k') if the intermediate value
// is non-zero. 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k1[j]
//				k[j] := ((a[i+31:i] AND b[i+31:i]) != 0) ? 1 : 0
//			ELSE
//				k[j] := 0
//			FI
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPTESTMD'. Intrinsic: '_mm512_mask_test_epi32_mask'.
// Requires KNCNI.
func M512MaskTestEpi32Mask(k1 x86.Mmask16, a x86.M512i, b x86.M512i) (dst x86.Mmask16) {
	panic("not implemented")
}


// M512TestEpi32Mask: Compute the bitwise AND of packed 32-bit integers in 'a'
// and 'b', producing intermediate 32-bit values, and set the corresponding bit
// in result mask 'k' if the intermediate value is non-zero. 
//
//		FOR j := 0 to 15
//			i := j*32
//			k[j] := ((a[i+31:i] AND b[i+31:i]) != 0) ? 1 : 0
//		ENDFOR
//		k[MAX:16] := 0
//
// Instruction: 'VPTESTMD'. Intrinsic: '_mm512_test_epi32_mask'.
// Requires KNCNI.
func M512TestEpi32Mask(a x86.M512i, b x86.M512i) (dst x86.Mmask16) {
	panic("not implemented")
}


// Tzcnti32: Counts the number of trailing bits in unsigned 32-bit integer 'x'
// starting at bit 'a' storing the result in 'dst'. 
//
//		count := 0
//		FOR j := a to 31
//			IF NOT(x[j]  1)
//				count := count + 1
//			FI
//		ENDFOR
//		dst := count
//
// Instruction: 'TZCNTI'. Intrinsic: '_mm_tzcnti_32'.
// Requires KNCNI.
func Tzcnti32(a int, x uint32) int {
	panic("not implemented")
}


// Tzcnti64: Counts the number of trailing bits in unsigned 64-bit integer 'x'
// starting at bit 'a' storing the result in 'dst'. 
//
//		count := 0
//		FOR j := a to 63
//			IF NOT(x[j]  1)
//				count := count + 1
//			FI
//		ENDFOR
//		dst := count
//
// Instruction: 'TZCNTI'. Intrinsic: '_mm_tzcnti_64'.
// Requires KNCNI.
func Tzcnti64(a int64, x uint64) int64 {
	panic("not implemented")
}


// M512MaskXorEpi32: Compute the bitwise XOR of packed 32-bit integers in 'a'
// and 'b', and store the results in 'dst' using writemask 'k' (elements are
// copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 15
//			i := j*32
//			IF k[j]
//				dst[i+31:i] := a[i+31:i] XOR b[i+31:i]
//			ELSE
//				dst[i+31:i] := src[i+31:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPXORD'. Intrinsic: '_mm512_mask_xor_epi32'.
// Requires KNCNI.
func M512MaskXorEpi32(src x86.M512i, k x86.Mmask16, a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512XorEpi32: Compute the bitwise XOR of packed 32-bit integers in 'a' and
// 'b', and store the results in 'dst'. 
//
//		FOR j := 0 to 15
//			i := j*32
//			dst[i+31:i] := a[i+31:i] XOR b[i+31:i]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPXORD'. Intrinsic: '_mm512_xor_epi32'.
// Requires KNCNI.
func M512XorEpi32(a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512MaskXorEpi64: Compute the bitwise XOR of packed 64-bit integers in 'a'
// and 'b', and store the results in 'dst' using writemask 'k' (elements are
// copied from 'src' when the corresponding mask bit is not set). 
//
//		FOR j := 0 to 7
//			i := j*64
//			IF k[j]
//				dst[i+63:i] := a[i+63:i] XOR b[i+63:i]
//			ELSE
//				dst[i+63:i] := src[i+63:i]
//			FI
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPXORQ'. Intrinsic: '_mm512_mask_xor_epi64'.
// Requires KNCNI.
func M512MaskXorEpi64(src x86.M512i, k x86.Mmask8, a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512XorEpi64: Compute the bitwise XOR of packed 64-bit integers in 'a' and
// 'b', and store the results in 'dst'. 
//
//		FOR j := 0 to 7
//			i := j*64
//			dst[i+63:i] := a[i+63:i] XOR b[i+63:i]
//		ENDFOR
//		dst[MAX:512] := 0
//
// Instruction: 'VPXORQ'. Intrinsic: '_mm512_xor_epi64'.
// Requires KNCNI.
func M512XorEpi64(a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}


// M512XorSi512: Compute the bitwise XOR of 512 bits (representing integer
// data) in 'a' and 'b', and store the result in 'dst'. 
//
//		dst[511:0] := (a[511:0] XOR b[511:0])
//		dst[MAX:512] := 0
//
// Instruction: 'VPXORD'. Intrinsic: '_mm512_xor_si512'.
// Requires KNCNI.
func M512XorSi512(a x86.M512i, b x86.M512i) (dst x86.M512i) {
	panic("not implemented")
}

