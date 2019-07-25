#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>
typedef struct Node
{
    int data;
    struct Node * pNext;
}NODE, *PNODE;
PNODE create_list(void);
void traverse_list(PNODE pHead);

int main(void)
{
    PNODE pHead=NULL;
    pHead=create_list();
    traverse_list(pHead);
    return 0;
};
PNODE create_list(void)
{
    int len;
    int val;
    int i;
    
    PNODE pHead=(PNODE)malloc(sizeof(NODE));
    if(pHead==NULL)
    {
        printf("非配失败");
        exit(-1);
    }
    PNODE pTail=pHead;
    printf("请输入要创建的个数");
    scanf("%d",&len);
    
    for(i=0;i<len;i++){
        printf("请输入一个数字");
        scanf("%d",&val);
        PNODE pNew=(PNODE)malloc(sizeof(NODE));
        if(NULL==pNew){
           printf("非配失败");
           exit(-1);
        }
        pNew->data=val;
        pNew->pNext=NULL;
        pTail->pNext=pNew;
        pTail=pNew;
    }
    return pHead;
};
void traverse_list(PNODE pHead){
    if (NULL==pHead){
        return;
    }
    PNODE pTemp=pHead->pNext;
    while(pTemp!=NULL){
        printf("%d \n",pTemp->data);
        pTemp=pTemp->pNext;
    }
}

