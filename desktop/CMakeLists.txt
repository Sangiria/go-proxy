cmake_minimum_required(VERSION 3.16)
project(desktop VERSION 0.1 LANGUAGES CXX)

find_package(Qt6 REQUIRED COMPONENTS Widgets)
find_package(Protobuf REQUIRED)
find_package(gRPC CONFIG REQUIRED)

set(CMAKE_AUTOUIC ON)
set(CMAKE_AUTOMOC ON)
set(CMAKE_AUTORCC ON)

set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_STANDARD_REQUIRED ON)

set(PROTO_DIR "${CMAKE_CURRENT_SOURCE_DIR}/..")
set(PROTO_FILE "${PROTO_DIR}/proxy.proto")
set(GEN_DIR "${CMAKE_CURRENT_SOURCE_DIR}/src/client/generated")

add_custom_command(
    OUTPUT "${GEN_DIR}/proxy.pb.cc" 
           "${GEN_DIR}/proxy.pb.h" 
           "${GEN_DIR}/proxy.grpc.pb.cc" 
           "${GEN_DIR}/proxy.grpc.pb.h"
    COMMAND ${Protobuf_PROTOC_EXECUTABLE}
    ARGS --cpp_out=${GEN_DIR}
         --grpc_out=${GEN_DIR}
         --plugin=protoc-gen-grpc=$<TARGET_FILE:gRPC::grpc_cpp_plugin>
         -I${PROTO_DIR}
         ${PROTO_FILE}
    DEPENDS "${PROTO_FILE}"
    VERBATIM
)

set(PROJECT_SOURCES
    src/main.cpp
    src/mainwindow.cpp
    src/mainwindow.h
    form/mainwindow.ui
    form/addsubscriptiondialog.ui
    "${GEN_DIR}/proxy.pb.cc" 
    "${GEN_DIR}/proxy.grpc.pb.cc"
)

qt_add_executable(desktop
    MANUAL_FINALIZATION
    ${PROJECT_SOURCES}
)

set_property(TARGET desktop PROPERTY AUTOUIC_SEARCH_PATHS
    "${CMAKE_CURRENT_SOURCE_DIR}/form"
)

target_link_libraries(desktop PRIVATE 
    Qt6::Widgets 
    gRPC::grpc++ 
    protobuf::libprotobuf
)

target_include_directories(desktop PRIVATE "${GEN_DIR}")

include(GNUInstallDirs)
install(TARGETS desktop
    BUNDLE DESTINATION .
    LIBRARY DESTINATION ${CMAKE_INSTALL_LIBDIR}
    RUNTIME DESTINATION ${CMAKE_INSTALL_BINDIR}
)

qt_finalize_executable(desktop)