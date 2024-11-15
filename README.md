## Installation  
1. download the dota-rich-presence archive.zip on GitHub Releases
2. Put `gamestate_integration_d2drp.cfg` file from the archive to the `STEAM\steamapps\common\dota 2 beta\game\dota\cfg\gamestate_integration` folder
3. unpack and run `dota2-discord-rich-presence.exe`

## Adding to Windows-startup
1. click on the `.exe` file  
2. click `Create shortcut`
3. click `Win + R`
4. in the window that opens, write `shell:startup`
5. move the shortcut to the folder that opens

![image](https://github.com/user-attachments/assets/2a37b43c-5504-4ab4-bbb3-a64673f4fb30)
![image](https://github.com/user-attachments/assets/d0e7bc80-e32e-4745-8f56-7c295181e44f)

## Building
`go get github.com/hugolgst/rich-go`  
`go get github.com/getlantern/systray`  
`go get gopkg.in/toast.v1`  
`go build -ldflags -H=windowsgui && go-dota-drpc.exe`  

