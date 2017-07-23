// THESE PACKAGES ARE FOR DEMONSTRATION PURPOSES ONLY!
//
// THEY DO NOT NOT CONTAIN WORKING INTRINSICS!
//
// See https://github.com/klauspost/intrinsics
package aes

import "github.com/klauspost/intrinsics/x86"

var _ = x86.M64{}  // Make sure we use x86 package


// AesdecSi128: Perform one round of an AES decryption flow on data (state) in
// 'a' using the round key in 'RoundKey', and store the result in 'dst'." 
//
//		state := a
//		a[127:0] := InvShiftRows(a[127:0])
//		a[127:0] := InvSubBytes(a[127:0])
//		a[127:0] := InvMixColumns(a[127:0])
//		dst[127:0] := a[127:0] XOR RoundKey[127:0]
//
// Instruction: 'AESDEC'. Intrinsic: '_mm_aesdec_si128'.
// Requires AES.
func AesdecSi128(a x86.M128i, RoundKey x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// AesdeclastSi128: Perform the last round of an AES decryption flow on data
// (state) in 'a' using the round key in 'RoundKey', and store the result in
// 'dst'." 
//
//		state := a
//		a[127:0] := InvShiftRows(a[127:0])
//		a[127:0] := InvSubBytes(a[127:0])
//		dst[127:0] := a[127:0] XOR RoundKey[127:0]
//
// Instruction: 'AESDECLAST'. Intrinsic: '_mm_aesdeclast_si128'.
// Requires AES.
func AesdeclastSi128(a x86.M128i, RoundKey x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// AesencSi128: Perform one round of an AES encryption flow on data (state) in
// 'a' using the round key in 'RoundKey', and store the result in 'dst'." 
//
//		state := a
//		a[127:0] := ShiftRows(a[127:0])
//		a[127:0] := SubBytes(a[127:0])
//		a[127:0] := MixColumns(a[127:0])
//		dst[127:0] := a[127:0] XOR RoundKey[127:0]
//
// Instruction: 'AESENC'. Intrinsic: '_mm_aesenc_si128'.
// Requires AES.
func AesencSi128(a x86.M128i, RoundKey x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// AesenclastSi128: Perform the last round of an AES encryption flow on data
// (state) in 'a' using the round key in 'RoundKey', and store the result in
// 'dst'." 
//
//		state := a
//		a[127:0] := ShiftRows(a[127:0])
//		a[127:0] := SubBytes(a[127:0])
//		dst[127:0] := a[127:0] XOR RoundKey[127:0]
//
// Instruction: 'AESENCLAST'. Intrinsic: '_mm_aesenclast_si128'.
// Requires AES.
func AesenclastSi128(a x86.M128i, RoundKey x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// AesimcSi128: Perform the InvMixColumns transformation on 'a' and store the
// result in 'dst'. 
//
//		dst[127:0] := InvMixColumns(a[127:0])
//
// Instruction: 'AESIMC'. Intrinsic: '_mm_aesimc_si128'.
// Requires AES.
func AesimcSi128(a x86.M128i) (dst x86.M128i) {
	panic("not implemented")
}


// AeskeygenassistSi128: Assist in expanding the AES cipher key by computing
// steps towards generating a round key for encryption cipher using data from
// 'a' and an 8-bit round constant specified in 'imm8', and store the result in
// 'dst'." 
//
//		X3[31:0] := a[127:96]
//		X2[31:0] := a[95:64]
//		X1[31:0] := a[63:32]
//		X0[31:0] := a[31:0]
//		RCON[31:0] := ZeroExtend(imm8[7:0]);
//		dst[31:0] := SubWord(X1)
//		dst[63:32] := (RotWord(SubWord(X1)) XOR RCON;
//		dst[95:64] := SubWord(X3)
//		dst[127:96] := RotWord(SubWord(X3)) XOR RCON;
//
// Instruction: 'AESKEYGENASSIST'. Intrinsic: '_mm_aeskeygenassist_si128'.
// Requires AES.
//
// FIXME: Requires compiler support (has immediate)
func AeskeygenassistSi128(a x86.M128i, imm8 byte) (dst x86.M128i) {
	panic("not implemented")
}

