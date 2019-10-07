--game.lua

local sharedApplication = CCApplication:sharedApplication()
local target = sharedApplication:getTargetPlatform()
if patchMe and patchMe("game") then return end  

FuncHelperUtil:loadChunkWithKeyAndSign("lib.zip", "6d2a4052c2", "ZY")

traceMem("before app run ")

print("USE_ENCRYPT_LUA is:"..(USE_ENCRYPT_LUA and "true" or "false")..", type is:"..type(USE_ENCRYPT_LUA))
if USE_ENCRYPT_LUA then 
    local key = "D8C092100C7F487B"
    local sign = "927CC05C-BE89-4356-B36E-374333E05387"
    FuncHelperUtil:setXXTeaKeyAndSign(key, #key, sign, #sign)
end

require("app.MyApp").new():run()

if SHOW_DEBUG_PANEL == 1 then 
	--require("upgrade.ConfigLayer").initConfig()
end
