TARGET = libgetcontrollerinfo.so
CC = gcc
CFLAGS = -std=gnu99 -c -Wall -Werror -fpic
LDFLAGS = -shared

SOURCES = control_get_info.c
OBJECTS = control_get_info.o

all: $(TARGET)
$(OBJECTS):
	$(CC) $(CFLAGS) -c $(SOURCES)
$(TARGET): $(OBJECTS)
	$(CC) ${LDFLAGS} -o $@ $^ 

.PHONY : clean
clean:
	rm $(TARGET) $(OBJECTS)