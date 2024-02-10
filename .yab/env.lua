local yab = require("yab")

yab.use("golang", "1.21.6")
yab.use("nodejs", "21.6.0")

if yab.os_type() == "windows" then
    yab.use("msys2", "2024-01-13")
    os.execute("pacman -S gcc")
elseif yab.os_type() == "linux" then
    print("Linux not supported")
elseif yab.os_type() == "darwin" then
    if (yab.check_exec("gcc") == false) and (yab.check_exec("clang") == false) then
        print("No C compiler found")
        print("Please install Xcode or Command Line Tools")
        os.exit(1)
    end
end
