@echo off
title �Զ���ȡ�ֿ����
set /p PAN=�������̷�:
set /p work_path=������Ŀ¼:
%PAN%:
cd %work_path%
for /d %%i in (*) do (
  echo.
  echo [%%i]
  echo.
  @cd %cd%\%%i && @git pull
)
pause