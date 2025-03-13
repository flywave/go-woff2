# Copyright 2017 Igalia S.L. All Rights Reserved.
#
# Distributed under MIT license.
# See file LICENSE for detail or copy at https://opensource.org/licenses/MIT

# Ubuntu 12.04 LTS has CMake 2.8.7, and is an important target since
# several CI services, such as Travis and Drone, use it.  Solaris 11
# has 2.8.6, and it's not difficult to support if you already have to
# support 2.8.7.
cmake_minimum_required(VERSION 2.8.6)

project(woff2)

include(GNUInstallDirs)

# Build options
option(BUILD_SHARED_LIBS "Build shared libraries" OFF)
option(CANONICAL_PREFIXES "Canonical prefixes" OFF)
option(NOISY_LOGGING "Noisy logging" ON)

# Version information
set(WOFF2_VERSION 1.0.2)

# When building shared libraries it is important to set the correct rpath
# See https://cmake.org/Wiki/CMake_RPATH_handling#Always_full_RPATH
set(CMAKE_SKIP_BUILD_RPATH FALSE)
set(CMAKE_BUILD_WITH_INSTALL_RPATH FALSE)
set(CMAKE_INSTALL_RPATH_USE_LINK_PATH TRUE)
list(FIND CMAKE_PLATFORM_IMPLICIT_LINK_DIRECTORIES "${CMAKE_INSTALL_LIBDIR}" isSystemDir)
if ("${isSystemDir}" STREQUAL "-1")
  set(CMAKE_INSTALL_RPATH "${CMAKE_INSTALL_LIBDIR}")
endif()

# Set compiler flags
if (NOT CANONICAL_PREFIXES)
    add_definitions(-no-canonical-prefixes)
  endif ()
if (NOISY_LOGGING)
    add_definitions(-DFONT_COMPRESSION_BIN)
endif ()
add_definitions(-D__STDC_FORMAT_MACROS)
set(COMMON_FLAGS -fno-omit-frame-pointer)

if(${CMAKE_SYSTEM_NAME} MATCHES "Darwin")
    add_definitions(-DOS_MACOSX)
else ()
    set(COMMON_FLAGS "${COMMON_FLAG} -fno-omit-frame-pointer")
endif()

set(CMAKE_C_FLAGS "${CMAKE_C_FLAGS} ${COMMON_FLAG}")

set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} ${COMMON_FLAG}")
set(CMAKE_CXX_STANDARD 11)

INCLUDE_DIRECTORIES(BEFORE
  ${CMAKE_CURRENT_SOURCE_DIR}/../brotli/c/include
)

# Set search path for our private/public headers as well as Brotli headers
include_directories("src" "include")

# Common part used by decoder and encoder
add_library(woff2common STATIC
            src/table_tags.cc
            src/variable_length.cc
            src/woff2_common.cc)

# WOFF2 Decoder
add_library(woff2dec STATIC
            src/woff2_dec.cc
            src/woff2_out.cc)
target_link_libraries(woff2dec woff2common "${BROTLIDEC_LIBRARIES}")

# WOFF2 Encoder
add_library(woff2enc STATIC
            src/font.cc
            src/glyph.cc
            src/normalize.cc
            src/transform.cc
            src/woff2_enc.cc)
target_link_libraries(woff2enc woff2common "${BROTLIENC_LIBRARIES}")

# Fuzzer libraries
add_library(convert_woff2ttf_fuzzer STATIC src/convert_woff2ttf_fuzzer.cc)
target_link_libraries(convert_woff2ttf_fuzzer woff2dec)
add_library(convert_woff2ttf_fuzzer_new_entry STATIC src/convert_woff2ttf_fuzzer_new_entry.cc)
target_link_libraries(convert_woff2ttf_fuzzer_new_entry woff2dec)


SET_TARGET_PROPERTIES(woff2common PROPERTIES
  ARCHIVE_OUTPUT_DIRECTORY_DEBUG ${CMAKE_CURRENT_BINARY_DIR}
  ARCHIVE_OUTPUT_DIRECTORY_RELEASE ${CMAKE_CURRENT_BINARY_DIR})
SET_TARGET_PROPERTIES(woff2common PROPERTIES 
  LIBRARY_OUTPUT_DIRECTORY_DEBUG ${CMAKE_CURRENT_BINARY_DIR}
  LIBRARY_OUTPUT_DIRECTORY_RELEASE ${CMAKE_CURRENT_BINARY_DIR})
SET_TARGET_PROPERTIES(woff2common PROPERTIES
  RUNTIME_OUTPUT_DIRECTORY_DEBUG ${CMAKE_CURRENT_BINARY_DIR}
  RUNTIME_OUTPUT_DIRECTORY_RELEASE ${CMAKE_CURRENT_BINARY_DIR})
  
SET_TARGET_PROPERTIES(woff2dec PROPERTIES
ARCHIVE_OUTPUT_DIRECTORY_DEBUG ${CMAKE_CURRENT_BINARY_DIR}
ARCHIVE_OUTPUT_DIRECTORY_RELEASE ${CMAKE_CURRENT_BINARY_DIR})
SET_TARGET_PROPERTIES(woff2dec PROPERTIES 
LIBRARY_OUTPUT_DIRECTORY_DEBUG ${CMAKE_CURRENT_BINARY_DIR}
LIBRARY_OUTPUT_DIRECTORY_RELEASE ${CMAKE_CURRENT_BINARY_DIR})
SET_TARGET_PROPERTIES(woff2dec PROPERTIES
RUNTIME_OUTPUT_DIRECTORY_DEBUG ${CMAKE_CURRENT_BINARY_DIR}
RUNTIME_OUTPUT_DIRECTORY_RELEASE ${CMAKE_CURRENT_BINARY_DIR})

SET_TARGET_PROPERTIES(woff2enc PROPERTIES
  ARCHIVE_OUTPUT_DIRECTORY_DEBUG ${CMAKE_CURRENT_BINARY_DIR}
  ARCHIVE_OUTPUT_DIRECTORY_RELEASE ${CMAKE_CURRENT_BINARY_DIR})
SET_TARGET_PROPERTIES(woff2enc PROPERTIES 
  LIBRARY_OUTPUT_DIRECTORY_DEBUG ${CMAKE_CURRENT_BINARY_DIR}
  LIBRARY_OUTPUT_DIRECTORY_RELEASE ${CMAKE_CURRENT_BINARY_DIR})
SET_TARGET_PROPERTIES(woff2enc PROPERTIES
  RUNTIME_OUTPUT_DIRECTORY_DEBUG ${CMAKE_CURRENT_BINARY_DIR}
  RUNTIME_OUTPUT_DIRECTORY_RELEASE ${CMAKE_CURRENT_BINARY_DIR})
  
SET_TARGET_PROPERTIES(convert_woff2ttf_fuzzer PROPERTIES
ARCHIVE_OUTPUT_DIRECTORY_DEBUG ${CMAKE_CURRENT_BINARY_DIR}
ARCHIVE_OUTPUT_DIRECTORY_RELEASE ${CMAKE_CURRENT_BINARY_DIR})
SET_TARGET_PROPERTIES(convert_woff2ttf_fuzzer PROPERTIES 
LIBRARY_OUTPUT_DIRECTORY_DEBUG ${CMAKE_CURRENT_BINARY_DIR}
LIBRARY_OUTPUT_DIRECTORY_RELEASE ${CMAKE_CURRENT_BINARY_DIR})
SET_TARGET_PROPERTIES(convert_woff2ttf_fuzzer PROPERTIES
RUNTIME_OUTPUT_DIRECTORY_DEBUG ${CMAKE_CURRENT_BINARY_DIR}
RUNTIME_OUTPUT_DIRECTORY_RELEASE ${CMAKE_CURRENT_BINARY_DIR})

SET_TARGET_PROPERTIES(convert_woff2ttf_fuzzer_new_entry PROPERTIES
  ARCHIVE_OUTPUT_DIRECTORY_DEBUG ${CMAKE_CURRENT_BINARY_DIR}
  ARCHIVE_OUTPUT_DIRECTORY_RELEASE ${CMAKE_CURRENT_BINARY_DIR})
SET_TARGET_PROPERTIES(convert_woff2ttf_fuzzer_new_entry PROPERTIES 
  LIBRARY_OUTPUT_DIRECTORY_DEBUG ${CMAKE_CURRENT_BINARY_DIR}
  LIBRARY_OUTPUT_DIRECTORY_RELEASE ${CMAKE_CURRENT_BINARY_DIR})
SET_TARGET_PROPERTIES(convert_woff2ttf_fuzzer_new_entry PROPERTIES
  RUNTIME_OUTPUT_DIRECTORY_DEBUG ${CMAKE_CURRENT_BINARY_DIR}
  RUNTIME_OUTPUT_DIRECTORY_RELEASE ${CMAKE_CURRENT_BINARY_DIR})
  