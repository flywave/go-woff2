package woff2

// #include "woff2_api.h"
import "C"
import "unsafe"

func MaxCompressedSize(data []byte) int {
	len := C.max_woff2_compressed_size((*C.uint8_t)(unsafe.Pointer(&data[0])), C.size_t(len(data)))
	return int(len)
}

func ConvertTTFToWOFF2(data []byte) []byte {
	size := MaxCompressedSize(data)
	result := make([]byte, size)
	resSize := C.size_t(size)
	ok := C.convert_ttf_to_woff2((*C.uint8_t)(unsafe.Pointer(&data[0])), C.size_t(len(data)), (*C.uint8_t)(unsafe.Pointer(&result[0])), &resSize)
	if ok {
		return result[:int(resSize)]
	}
	return nil
}
