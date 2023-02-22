/*
# Week 08, 2023 - queue (hard)

This problem was asked by **Netflix**.

Implement a queue using a set of fixed-length arrays.

The queue should support `enqueue`, `dequeue`, and `get_size` operations.
*/

#include <stdio.h>
#include <stdlib.h>

#define MAX 10

typedef struct
{
    int *data;
    int size;
    int head;
    int tail;
} Queue;

Queue *queue_create()
{
    Queue *q = malloc(sizeof(Queue));
    q->data = malloc(sizeof(int) * MAX);
    q->size = 0;
    q->head = 0;
    q->tail = 0;
    return q;
}

void queue_destroy(Queue *q)
{
    free(q->data);
    free(q);
}

void queue_enqueue(Queue *q, int value)
{
    if (q->size == MAX)
    {
        printf("Queue is full\n");
        return;
    }
    q->data[q->tail] = value;
    q->tail = (q->tail + 1) % MAX;
    q->size++;
}

int queue_dequeue(Queue *q)
{
    if (q->size == 0)
    {
        printf("Queue is empty\n");
        return -1;
    }
    int value = q->data[q->head];
    q->head = (q->head + 1) % MAX;
    q->size--;
    return value;
}

void queue_print(Queue *q)
{
    printf("Queue: ");
    for (int i = 0; i < q->size; i++)
    {
        printf("%d ", q->data[(q->head + i) % MAX]);
    }
    printf("\n");
}

int main()
{
    Queue *q = queue_create();
    for (;;)
    {
        int choice;
        printf("1. Enqueue 2. Dequeue 3. Get size 4. Print 5. Exit): ");
        scanf("%d", &choice);
        switch (choice)
        {
        case 1:
        {
            int value;
            printf("Enter value: ");
            scanf("%d", &value);
            queue_enqueue(q, value);
            break;
        }
        case 2:
        {
            int value = queue_dequeue(q);
            printf("Dequeued value: %d\n", value);
            break;
        }
        case 3:
            printf("Size: %d\n", q->size);
            break;
        case 4:
            queue_print(q);
            break;
        case 5:
            queue_destroy(q);
            return 0;
        default:
            break;
        }
    }
}