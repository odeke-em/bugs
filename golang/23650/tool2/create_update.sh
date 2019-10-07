#加密代码 temp
rm -r temp/scripts
./bin/compile_scripts.sh -m files -i ../test/pc_test/sanguohero/scripts -o temp/scripts -e xxtea_chunk -ek MYKEY -es XT
# 导出更新包 temp => update
./tools/create_update_zip -mode=1  -version=1
