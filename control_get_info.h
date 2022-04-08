// C library headers
#include <stdio.h>
#include <string.h>

// Linux headers
#include <fcntl.h>     // Contains file controls like O_RDWR
#include <errno.h>     // Error integer and strerror() function
#include <termios.h>   // Contains POSIX terminal control definitions
#include <unistd.h>    // write(), read(), close()
#include <sys/time.h>  // get time to implement timeout
#include <sys/types.h> 

int init_serial(void);              // Initialize serial port
int read_data(double timeout_ms);   // Blocking (has timeout) get data api

extern char read_buf[256];
