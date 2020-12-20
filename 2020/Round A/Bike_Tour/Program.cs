using System;
using System.Linq;

namespace Bike_Tour
{
    class Program
    {
        static void Main(string[] args)
        {
          int cases = int.Parse(Console.ReadLine());
          for(int testCase = 1; testCase < cases + 1; testCase++)
          {
              var n  = int.Parse(Console.ReadLine());;
              var h = Console.ReadLine().Split().Select(x => int.Parse(x)).ToList();
              int ans=0;
              for(int i=0; i< n ; i++){
                if(i!=0 && i!=n-1 && h[i-1]<h[i] && h[i]>h[i+1]){
                  ans+=1;
                }
              }
              Console.WriteLine($"Case #{testCase}: {ans}");
          }
        }
    }
}
