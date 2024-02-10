local yab = require("yab")

yab.use("golang", "1.21.6")
yab.use("nodejs", "21.6.0")

if yab.os_type() == "windows" then
    yab.use("msys2", "2024-01-13")
    os.execute("pacman -S gcc")
    os.execute("pacman -S opencv4")
elseif yab.os_type() == "linux" then
    print("Linux not supported")
elseif yab.os_type() == "darwin" then
    if (yab.check_exec("gcc") == false) and (yab.check_exec("clang") == false) then
        print("No C compiler found")
        print("Please install Xcode or Command Line Tools")
        os.exit(1)
    end
    if yab.check_exec("brew") == false then
        print("Homebrew not found")
        print("Please install Homebrew")
        os.exit(1)
    end
    if yab.check_exec("pkg-config") == false then
        print("pkg-config not found")
        print("Please install pkg-config")
        print("brew install pkg-config")
        os.exit(1)
    end
    if os.execute("pkg-config --exists opencv4") ~= 0 then
        print("OpenCV not found")
        print("Please install OpenCV")
        os.exit(1)
    end
end
