# cd build
# You may need to adjust this next file:
# cmake .. -DCMAKE_TOOLCHAIN_FILE=/Users/josh/Code/Retro68-build/build-target/cmake/retro68.toolchain.cmake
# make
# Then re-run make when you want a fresh build.

cmake_minimum_required(VERSION 2.8)

add_application(Limes CREATOR LIMZ
    main.c
    limes.rsrc.bin
)