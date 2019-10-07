
function __G__TRACKBACK__(errorMessage)

    print("----------------------------------------")
    print("LUA ERROR: " .. tostring(errorMessage) .. "\n")
    local traceback = debug.traceback("", 2)
    print(traceback)
    print("----------------------------------------")

    --只有G_Report初始过后才会对错误日志做处理
    -- 关闭错误日志提交
    if false and G_Report ~= nil then
        G_Report:onTrackBack(errorMessage, traceback)
    end

    if SHOW_EXCEPTION_TIP and sm_notifyLayer ~= nil then
        sm_notifyLayer:getDebugNode():removeChildByTag(10000)
        local text = tostring(errorMessage)
        require("upgrade.ErrMsgBox").showErrorMsgBox(text)
    end
end

function traceMem(desc)
    if desc == nil then
        desc = "memory:"
    end

    if CCLuaObjcBridge then
        local callStaticMethod = CCLuaObjcBridge.callStaticMethod

        local ok, ret = callStaticMethod("NativeProxy", "getUsedMemory", nil)

        if ok then
            print(desc .. tostring(ret) .."KB")
        else
            print("call memory failed..." .. tostring(ret))

        end
    elseif CCLuaJavaBridge then

        local methodName = "getUsedMemory"
        local callJavaStaticMethod = CCLuaJavaBridge.callStaticMethod

        --local ok, ret = callJavaStaticMethod("com.youzu.sanguohero.platform.NativeProxy", "getUsedMemory", nil, "()I")
        local ok, ret = callJavaStaticMethod("com.mobile-ease.platform.NativeProxy", "getUsedMemory", nil, "()I")

        if ok then
            print(desc .. tostring(ret) .."KB")
        else
            print("call memory failed...".. tostring(ret))
        end
    end
end

--设置资源加密
FuncHelperUtil:setEncryptConfig(Default_Json_Encrypt_Key, "6d2a4052c2")
FuncHelperUtil:setEncryptConfig(Default_Png_Encrypt_Key, "6d2a4052c2")

local sharedApplication = CCApplication:sharedApplication()
local target = sharedApplication:getTargetPlatform()

local fileUtils = nil
if target < 10 then
    fileUtils = CCFileUtils:sharedFileUtils()
else
    fileUtils = cc.FileUtils:sharedFileUtils()
end

function getUpgradeDir()
    local upgradeFolder = CCFileUtils:sharedFileUtils():getWritablePath()
    upgradeFolder = upgradeFolder.. "upgrade/"
    return upgradeFolder
end

AutoUpgradeHelper:setUpgradeFolder(getUpgradeDir())
AutoUpgradeHelper:setUpgradeFolder(getUpgradeDir().."scripts/")
require("upgrade.config")

if USE_FLAT_LUA == nil or USE_FLAT_LUA == "0" then
  if target == kTargetIphone or target == kTargetIpad or target == kTargetAndroid then
    --FuncHelperUtil:loadChunkWithDefaultKeyAndSign("upgrade.zip")
  end
end

require("upgrade.upgrade")
