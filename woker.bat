@echo off


:start
php D:code\life.hallo.site\artisan schedule:run
choice /t 5 /d y /n >nul

goto start