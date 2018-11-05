# server.py
# Fib microservice

from socket import *
from threading import Thread

def fib(n):
    if n <= 2:
        return 1
    else:
        return fib(n-1)+fib(n-2)

def fib_server(address):
    sock = socket(AF_INET, SOCK_STREAM)
    sock.setsockopt(SOL_SOCKET, SO_REUSEADDR, 1)
    sock.bind(address)
    sock.listen(5)
    while True:
        client,addr = sock.accept()
        print("Connnection",addr)
        Thread(target=fib_handler,args=(client,)).start()
        #fib_handler(client)


def fib_handler(client):
    while True:
        req = client.recv(100)
        if not req:
            break
        n = int(req)
        result = fib(n)
        resp = str(result).encode('ascii')+ b'\n'
        client.send(resp)
    print("Closed")

fib_server(('',25000))
