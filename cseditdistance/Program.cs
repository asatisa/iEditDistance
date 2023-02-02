// See https://aka.ms/new-console-template for more information
using System.Xml.Linq;

//var name = Console.ReadLine();
var currentDate = DateTime.Now;

Console.WriteLine("Hello, World!");
Console.WriteLine($"{Environment.NewLine}On {currentDate:d} at {currentDate:t}!");

Console.Write($"{Environment.NewLine}Press any key to exit...");
Console.ReadKey(true);