@echo off
title win laravel queue
set /p pan=�����̷�:
set /p work_path=������ĿĿ¼:
%pan%:
:start
php %work_path%\artisan schedule:run
choice /t 60 /d y /n >nul

goto start