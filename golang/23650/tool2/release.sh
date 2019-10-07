#!/bin/bash
#make
cd $(dirname ${0})
# 加密代码 src => temp
InPath=../../k007_client/test_1.1/pc_test/sanguohero/
TempPath=temp_1.1
OutPath=out_1.1
#rm -r $TempPath
#cp -r $InPath $TempPath
rm -r $TempPath/scripts
#cp -r $InPath/scripts $TempPath/scripts
rm -r $TempPath/scripts/UF
./bin/compile_scripts.sh -m files -i $InPath/scripts -o $TempPath/scripts -e xxtea_chunk -ek 6d2a4052c2 -es ZY 
# 压缩资源
rm -r $TempPath/res_src
cp -r $InPath/res $TempPath/res_src
rm $TempPath/res_src/HDRes.pack
rm $TempPath/res_src/upgrade.zip
cp -r lib_zip/lib.zip $TempPath/res_src/lib.zip
cp -r lib_zip/framework_precompiled.zip $TempPath/res_src/framework_precompiled.zip
#./compress_png.sh $TempPath/res_src $TempPath/res_src
./tools/compress_png -i $TempPath/res_src -o $TempPath/res -t ./tools -lib ./pic_lib
cp -r $InPath/res/ui/common/common_1.png $TempPath/res_src/ui/common/common_1.png
cp -r $InPath/res/ui/background/bg_common.png $TempPath/res_src/ui/background/bg_common.png
# 加密资源 src => temp
rm -r $TempPath/res
./tools/encrypt_res -i $TempPath/res_src -o $TempPath/res
# 导出更新包 temp => update
rm -r $OutPath
./tools/md5_res -i $TempPath/ -o $OutPath/ -m all
