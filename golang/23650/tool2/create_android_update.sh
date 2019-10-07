#!/bin/sh
cd `dirname $0`
basepath=$( pwd)
inputPath=E:/BH/pc_test/OnePunchMan
outPath=$basepath/android_update
action=$1
curVersion=$2
function create_zip()
{
    curDir=$1
    updateName=$2
    updatePath=$curDir/$updateName
    tempDir=$updatePath/temp
    outDir=$updatePath/out
    rm -rf $tempDir && mkdir $tempDir
    rm -rf $outDir && mkdir $outDir 
    #mv $updatePath/scripts $updatePath/scripts_res 
    ./bin/compile_scripts.sh -m files -i $updatePath/scripts -o $tempDir/scripts -e xxtea_chunk -ek 6d2a4052c2 -es ZY 
    if [ -d $updatePath/res ]; then
        cp -r $updatePath/res $tempDir/res_src
        ./compress_png.sh $updatePath/res_src $tempDir/res_src
        cp -r $updatePath/res/ui/common/common_1.png $tempDir/ui/common/common_1.png
        cp -r $updatePath/res/ui/background/bg_common.png $tempDir/ui/background/bg_common.png
        #加密资源
        ./tools/encrypt_res -i $tempDir/res_src -o $tempDir/res
    else
        mkdir $tempDir/res 
    fi
    ./tools/md5_res -i $tempDir/ -o $outDir/ -m all
    cd $outDir
    zip -r $updateName.zip .
    mv *.zip $curDir #&& rm -rf $updatePath
}


if [ $action = "create" ]; then
    ./tools/md5_res -action md5_content -i $inputPath -o $outPath -v $curVersion
else
    oldVersion=$3
    #./tools/md5_res -action export_update -i $inputPath -o $outPath -v $curVersion -old_version $oldVersion
    # 加密
    #curP=$outPath/update_$oldVersion"_to_"$curVersion
    updateName=update_$oldVersion"_to_"$curVersion
    rm $outPath/hd_res_md5_$curVersion.txt
    create_zip $outPath update_$oldVersion"_to_"$curVersion
fi
