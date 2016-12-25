@echo off
echo building...
gox -output="bin/{{.Dir}}_{{.OS}}_{{.Arch}}"
echo uploading to ftp...
"C:\path\WinSCP.com" ^
  /ini=nul ^
  /command ^
    "open sftp://lhunter@lukehunter.net/ -hostkey=""ssh-rsa 2048 f1:1f:a8:45:ed:ac:74:37:fc:ea:3b:d7:19:7b:ea:53""" ^
    "lcd C:\projects\go\src\github.com\lukehunter\ssgallery\bin" ^
    "cd /home/lhunter/lukehunter.net/ssgallery_release" ^
    "put *" ^
    "exit" 

set WINSCP_RESULT=%ERRORLEVEL%
if %WINSCP_RESULT% equ 0 (
  echo Success
) else (
  echo Error
)

echo done
exit /b %WINSCP_RESULT%