@echo off

title ��ѯPID��ռ��

set /p PID=������PID:

tasklist|findstr %PID%

pause