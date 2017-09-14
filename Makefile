all:
	make -C client
	make -C server
	mv client/ChatClient .
	mv server/ChatServer .

clean:
	rm ChatClient
	rm ChatServer

re: clean all

.PHONY: all clean re
