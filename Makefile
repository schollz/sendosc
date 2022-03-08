
sendosc:
	go build -ldflags="-s -w" 
	upx --brute sendosc