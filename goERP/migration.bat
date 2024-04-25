@echo off
if "%~1"=="" (
    echo The migration name must be provided!
    exit /b
)
for /f "tokens=2 delims==" %%I in ('wmic os get localdatetime /format:list') do set datetime=%%I
set datetime=%datetime:~0,4%-%datetime:~4,2%-%datetime:~6,2%_%datetime:~8,2%-%datetime:~10,2%-%datetime:~12,2%
set datetime=%datetime::=-%
set nome=%datetime%_%1.sql
echo [SQL SCRIPT HERE] > config/migration/queries/%nome%