#ifndef WOFF2_C_API_H_
#define WOFF2_C_API_H_

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

#if defined(WIN32) || defined(WINDOWS) || defined(_WIN32) || defined(_WINDOWS)
#define FLYWAVE_WOFF2_API __declspec(dllexport)
#else
#define FLYWAVE_WOFF2_API
#endif

FLYWAVE_WOFF2_API size_t max_woff2_compressed_size(const uint8_t *data,
                                                   size_t length);
FLYWAVE_WOFF2_API size_t compute_woff2_final_size(const uint8_t *data,
                                                  size_t length);

FLYWAVE_WOFF2_API bool convert_ttf_to_woff2(const uint8_t *data, size_t length,
                                            uint8_t *result,
                                            size_t *result_length);
FLYWAVE_WOFF2_API bool convert_woff2_to_ttf(uint8_t *result,
                                            size_t result_length,
                                            const uint8_t *data, size_t length);

#ifdef __cplusplus
}
#endif

#endif // WOFF2_C_API_H_