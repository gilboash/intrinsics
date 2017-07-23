// +build ignore

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

type Param struct {
	Name  string
	Type  string
	Const bool
}

type Params []Param

type Intrinsic struct {
	OrgName     string
	Name        string
	AsmName     string
	Instruction string
	Description string
	RetType     string
	Params      Params
	CpuID       string
	Operation   string
	Performance Perf
	Package     string
	cParam      *Param
}

type Timing struct {
	Arch       string
	Latency    float64
	Throughput float64
}

type Perf map[string]Timing

//var m64, m128, m256, m512, gen *os.File
//var m64a, m128a, m256a, m512a, gena *os.File
type pack struct {
	goFile  *os.File
	asmFile *os.File
}

const genimport = `import . "github.com/klauspost/intrinsics/arm"`

func main() {
	f, err := os.Open("arm_neon-processed.h")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	parseHeader(f)
	for _, pk := range packages {
		pk.goFile.Close()
		//pk.asmFile.Close()
	}
}

func parseDesc(f io.Reader) error {

}

func NewIntrinsic() *Intrinsic {
	return &Intrinsic{}
}

func parseHeader(r *os.File) {
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanWords)
	for {
		if !s.Scan() {
			return
		}
		if s.Text() == "DEF:" {
			n := 0
			var in Intrinsic
			in.CpuID = "NEON"
			in.Description = "ARM NEON intrinsic"
			var cParam Param
			for {
				s.Scan()
				t := s.Text()
				end := strings.Contains(t, "{")
				switch n {
				case 0:
					in.RetType = fixType(t)
					n++
				case 1:
					in.OrgName = t
					in.Name = fixFuncName(in.OrgName)
					fmt.Print(in.Name)
					n++
				case 2:
					t = strings.TrimPrefix(t, "(")
					t = strings.TrimPrefix(t, ",")
					if strings.Contains(t, "const") {
						cParam.Const = true
						continue
					}
					cParam.Type = fixType(t)
					n++
				case 3:
					if t == "const" {
						cParam.Const = true
						continue
					}
					if strings.HasPrefix(t, "*") {
						cParam.Type = "*" + cParam.Type
						t = strings.Trim(t, "*")
						if t == "" {
							continue
						}
					}
					t = strings.Trim(t, ",")
					n = 2
					if strings.HasSuffix(t, ")") {
						n = 4
					}
					t = strings.Trim(t, ")")
					cParam.Name = t
					in.Params = append(in.Params, cParam)
				case 4:
					break
				}
				//fmt.Print(t, "|")
				if end {
					fmt.Print("\n")
					break
				}
			}
			in.Finish()
		}
	}
}

func (i Intrinsic) getPackage() string {
	pk := strings.ToLower(i.CpuID)
	if pk == "" {
		pk = "misc"
	}
	if pk == "sse4.1" || pk == "sse4.2" {
		pk = "sse4"
	}
	return pk
}

var packages = make(map[string]pack)

func (i Intrinsic) getFiles() pack {
	pk := i.getPackage()
	p, ok := packages[pk]
	if !ok {
		var err error
		err = os.MkdirAll("arm/"+pk, 666)
		if err != nil {
			panic(err)
		}
		p.goFile, err = os.Create("arm/" + pk + "/" + pk + ".go")
		if err != nil {
			panic(err)
		}
		p.goFile.WriteString("// THESE PACKAGES ARE FOR DEMONSTRATION PURPOSES ONLY!\n")
		p.goFile.WriteString("//\n// THEY DO NOT NOT CONTAIN WORKING INTRINSICS!\n")
		p.goFile.WriteString("//\n// See https://github.com/klauspost/intrinsics\n")
		p.goFile.WriteString("package " + pk + "\n\n")
		p.goFile.WriteString(`import "github.com/klauspost/intrinsics/arm"` + "\n\n")

		p.goFile.WriteString(`var _ = arm.Int8x8{}  // Make sure we use arm package` + "\n\n")

		//p.asmFile, err = os.Create("arm/" + pk + "/" + pk + "_amd64.s")
		if err != nil {
			panic(err)
		}
		packages[pk] = p
	}
	return p
}

func fixFuncName(in string) string {
	out := CamelCase(in)
	if strings.HasPrefix(in, "v") {
		out = strings.TrimPrefix(out, "v")
	}
	out = strings.Title(out)
	return out
}

func fixType(r string) string {
	pointer := false
	if strings.Contains(r, "*") {
		pointer = true
	}
	if len(r) == 0 {
		return ""
	}

	r = strings.TrimSuffix(r, "_t")
	exxes := strings.Split(r, "x")
	nelems := 1
	if len(exxes) == 3 {
		nelems, _ = strconv.Atoi(exxes[2])
		r = strings.Join(exxes[:2], "x")
	}
	switch r {
	case "int8x8":
		r = "arm.Int8x8"
	case "int16x4":
		r = "arm.Int16x4"
	case "int32x2":
		r = "arm.Int32x2"
	case "int64x1":
		r = "arm.Int64x1"
	case "float32x2":
		r = "arm.Float32x2"
	case "poly8x8":
		r = "arm.Poly8x8"
	case "poly16x4":
		r = "arm.Poly16x4"
	case "uint8x8":
		r = "arm.Uint8x8"
	case "uint16x4":
		r = "arm.Uint16x4"
	case "uint32x2":
		r = "arm.Uint32x2"
	case "float64x1":
		r = "arm.Float64x1"
	case "uint64x1":
		r = "arm.Uint64x1"
	case "int8x16":
		r = "arm.Int8x16"
	case "int16x8":
		r = "arm.Int16x8"
	case "int32x4":
		r = "arm.Int32x4"
	case "int64x2":
		r = "arm.Int64x2"
	case "float32x4":
		r = "arm.Float32x4"
	case "float64x2":
		r = "arm.Float64x2"
	case "poly8x16":
		r = "arm.Poly8x16"
	case "poly16x8":
		r = "arm.Poly16x8"
	case "poly64x2":
		r = "arm.Poly64x2"
	case "uint8x16":
		r = "arm.Uint8x16"
	case "uint16x8":
		r = "arm.Uint16x8"
	case "uint32x4":
		r = "arm.Uint32x4"
	case "uint64x2":
		r = "arm.Uint64x2"
	case "poly8":
		r = "arm.Poly8"
	case "poly16":
		r = "arm.Poly16"
	case "poly64":
		r = "arm.Poly64"
	case "poly128":
		r = "arm.Poly128"
	}
	if strings.HasPrefix(r, "arm.") {
		if nelems > 1 {
			r = "[" + strconv.Itoa(nelems) + "]" + r
		}
		return r
	}
	r = CamelCase(r)

	switch r {
	case "void":
		if pointer {
			r = "uintptr"
		} else {
			r = ""
		}
	case "char", "charConst":
		r = "byte"
	case "unsignedChar", "poly8T", "uint8T":
		r = "uint8"
	case "unsignedShort", "poly16T", "uint16T":
		r = "uint16"
	case "sizeT", "constInt", "intConst":
		r = "int"
	case "int64Const":
		r = "int"
	case "unsignedInt64":
		r = "uint64"
	case "unsigned", "constUnsignedInt":
		r = "uint"
	case "unsignedInt", "unsignedLong", "uint32T":
		r = "uint32"
	case "unsignedInt32":
		r = "uint32"
	case "voidConst":
		r = "uintptr"
	case "constVoid":
		r = "uintptr"
	case "mem_addr":
		r = "uintptr"
	case "float", "constFloat", "float32T":
		r = "float32"
	case "double", "constDouble", "float64T":
		r = "float64"
	case "int8T":
		r = "int8"
	case "short", "int16T":
		r = "int16"
	case "int32T":
		r = "int32"
	case "int64T":
		r = "int64"
	case "uint64T":
		r = "uint64"
	case "floatConst":
		r = "uintptr"
	case "doubleConst":
		r = "uintptr"
	case "longLong":
		r = "int64"
	case "constMMCMPINTENUM":
		r = "uint8"
	}

	if r == "uintptr" || r == "" {
		return r
	}

	if pointer {
		r = "*" + r
	}
	return r

}

var camelingRegex = regexp.MustCompile("[0-9A-Za-z]+")

func CamelCase(src string) string {
	byteSrc := []byte(src)
	chunks := camelingRegex.FindAll(byteSrc, -1)
	for idx, val := range chunks {
		if idx > 0 {
			chunks[idx] = bytes.Title(val)
		}
	}
	return string(bytes.Join(chunks, nil))
}

func (p Params) HasParam(s string) bool {
	for _, param := range p {
		if param.Name == s {
			return true
		}
	}
	return false
}

var usedNames = make(map[string]struct{})

func (in *Intrinsic) fixup() {
	in.Package = in.getPackage()
	name := in.Name
	next := 2
	for {
		_, ok := usedNames[in.Package+name]
		if !ok {
			break
		}
		fmt.Println("Warning: Name conflict:", in.Package+"/"+name+"(...)")
		name = fmt.Sprintf("%s%d", in.Name, next)
		next++

	}
	usedNames[in.Package+name] = struct{}{}
	in.Name = name
	a := []rune(name)
	a[0] = unicode.ToLower(a[0])

	in.AsmName = string(a)
	for i, param := range in.Params {
		switch param.Name {
		case "type":
			in.Params[i].Name = "typ"
		case "func":
			in.Params[i].Name = "fnc"
		case "imm8":
			in.Params[i].Type = "byte"
		}
		if param.Type == "" {
			in.Params[i].Type = "uintptr"
		}
	}

}

func (in Intrinsic) shouldSkip() bool {
	if in.RetType == "uintptr" {
		return true
	}
	if strings.HasPrefix(in.RetType, "*") {
		return true
	}
	for _, param := range in.Params {
		if param.Type == "uintptr" {
			return true
		}
	}
	return false
}

func (in Intrinsic) shouldRework() bool {
	if strings.HasPrefix(in.RetType, "*") {
		return true
	}
	for _, param := range in.Params {
		if strings.HasPrefix(param.Type, "*") {
			return true
		}
	}
	return false
}

func (in Intrinsic) hasImmediate() bool {
	for _, param := range in.Params {
		if param.Const {
			return true
		}
	}
	return false
}

func (in Intrinsic) Finish() {
	if in.Name == "" {
		return
	}

	in.fixup()

	out := in.getFiles().goFile

	if in.shouldSkip() {
		fmt.Fprintf(out, "\n// Skipped: %s. Contains pointer parameter.\n\n", in.OrgName)
		return
	}
	rework := in.shouldRework()

	//fmt.Println("Finished:", in)
	desc := in.Name + ": " + in.Description
	desc = strings.Replace(desc, "\r\n", "\n", -1)
	desc = strings.Replace(desc, "\r", "\n", -1)
	desc = WrapString(desc, 76)
	desc = strings.Replace(desc, "\n", "\n// ", -1)
	fmt.Fprintln(out, "\n//", desc, "\n//")
	//op := strings.Replace(strings.TrimSpace(in.Operation), "\n", "\n//\t\t", -1)
	//fmt.Fprint(out, "//\t\t", op, "\n//\n")
	fmt.Fprintln(out, "// Intrinsic:", "'"+in.OrgName+"'.")
	if len(in.CpuID) > 0 {
		fmt.Fprintln(out, "// Requires", in.CpuID+".")
	}
	if rework {
		fmt.Fprintln(out, "//\n// FIXME: Will likely need to be reworked (has pointer parameter).")
	}
	if in.hasImmediate() {
		fmt.Fprintln(out, "//\n// FIXME: Requires compiler support (has immediate)")
	}
	fmt.Fprint(out, "func ", in.Name, "(")
	params := []string{}
	for _, param := range in.Params {
		params = append(params, fmt.Sprint(param.Name, " ", param.Type))
	}

	retparam := Param{Type: in.RetType}
	rettypeprint := in.RetType
	if strings.Contains(rettypeprint, ".") {
		rettypeprint = "(dst " + rettypeprint + ")"
	}

	fmt.Fprint(out, strings.Join(params, ", "), ") ", rettypeprint, " {\n")

	// We have a receiver pointer,
	// ALWAYS SKIP
	if true {
		//fmt.Fprintln(out, "\t// FIXME: Rework to avoid possible return value as parameter.")
		if strings.Contains(in.RetType, "arm.") {
			fmt.Fprint(out, "\treturn "+in.RetType+"{}")
		} else if in.RetType != "" {
			fmt.Fprint(out, "\treturn 0")
		}
		fmt.Fprint(out, "\n}\n")
		return
	}

	if in.RetType != "" {
		fmt.Fprint(out, "\treturn "+in.RetType+"("+in.AsmName+"(")
	} else {
		fmt.Fprint(out, "\t"+in.AsmName+"(")
	}
	for i, param := range in.Params {
		pn := param.Name
		if param.getNative() != "" {
			// cast
			pn = param.getNative() + "(" + pn + ")"
		}
		fmt.Fprint(out, pn)
		if i != len(in.Params)-1 {
			fmt.Fprint(out, ", ")
		}
	}
	if in.RetType != "" {
		fmt.Fprint(out, ")")
	}
	fmt.Fprintln(out, ")\n}")
	fmt.Fprintln(out, "")

	/* SKIP ASM DECL */
	return

	// ASM declaration
	fmt.Fprint(out, "func ", in.AsmName, "(")
	params = []string{}
	for _, param := range in.Params {
		typ := param.Type
		ot := param.getNative()
		if ot != "" {
			typ = ot
		}
		params = append(params, fmt.Sprint(param.Name, " ", typ))
	}
	if retparam.getNative() != "" {
		retparam = Param{Type: retparam.getNative()}
	}
	fmt.Fprint(out, strings.Join(params, ", "), ") ", retparam.Type, "\n\n")

	//emptyType(in.RetType)
	// ASM
	out = in.getFiles().asmFile
	fmt.Fprint(out, "// func ", in.AsmName, "(")
	fmt.Fprint(out, strings.Join(params, ", "), ") ", retparam.Type, "\n")
	fmt.Fprint(out, "TEXT ·"+in.AsmName+"(SB),7,$0\n")
	load, off := in.Params.getAsm()
	fmt.Fprintln(out, load)
	retparam = Param{Type: in.RetType}

	fmt.Fprint(out, "\t// TODO: Code missing")
	if in.Instruction != "..." && in.Instruction != "" {
		if len(params) == 1 {
			if retparam.getReg(0) == "X0" && in.Params[0].getReg(0) == "X0" {
				fmt.Fprintf(out, " - could be:\n\t// %s %s, %s", in.Instruction, in.Params[0].getReg(0), in.Params[0].getReg(0))
			} else if retparam.getReg(0) == "M0" && in.Params[0].getReg(0) == "M0" {
				fmt.Fprintf(out, " - could be:\n\t// %s %s, %s", in.Instruction, in.Params[0].getReg(0), in.Params[0].getReg(0))
			} else if in.Params[0].getReg(0) != "" {
				fmt.Fprintf(out, " - could be:\n\t// %s %s", in.Instruction, in.Params[0].getReg(0))
			} else {
				fmt.Fprintf(out, " - uses instrunction: %s", in.Instruction)
			}
		} else if len(params) == 2 {
			fmt.Fprintf(out, " - could be:\n\t// %s %s, %s", in.Instruction, in.Params[0].getReg(0), in.Params[1].getReg(1))
		} else {
			fmt.Fprintf(out, " - uses instrunction: %s", in.Instruction)
		}
	}
	fmt.Fprintln(out, "\n")

	// Attempts to write a return value
	if in.RetType != "" {
		// Guess a return reg
		rreg := retparam.getReg(0)

		if len(params) > 0 {
			rtry := retparam.getReg(len(params) - 1)
			if rtry != "" {
				rreg = rtry
			}
		}

		if retparam.getSize() > 0 && retparam.getReg(0) == "R8" {
			fmt.Fprint(out, "\tMOV"+retparam.getPostfix()+" $0, ret+"+strconv.Itoa(off)+"(FP)\n")
		} else if retparam.getSize() > 8 {
			fmt.Fprint(out, "\tMOV"+retparam.getPostfix()+" "+rreg+", ret+"+strconv.Itoa(off)+"(FP)\n")
		} else {
			fmt.Fprintf(out, "\t// Return size: %d\n", retparam.getSize())
		}

	}
	fmt.Fprint(out, "\tRET\n")
	fmt.Fprint(out, "\n")
}

func (p Param) getSize() int {
	t := p.Type
	tsplit := strings.Split(t, ".")
	if len(tsplit) > 1 {
		t = tsplit[1]
	}

	switch t {
	case "M64", "M64i":
		return 8
	case "M128", "M128i", "M128d":
		return 16
	case "M256", "M256i", "M256d":
		return 32
	case "M512", "M512i", "M512d":
		return 64
	case "int", "uint", "int64", "uint64", "float64", "uintptr", "Mmask64":
		return 8
	case "int32", "uint32", "float32", "Mmask32":
		return 4
	case "int16", "uint16", "Mmask16":
		return 2
	case "int8", "uint8", "Mmask8", "byte":
		return 1
	}
	return 0
}

func (p Param) getNative() string {
	t := p.Type
	tsplit := strings.Split(t, ".")
	if len(tsplit) > 1 {
		t = tsplit[1]
	}

	switch t {
	case "Int8x8":
		return "[8]int8"
	case "Int16x4":
		return "[4]int16"
	case "Int32x2":
		return "[2]int32"
	case "Int64x1":
		return "int64"
	case "Float32x2":
		return "[2]float32"
	case "Poly8x8":
		return "[8]byte"
	case "Poly16x4":
		return "[4]uint16"
	case "Uint8x8":
		return "[8]byte"
	case "Uint16x4":
		return "[4]uint16"
	case "Uint32x2":
		return "[2]uint32"
	case "Float64x1":
		return "float64"
	case "Uint64x1":
		return "uint64"
	case "Int8x16":
		return "[16]int8"
	case "Int16x8":
		return "[8]int16"
	case "Int32x4":
		return "[4]int32"
	case "Int64x2":
		return "[2]int64"
	case "Float32x4":
		return "[4]float32"
	case "Float64x2":
		return "[2]float64"
	case "Poly8x16":
		return "[16]uint8"
	case "Poly16x8":
		return "[8]uint16"
	case "Poly64x2":
		return "[2]uint64"
	case "Uint8x16":
		return "[16]uint8"
	case "Uint16x8":
		return "[8]uint16"
	case "Uint32x4":
		return "[4]uint32"
	case "Uint64x2":
		return "[2]uint64"
	case "Poly8":
		return "uint8"
	case "Poly16":
		return "uint16"
	case "Poly64":
		return "uint64"
	case "Poly128":
		return "[2]uint64"

	}
	return ""
}

func (p Param) getPostfix() string {
	size := p.getSize()
	switch size {
	case 1:
		return "B"
	case 2:
		return "W"
	case 4:
		return "L"
	case 8:
		return "Q"
	case 16:
		return "OU"
	}
	return ""
}

func (p Param) getReg(n int) string {
	if n > 7 {
		return ""
	}
	t := p.Type
	tsplit := strings.Split(t, ".")
	if len(tsplit) > 1 {
		t = tsplit[1]
	}

	switch t {
	case "M64", "M64i":
		return "M" + strconv.Itoa(n)
	case "M128", "M128i", "M128d":
		return "X" + strconv.Itoa(n)
	case "M256", "M256i", "M256d":
		return "Y" + strconv.Itoa(n)
	case "M512", "M512i", "M512d":
		return "Z" + strconv.Itoa(n)
	case "int", "uint", "int64", "uint64", "float64", "uintptr", "Mmask64":
		return "R" + strconv.Itoa(8+n)
	case "int32", "uint32", "float32", "Mmask32":
		return "R" + strconv.Itoa(8+n)
	case "int16", "uint16", "Mmask16":
		return "R" + strconv.Itoa(8+n)
	case "int8", "uint8", "Mmask8", "byte":
		return "R" + strconv.Itoa(8+n)
	}
	return ""
}

func (p Params) getAsm() (string, int) {
	out := ""
	atbytes := 0
	for i, param := range p {
		size := param.getSize()
		if size == 0 {
			return "\t// FIXME: Unimplemented. Unknown size of type " + param.Type + "\n", 0
		}
		reg := param.getReg(i)
		if reg == "" {
			return "\t// FIXME: Unimplemented. Unknown register for type " + param.Type + "\n", 0
		}
		postfix := param.getPostfix()
		if postfix == "" {
			out += "\t// FIXME: Unimplemented. Unknown MOVE postfix for type " + param.Type + "\n\t//"
		}
		if strings.Contains(param.Name, "imm") {
			out += "\t// FIXME: Immediate parameter should be removed (" + param.Name + " " + param.Type + ")\n"
		} else {
			out += "\tMOV" + postfix + " " + param.Name + "+" + strconv.Itoa(atbytes) + "(FP)," + reg + "\n"
		}
		if size < 4 {
			size = 4
		}
		atbytes += size
	}
	return out, atbytes
}

func emptyType(s string) string {
	if strings.HasPrefix(s, "int") {
		return "0"
	}
	if strings.HasPrefix(s, "uint") {
		return "0"
	}
	if strings.HasPrefix(s, "Mmask") {
		return s + "(0)"
	}
	if strings.HasPrefix(s, "float") {
		return "0.0"
	}
	switch s {
	case "string":
		return `""`
	case "byte":
		return "0"
	case "unsafe.Pointer":
		return s + "(uintptr(0))"
	default:
		return s + "{}"
	}
}

// WrapString wraps the given string within lim width in characters.
//
// Wrapping is currently naive and only happens at white-space. A future
// version of the library will implement smarter wrapping. This means that
// pathological cases can dramatically reach past the limit, such as a very
// long word.
func WrapString(s string, lim uint) string {
	// Initialize a buffer with a slightly larger size to account for breaks
	init := make([]byte, 0, len(s))
	buf := bytes.NewBuffer(init)

	var current uint
	var wordBuf, spaceBuf bytes.Buffer
	var skipline bool

	for _, char := range s {
		if char == '\n' {
			if wordBuf.Len() == 0 {
				if current+uint(spaceBuf.Len()) > lim {
					current = 0
				} else {
					current += uint(spaceBuf.Len())
					spaceBuf.WriteTo(buf)
				}
				spaceBuf.Reset()
			} else {
				current += uint(spaceBuf.Len() + wordBuf.Len())
				spaceBuf.WriteTo(buf)
				spaceBuf.Reset()
				wordBuf.WriteTo(buf)
				wordBuf.Reset()
			}
			buf.WriteRune(char)
			current = 0
			skipline = false
		} else if unicode.IsSpace(char) {
			if current == 0 && spaceBuf.Len() > 2 {
				skipline = true
			}

			if spaceBuf.Len() == 0 || wordBuf.Len() > 0 {
				current += uint(spaceBuf.Len() + wordBuf.Len())
				spaceBuf.WriteTo(buf)
				spaceBuf.Reset()
				wordBuf.WriteTo(buf)
				wordBuf.Reset()
			}

			spaceBuf.WriteRune(char)
		} else {

			wordBuf.WriteRune(char)

			if current+uint(spaceBuf.Len()+wordBuf.Len()) > lim && uint(wordBuf.Len()) < lim && !skipline {
				buf.WriteRune('\n')
				current = 0
				spaceBuf.Reset()
			}
		}
	}

	if wordBuf.Len() == 0 {
		if current+uint(spaceBuf.Len()) <= lim {
			spaceBuf.WriteTo(buf)
		}
	} else {
		spaceBuf.WriteTo(buf)
		wordBuf.WriteTo(buf)
	}

	return buf.String()
}
