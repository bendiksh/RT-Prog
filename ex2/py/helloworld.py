import threading

global i
i = 0

myLock = threading.Lock()
#lock = threading.Lock

def Thread1():
	global i
	for j in range(1000000):
		myLock.acquire()
		i=i+1
		myLock.release()


def Thread2():
	global i	
	for j in range(1000000):
		myLock.acquire()
		i=i-1
		myLock.release()

# Potentially useful thing:
# In Python you "import" a global variable, instead of "export"ing it when you declare it
# (This is probably an effort to make you feel bad about typing the word "global")
	


def main():
	thread1 = threading.Thread(target = Thread1, args = (),)
	thread2 = threading.Thread(target = Thread2, args = (),)
	thread1.start()
	thread2.start()
	
	thread1.join()
	thread2.join()
	print(i)

main()
