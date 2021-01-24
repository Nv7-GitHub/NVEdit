# Fyne-cross failed with linux, so this is a script for linux users to compile it themselves
# linux output: NVEdit.tar.gz
go build -o NVEdit -ldflags="-s -w"
fyne package -os linux -icon Icon.png -appID com.nvcode.nvedit
