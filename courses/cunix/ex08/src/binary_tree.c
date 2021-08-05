#include "binary_tree.h"
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

node_t  *allocnode(){
    node_t* node = malloc(sizeof(node_t));
    node->left = NULL;
    node->right = NULL;
    node->key = NULL;
    node->data = NULL;
    return node;
}

node_t  *insert(node_t *root, char *key, void *data){
    if(root == NULL){
        node_t* new_node = allocnode();
        new_node->data = data;
        new_node->key = key;
        return new_node;
    }
    if(strcmp(root->key, key) < 0){
        root->right = insert(root->right, key, data);
        return root;
    }
    if(strcmp(root->key, key) > 0){
        root->left = insert(root->left, key, data);
        return root;
    }
    root->key = key;
    root->data = data;
    return root;
}

void print_node(node_t *node){
    if(node == NULL) return;

    printf("%s: %s\n", node->key, (char*)node->data);
}

void visit_tree(node_t *node, void (*fp)(node_t *root)){
    if(node == NULL) return;

    visit_tree(node->left, fp);
    visit_tree(node->right, fp);
    fp(node);
}

void destroy_tree(node_t *node, void (*fdestroy)(node_t *root)){
    if(node == NULL) return;
    
    destroy_tree(node->left, fdestroy);
    destroy_tree(node->right, fdestroy);
    fdestroy(node);
    free(node);
}
