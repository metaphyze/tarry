# tarry
A simple command line tool for waiting until a specific time.  This is not the same as "sleep" which will wait for duration of time.
This is useful if you want to execute something at a specific time.  You would probably use it like this:

```
   tarry -until=16:03:04 && ./someOtherCommand
```

This would wait until 4:03:04 PM and then execute ./someOtherCommand


### Command line options
Here's the complete list of command line options.  
```
  -help Print out all the options
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
```
