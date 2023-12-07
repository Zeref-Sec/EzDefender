# EzDefender
Tips and Trick for easy bypassing of Windows Defender - 2023

As a penetration tester we often need to deal Windows Defender and long gone are the days were you could use a default meterpreter payload and "getsystem".
Most default meterpreter payloads are going to be discovered by CTI and SOC teams and signatures will be generated to make sure the payloads are caught.

So here I have curated a few tips to get around Windows defender without spending too much time playing around with Metasploit Templates or developing new payloads from scratch.

## Metasploit

As mentioned before Meterpreter doesn't tend to work and default metasploit payloads will get caught 90% of the time. However, there is still utility within Metasploit's payloads that can be of use for workarounds of Windows Defender.

RC4 Encryption for Basic Payloads:

![image](https://github.com/Zeref-Sec/EzDefender/assets/77666064/8d2f5b80-4de5-4961-93bd-1497610a779d)

This payload will work for most Windows machines before Windows Server 2019. However more modern systems require something a bit more customized.

## Golang Payloads

Golang is a great language as payloads can cross compiled for most common operating systems and have low signature detections. In this repo you'll find a basic outline for a payload created in Golang:

Windows 11 - 07/12/2023 - Real Time Protect Enabled Results
![image](https://github.com/Zeref-Sec/EzDefender/assets/77666064/469a99da-4e36-483c-900a-5c26f3df3bd7)
