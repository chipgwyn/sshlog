'go' is a little script that will connect to a remote host and log all output

Currently it uses 'ssh' and the 'script' utility.  It places the logfiles in
a YYYY/MM/DD dir and save the file as $hostname.timestamp.  

This allows you to go back in time and see what you did at some point, view the
output from some command, or even just view what you were up to that day.

It should work on most unix-like systems.

*NOTE*  OSX and Linux use 2 different versions of the 'script' command, so yeah...

I'm hoping to generate some additional features in the next month or so, but pull requests
are most welcome!

--chip


