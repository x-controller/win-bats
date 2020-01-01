@echo off

title 查询端口被占用

set /p PORT=请输入端口号:

netstat -aon|findstr %PORT%

pause