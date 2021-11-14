CC=g++
CFLAGS=-Wall -std=c++11 

TARGET=main

SRC_FILES := $(wildcard *.cpp)
OBJ_FILES := $(patsubst %.cpp,%.o,$(SRC_FILES))

all: $(TARGET)

$(TARGET): $(OBJ_FILES)
	$(CC) $(CFLAGS) -o $@ $^

%.o: %.cpp
	$(CC) $(CFLAGS) -g -c -o $@ $<

run: $(TARGET)
	chmod +x $(TARGET)
	./$(TARGET)

clean:
	$(RM) $(TARGET) $(OBJ_FILES)

.PHONY: all clean