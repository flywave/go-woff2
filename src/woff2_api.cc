#include <cstring>

#include <woff2/decode.h>
#include <woff2/encode.h>

#include "woff2_api.h"

size_t max_woff2_compressed_size(const uint8_t *data, size_t length) {
  return woff2::MaxWOFF2CompressedSize(data, length);
}

size_t compute_woff2_final_size(const uint8_t *data, size_t length) {
  return woff2::ComputeWOFF2FinalSize(data, length);
}

bool convert_ttf_to_woff2(const uint8_t *data, size_t length, uint8_t *result,
                          size_t *result_length) {
  return woff2::ConvertTTFToWOFF2(data, length, result, result_length);
}

bool convert_woff2_to_ttf(uint8_t *result, size_t result_length,
                          const uint8_t *data, size_t length) {
  return woff2::ConvertWOFF2ToTTF(result, result_length, data, length);
}
