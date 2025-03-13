package woff2

// #include <stdlib.h>
// #include <string.h>
// #include "woff2_api.h"
// #cgo CFLAGS: -I ./lib
// #cgo CXXFLAGS: -I ./lib
// #cgo linux LDFLAGS:  -L ./lib/linux -Wl,--start-group  -lstdc++ -lm -pthread -lbrotlicommon -lbrotlidec -lbrotlienc -lwoff2common -lwoff2dec -lwoff2enc -lc_woff2 -Wl,--end-group
// #cgo windows LDFLAGS: -L ./lib/windows  -Wl,--start-group  -lstdc++ -lbrotlicommon -lbrotlidec -lbrotlienc -lwoff2common -lwoff2dec -lwoff2enc -lc_woff2 -Wl,--end-group
// #cgo darwin,amd64 LDFLAGS: -L　./lib/darwin -lbrotlicommon -lbrotlidec -lbrotlienc -lwoff2common -lwoff2dec -lwoff2enc -lc_woff2 -lc++
// #cgo darwin,arm64 LDFLAGS: -L　./lib/darwin_arm -lbrotlicommon -lbrotlidec -lbrotlienc -lwoff2common -lwoff2dec -lwoff2enc -lc_woff2 -lc++
import "C"
