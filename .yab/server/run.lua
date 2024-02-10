require("env")
local yab = require("yab")

yab.setenv("CREDENTIALS_LOCATION", yab.stdout("pwd") .. "/serviceAccountKey.json")
yab.setenv("FULLSCREEN", "true")
yab.cd("server", function()
	os.execute("go run ./cmd/server")
end)
