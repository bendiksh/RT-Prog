from threading import Thread

i = 0

def Thread1():
	for j in range(1000000):
		i=i+1

def Thread1():
	for j in range(1000000):
		i=i-1

# Potentially useful thing:
# In Python you "import" a global variable, instead of "export"ing it when you declare it
# (This is probably an effort to make you feel bad about typing the word "global")
	global i


def main():
	thread1 = Thread(target = Thread1, args = (),)
	thread2 = Thread(target = Thread1, args = (),)
	thread1.start()
	thread2.start()
	
	thread1.join()
	thread2.join()
	print(i)

main()
