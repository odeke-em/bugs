What steps will reproduce the problem?
1. on os x generate file with umlaut (e. g. $touch nur_übung.txt)
2. start vim
3. type :n nur_ü<Tab>

What is the expected output?
  completion of filename to
  :n nur_übung.txt

 What do you see instead?
  Bell and no completion.

What version of the product are you using? On what operating system?
  7.4 on os x 10.9

Please provide any additional information below.

The problem is most likely, that os x handles umlauts different than other *nix 
oses: Umlauts are saved as single unicode character for example on linux. Os x 
uses Unicode U+0308 (combining diaeresis) with a letter (here u). Anyway I am 
wondering that open works correctly (e. g. typing :n nur_übung.txt<CR> works 
as expected).
