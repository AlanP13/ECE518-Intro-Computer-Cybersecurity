// pi.cpp
#include <omp.h>
#include <stdio.h>

int main()
{
    const size_t num_steps= 4000000000LL;
    const double step = 1.0/num_steps;

    double sec = omp_get_wtime();

    double sum = 0;
    for (size_t i = 0; i < num_steps; ++i)
    {
        double x = (i+0.5)*step;
        sum = sum+4.0/(1.0+x*x);
    }
    double pi = step*sum;
    
    sec = omp_get_wtime()-sec;

    printf("pi = %.16f, time %.3f\n", pi, sec);

    return 0;
}
