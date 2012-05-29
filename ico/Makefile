exe: clean	
	windres -o demo_res.o demo.rc
	go tool 6g demo.go
	go tool pack grc _demo_go_.6 demo.6 demo_res.o
	go tool 6l -o demo.exe _demo_go_.6
	rm  *.6 *.o
clean:
	rm -f *.6 *.o *.exe