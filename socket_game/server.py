import socket
from random import randint

HOST = "127.0.0.1"
PORT = 9999

def createConnect(): 
    s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    s.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEADDR, 1)
    s.bind((HOST, PORT))
    print("Listening...")

    s.listen(1)
    conn, addr = s.accept()
    return conn


    
def handlConnect(conn):
    print("Accepted connection!")
    conn.send(b"Guess the number GAME\n")
    random_number = randint(1, 100)
    print(random_number)

    guesGame(conn, random_number)
    print("Game end")
    conn.close()

def guesGame(conn, random_number):
    while 1:
        data = conn.recv(1024)
        decodeData = data.decode("utf-8")
        decodeData = parseInput(decodeData)
        if not decodeData:
            conn.send(b"END")
            print("Invalid input")
            break
        if decodeData == random_number:
            conn.send(b"EQUAL")
            break
        elif decodeData < random_number:
            conn.send(b"MORE")
        elif decodeData > random_number:
            conn.send(b"LESS")
    
def parseInput(decodeData):
    if decodeData.find("GUESS:") != -1:
        decodeData = decodeData.replace("GUESS:", "")
        if decodeData.isnumeric():
            return int(decodeData)
        return False
    return False
    

def main():
	conn = createConnect()
	handlConnect(conn)


if __name__ == "__main__":
	main()