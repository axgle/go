cd ./sqlite3 (which have sqlite3.c)
gcc -shared -s -o sqlite3.dll -Wl,--out-implib,libsqlite.dll.a 
(create `sqlite3.dll` and `libsqlite.dll.a`)

cp `libsqlite.dll.a` to `C:\MinGW64\lib`
cp `sqlite3.c` to `C:\MinGW64\include`
cp `sqlite3.dll` to your app directory 


config `pkg-config` file(sqlite3.pc) if need

