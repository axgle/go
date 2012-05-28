exe: clean
	windres -o controls-rc.o controls.rc
	go tool 6g controls.go
	go tool pack grc _go_.6 controls.6 controls-rc.o
	go tool 6l -o controls.exe _go_.6
	rm *.6 *.o
clean:
	rm -f *.6 *.o *.exe