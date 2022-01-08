
'''
week 13, 2021
This problem was asked by Facebook.

Given a array of numbers representing the stock prices of a company in chronological order, write a function that calculates the maximum profit you could have made from buying and selling that stock once. You must buy before you can sell it.

For example, given   [9, 11, 8, 5, 7, 10], you should return 5, since you could buy the stock at 5 dollars and sell it at 10 dollars.
'''

def highestProfit(stocks):
    max = 0
    for i in range(len(stocks) - 1, -1, -1):
        for j in range(i, -1, -1):
            if stocks[j] < stocks[i] and (stocks[i] - stocks[j]) > max:
                max = stocks[i] - stocks[j]
    return max

stocks = [9, 11, 8, 5, 7, 10]
print(highestProfit(stocks))
