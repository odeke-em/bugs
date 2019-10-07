#make
# 加密代码 src => temp
source=../../k007_client/test_1.1/pc_test/sanguohero/scripts
temp=temp_1.1
out=out_1.1

rm -r $temp/scripts
./bin/compile_scripts.sh -m files -i $source -o $temp/scripts -e xxtea_chunk -ek 6d2a4052c2 -es ZY 
#cp -r ../test/pc_test/sanguohero/scripts temp/scripts
# 加密资源 src => temp
#./tools/encrypt_res -i ../test/pc_test/sanguohero/
# 导出更新包 temp => update
#./tools/md5_res -i /Volumes/hd2/k007/test/pc_test/sanguohero/ -m all
./tools/md5_res -i $temp/ -o $out/ -m code
