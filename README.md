# tarry
A simple command line tool for waiting until a specific time.  This is not the same as "sleep" which will wait for a duration of time.
This is useful if you want to execute something at a specific time or more likely execute several things at exactly the same time such as testing if a server can handle multiple <em>very</em> simultaneous requests.  You could use it like this with "&&" on Linux, Mac, or Windows:

```
   tarry -until=16:03:04 && someOtherCommand
```
This would wait until 4:03:04 PM and then execute someOtherCommand.  Here's a Linux/Mac example of how to run multiple requests all scheduled to start at the same time:

```
   for request in 1 2 3 4 5 6 7 8 9 10
   do
       tarry -until=16:03:04 && date > results.$request &
   done
```

### You can get it here
You can download the [Ubuntu version here](https://metaphyze-public.s3.amazonaws.com/tarry/releases/1.0/ubuntu/tarry) (built on Ubuntu 18.04LTS), the [Mac version here](https://metaphyze-public.s3.amazonaws.com/tarry/releases/1.0/macos/tarry) (built on macOS Mojave), and the [Windows version here](https://metaphyze-public.s3.amazonaws.com/tarry/releases/1.0/windows/tarry.exe)  These are all standalone executables.  You don't need to install any libraries to run them, but on Linux/Mac you'll need to 
```
chmod +x tarry
```
to make them executable.

### Or you can build it yourself
If you want to build tarry yourself, you'll need to install Go.  Many sites give instructions on this so I won't repeat them.  
Here's a good one: [How to Install Go on Ubuntu 18.04](https://linuxize.com/post/how-to-install-go-on-ubuntu-18-04/).
Once Go is installed, you can build it by simply typing:

    /home/ubuntu> go build YOUR_GO_WORK_DIR/src/tarry/tarry.go


### Command line options
Here's the complete list of command line options.  
```
  -help 
        Print out all the options
  -until string
        Required. Wait (tarry) until this specific time (hours, minutes, seconds).  
        Format: 17:59:01
        For example,
                tarry -until=17:59:01
        Use -whatTimeIsIt to find out what the current system time is.  
        If you need millisecond precision, use the -plusMs flag.
  -plusMs int
        Optional (defaults to 0).  Add a few milliseconds onto your time (0-999).  For example,
                tarry -until=17:59:01 -plusMs=250
        would wait until 17:59:01:250, that is, 1/4 second beyond 17:59:01
  -inDaysFromNow int
        Optional (defaults to 0). Specifies the number of days (0-365) from today.  
        For example, tomorrow would be -inDaysFromNow=1
  -whatTimeIsIt
        Show the current system time.  You'll probably need to use this to set -until.  For example,
                tarry -whatTimeIsIt
                Current time: 17:22:03
  -printTimeWhenDone
        Print out the current time when done (includes milliseconds)
              
```
