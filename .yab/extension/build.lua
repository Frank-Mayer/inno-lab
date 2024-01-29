require("env")
local yab = require("yab")

yab.task("./extension/package.json", "./extension/node_modules/", function()
	yab.cd("./extension/", function()
		os.execute("npm install")
	end)
end)

yab.task(yab.find("./extension/src/", "**.ts"), "./extension/content.js", function()
	yab.cd("./extension/", function()
		os.execute("npm run build")
	end)
end)

os.execute("mkdir -p ./out/")
os.execute("rm -f ./out/extension.zip")
os.execute("zip -r ./out/extension.zip ./extension/ -x *.git*")
