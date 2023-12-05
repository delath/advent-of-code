using System;
using System.IO;
using System.Collections.Generic;

class Program
{
    static void Main()
    {
        string filePath = "input.txt";
        string input = File.ReadAllText(filePath);

        int idsSum = 0;
        int counter = 0;

        foreach(var line in input.Split("\n"))
        {
            if (line.Length < 9)
            {
                continue;
            }

            string trimmedLine = line.Split(":")[1];

            counter++;

            bool isValid = true;

            foreach(var extraction in trimmedLine.Split(";"))
            {
                foreach(var rawColor in extraction.Split(","))
                {
                    string rawTrimmedColor = rawColor.Substring(1);

                    string[] trimmedColor = rawTrimmedColor.Split(" ");

                    int instances = int.Parse(trimmedColor[0]);
                    string color = trimmedColor[1];

                    if ((color == "red" && instances > 12) || (color == "green" && instances > 13) || (color == "blue" && instances > 14))
                    {
                        isValid = false;
                    }
                }
            }

            if (isValid)
            {
                idsSum += counter;
            }

        }

        Console.Write("Ids sum: " + idsSum);

    }
}