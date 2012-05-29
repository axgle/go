#http://stackoverflow.com/questions/2627004/embedding-binary-blobs-using-gcc-mingw
exe: 
	ld -r -b binary -o binary.o hello.txt
	gcc -o demo.exe demo.c binary.o
	#rm *.o	
