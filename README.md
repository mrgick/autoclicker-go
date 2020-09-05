# autoclicker-go
Autoclicker wroted on go (golang) and based on library robotgo (https://github.com/go-vgo/robotgo)
# How to build?
1) Install go -> https://golang.org/   and  gcc -> https://gcc.gnu.org/
2) Install library robotgo -> in terminal (cmd) run -> go get github.com/go-vgo/robotgo
3) Install special requirements for your system -> https://github.com/go-vgo/robotgo#requirements
4) Open terminal (cmd) in directory with main.go and run -> go build main.go

# How to use?
1) Open terminal (cmd) with your build file
2) Run your build file in terminal (cmd)
3) To run clicking press ctrl+shift+x
4) To stop clicking press ctrl+shift+a

#UPD:
added flags:
To set parametres that you need -> run in terminal(cmd) with flags:                  
--clicking_speed=60 --start_keys=x+ctrl+shift --stop_keys=a+ctrl+shift

YOU NEED SET ONLY 3 KEYS in set_start_keys and set_stop_keys!!!
All keys you can find on https://github.com/go-vgo/robotgo/blob/master/docs/keys.md

Good luck)

