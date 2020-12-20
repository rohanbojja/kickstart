using System;
using System.Linq;

namespace Allocation
{
    class Program
    {
        static void Main(string[] args)
        {
            int cases = int.Parse(Console.ReadLine());
            for(int testCase = 1; testCase < cases + 1; testCase++)
            {
                var nb = Console.ReadLine().Split().Select(x => int.Parse(x)).ToList();
                var a = Console.ReadLine().Split().Select(x => int.Parse(x)).ToList();
                a.Sort();
                int ans = 0;
                int b2 = 0;
                foreach(int price in a)
                {
                    if (b2 + price > nb[1])
                    {
                        break;
                    }
                    else
                    {
                        ans += 1;
                        b2 += price;
                    }
                }
                Console.WriteLine($"Case #{testCase}: {ans}");
            }
        }
    }
}
