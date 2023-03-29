import socket


def createConnect():
    conn = socket.socket()
    conn.connect( ("127.0.0.1", 9999) )
    guesGame(conn)
    conn.close()

def guesGame(conn):
    decodeData = ''
    while 1:
        data = conn.recv(1024)
        decodeData = data.decode("utf-8")
        if str(decodeData) == "EQUAL" or str(decodeData) == "END":
            break
        print(decodeData)
        number = input()
        conn.send(number.encode("utf-8"))


def main():
	conn = createConnect()


if __name__ == "__main__":
	main()