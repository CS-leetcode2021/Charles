long shift_left(long x ,long n)
{
    x <<= 4;
    x >>= n;
    return x ;
}