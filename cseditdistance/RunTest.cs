using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace cseditdistance
{
  public class RunTest
  {
    public int Calculate(string source1, string source2) //O(n*m)
    {  
      int return_int = 0;
      var currentDate = DateTime.Now;

      return_int = LevenshteinDistance.Calculate(source1, source2);

      Console.WriteLine($"{currentDate:d}-{currentDate:t}: Result = {return_int} #: When text1 = {source1} , text2 = {source2}");
      return return_int;
     }
  }
}
