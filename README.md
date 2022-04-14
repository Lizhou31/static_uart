## Simple get data from UART (JetsonNANO /dev/ttyTHS1)
* Compile : in root folder

		make

* Export environment variables :

		export LD_LIBRARY_PATH=./
	or use ldconfig.
	[reference](https://tldp.org/HOWTO/Program-Library-HOWTO/shared-libraries.html)

* API

		int init_serial(void);              // Initialize serial port
		int read_data(double timeout_ms);   // Blocking (has timeout, in milisecond) get data api
	timeout shoud not be less than 100ms, because of some io delay?
