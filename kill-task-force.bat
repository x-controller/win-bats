@echo off

title 强制关闭占用PID的进程

set /p PID=请输入PID:

taskkill /pid %PID% /f

pause