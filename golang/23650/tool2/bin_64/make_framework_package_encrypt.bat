@echo off
set DIR=%~dp0
cd "%DIR%\..\"
%DIR%win32\php.exe %DIR%\lib\compile_scripts.php -i framework -o lib\framework_precompiled\framework_precompiled_encrypt.zip -p framework -m zip -es 9D10F6DE-4E81-4ADC-817A-B381783AE639 -ek 7631A09609C74180 -e xxtea_zip

copy lib\framework_precompiled\framework_precompiled_encrypt.zip template\PROJECT_TEMPLATE_01\res\
