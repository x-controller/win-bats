@echo off
title �Զ���ȡ�ֿ����
set /p PAN=�������̷�:
set /p work_path=������Ŀ¼:
%PAN%:
cd %work_path%
for /d %%i in (*) do (
  @cd %cd%\%%i
  IF EXIST .git (
    echo [%%i]
    @git pull
  )
)
pause