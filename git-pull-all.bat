@echo off
title 自动拉取仓库更新
set /p PAN=请输入盘符:
set /p work_path=请输入目录:
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