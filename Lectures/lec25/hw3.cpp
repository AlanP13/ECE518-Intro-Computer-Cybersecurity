// hw3.cpp
#include <omp.h>
#include <stdio.h>

int main()
{
    #pragma omp parallel
    {
        int n = omp_get_num_threads();
        int i = omp_get_thread_num();
        printf("Thread %d of %d: ", i, n);
        printf("Hello, world!\n");
    }
    return 0;
}
