// pi_omp.cpp
#include <omp.h>
#include <stdio.h>
#include <algorithm>

int main()
{
    const size_t num_steps= 4000000000LL;
    const double step = 1.0/num_steps;

    double sec = omp_get_wtime();

    const int max_threads = 100;
    double sum[max_threads];
    int num_threads = 0;
    #pragma omp parallel
    {
        int n = std::min(omp_get_num_threads(), max_threads);
        int k = omp_get_thread_num();
        if (k < max_threads) {
            sum[k] = 0;
            for (size_t i = k; i < num_steps; i += n) {
                double x = (i+0.5)*step;
                sum[k] = sum[k]+4.0/(1.0+x*x);
            }
        }
        if (k == 0) num_threads = n;
    }

    double pi = 0;
    for (size_t k = 0; k < num_threads; ++k)
        pi += step*sum[k];
    
    sec = omp_get_wtime()-sec;

    printf("pi = %.16f, threads %d, time %.3f\n", pi, num_threads, sec);

    return 0;
}
