/* Sum rows of n X n matrix a
   and store in vector b */
void sum_rows1(double *a, double *b, long n) {
    long i, j;
    for (i = 0; i < n; i++) {
	    b[i] = 0; // 当b进行更新的时候，因为B[3] = A +3，相当于两个不同的变量好指向同一个内存地址，会改变中间的值，计算时候会出现差异结果
                    // 两个内存地址的值会互相叠加
	for (j = 0; j < n; j++)
	    b[i] += a[i*n + j];
    }
}

/* Sum rows of n X n matrix a
   and store in vector b */
void sum_rows2(double *a, double *b, long n) {
    long i, j;
    for (i = 0; i < n; i++) {
	    double val = 0;     // 引入局部变量来改变
	for (j = 0; j < n; j++)
	    val += a[i*n + j];
	    b[i] = val;
    }
}

