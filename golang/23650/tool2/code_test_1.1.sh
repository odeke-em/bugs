#make
# 加密代码 src => temp
rm -r temp/scripts
#./bin/compile_scripts.sh -m files -i ../test/pc_test/sanguohero/scripts -o temp/scripts -e xxtea_chunk -ek 6d2a4052c2 -es ZY 
cp -r ../../k007_client/test_1.1/pc_test/sanguohero/scripts temp/scripts
# 加密资源 src => temp
#./tools/encrypt_res -i ../test/pc_test/sanguohero/
# 导出更新包 temp => update
#./tools/md5_res -i /Volumes/hd2/k007/test/pc_test/sanguohero/ -m all
./tools/md5_res -i temp/ -o out_1.1/ -m code


#protoc -I ../test/pc_test/sanguohero/res  -o cs.pb ../test/pc_test/sanguohero/res/cs.proto
