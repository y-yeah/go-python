from storage import (
    storage
)

def printOne():
    print(storage.getData())
    print(storage.setData(10))

def printTwo():
    print(storage.getData())
