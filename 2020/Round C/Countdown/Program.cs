using System;
using System.Linq;

namespace Countdown
{
    class Program
    {
        static void Main(string[] args)
        {
            int cases = int.Parse(Console.ReadLine());
            for (int testCase = 1; testCase < cases + 1; testCase++)
            {
                var nk = Console.ReadLine().Split().Select(x => int.Parse(x)).ToList();
                var a = Console.ReadLine().Split().Select(x => int.Parse(x)).ToList();
                int ans = 0;
                int n = nk[0];
                int k = nk[1];
                for (int i = 0; i <= n - k; i++)
                {
                    bool flag = true;
                    for (int j = i; j < i + k; j++)
                    {
                        if (a[j] != k + i - j)
                        {
                            flag = false;
                            break;
                        }
                        // Console.WriteLine($"{a[j]} {j}");
                    }
                    if (flag)
                    {
                        ans += 1;
                    }
                }
                Console.WriteLine($"Case #{testCase}: {ans}");
            }
        }
    }
}
