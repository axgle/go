exe: clean	
	windres -o res.o resource.rc
	go tool 6g dhv.go
	go tool pack grc _go_.6 dhv.6 res.o
	go tool 6l -o dhv.exe _go_.6
	rm  *.6 *.o
clean:
	rm -f *.6 *.o *.exe