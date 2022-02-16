package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	wasmer "github.com/wasmerio/wasmer-go/wasmer"
)

const (
	Exceptions     = 1
	MutableGlobals = 2
	SatFloatToInt  = 4
	SignExtension  = 8
	SIMD           = 16
	Threads        = 32
	MultiValue     = 64
	TailCall       = 128
	BulkMemory     = 256
	ReferenceTypes = 512
	Annotations    = 1024
	GC             = 2048
)

const WASM_BINARY_PATH = "../../out/wasi/Debug/libwabt.wasm"

func wasm2wat(wasm string, flags int) (string, error) {
	wasmBytes, _ := ioutil.ReadFile(WASM_BINARY_PATH)

	engine := wasmer.NewEngine()
	store := wasmer.NewStore(engine)
	module, err := wasmer.NewModule(store, wasmBytes)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	wasiENV, err := wasmer.NewWasiStateBuilder("wasi_test_program").Finalize()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	importObject, err := wasiENV.GenerateImportObject(store, module)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	instance, err := wasmer.NewInstance(module, importObject)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	memory, _ := instance.Exports.GetMemory("memory")
	realloc, _ := instance.Exports.GetFunction("canonical_abi_realloc")
	free, _ := instance.Exports.GetFunction("canonical_abi_free")
	val0 := wasm
	len0 := int32(len(val0))
	_ptr0, _ := realloc(0, 0, 1, len0*1)
	ptr0 := _ptr0.(int32)
	buffer := memory.Data()
	copy(buffer[ptr0:ptr0+len0*1], []byte(val0))
	wasm2wat, _ := instance.Exports.GetFunction("wasm2wat")
	_ret, _ := wasm2wat(ptr0, len0, flags)
	ret := _ret.(int32)
	ret = int32(ret / int32(4))

	if buffer[ret] == 0 {
		ptr1 := buffer[ret+2]
		len1 := buffer[ret+4]
		list1 := string(buffer[ptr1 : ptr1+len1])
		free(ptr1, len1, 1)
		fmt.Println(list1)
		return list1, nil
	} else if buffer[ret] == 1 {
		ptr2 := buffer[ret+2]
		len2 := buffer[ret+4]
		list2 := string(buffer[ptr2 : ptr2+len2])
		free(ptr2, len2, 1)
		fmt.Println(list2)
		return "", errors.New(list2)
	}

	return "", errors.New("Foreign wasm2wat() returned neither 0 nor 1")
}

// func wat2wasm(wasm string, flags int) (string, error) {
// 	memory = self.instance.exports.memory
//         realloc = self.instance.exports.canonical_abi_realloc
//         free = self.instance.exports.canonical_abi_free
//         ptr0, len0 = encode_utf8(wat, realloc, memory)
//         ret = self.instance.exports.wat2wasm(ptr0, len0, flags)
//         ret = int(ret / 4)
//         buf32 = memory.int32_view()
//         if buf32[ret] == 0:
//             ptr1 = buf32[ret + 2]
//             len1 = buf32[ret + 4]
//             buf8 = memory.int8_view()
//             wasm = bytearray(buf8[ptr1:ptr1+len1])
//             free(ptr1, len1, 1)
//             return wasm
//         elif buf32[ret] == 1:
//             ptr2 = buf32[ret + 2]
//             len2 = buf32[ret + 4]
//             message = decode_utf8(memory, ptr2, len2)
//             free(ptr2, len2, 1)
//             raise Exception(message)
// }

func main() {
	// wat2wasm("(module)", MutableGlobals)
	// wat3wasm("xyz", MutableGlobals)
	// wasm2wat(bytearray([
	//    0, 97, 115, 109,   1,   0,   0,   0,   0,
	//    8,  4, 110,  97, 109, 101,   2,   1,   0,
	//    0,  9,   7, 108, 105, 110, 107, 105, 110,
	// 	103,  2
	// ]))
}
