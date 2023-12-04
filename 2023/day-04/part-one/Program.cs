using System;
using System.IO;
using System.Collections.Generic;

class Program
{
    static void Main()
    {
        string filePath = "input.txt";
        string input = File.ReadAllText(filePath);

        int points = 0;

        foreach(var line in input.Split("\n"))
        {
            if (line.Length < 9)
            {
                continue;
            }

            string trimmedLine = line.Substring(9);

            string[] card = trimmedLine.Split(' ');

            bool areWinningNumbers = true;

            List<int> winningNumbers = new List<int>();

            bool isFirstWin = true;

            int cardPoints = 0;

            foreach(var number in card)
            {
                if(number == "|")
                {
                    areWinningNumbers = false;
                }
                else if (number != "")
                {
                    if (areWinningNumbers) {
                        winningNumbers.Add(int.Parse(number));
                    } else {
                        if (winningNumbers.Contains(int.Parse(number)))
                        {
                            if (isFirstWin) {
                                isFirstWin = false;
                                cardPoints++;
                            } else {
                                cardPoints = cardPoints*2;
                            }
                        }
 
                    }
                }
            }

            points = points+cardPoints;

        }

        Console.Write("Total points: " + points);

    }
}