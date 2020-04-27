@echo off
title win laravel queue
set /p pan=输入盘符:
set /p work_path=输入项目目录:
%pan%:
:start
php %work_path%\artisan schedule:run
choice /t 60 /d y /n >nul

goto start