@echo off

title 查询PID被占用

set /p PID=请输入PID:

tasklist|findstr %PID%

pause