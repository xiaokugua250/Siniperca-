#ifndef TLPI_HDR_H
#define TLPI_HDR_H //prevent accidental double inclusion

#include <sys/types.h> //type define used by many programs
#include <stdio.h> //stardard I/O functions
#include <stdlib.h> //prototypes of commony used library functions plus EXIT_SUCCESS and EXIT_FAILURE constants
#include <unistd.h> //prototypes for many system calls
#include <errno.h> // decare errno and defines error constants
#include <string.h> // commonly used string-handling functions

#include "get_num.h" //decarles our functions for handling numeric arguments (getInt(),getLong())

#include "error_functions.h" // decalres our error-handling functions

typedef enum {FALSE,TRUE} Boolean;

#define min(m,n) ((m) < (n) ? (m):(n))
#define max(m,n) ((m) > (n) ? (m):(n))