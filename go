#!/usr/bin/env python

import os
import sys
import datetime
from subprocess import call
from distutils.spawn import find_executable

#######################################
# User configurable stuff, shouldn't
# NEED to modify anything here, but
# here it is if you want
#######################################
LOGGING_DIR 	= os.path.join(os.getenv('HOME'), 'daily_log')
SSH_EXEC 		= find_executable('ssh')
TEE_EXEC 		= find_executable('tee')
SCRIPT_EXEC 	= find_executable('script')



def runCommand(cmd):
    """
    Function: runCommand(cmd):
    Descr: given a list object of commands and args, pass it to the OS to be
            run.  Once the command returns, the script will exit.
    Args: cmd = list of commands and args.  Given a list: [ssh, remote_host],
            the command "ssh remote_host" will be executed
    """
    call(cmd)
    sys.exit(0)



def datePath():
	"""
	Function: datePath()
	Descr: returns a string to be used for a directory path YYYY/MM/DD, 0 padded
	Args: none
	"""
	return datetime.date.today().strftime('%Y/%m/%d')



def timeString():
	"""
	Function: timeString()
	Descr: returns a string in the format of HH.MM.SS.MS, 0 padded, 24 hr format
	Args: none
	"""
	return datetime.datetime.now().strftime('%H.%M.%S.%f')



def timeDateString():
	"""
	Function: timeDateString()
	Descr: returns a string including the date along with the time: YYMMDD-HH.MM.SS.MS
	Args: none
	"""
	return datetime.datetime.now().strftime('%Y%m%d-%H.%M.%S.%f')



def logmethod():
	"""
	Function: logmethod():
	Descr: returns a best guess at how to log the ssh session.  Uses the linux-util 'script' if 
			the os is linux or darwin, otherwise uses the 'tee' command.
	Args: none
	"""
	logmethods = { 	'linux2'   : 'new', 
					'darwin'  : 'old',
				}

	ostype = sys.platform
	if ostype in logmethods.keys():
		return logmethods[ostype]
	else:
		print "ERROR: No 'script' command available, exiting..."
		sys.exit(1)

		
def gencommand(**kwargs):
	"""
	Function gencommand():
	Descr:	returns a list of commands to be executed to connect to the host and
			implements the correct logging method
	Args:	conn_exec:	the command to execute a connection to the remote host.  
						Normally, this is just 'ssh', could be 'telnet'
			logmethod:	the method used for logging the connection output, at the moment,
						just supports the different versions of the 'script' command
			logfile:	the filename to log the output to
			host:		the host to connect to
	Returns:	list of arguments to connect to the remote host and log the output
	"""
	if 'old' in kwargs['logmethod']:
		return [SCRIPT_EXEC, kwargs['logfile'], kwargs['conn_exec'], 
				kwargs['host']]

	if 'new' in kwargs['logmethod']:
		return [SCRIPT_EXEC, '-c',  kwargs['conn_exec'] + '  ' + kwargs['host'],
			kwargs['logfile']]



def printhelp():
	print "Gimme somewhere to go!"
	sys.exit(1)


def transformhost(hostname):
  """
  Function: transformhost():
  Descr:  returns a new hostname after performing some transformation on it.  This is
      normally used to create a 'shortcut' for hostnames that are rather long.
      Say, use 'nms.site' as a shortcut for 'mysuperlongnmshost.site'
  Args: hostname: this is the hostname received as an argument on the command line
  Returns:  newhost, a string containing the transformed hostname
  """
	newhost = str()

	if 'nms.' in hostname:
		newhost = hostname.replace('nms.', 'nmsspare.')

	return newhost

	
#######################################
# MAIN!!
#######################################
def main():

	remote_host = str()
	if len(sys.argv) < 2:
		printhelp()
	else:
		remote_host = sys.argv[1] or printhelp()

	remote_host = transformhost(remote_host)

	current_log_dir = os.path.join(LOGGING_DIR, datePath())
	if not os.path.isdir(current_log_dir):
		os.makedirs(current_log_dir)

	current_log_file = '.'.join([remote_host, timeDateString()])
	current_log_file_path = os.path.join(current_log_dir, current_log_file)
	
	print("-Connecting to: {0}".format(remote_host))
	print("-Logging to: {0}".format(current_log_file_path))

	cmd = gencommand(conn_exec=SSH_EXEC, logmethod=logmethod(), 
					logfile=current_log_file_path, host=remote_host)
	runCommand(cmd) 



if __name__ == '__main__':
	main()



