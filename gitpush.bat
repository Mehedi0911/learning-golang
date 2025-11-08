@echo off
if "%~1"=="" (
  echo ❌ Please provide a commit message.
  echo Usage: gitpush "your commit message"
  exit /b 1
)

git add .
git commit -m "%~1"
git push

echo ✅ Code pushed successfully with message: %~1
