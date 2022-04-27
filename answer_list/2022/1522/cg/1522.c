// Weekly Coding Problem
// Week 15, 2022 - (NR)
// Chris Georgiades

/*******************************
 * compile: gcc -o 1522 1522.c *
 * run:     ./1522 < infile    *
 *******************************/

#include <stdio.h>
#include <stdbool.h>

void printHeader(int, bool);
void printBody(int, bool);
void printPunchcard(int, int);

int main()
{
    int records;
    char line[80];
    scanf("%d", &records);

    for (int i = 0; i < records; ++i)
    {
        int h, w;
        scanf("%d %d", &h, &w);
        printf("Case #%d\n", i + 1);
        printPunchcard(h, w);
    }
    return 0;
}

void printPunchcard(int h, int w)
{
    if (h > 1)
    {
        printHeader(w, true);
        printBody(w, true);
    }

    for (int i = 1; i < h; ++i)
    {
        printHeader(w, false);
        printBody(w, false);
    }

    printHeader(w, false);
}

void printHeader(int w, bool isRow1)
{
    for (int i = 0; i < w; ++i)
    {
        (isRow1 && i == 0) ? printf("..") : printf("+-");
        if (i == w - 1)
            printf("+\n");
    }
}

void printBody(int w, bool isRow1)
{
    for (int i = 0; i < w; ++i)
    {
        (isRow1 && i == 0) ? printf("..") : printf("|.");
        if (i == w - 1)
            printf("|\n");
    }
}