#include<pthread.h>
#include<stdio.h>

// Note the return type: void*

int i = 0;
pthread_mutex_t mut;

void* Thread1(){
	int j;
	for(j=0;j<1000000;j++)
	{
		pthread_mutex_lock(&mut);
		i++;
		pthread_mutex_unlock(&mut);
	}
	return &i;
}
void* Thread2(){
	int j;	
	for(j=0;j<1000001;j++)
	{
		pthread_mutex_lock(&mut);
		i--;
		pthread_mutex_unlock(&mut);
	}
	return &i;
}

int main(){
	pthread_t thread1, thread2;
	pthread_mutex_init(&mut, NULL);
	pthread_create(&thread1, NULL, &Thread1, NULL);
	pthread_create(&thread2, NULL, &Thread2, NULL);
	// Arguments to a thread would be passed here ---------^
	pthread_join(thread1, NULL);
	pthread_join(thread2, NULL);
	printf("%d\n",i);
	return 0;
}
