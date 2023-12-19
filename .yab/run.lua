require("proto")

local function panic(msg)
	print(msg)
	os.exit(1)
end

local js_pacman = Yab.check_exec("pnpm") and "pnpm" or Yab.check_exec("npm") and "npm" or panic("npm or pnpm not found")

print("using " .. js_pacman .. " as js package manager")

if Yab.os_type() == "windows" then
	os.execute("start /D extension " .. js_pacman .. " install")
	os.execute("start /D extension " .. js_pacman .. " run build")
else
	os.execute("cd extension && " .. js_pacman .. " install")
    os.execute("cd extension && " .. js_pacman .. " run build")
end

if Yab.os_type() == "windows" then
	os.execute("start /D server go mod tidy")
	os.execute("start /D server go run ./cmd/server/")
else
	os.execute("cd server && go mod tidy")
	os.execute("cd server && go run ./cmd/server/")
end
