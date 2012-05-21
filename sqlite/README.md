
	wget http://sqlite.org/sqlite-amalgamation-3071200.zip
	unzip sqlite-amalgamation-3071200.zip
	cd sqlite-amalgamation-3071200
	rm shell.c
	gcc -c sqlite3.c
	gcc -shared -o sqlite3.dll  sqlite3.o

	wget http://sqlite.org/sqlite-dll-win32-x86-3071200.zip
	unzip sqlite-dll-win32-x86-3071200.zip
	rm sqlite3.dll (dont use the sqlite-dll-win32-x86-3071200.zip's dll,just use sqlite3.def )

`libsqlite3.dll.a` can create by `dlltool -D sqlite3.dll -d sqlite3.def -l libsqlite3.dll.a`

cp `libsqlite.dll.a` to `C:\MinGW64\lib`
cp `sqlite3.c` `sqlite3.h` to `C:\MinGW64\include`
cp `sqlite3.dll` to your app directory 


config `pkg-config` file(sqlite3.pc) if need 
or modify mattn\go-sqlite3\sqlite3.go:

	//#cgo pkg-config: sqlite3
	#cgo CFLAGS: -IC:/MinGW64/include
	#cgo LDFLAGS: -lsqlite3

