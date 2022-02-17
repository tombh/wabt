package main

// func UTF8_encode(str string, realloc wasmer.Function, memory wasmer.Memory) (int32, int32) {
//
//   string_bytes = string.encode('utf8')
// 	string_len = len(string_bytes)
// 	ptr = realloc(0, 0, 1, string_len) & 0xffffffff
// 	if ptr + string_len > memory.data_size:
// 			raise IndexError("The string is out of bounds")
// 	buffer = memory.int8_view()
// 	buffer[ptr:ptr+string_len] = string_bytes
// 	return ptr, string_len
// }
