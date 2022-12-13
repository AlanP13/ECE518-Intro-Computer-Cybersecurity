// hw4.cpp
#include <omp.h>
#include <stdio.h>

int main()
{
    #pragma omp parallel
    {
        int n = omp_get_num_threads();
        int i = omp_get_thread_num();
        printf("Thread %d of %d: working on task 1\n", i, n);
    }

    printf("Done with task 1\n");
    
    #pragma omp parallel
    {
        int n = omp_get_num_threads();
        int i = omp_get_thread_num();
        printf("Thread %d of %d: working on task 2\n", i, n);
    }
    
    printf("Done with task 2\n");
    
    return 0;
}
