for /f "usebackq delims== tokens=1,2" %%i in (`set`) do @echo %%i:%%j
@pause