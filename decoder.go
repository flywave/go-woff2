package woff2

// #include "woff2_api.h"
import "C"
import "unsafe"

func ComputeFinalSize(data []byte) int {
	len := C.compute_woff2_final_size((*C.uint8_t)(unsafe.Pointer(&data[0])), C.size_t(len(data)))
	return int(len)
}

func ConvertWOFF2ToTTF(data []byte) []byte {
	size := ComputeFinalSize(data)
	result := make([]byte, size)
	ok := C.convert_woff2_to_ttf((*C.uint8_t)(unsafe.Pointer(&result[0])), C.size_t(size), (*C.uint8_t)(unsafe.Pointer(&data[0])), C.size_t(len(data)))
	if ok {
		return result
	}
	return nil
}
